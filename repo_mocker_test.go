package txtproc

import (
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestRepoMocker_Get(t *testing.T) {
	want := []*ReplacerData{
		{
			StringToCompare:   "word",
			StringReplacement: "replacement",
		},
	}

	m := RepoMocker()
	m.On("Get", mock.Anything, int64(1)).Return(want, nil).Once()
	got, _ := m.Get(context.Background(), 1)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("RepoMocker.Get want %v got %v", want, got)
		t.Fail()
	}
}

func TestRepoMocker_Total(t *testing.T) {
	m := RepoMocker()
	m.On("Total", mock.Anything).Return(int64(1)).Once()
	got := m.Total(context.Background())
	if got != 1 {
		t.Errorf("RepoMocker.Total want %v got %v", 1, got)
		t.Fail()
	}
}

func TestRepoMocker_PerBatch(t *testing.T) {
	m := RepoMocker()
	m.On("PerBatch", mock.Anything).Return(int64(1)).Once()
	got := m.PerBatch(context.Background())
	if got != 1 {
		t.Errorf("RepoMocker.PerBatch want %v got %v", 1, got)
		t.Fail()
	}
}

func TestRepoMocker(t *testing.T) {
	if RepoMocker() == nil {
		t.Error("RepoMocker should return not nil")
		t.Fail()
	}
}
