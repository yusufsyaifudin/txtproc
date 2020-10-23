package txtproc

import (
	"context"
	"ysf/txtproc/word"

	"github.com/opentracing/opentracing-go"
)

// ProfanityFilter will return `mappedStrings` and replace some text based on your algorithm,
// then set it into `GetReplaced()` on the `Word`.
// The operation takes O(N) + O(P*M) + O(Q*M)
// So, it O(N) + O(N^2) + O(N^2) because M is more dominant, and substitute M to N.
// O(N) + 2O(N^2), This is O(max(2(N^2), N)) which is 2O(N^2)
//
// Where N is the length of text.
// M is the number of words in text (keep in mind that space is one word).
// P is the number of batch (page) of good words = Total good words / Per batch of good words
// Q is the number of batch (page) of bad words = Total bad words / Per batch of bad words
//
// It will replace the first match of data based on computed similarity score.
// When comparing string from your database and score return higher or equal
// than minimum score stated in `WordReplacerConfig`, it will be replaced.
// It's up to you to order the string when `ReplacerDataSeeder.Get` is called.
func ProfanityFilter(ctx context.Context, text string, filterConfig ProfanityFilterConfig) (mappedStrings *word.Text, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "txtproc.ProfanityFilter")
	defer func() {
		ctx.Done()
		span.Finish()
	}()

	mappedStrings, err = WordSeparator(ctx, text)
	if err != nil {
		return
	}

	err = WordReplacer(ctx, mappedStrings, filterConfig.GoodWordsReplacerConfig)
	if err != nil {
		return
	}

	err = WordReplacer(ctx, mappedStrings, filterConfig.BadWordsReplacerConfig)
	if err != nil {
		return
	}

	return
}
