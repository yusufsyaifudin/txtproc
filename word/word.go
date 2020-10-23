package word

import (
	"sync"
)

// Word is a struct represent the word
type Word struct {
	lock       sync.RWMutex
	original   string
	normalized string
	replaced   string
	isReplaced bool
}

// GetOriginal get original string
func (m *Word) GetOriginal() string {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.original
}

// GetNormalized get normalized string, it will contain alphanumeric character only
func (m *Word) GetNormalized() string {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.normalized
}

// GetReplaced get string that already replaced, will same as GetOriginal when needed
func (m *Word) GetReplaced() string {
	m.lock.RLock()
	defer m.lock.RUnlock()

	if m.replaced == "" {
		return m.original
	}

	return m.replaced
}

// SetReplaced replaces the string in GetReplaced
func (m *Word) SetReplaced(str string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.replaced = str
	m.isReplaced = true
}

// IsReplaced whether the current string already replaced or not
func (m *Word) IsReplaced() bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.isReplaced
}

// NewWord return new word. So, original and normalized cannot be changed at runtime.
func NewWord(original, normalized string) *Word {
	return &Word{
		original:   original,
		normalized: normalized,
		replaced:   original,
		isReplaced: false,
	}
}
