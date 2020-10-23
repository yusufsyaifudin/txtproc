package repo

import (
	"context"
	"reflect"
	"testing"
	"ysf/txtproc"
)

var inMemoryDataTest = map[string]string{
	"word": "replacement",
}

const inMemoryPerBatchTest = 100

func TestInMemory_Get(t *testing.T) {
	r := InMemory(inMemoryDataTest, inMemoryPerBatchTest)
	data, err := r.Get(context.Background(), 1)
	if err != nil {
		t.Errorf("InMemory.Get should not be error, but got %v", err)
		t.Fail()
		return
	}

	if len(data) != len(inMemoryDataTest) {
		t.Errorf("InMemory.Get should return same length of input")
		t.Fail()
		return
	}
}

func TestInMemory_Total(t *testing.T) {
	r := InMemory(inMemoryDataTest, inMemoryPerBatchTest)
	total := r.Total(context.Background())
	if total != int64(len(inMemoryDataTest)) {
		t.Errorf("InMemory.Total should return same length of input")
		t.Fail()
		return
	}
}

func TestInMemory_PerBatch(t *testing.T) {
	r := InMemory(inMemoryDataTest, inMemoryPerBatchTest)
	perBatch := r.PerBatch(context.Background())
	if perBatch != inMemoryPerBatchTest {
		t.Errorf("InMemory.PerBatch should return per batch as specified when initating struct")
		t.Fail()
		return
	}
}

func TestInMemory(t *testing.T) {
	if InMemory(map[string]string{}, inMemoryPerBatchTest) == nil {
		t.Errorf("InMemory should not be nit")
		t.Fail()
	}
}

// Test_paginate when offset is more than length of data, so end will be more than data length too
func Test_paginate(t *testing.T) {
	data := []txtproc.ReplacerData{
		{
			StringToCompare:   "word",
			StringReplacement: "replacement",
		},
	}

	x := paginate(data, 2, 1)
	want := make([]txtproc.ReplacerData, 0)
	if !reflect.DeepEqual(x, want) {
		t.Errorf("want %v got %v", want, x)
		t.Fail()
	}
}
