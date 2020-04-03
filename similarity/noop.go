package similarity

import (
	"context"
)

// noopMatcher default matcher when replacing word
type noopMatcher struct{}

// Compare will always return -1 and no error
func (n noopMatcher) Compare(_ context.Context, _ string, _ string) (score float64, err error) {
	return -1, nil
}

// Noop similarity checker that don't do anything
func Noop() Similarity {
	return &noopMatcher{}
}
