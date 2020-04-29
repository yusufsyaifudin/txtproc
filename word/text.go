package word

import (
	"bytes"
)

// Text is collection of word
type Text struct {
	originalText string
	data         *Words
}

// GetWords returns the collection of mapped string
func (t *Text) GetWords() *Words {
	return t.data
}

// GetOriginalText return original text as is
func (t *Text) GetOriginalText() string {
	return t.originalText
}

// GetNormalizedText return normalized text
func (t *Text) GetNormalizedText() string {
	var buf = bytes.Buffer{}
	defer buf.Reset()

	for _, t := range t.data.Get() {
		buf.WriteString(t.GetNormalized())
	}

	return buf.String()
}

// GetReplacedText return text with some collection of text (word) already been replaced
func (t *Text) GetReplacedText() string {
	var buf = bytes.Buffer{}
	defer buf.Reset()

	for _, t := range t.data.Get() {
		buf.WriteString(t.GetReplaced())
	}

	return buf.String()
}

// NewText return new collection of words
func NewText(originalText string, words *Words) *Text {
	return &Text{
		originalText: originalText,
		data:         words,
	}
}
