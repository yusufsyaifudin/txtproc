package txtproc

import (
	"context"
	"fmt"
	"math"
	"strings"
	"ysf/txtproc/similarity"
	"ysf/txtproc/word"

	"github.com/opentracing/opentracing-go"
)

// WordReplacer will replace mappedStrings using config.
// This takes O(P*M) = O(N*N) = O(N^2)
// M is the number of words in text (keep in mind that space is one word).
// P is the number of batch (page) of replacer data = Total data / Per batch of data
func WordReplacer(ctx context.Context, mappedStrings *word.Text, conf WordReplacerConfig) (err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "txtproc.WordReplacer")
	defer func() {
		ctx.Done()
		span.Finish()
	}()

	if mappedStrings == nil {
		err = ErrMappedStringsNil
		return
	}

	if conf.ReplacerMinimumScore < 0 || conf.ReplacerMinimumScore > 1 {
		err = ErrMinReplacerScore
		return
	}

	replacerSeeder := newReplacerDataDefault()
	if conf.ReplacerDataSeeder != nil {
		replacerSeeder = conf.ReplacerDataSeeder
	}

	simFunc := similarity.Noop()
	if conf.SimilarityFunc != nil {
		simFunc = conf.SimilarityFunc
	}

	var (
		replacerTotal       = replacerSeeder.Total(ctx)
		replacerGetPerBatch = replacerSeeder.PerBatch(ctx)
		replacerNumOfBatch  = int64(math.Ceil(float64(replacerTotal) / float64(replacerGetPerBatch)))
	)

	for i := int64(1); i <= replacerNumOfBatch; i++ {

		// call get data here to minimize db call if the data from db or external storage
		var replacerData = make([]ReplacerData, 0)
		replacerData, err = replacerSeeder.Get(ctx, i)
		if err != nil {
			return
		}

		for _, w := range mappedStrings.GetWords().Get() {
			if w.IsReplaced() {
				continue
			}

			var currentWord = w.GetOriginal()
			if conf.WordToCompare == WordNormalized {
				currentWord = w.GetNormalized()
			}

			if !conf.CaseSensitive {
				currentWord = strings.ToLower(currentWord)
			}

			for _, replacerDatum := range replacerData {
				var score float64 = 0
				score, err = simFunc.Compare(ctx, currentWord, replacerDatum.StringToCompare)
				if err != nil {
					err = fmt.Errorf(
						"error compare string '%s' vs '%s': %s",
						currentWord, replacerDatum.StringToCompare, err.Error(),
					)
					return
				}

				// when score is higher or equal than minimum score in config, replace it
				if score >= conf.ReplacerMinimumScore {
					w.SetReplaced(replacerDatum.StringReplacement)
					continue
				}
			}

		}
	}

	return
}
