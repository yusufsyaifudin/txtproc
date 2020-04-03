package txtproc

import (
	"context"

	"github.com/opentracing/opentracing-go"
)

// ReplacerData is a collection of string (bad words) that need to be replaced.
// This is like an `map[string]string` where string to compare and string replacement can be fetch using `Get` method.
// By defining the interface (rather than using Go map type), we can implement using the best approach we can imagine.
// For example, by always querying into Redis rather than
// This can be implemented using database, in-memory, or anything.
// By having this interface, we can just load/query what we need in that time, not pre-loading all data from database
// to Go map (but you can still do it).
type ReplacerData interface {
	// Get data from database (in-memory, or SQL or noSql) by sequence.
	// This function will be called `Total()` times,
	// and the `dataNumber` will be contain value range between 1 - `Total()`
	// You can treat `dataNumber` as the offset limit, for example:
	// offset = (page - 1) * limit, where limit always 1
	// So, for page 3, the limit = 1, and offset = 2.
	// Function will return `strToCompare` as the compared string, and `replacement` as the replacement string.
	// It is up to you whether you want to return partial replacement like 'f*ck' or full replacement like '****'
	Get(ctx context.Context, dataNumber int64) (strToCompare, replacement string, err error)

	// Total returns the number of total data that need to be checked in collection of bad words.
	Total(ctx context.Context) int64
}

// replacerDataDefault will be used when no `ReplacerData` passed in `ReplacerConfig`
type replacerDataDefault struct{}

// Get will not return any string
func (r replacerDataDefault) Get(ctx context.Context, _ int64) (strToCompare, replacement string, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "replacerDataDefault.Get")
	defer func() {
		ctx.Done()
		span.Finish()
	}()

	return
}

// Total will always return 1 on default, so only 1 get data (one iteration) to call
func (r replacerDataDefault) Total(ctx context.Context) int64 {
	span, ctx := opentracing.StartSpanFromContext(ctx, "replacerDataDefault.Total")
	defer func() {
		ctx.Done()
		span.Finish()
	}()

	return 1
}

// newReplacerDataDefault default value for replace the data
func newReplacerDataDefault() ReplacerData {
	return &replacerDataDefault{}
}
