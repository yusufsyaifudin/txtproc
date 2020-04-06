package txtproc

import (
	"context"
	"fmt"
	"strings"
	"ysf/txtproc/similarity"

	"github.com/opentracing/opentracing-go"
)

// ProfanityFilter will return `mappedStrings` and replace some text based on your algorithm,
// then set it into `GetReplaced()` on the `MappedString`.
// The operation will takes O(N + M + P) where N is the length of the `text`,
// M is the length of the good words data (whitelisted data), and P is the length of the bad words data.
//
// It will replace the first match of data based on computed similarity score.
// When comparing string from your database and score return higher or equal
// than minimum score stated in `ProfanityFilterConfig`, it will be replaced.
// It's up to you to order the string when `ReplacerDataSeeder.Get` is called.
func ProfanityFilter(ctx context.Context, text string, filterConfig ProfanityFilterConfig) (mappedStrings MappedStrings, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "txtproc.ProfanityFilter")
	defer func() {
		ctx.Done()
		span.Finish()
	}()

	if filterConfig.GoodWordsData == nil {
		filterConfig.GoodWordsData = newReplacerDataDefault()
	}

	var goodWordSimFunc = filterConfig.GoodWordsSimFunc
	if goodWordSimFunc == nil {
		goodWordSimFunc = similarity.Noop()
	}

	if filterConfig.GoodWordsMinimumScore < 0 {
		panic("minimum score for good words filter must greater or equal than 0")
	}

	if filterConfig.GoodWordsMinimumScore > 1 {
		panic("minimum score for good words filter  must lower or equal than 1")
	}

	if filterConfig.BadWordsData == nil {
		filterConfig.BadWordsData = newReplacerDataDefault()
	}

	var badWordSimFunc = filterConfig.BadWordsSimFunc
	if badWordSimFunc == nil {
		badWordSimFunc = similarity.Noop()
	}

	if filterConfig.BadWordsMinimumScore < 0 {
		panic("minimum score for bad words filter must greater or equal than 0")
	}

	if filterConfig.BadWordsMinimumScore > 1 {
		panic("minimum score for bad words filter  must lower or equal than 1")
	}

	mappedStrings, err = wordSeparator(ctx, text)
	if err != nil {
		return
	}

	var i = int64(0)
	for i < filterConfig.GoodWordsData.Total(ctx) {
		i++

		// call get data here to minimize db call if the data from db or external storage
		var goodWordDataReplacer = &ReplacerData{}
		goodWordDataReplacer, err = filterConfig.GoodWordsData.Get(ctx, i)
		if err != nil {
			return
		}

		for _, w := range mappedStrings.GetMappedString() {
			if w.IsReplaced() {
				continue
			}

			var currentWord = w.GetOriginal()
			if filterConfig.WordToCompare == WordNormalized {
				currentWord = w.GetNormalized()
			}

			if !filterConfig.CaseSensitive {
				currentWord = strings.ToLower(currentWord)
			}

			var score float64 = 0
			score, err = badWordSimFunc.Compare(ctx, currentWord, goodWordDataReplacer.StringToCompare)
			if err != nil {
				err = fmt.Errorf(
					"error compare string '%s' vs '%s': %s",
					currentWord, goodWordDataReplacer.StringToCompare, err.Error(),
				)
				return
			}

			// when score is higher or equal than minimum score in config, replace it
			if score >= filterConfig.BadWordsMinimumScore {
				w.replaced = goodWordDataReplacer.StringReplacement
				w.isReplaced = true
				continue
			}
		}
	}

	var j = int64(0)
	for j < filterConfig.BadWordsData.Total(ctx) {
		j++

		// call get data here to minimize db call if the data from db or external storage
		var badWordDataReplacer = &ReplacerData{}
		badWordDataReplacer, err = filterConfig.BadWordsData.Get(ctx, j)
		if err != nil {
			return
		}

		for _, w := range mappedStrings.GetMappedString() {
			if w.IsReplaced() {
				continue
			}

			var currentWord = w.GetOriginal()
			if filterConfig.WordToCompare == WordNormalized {
				currentWord = w.GetNormalized()
			}

			if !filterConfig.CaseSensitive {
				currentWord = strings.ToLower(currentWord)
			}

			var score float64 = 0
			score, err = badWordSimFunc.Compare(ctx, currentWord, badWordDataReplacer.StringToCompare)
			if err != nil {
				err = fmt.Errorf(
					"error compare string '%s' vs '%s': %s",
					currentWord, badWordDataReplacer.StringToCompare, err.Error(),
				)
				return
			}

			// when score is higher or equal than minimum score in config, replace it
			if score >= filterConfig.BadWordsMinimumScore {
				w.replaced = badWordDataReplacer.StringReplacement
				w.isReplaced = true
				continue
			}
		}

	}

	return
}
