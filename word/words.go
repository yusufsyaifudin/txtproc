package word

import (
	"fmt"
	"sort"
	"sync"
)

type Words struct {
	lock  sync.RWMutex
	words map[int]*Word
}

func (w *Words) Append(idx int, word *Word) {
	w.lock.Lock()
	defer w.lock.Unlock()

	if v, exist := w.words[idx]; exist {
		panic(fmt.Errorf("index %d already occupied by word %v", idx, v))
	}

	w.words[idx] = word
}

func (w *Words) Get() []*Word {
	w.lock.RLock()
	defer w.lock.RUnlock()

	keys := make([]int, 0)
	for k, _ := range w.words {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	words := make([]*Word, 0)
	for _, k := range keys {
		words = append(words, w.words[k])
	}

	return words
}

// MergeAndReplaceWordBetweenIdx this can be called as word compaction.
// idxStart and idxEnd must be between 0 - length of current word slices
// When ...
func (w *Words) MergeAndReplaceWordBetweenIdx(idxStart, idxEnd int, wordReplacement *Word) {
	w.lock.Lock()
	defer w.lock.Unlock()

	if _, exist := w.words[idxStart]; !exist {
		panic(fmt.Errorf("index start %d is not occupied by any word", idxStart))
	}

	if _, exist := w.words[idxEnd]; !exist {
		panic(fmt.Errorf("index end %d is not occupied by any word", idxEnd))
	}

	var newWords = make(map[int]*Word)
	var i = 1
	for idx, word := range w.words {
		// if current index is idxStart (index that want to replace), then append new value
		// then continue to next iteration
		if idx == idxStart {
			newWords[i] = wordReplacement
			i++
			continue
		}

		// when the index within the replacement range, just skip it.
		if idx >= idxStart && idx <= idxEnd {
			continue
		}

		newWords[i] = word
		i++
	}
}

func NewWords() *Words {
	return &Words{
		words: make(map[int]*Word),
	}
}
