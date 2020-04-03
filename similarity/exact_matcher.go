package similarity

import (
	"context"

	"github.com/opentracing/opentracing-go"
)

type exactMatcher struct{}

func (e exactMatcher) Compare(ctx context.Context, str1 string, str2 string) (score float64, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "similarity.Compare")
	defer func() {
		ctx.Done()
		span.Finish()
	}()

	if str1 == str2 {
		return 1, nil
	}

	return 0, nil
}

func ExactMatcher() Similarity {
	return &exactMatcher{}
}
