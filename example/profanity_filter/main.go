package main

import (
	"context"
	"fmt"
	"ysf/txtproc"
	"ysf/txtproc/similarity"
)

type repo struct {
	badWordsWithPage map[int64]map[string]string
}

func (r repo) Get(_ context.Context, dataNumber int64) (dataReplacer *txtproc.ReplacerData, err error) {
	m, exist := r.badWordsWithPage[dataNumber]
	if !exist {
		return &txtproc.ReplacerData{}, nil
	}

	for k, v := range m {
		fmt.Println(k, v)
		return &txtproc.ReplacerData{
			StringToCompare:   k,
			StringReplacement: v,
		}, nil
	}

	return &txtproc.ReplacerData{}, nil
}

func (r repo) Total(_ context.Context) int64 {
	return int64(len(r.badWordsWithPage))
}

func newRepos(data map[string]string) txtproc.ReplacerDataSeeder {
	var badWordsWithPage = make(map[int64]map[string]string)
	var idx = int64(0)
	for k, v := range data {
		idx++

		badWordsWithPage[idx] = map[string]string{
			k: v,
		}
	}

	return &repo{
		badWordsWithPage: badWordsWithPage,
	}
}

func main() {

	replacerData := newRepos(map[string]string{
		"anjing": "*nj*ng",
		"asu":    "***",
	})

	config := txtproc.ReplacerConfig{
		WordToCompare: txtproc.WordNormalized,
		ReplacerData:  replacerData,
		CaseSensitive: false,
		Similarity:    similarity.ExactMatcher(),
		MinimumScore:  1,
	}

	text := "Dasar anjing asu!"
	words, _ := txtproc.WordReplacer(context.Background(), text, config)
	fmt.Println(words.GetOriginalText())
	fmt.Println(words.GetNormalizedText())
	fmt.Println(words.GetReplacedText())
}
