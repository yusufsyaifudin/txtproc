package txtproc

import (
	"ysf/txtproc/similarity"
)

// WordReplacerConfig configures for replace word
type WordReplacerConfig struct {
	WordToCompare WordType // which word to compare
	CaseSensitive bool     // default false

	ReplacerDataSeeder   ReplacerDataSeeder    // seed data for whitelisted data
	SimilarityFunc       similarity.Similarity // similarity distance algorithm
	ReplacerMinimumScore float64               // minimum score to make the filter replaces the string, between 0 to 1
}
