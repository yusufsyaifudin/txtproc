package txtproc

import (
	"context"
	"testing"
)

func TestReplacerDataDefault_Get(t *testing.T) {
	replacer := new(replacerDataDefault)

	_, err := replacer.Get(context.Background(), 1)
	if err != nil {
		t.Error("replacerDataDefault.Get should return not nil")
		t.Fail()
	}
}

func TestReplacerDataDefault_Total(t *testing.T) {
	replacer := new(replacerDataDefault)

	total := replacer.Total(context.Background())
	if total != 0 {
		t.Error("replacerDataDefault.Total should return 1")
		t.Fail()
	}
}

func TestReplacerDataDefault_PerBatch(t *testing.T) {
	replacer := new(replacerDataDefault)

	batch := replacer.PerBatch(context.Background())
	if batch != 1 {
		t.Error("replacerDataDefault.PerBatch should return 1")
		t.Fail()
	}
}

func Test_newReplacerDataDefault(t *testing.T) {
	if newReplacerDataDefault() == nil {
		t.Error("newReplacerDataDefault should return not nil")
		t.Fail()
	}
}
