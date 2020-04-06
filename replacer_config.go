package txtproc

import (
	"ysf/txtproc/similarity"
)

// ProfanityFilterConfig configures the profanity filter struct.
type ProfanityFilterConfig struct {
	WordToCompare WordType // which word to compare
	CaseSensitive bool     // default false

	GoodWordsData         ReplacerDataSeeder    // seed data for whitelisted data
	GoodWordsSimFunc      similarity.Similarity // similarity distance algorithm
	GoodWordsMinimumScore float64               // minimum score to make the filter replaces the string, between 0 to 1

	BadWordsData         ReplacerDataSeeder    // seed data for the replacement of bad words
	BadWordsSimFunc      similarity.Similarity // similarity distance algorithm
	BadWordsMinimumScore float64               // minimum score to make the filter replaces the string, between 0 to 1
}
