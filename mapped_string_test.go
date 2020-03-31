package txtproc

import (
	"testing"
)

func TestMappedString_GetOriginal(t *testing.T) {
	m := MappedString{
		original:   "word!",
		normalized: "word",
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
	}

	if m.GetNormalized() != "word" {
		t.Errorf("want %s got %s", "word!", m.GetNormalized())
		t.Fail()
		return
	}
}
