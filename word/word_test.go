package word

import (
	"reflect"
	"testing"
)

func TestMappedString_GetOriginal(t *testing.T) {
	m := &Word{
		original:   "word!",
		normalized: "word",
		replaced:   "****",
	}

	if m.GetOriginal() != "word!" {
		t.Errorf("want %s got %s", "word!", m.GetOriginal())
		t.Fail()
		return
	}
}

func TestMappedString_GetNormalized(t *testing.T) {
	m := &Word{
		original:   "word!",
		normalized: "word",
		replaced:   "****",
	}

	if m.GetNormalized() != "word" {
		t.Errorf("want %s got %s", "word!", m.GetNormalized())
		t.Fail()
		return
	}
}

func TestMappedString_GetReplaced(t *testing.T) {
	m := &Word{
		original:   "word!",
		normalized: "word",
		replaced:   "****",
	}

	if m.GetReplaced() != "****" {
		t.Errorf("want %s got %s", "****", m.GetReplaced())
		t.Fail()
		return
	}
}

// TestMappedString_GetReplaced2 when replaced is empty
func TestMappedString_GetReplaced2(t *testing.T) {
	m := &Word{
		original:   "word!",
		normalized: "word",
	}

	if m.GetReplaced() != "word!" {
		t.Errorf("want %s got %s", "word!", m.GetReplaced())
		t.Fail()
		return
	}
}

func TestMappedString_IsReplaced(t *testing.T) {
	m := &Word{
		original:   "word!",
		normalized: "word",
		replaced:   "",
		isReplaced: true,
	}

	if !m.IsReplaced() {
		t.Errorf("want %v got %v", true, m.IsReplaced())
		t.Fail()
		return
	}
}

func TestMappedStrings_GetMappedString(t *testing.T) {
	var mappedStrings = Text{
		data: &Words{
			words: map[int]*Word{
				1: {
					original:   "word!",
					normalized: "word",
					replaced:   "****",
				},
			},
		},
		originalText: "word!",
	}

	var want = &Word{
		original:   "word!",
		normalized: "word",
		replaced:   "****",
	}

	if !reflect.DeepEqual(mappedStrings.GetWords().words[1], want) {
		t.Errorf("want %v got %v", want, mappedStrings.GetWords())
		t.Fail()
		return
	}
}

func TestMappedStrings_GetOriginalText(t *testing.T) {
	var originalText = "word!"

	var mappedStrings = Text{
		data: &Words{
			words: map[int]*Word{
				1: {
					original:   "word!",
					normalized: "word",
					replaced:   "****",
				},
			},
		},
		originalText: "word!",
	}

	if mappedStrings.GetOriginalText() != originalText {
		t.Errorf("want %v got %v", originalText, mappedStrings.GetOriginalText())
		t.Fail()
		return
	}
}

func TestMappedStrings_GetNormalizedText(t *testing.T) {
	var mappedStrings = Text{
		data: &Words{
			words: map[int]*Word{
				1: {
					original:   "word!",
					normalized: "word",
					replaced:   "****",
				},
			},
		},
		originalText: "word!",
	}

	var want = "word"
	if mappedStrings.GetNormalizedText() != want {
		t.Errorf("want %v got %v", want, mappedStrings.GetNormalizedText())
		t.Fail()
		return
	}
}

func TestMappedStrings_GetReplacedText(t *testing.T) {
	var mappedStrings = Text{
		data: &Words{
			words: map[int]*Word{
				1: {
					original:   "word!",
					normalized: "word",
					replaced:   "****",
				},
			},
		},
		originalText: "word!",
	}

	var want = "****"
	if mappedStrings.GetReplacedText() != want {
		t.Errorf("want %v got %v", want, mappedStrings.GetReplacedText())
		t.Fail()
		return
	}
}
