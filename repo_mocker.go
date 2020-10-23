package txtproc

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// repoMocker using github.com/stretchr/testify
type repoMocker struct {
	mock.Mock
}

// Get mock the returning data each batch
func (m *repoMocker) Get(ctx context.Context, batch int64) (dataReplacer []ReplacerData, err error) {
	args := m.Called(ctx, batch)
	return args.Get(0).([]ReplacerData), args.Error(1)
}

// Total mock the number of total data
func (m *repoMocker) Total(ctx context.Context) int64 {
	args := m.Called(ctx)
	return args.Get(0).(int64)
}

// PerBatch mock the number of per batch
func (m *repoMocker) PerBatch(ctx context.Context) int64 {
	args := m.Called(ctx)
	return args.Get(0).(int64)
}

// RepoMocker return testify mock instance
func RepoMocker() *repoMocker {
	return &repoMocker{}
}
