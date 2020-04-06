package txtproc

import (
	"bytes"
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
)

// WordSeparator will split text into slice of word. Word itself is a group of character.
// For example: "abc 123 a b 1" will converted into ["abc", " ", "123", " ", "a", " ", "b", " ", "1"]
func WordSeparator(ctx context.Context, text string) (strCollection MappedStrings, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "txtproc.WordSeparator")
	defer func() {
		ctx.Done()
		span.Finish()
	}()

	return wordSeparator(ctx, text)
}

// wordSeparator will split text while holding its structure (spaces, punctuation, etc).
func wordSeparator(ctx context.Context, text string) (strCollection MappedStrings, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "wordSeparator")
	defer func() {
		ctx.Done()
		span.Finish()
	}()

	strCollection = MappedStrings{
		data:         []*MappedString{},
		originalText: text,
	}

	if text == "" {
		err = fmt.Errorf("empty text")
		return
	}

	var words = make([]*MappedString, 0)

	var bufOriginal = bytes.Buffer{}
	var bufNormalized = bytes.Buffer{}
	var textLength = len(text)

	span.SetTag("length", textLength)

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
				words = append(words, &MappedString{
					original:   bufOriginal.String(),
					normalized: bufNormalized.String(),
				})
			}

			words = append(words, &MappedString{
				original:   string(char),
				normalized: string(char),
			})

			bufOriginal.Reset()
			bufNormalized.Reset()
			continue
		}

		if i+1 == textLength {
			if bufOriginal.Len() > 0 {
				words = append(words, &MappedString{
					original:   bufOriginal.String(),
					normalized: bufNormalized.String(),
				})
			}

			bufOriginal.Reset()
			bufNormalized.Reset()
			continue
		}
	}

	// data to return
	strCollection.data = words
	return
}
