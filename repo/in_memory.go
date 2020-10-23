package repo

import (
	"context"
	"ysf/txtproc"

	"github.com/opentracing/opentracing-go"
)

// inMemory is a struct handling inMemory in-memory database
type inMemory struct {
	length       int64
	perBatch     int
	sliceOfWords []txtproc.ReplacerData
}

// Get return the data in slice of registered word in inMemory database
func (n inMemory) Get(ctx context.Context, batch int64) (dataReplacer []txtproc.ReplacerData, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "inMemory.Get")
	defer func() {
		ctx.Done()
		span.Finish()
	}()

	offset := (batch - 1) * int64(n.perBatch)
	dataReplacer = paginate(n.sliceOfWords, offset, int64(n.perBatch))
	return
}

// Total return total length of data
func (n inMemory) Total(ctx context.Context) int64 {
	span, ctx := opentracing.StartSpanFromContext(ctx, "inMemory.Total")
	defer func() {
		ctx.Done()
		span.Finish()
	}()

	return n.length
}

// PerBatch return the length of slices, makes it 1 iteration only
func (n inMemory) PerBatch(ctx context.Context) int64 {
	span, ctx := opentracing.StartSpanFromContext(ctx, "inMemory.PerBatch")
	defer func() {
		ctx.Done()
		span.Finish()
	}()

	return int64(n.perBatch)
}

// InMemory is repository using Go native slice type
func InMemory(data map[string]string, perBatch int) txtproc.ReplacerDataSeeder {
	var dataSlices = make([]txtproc.ReplacerData, 0)

	for k, v := range data {
		dataSlices = append(dataSlices, txtproc.ReplacerData{
			StringToCompare:   k,
			StringReplacement: v,
		})
	}

	return &inMemory{
		length:       int64(len(dataSlices)),
		perBatch:     perBatch,
		sliceOfWords: dataSlices,
	}
}

// paginate paginate replacer data
var paginate = func(x []txtproc.ReplacerData, offset int64, limit int64) []txtproc.ReplacerData {
	if offset > int64(len(x)) {
		offset = int64(len(x))
	}

	end := offset + limit
	if end > int64(len(x)) {
		end = int64(len(x))
	}

	return x[offset:end]
}
