package txtproc

// MappedString is a struct represent the word
type MappedString struct {
	original   string
	normalized string
}

// GetOriginal get original string
func (m *MappedString) GetOriginal() string {
	return m.original
}

// GetNormalized get normalized string, it will contain alphanumeric character only
func (m *MappedString) GetNormalized() string {
	return m.normalized
}
