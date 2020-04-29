package similarity

import (
	"context"

	"github.com/opentracing/opentracing-go"
)

// exactMatcher matcher with == comparison
type exactMatcher struct{}

// Compare return 1 when str1 == str2, otherwise 0
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

// ExactMatcher return similarity algorithm with exact comparison between two string
func ExactMatcher() Similarity {
	return &exactMatcher{}
}
