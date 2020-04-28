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
	wordsMapWithPage map[int64]*dataCollection
}

func (r repo) Get(_ context.Context, dataNumber int64) (dataReplacer *txtproc.ReplacerData, err error) {
	m, exist := r.wordsMapWithPage[dataNumber]
	if !exist {
		return &txtproc.ReplacerData{}, nil
	}

	return &txtproc.ReplacerData{
		StringToCompare:   m.str,
		StringReplacement: m.strReplacement,
	}, nil
}

func (r repo) Total(_ context.Context) int64 {
	return int64(len(r.wordsMapWithPage))
}

func newRepos(data map[string]string) txtproc.ReplacerDataSeeder {
	var dataPaginate = make(map[int64]*dataCollection)
	var idx = int64(0)
	for k, v := range data {
		idx++

		dataPaginate[idx] = &dataCollection{
			str:            k,
			strReplacement: v,
		}
	}

	return &repo{
		wordsMapWithPage: dataPaginate,
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
