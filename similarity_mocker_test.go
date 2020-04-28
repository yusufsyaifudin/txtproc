package txtproc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestSimilarityMocker_Compare(t *testing.T) {
	m := SimilarityMocker()
	m.On("Compare", mock.Anything, "x", "x").Return(0.0, nil).Once()
	score, _ := m.Compare(context.Background(), "x", "x")
	if score != 0 {
		t.Error("SimilarityMocker must return 0 as created expectation")
		t.Fail()
	}
}

func TestSimilarityMocker(t *testing.T) {
	if SimilarityMocker() == nil {
		t.Error("SimilarityMocker should return not nil")
		t.Fail()
	}
}
