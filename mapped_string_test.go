package txtproc

import (
	"reflect"
	"testing"
)

func TestMappedString_GetOriginal(t *testing.T) {
	m := MappedString{
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
	m := MappedString{
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
	m := MappedString{
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
	m := MappedString{
		original:   "word!",
		normalized: "word",
	}

	if m.GetReplaced() != "word!" {
		t.Errorf("want %s got %s", "word!", m.GetReplaced())
		t.Fail()
		return
	}
}

func TestMappedStrings_GetMappedString(t *testing.T) {
	var mappedStrings = MappedStrings{
		data: []MappedString{
			{
				original:   "word!",
				normalized: "word",
				replaced:   "****",
			},
		},
		originalText: "word!",
	}

	var want = []MappedString{
		{
			original:   "word!",
			normalized: "word",
			replaced:   "****",
		},
	}

	if !reflect.DeepEqual(mappedStrings.GetMappedString(), want) {
		t.Errorf("want %v got %v", want, mappedStrings.GetMappedString())
		t.Fail()
		return
	}
}

func TestMappedStrings_GetOriginalText(t *testing.T) {
	var originalText = "word!"

	var mappedStrings = MappedStrings{
		data: []MappedString{
			{
				original:   "word!",
				normalized: "word",
				replaced:   "****",
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
	var mappedStrings = MappedStrings{
		data: []MappedString{
			{
				original:   "word!",
				normalized: "word",
				replaced:   "****",
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
	var mappedStrings = MappedStrings{
		data: []MappedString{
			{
				original:   "word!",
				normalized: "word",
				replaced:   "****",
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
