package txtproc

import (
	"context"
	"fmt"
	"math"
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

	var (
		goodWordTotal    = filterConfig.GoodWordsData.Total(ctx)
		goodWordPerBatch = filterConfig.GoodWordsData.PerBatch(ctx)
		goodWordNumBatch = int64(math.Ceil(float64(goodWordTotal) / float64(goodWordPerBatch)))
	)

	for i := int64(1); i <= goodWordNumBatch; i++ {

		// call get data here to minimize db call if the data from db or external storage
		var goodWordDataReplacer = make([]*ReplacerData, 0)
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

			for _, goodWord := range goodWordDataReplacer {
				if goodWord == nil {
					continue
				}

				var score float64 = 0
				score, err = badWordSimFunc.Compare(ctx, currentWord, goodWord.StringToCompare)
				if err != nil {
					err = fmt.Errorf(
						"error compare string '%s' vs '%s': %s",
						currentWord, goodWord.StringToCompare, err.Error(),
					)
					return
				}

				// when score is higher or equal than minimum score in config, replace it
				if score >= filterConfig.BadWordsMinimumScore {
					w.replaced = goodWord.StringReplacement
					w.isReplaced = true
					continue
				}
			}

		}
	}

	var (
		badWordTotal    = filterConfig.BadWordsData.Total(ctx)
		badWordPerBatch = filterConfig.BadWordsData.PerBatch(ctx)
		badWordNumBatch = int64(math.Ceil(float64(badWordTotal) / float64(badWordPerBatch)))
	)

	for j := int64(1); j <= badWordNumBatch; j++ {

		// call get data here to minimize db call if the data from db or external storage
		var badWordDataReplacer = make([]*ReplacerData, 0)
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

			for _, badWord := range badWordDataReplacer {
				if badWord == nil {
					continue
				}

				var score float64 = 0
				score, err = badWordSimFunc.Compare(ctx, currentWord, badWord.StringToCompare)
				if err != nil {
					err = fmt.Errorf(
						"error compare string '%s' vs '%s': %s",
						currentWord, badWord.StringToCompare, err.Error(),
					)
					return
				}

				// when score is higher or equal than minimum score in config, replace it
				if score >= filterConfig.BadWordsMinimumScore {
					w.replaced = badWord.StringReplacement
					w.isReplaced = true
					continue
				}
			}
		}

	}

	return
}
