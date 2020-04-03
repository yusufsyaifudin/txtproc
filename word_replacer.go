package txtproc

import (
	"context"

	"github.com/opentracing/opentracing-go"
)

// WordReplacer will return `mappedStrings` and replace some text based on your algorithm,
// then set it into `GetReplaced()` on the `MappedString`
func WordReplacer(ctx context.Context, text string, replacer ReplacerConfig) (mappedStrings MappedStrings, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "txtproc.WordReplacer")
	defer func() {
		ctx.Done()
		span.Finish()
	}()

	return wordSeparator(ctx, text, replacer)
}