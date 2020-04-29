package txtproc

// ProfanityFilterConfig configures the profanity filter struct.
type ProfanityFilterConfig struct {
	GoodWordsReplacerConfig WordReplacerConfig // seed data for whitelisted data
	BadWordsReplacerConfig  WordReplacerConfig // seed data for the replacement of bad words
}
