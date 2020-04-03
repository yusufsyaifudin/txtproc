package txtproc

import (
	"bytes"
)

// MappedString is a struct represent the word
type MappedString struct {
	original   string
	normalized string
	replaced   string
	isReplaced bool
}

// GetOriginal get original string
func (m *MappedString) GetOriginal() string {
	return m.original
}

// GetNormalized get normalized string, it will contain alphanumeric character only
func (m *MappedString) GetNormalized() string {
	return m.normalized
}

// GetReplaced get string that already replaced, will same as GetOriginal when needed
func (m *MappedString) GetReplaced() string {
	if m.replaced == "" {
		return m.GetOriginal()
	}

	return m.replaced
}

// IsReplaced whether the current string already replaced or not
func (m *MappedString) IsReplaced() bool {
	return m.isReplaced
}

// MappedStrings is collection of mapped string
type MappedStrings struct {
	originalText string
	data         []MappedString
}

// GetMappedString returns the collection of mapped string
func (ms *MappedStrings) GetMappedString() []MappedString {
	return ms.data
}

// GetOriginalText return original text as is
func (ms *MappedStrings) GetOriginalText() string {
	return ms.originalText
}

// GetNormalizedText return normalized text
func (ms *MappedStrings) GetNormalizedText() string {
	var buf = bytes.Buffer{}
	defer buf.Reset()

	for _, t := range ms.data {
		buf.WriteString(t.GetNormalized())
	}

	return buf.String()
}

// GetReplacedText return text with some collection of text (word) already been replaced
func (ms *MappedStrings) GetReplacedText() string {
	var buf = bytes.Buffer{}
	defer buf.Reset()

	for _, t := range ms.data {
		buf.WriteString(t.GetReplaced())
	}

	return buf.String()
}
