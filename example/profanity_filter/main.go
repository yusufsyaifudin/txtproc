package main

import (
	"context"
	"fmt"
	"ysf/txtproc"
	"ysf/txtproc/repo"
	"ysf/txtproc/similarity"
)

func main() {

	goodWordsData := repo.InMemory(map[string]string{
		"assassin": "assassin", // assassin contains 'ass' but it must be skipped
	}, 100)

	badWordsData := repo.InMemory(map[string]string{
		"anjing": "*nj*ng",
		"asu":    "***",
		"kirik":  "***",
	}, 100)

	config := txtproc.ProfanityFilterConfig{
		GoodWordsReplacerConfig: txtproc.WordReplacerConfig{
			WordToCompare:        txtproc.WordNormalized,
			CaseSensitive:        true,
			ReplacerDataSeeder:   goodWordsData,
			SimilarityFunc:       similarity.ExactMatcher(),
			ReplacerMinimumScore: 1,
		},
		BadWordsReplacerConfig: txtproc.WordReplacerConfig{
			WordToCompare:        txtproc.WordNormalized,
			CaseSensitive:        true,
			ReplacerDataSeeder:   badWordsData,
			SimilarityFunc:       similarity.ExactMatcher(),
			ReplacerMinimumScore: 1,
		},
	}

	text := "Dasar anjing asu! kirik"
	words, _ := txtproc.ProfanityFilter(context.Background(), text, config)
	fmt.Println(words.GetOriginalText())
	fmt.Println(words.GetNormalizedText())
	fmt.Println(words.GetReplacedText())

	words.GetWords().Get()[0].SetReplaced("anu")
	fmt.Println(words.GetOriginalText())   // won't change
	fmt.Println(words.GetNormalizedText()) // won't change
	fmt.Println(words.GetReplacedText())
}
