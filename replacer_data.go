package txtproc

import (
	"context"

	"github.com/opentracing/opentracing-go"
)

// BadWordsData is a struct to map what data to be compared, replaced to what
// and other information you want to carry for example the primary key, etc.
type ReplacerData struct {
	StringToCompare   string
	StringReplacement string
}

// ReplacerDataSeeder is a collection of string (bad words) that need to be replaced.
// This is like an `map[string]string` where string to compare and string replacement can be fetch using `Get` method.
// By defining the interface (rather than using Go map type), we can implement using the best approach we can imagine.
// For example, by always querying into Redis rather than
// This can be implemented using database, in-memory, or anything.
// By having this interface, we can just load/query what we need in that time, not pre-loading all data from database
// to Go map (but you can still do it).
type ReplacerDataSeeder interface {
	// Get data from database (in-memory, or SQL or noSql) by sequence.
	// This function will be called `Total()/PerBatch()` times,
	// and the `batch` will be contain value range between 1 - `Total()/PerBatch()`
	// You can treat `batch` as the offset limit, for example:
	// Total = 100, PerBatch = 10, it will iterate between 1 to (100/10 = 10).
	// More example, when Total = 100, PerBatch = 100, it will iterate once only.
	// It your responsibility to tune how many data fetched per batch, as it will only matter with query speed.
	// You can query SQL something like this:
	// offset = (batch - 1) * PerBatch()
	// SELECT * FROM bad_words LIMIT :PerBatch() OFFSET :offset ORDER BY id DESC;
	// Function will return `strToCompare` as the compared string, and `replacement` as the replacement string.
	// It is up to you whether you want to return partial replacement like 'f*ck' or full replacement like '****'
	Get(ctx context.Context, batch int64) (dataReplacer []ReplacerData, err error)

	// Total returns the number of total data that need to be checked in collection of bad words.
	Total(ctx context.Context) int64

	// PerBatch will return how many data retrieved per `Get`.
	PerBatch(ctx context.Context) int64
}

// replacerDataDefault will be used when no `ReplacerDataSeeder` passed in `WordReplacerConfig`
type replacerDataDefault struct{}

// Get will not return any string
func (r replacerDataDefault) Get(ctx context.Context, _ int64) (dataReplacer []ReplacerData, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "replacerDataDefault.Get")
	defer func() {
		ctx.Done()
		span.Finish()
	}()

	dataReplacer = []ReplacerData{}
	return
}

// Total will always return 1 on default, so it will not do any iteration
func (r replacerDataDefault) Total(ctx context.Context) int64 {
	span, ctx := opentracing.StartSpanFromContext(ctx, "replacerDataDefault.Total")
	defer func() {
		ctx.Done()
		span.Finish()
	}()

	return 0
}

// PerBatch return 1 since Total only return 1. So the iteration will be 0 (total/per batch = 0/1 = 0).
func (r replacerDataDefault) PerBatch(ctx context.Context) int64 {
	span, ctx := opentracing.StartSpanFromContext(ctx, "replacerDataDefault.PerBatch")
	defer func() {
		ctx.Done()
		span.Finish()
	}()

	return 1
}

// newReplacerDataDefault default value for replace the data
func newReplacerDataDefault() ReplacerDataSeeder {
	return &replacerDataDefault{}
}
