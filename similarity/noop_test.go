package similarity

import (
	"context"
	"testing"
)

func TestNoopMatcher_Compare1(t *testing.T) {
	m := Noop()
	score, _ := m.Compare(context.Background(), "", "x")
	if score != -1 {
		t.Error("Noop must return -1 on different string")
		t.Fail()
	}
}

func TestNoopMatcher_Compare2(t *testing.T) {
	m := Noop()
	score, _ := m.Compare(context.Background(), "x", "x")
	if score != -1 {
		t.Error("Noop must return -1 on similar string")
		t.Fail()
	}
}

func TestNoop(t *testing.T) {
	if Noop() == nil {
		t.Error("Noop should return not nil")
		t.Fail()
	}
}
