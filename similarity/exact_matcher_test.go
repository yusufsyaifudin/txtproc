package similarity

import (
	"context"
	"testing"
)

func TestExactMatcher_Compare1(t *testing.T) {
	m := ExactMatcher()
	score, _ := m.Compare(context.Background(), "", "x")
	if score != 0 {
		t.Error("ExactMatcher must return 0 on different string")
		t.Fail()
	}
}

func TestExactMatcher_Compare2(t *testing.T) {
	m := ExactMatcher()
	score, _ := m.Compare(context.Background(), "x", "x")
	if score != 1 {
		t.Error("ExactMatcher must return 1 on similar string")
		t.Fail()
	}
}

func TestExactMatcher(t *testing.T) {
	if ExactMatcher() == nil {
		t.Error("ExactMatcher should return not nil")
		t.Fail()
	}
}
