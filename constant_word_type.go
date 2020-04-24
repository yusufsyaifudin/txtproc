package txtproc

type WordType float64

const (
	_ = iota // ignore first value by assigning to blank identifier

	// is a original word that shown in the sequence
	WordOriginal WordType = 1 << (10 * iota)

	// is a word that consist only alphanumeric character. Other char like punctuation is removed.
	WordNormalized
)
