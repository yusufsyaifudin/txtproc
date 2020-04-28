package main

import (
	"context"
	"fmt"
	"ysf/txtproc"
	"ysf/txtproc/similarity"
)

type dataCollection struct {
	str            string
	strReplacement string
}

type repo struct {
	sliceOfWords []*dataCollection
}

func (r repo) Get(ctx context.Context, batch int64) (dataReplacer []*txtproc.ReplacerData, err error) {
	dataReplacer = []*txtproc.ReplacerData{}

	var paginate = func(x []*dataCollection, offset int64, limit int64) []*dataCollection {
		if offset > int64(len(x)) {
			offset = int64(len(x))
		}

		end := offset + limit
		if end > int64(len(x)) {
			end = int64(len(x))
		}

		return x[offset:end]
	}

	offset := (batch - 1) * int64(len(r.sliceOfWords))
	data := paginate(r.sliceOfWords, offset, r.PerBatch(ctx))
	for _, d := range data {
		dataReplacer = append(dataReplacer, &txtproc.ReplacerData{
			StringToCompare:   d.str,
			StringReplacement: d.strReplacement,
		})
	}

	return
}

func (r repo) Total(_ context.Context) int64 {
	return int64(len(r.sliceOfWords))
}

func (r repo) PerBatch(_ context.Context) int64 {
	return int64(len(r.sliceOfWords))
}

func newRepos(data map[string]string) txtproc.ReplacerDataSeeder {
	var dataSlices = make([]*dataCollection, 0)

	for k, v := range data {
		dataSlices = append(dataSlices, &dataCollection{
			str:            k,
			strReplacement: v,
		})
	}

	return &repo{
		sliceOfWords: dataSlices,
	}
}

type comparer struct{}

// Compare will compare the str1 and str2.
// str1 is always current word to compare while str2 is the word from db.
// When we have some list of whitelisted words, we can always compare with str1
func (c comparer) Compare(_ context.Context, str1 string, str2 string) (score float64, err error) {
	if str1 == str2 {
		return 1, nil
	}
	return 0, nil
}

func newComparer() similarity.Similarity {
	return &comparer{}
}

func main() {

	goodWordsData := newRepos(map[string]string{
		"assassin": "assassin", // assassin contains 'ass' but it must be skipped
	})

	badWordsData := newRepos(map[string]string{
		"anjing": "*nj*ng",
		"asu":    "***",
		"kirik":  "***",
	})

	config := txtproc.ProfanityFilterConfig{
		WordToCompare:         txtproc.WordNormalized,
		CaseSensitive:         false,
		GoodWordsData:         goodWordsData,
		GoodWordsSimFunc:      newComparer(),
		GoodWordsMinimumScore: 1,
		BadWordsData:          badWordsData,
		BadWordsSimFunc:       newComparer(),
		BadWordsMinimumScore:  1,
	}

	text := "Dasar anjing asu! kirik"
	words, _ := txtproc.ProfanityFilter(context.Background(), text, config)
	fmt.Println(words.GetOriginalText())
	fmt.Println(words.GetNormalizedText())
	fmt.Println(words.GetReplacedText())
}
