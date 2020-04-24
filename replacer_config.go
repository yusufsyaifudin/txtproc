package txtproc

import "ysf/txtproc/similarity"

// ReplacerConfig configures the profanity filter struct.
type ReplacerConfig struct {
	// this to ensure that we don't call the replacer when not needed
	// will be true on function `WordReplacer` false by default.
	enabled bool

	WordToCompare WordType              // which word to compare
	ReplacerData  ReplacerDataSeeder    // seed data for the replacement
	CaseSensitive bool                  // default false
	Similarity    similarity.Similarity // similarity distance algorithm
	MinimumScore  float64               // minimum score to make the filter replaces the string, between 0 to 1
}
