package similarity

import (
	"context"
)

// Similarity interface for creating your own similarity algorithm
type Similarity interface {
	Compare(ctx context.Context, str1 string, str2 string) (score float64, err error)
}
