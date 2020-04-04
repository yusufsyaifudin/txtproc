package main

import (
	"context"
	"fmt"
	"ysf/txtproc"
	"ysf/txtproc/similarity"
)

type badWord struct {
	str            string
	strReplacement string
}

type repo struct {
	badWordsWithPage map[int64]*badWord
}

func (r repo) Get(_ context.Context, dataNumber int64) (dataReplacer *txtproc.ReplacerData, err error) {
	m, exist := r.badWordsWithPage[dataNumber]
	if !exist {
		return &txtproc.ReplacerData{}, nil
	}

	return &txtproc.ReplacerData{
		StringToCompare:   m.str,
		StringReplacement: m.strReplacement,
	}, nil
}

func (r repo) Total(_ context.Context) int64 {
	return int64(len(r.badWordsWithPage))
}

func newRepos(data map[string]string) txtproc.ReplacerDataSeeder {
	var badWordsWithPage = make(map[int64]*badWord)
	var idx = int64(0)
	for k, v := range data {
		idx++

		badWordsWithPage[idx] = &badWord{
			str:            k,
			strReplacement: v,
		}
	}

	return &repo{
		badWordsWithPage: badWordsWithPage,
	}
}

type comparer struct {
	whitelist map[string]bool
}

// Compare will compare the str1 and str2.
// str1 is always current word to compare while str2 is the word from db.
// When we have some list of whitelisted words, we can always compare with str1
func (c comparer) Compare(_ context.Context, str1 string, str2 string) (score float64, err error) {
	// if word is in whitelist, then skip to replace
	if ok, exist := c.whitelist[str1]; ok && exist {
		return 0, nil
	}

	if str1 == str2 {
		return 1, nil
	}

	return 0, nil
}

func newComparerWithWhitelistedWords(whiteList []string) similarity.Similarity {
	var whiteListMap = make(map[string]bool)
	for _, w := range whiteList {
		whiteListMap[w] = true
	}

	return &comparer{
		whitelist: whiteListMap,
	}
}

func main() {

	replacerData := newRepos(map[string]string{
		"anjing": "*nj*ng",
		"asu":    "***",
		"kirik":  "***",
	})

	simAlg := newComparerWithWhitelistedWords([]string{
		"kirik",
	})

	config := txtproc.ReplacerConfig{
		WordToCompare: txtproc.WordNormalized,
		ReplacerData:  replacerData,
		CaseSensitive: false,
		Similarity:    simAlg,
		MinimumScore:  1,
	}

	text := "Dasar anjing asu! kirik"
	words, _ := txtproc.WordReplacer(context.Background(), text, config)
	fmt.Println(words.GetOriginalText())
	fmt.Println(words.GetNormalizedText())
	fmt.Println(words.GetReplacedText())
}
