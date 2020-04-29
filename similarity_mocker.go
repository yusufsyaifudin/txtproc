package txtproc

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// similarityMocker using github.com/stretchr/testify
type similarityMocker struct {
	mock.Mock
}

// Compare using testify framework to compare
func (m similarityMocker) Compare(ctx context.Context, str1 string, str2 string) (score float64, err error) {
	args := m.Called(ctx, str1, str2)
	return args.Get(0).(float64), args.Error(1)
}

// SimilarityMocker return mocking object with testify
func SimilarityMocker() *similarityMocker {
	return &similarityMocker{}
}
