package txtproc

import (
	"bytes"
	"context"
	"ysf/txtproc/word"

	"github.com/opentracing/opentracing-go"
)

// WordSeparator will split text into slice of word. Word itself is a group of character.
// This will split text while holding its structure (spaces, punctuation, etc).
// For example: "abc 123 a b 1" will converted into ["abc", " ", "123", " ", "a", " ", "b", " ", "1"]
// It has time complexity of O(N) where N is the length of text.
func WordSeparator(ctx context.Context, text string) (strCollection *word.Text, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "txtproc.WordSeparator")
	defer func() {
		ctx.Done()
		span.Finish()
	}()

	if text == "" {
		err = ErrEmptyText
		return
	}

	var bufOriginal = bytes.Buffer{}
	var bufNormalized = bytes.Buffer{}
	var textLength = len(text)

	span.SetTag("length", textLength)

	var words = word.NewWords()
	idx := 1
	for i, char := range text {
		_, isCharSeparator := separatorChar[char]
		_, isCharWord := whitelistedCharWord[char]

		if !isCharSeparator {
			bufOriginal.WriteRune(char)
		}

		if isCharWord {
			bufNormalized.WriteRune(char)
		}

		// if current is separator or the end of string, then append the buffer and reset it
		if isCharSeparator {
			// this is for the word
			if bufOriginal.Len() > 0 {
				w := word.NewWord(bufOriginal.String(), bufNormalized.String())
				words.Append(idx, w)
				idx++
			}

			w := word.NewWord(string(char), string(char))
			words.Append(idx, w)

			bufOriginal.Reset()
			bufNormalized.Reset()
			idx++
			continue
		}

		if i+1 == textLength {
			if bufOriginal.Len() > 0 {
				w := word.NewWord(bufOriginal.String(), bufNormalized.String())
				words.Append(idx, w)
				idx++
			}

			bufOriginal.Reset()
			bufNormalized.Reset()
			continue
		}
	}

	// data to return
	return word.NewText(text, words), nil
}
