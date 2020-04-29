package txtproc

import (
	"context"
	"reflect"
	"testing"
	"ysf/txtproc/word"
)

func BenchmarkWordSeparator_1Word(b *testing.B) {
	text := `word`

	for n := 0; n < b.N; n++ {
		_, _ = WordSeparator(context.Background(), text)
	}
}

func BenchmarkWordSeparator_100Words(b *testing.B) {
	text := `accumsan tortor posuere ac ut consequat semper viverra nam libero justo laoreet sit amet cursus sit amet dictum sit amet justo donec enim diam vulputate ut pharetra sit amet aliquam id diam maecenas ultricies mi eget mauris pharetra et ultrices neque ornare aenean euismod elementum nisi quis eleifend quam adipiscing vitae proin sagittis nisl rhoncus mattis rhoncus urna neque viverra justo nec ultrices dui sapien eget mi proin sed libero enim sed faucibus turpis in eu mi bibendum neque egestas congue quisque egestas diam in arcu cursus euismod quis viverra nibh cras pulvinar mattis nunc sed blandit libero volutpat sed`

	for n := 0; n < b.N; n++ {
		_, _ = WordSeparator(context.Background(), text)
	}
}

func BenchmarkWordSeparator_200Words(b *testing.B) {
	text := `fringilla ut morbi tincidunt augue interdum velit euismod in pellentesque massa placerat duis ultricies lacus sed turpis tincidunt id aliquet risus feugiat in ante metus dictum at tempor commodo ullamcorper a lacus vestibulum sed arcu non odio euismod lacinia at quis risus sed vulputate odio ut enim blandit volutpat maecenas volutpat blandit aliquam etiam erat velit scelerisque in dictum non consectetur a erat nam at lectus urna duis convallis convallis tellus id interdum velit laoreet id donec ultrices tincidunt arcu non sodales neque sodales ut etiam sit amet nisl purus in mollis nunc sed id semper risus in hendrerit gravida rutrum quisque non tellus orci ac auctor augue mauris augue neque gravida in fermentum et sollicitudin ac orci phasellus egestas tellus rutrum tellus pellentesque eu tincidunt tortor aliquam nulla facilisi cras fermentum odio eu feugiat pretium nibh ipsum consequat nisl vel pretium lectus quam id leo in vitae turpis massa sed elementum tempus egestas sed sed risus pretium quam vulputate dignissim suspendisse in est ante in nibh mauris cursus mattis molestie a iaculis at erat pellentesque adipiscing commodo elit at imperdiet dui accumsan sit amet nulla facilisi morbi tempus iaculis urna id volutpat lacus laoreet non curabitur gravida arcu ac`

	for n := 0; n < b.N; n++ {
		_, _ = WordSeparator(context.Background(), text)
	}
}

func TestWordSeparators(t *testing.T) {
	text := ` `
	words, _ := WordSeparator(context.Background(), text)

	if words.GetOriginalText() != text {
		t.Errorf("returned '%s' want '%s'", words.GetOriginalText(), text)
		t.Fail()
		return
	}
}

// TestWordSeparators1 check that new line still included
func TestWordSeparators1(t *testing.T) {
	text := `nama! kata1 1234??
`
	words, _ := WordSeparator(context.Background(), text)

	if words.GetOriginalText() != text {
		t.Errorf("returned '%s' want '%s'", words.GetOriginalText(), text)
		t.Fail()
		return
	}
}

// TestWordSeparators2 check that complex space and multi newline still included
func TestWordSeparators2(t *testing.T) {
	text := `nama!     kata1 1234??


`
	words, _ := WordSeparator(context.Background(), text)

	if words.GetOriginalText() != text {
		t.Errorf("returned '%s' want '%s'", words.GetOriginalText(), text)
		t.Fail()
		return
	}
}

func TestWordSeparators3(t *testing.T) {
	type testData struct {
		text string
		want []*word.Word
	}

	table := []testData{
		// space only
		{
			text: " ",
			want: []*word.Word{
				word.NewWord(" ", " "),
			},
		},

		{
			text: "nama!",
			want: []*word.Word{
				word.NewWord("nama!", "nama"),
			},
		},

		// check that space in between is okay
		{
			text: "nama! 123",
			want: []*word.Word{
				word.NewWord("nama!", "nama"),
				word.NewWord(" ", " "),
				word.NewWord("123", "123"),
			},
		},

		// space in prefix
		{
			text: " nama!",
			want: []*word.Word{
				word.NewWord(" ", " "),
				word.NewWord("nama!", "nama"),
			},
		},

		// space in suffix
		{
			text: "nama! ",
			want: []*word.Word{
				word.NewWord("nama!", "nama"),
				word.NewWord(" ", " "),
			},
		},
	}

	for i, data := range table {
		words, _ := WordSeparator(context.Background(), data.text)
		if !reflect.DeepEqual(data.want, words.GetWords().Get()) {
			t.Errorf("%d want %v but return %v", i, data.want, words.GetWords().Get())
			t.Fail()
			return
		}
	}
}

// TestWordSeparator4 one prefix, one suffix and two spaces in between
func TestWordSeparator4(t *testing.T) {
	text := " c  c "
	want := []*word.Word{
		word.NewWord(" ", " "),
		word.NewWord("c", "c"),
		word.NewWord(" ", " "),
		word.NewWord(" ", " "),
		word.NewWord("c", "c"),
		word.NewWord(" ", " "),
	}

	words, _ := WordSeparator(context.Background(), text)
	if !reflect.DeepEqual(want, words.GetWords().Get()) {
		t.Errorf("want %v but return %v", want, words)
		t.Fail()
		return
	}
}

// TestWordSeparator5 multi line
func TestWordSeparator5(t *testing.T) {
	text := `a
`
	want := []*word.Word{
		word.NewWord("a", "a"),
		word.NewWord("\n", "\n"),
	}

	words, _ := WordSeparator(context.Background(), text)
	if !reflect.DeepEqual(want, words.GetWords().Get()) {
		t.Errorf("want %v but return %v", want, words)
		t.Fail()
		return
	}
}

// TestWordSeparator6 multi line
func TestWordSeparator6(t *testing.T) {
	text := `
a
`
	want := []*word.Word{
		word.NewWord("\n", "\n"),
		word.NewWord("a", "a"),
		word.NewWord("\n", "\n"),
	}

	words, _ := WordSeparator(context.Background(), text)
	if !reflect.DeepEqual(want, words.GetWords().Get()) {
		t.Errorf("want %v but return %v", want, words)
		t.Fail()
		return
	}
}

// TestWordSeparator7 Carriage Return line
func TestWordSeparator7(t *testing.T) {
	text := "\ra\n"
	want := []*word.Word{
		word.NewWord("\r", "\r"),
		word.NewWord("a", "a"),
		word.NewWord("\n", "\n"),
	}

	words, _ := WordSeparator(context.Background(), text)
	if !reflect.DeepEqual(want, words.GetWords().Get()) {
		t.Errorf("want %v but return %v", want, words)
		t.Fail()
		return
	}
}

func TestWordSeparator8(t *testing.T) {
	_, err := WordSeparator(context.Background(), "")
	if err == nil {
		t.Errorf("should return error: '%s'", "empty text")
		t.Fail()
		return
	}
}

// TestWordSeparator9 check tabs
func TestWordSeparator9(t *testing.T) {
	text := `a	b`
	want := []*word.Word{
		word.NewWord("a", "a"),
		word.NewWord("	", "	"),
		word.NewWord("b", "b"),
	}

	words, _ := WordSeparator(context.Background(), text)
	if !reflect.DeepEqual(want, words.GetWords().Get()) {
		t.Errorf("want %v but return %v", want, words)
		t.Fail()
		return
	}
}
