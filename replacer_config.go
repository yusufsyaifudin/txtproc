package txtproc

import "ysf/txtproc/similarity"

// ReplacerConfig configures the profanity filter struct.
type ReplacerConfig struct {
	WordToCompare WordType              // which word to compare
	ReplacerData  ReplacerData          // seed data for the replacement
	CaseSensitive bool                  // default false
	Similarity    similarity.Similarity // similarity distance algorithm
	MinimumScore  float64               // minimum score to make the filter replaces the string, between 0 to 1
}
