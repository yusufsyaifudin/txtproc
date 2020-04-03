package txtproc

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"ysf/txtproc/similarity"

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

	return wordSeparator(ctx, text, ReplacerConfig{})
}

// wordSeparator will split text while holding its structure (spaces, punctuation, etc)
// it will replace the first match of data when the score is higher
// or equal than minimum score stated in `ReplacerConfig`
func wordSeparator(ctx context.Context, text string, replacerCfg ReplacerConfig) (strCollection MappedStrings, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "wordSeparator")
	defer func() {
		ctx.Done()
		span.Finish()
	}()

	strCollection = MappedStrings{
		originalText: text,
		data:         []MappedString{},
	}

	if text == "" {
		err = fmt.Errorf("empty text")
		return
	}

	if replacerCfg.ReplacerData == nil {
		replacerCfg.ReplacerData = newReplacerDataDefault()
	}

	var similarityFunc = replacerCfg.Similarity
	if similarityFunc == nil {
		similarityFunc = similarity.Noop()
	}

	if replacerCfg.MinimumScore < 0 {
		panic("minimum score must greater or equal than 0")
	}

	if replacerCfg.MinimumScore > 1 {
		panic("minimum score must lower or equal than 1")
	}

	var words = make([]MappedString, 0)
	var bufOriginal = bytes.Buffer{}
	defer bufOriginal.Reset()

	var bufNormalized = bytes.Buffer{}
	defer bufNormalized.Reset()

	var textLength = len(text)

	span.SetTag("length", textLength)

	var i = int64(0)
	for i < replacerCfg.ReplacerData.Total(ctx) {
		i++

		// call get data here to minimize db call if the data from db or external storage
		var wordFromDbToCompare, wordFromDbReplacement string
		wordFromDbToCompare, wordFromDbReplacement, err = replacerCfg.ReplacerData.Get(ctx, i)
		if err != nil {
			words = []MappedString{}
			return
		}

		// iterate over text character
		for j, char := range text {
			_, isCharSeparator := separatorChar[char]
			_, isCharWord := whitelistedCharWord[char]

			if !isCharSeparator {
				bufOriginal.WriteRune(char)
			}

			if isCharWord {
				bufNormalized.WriteRune(char)
			}

			// if current is separator or the end of string, then append the buffer and reset it
			// `bufOriginal` and `bufNormalized` will return the
			if isCharSeparator {
				// this is for the word itself
				if bufOriginal.Len() > 0 {
					var currentWord = bufOriginal.String()
					if replacerCfg.WordToCompare == WordNormalized {
						currentWord = bufNormalized.String()
					}

					if !replacerCfg.CaseSensitive {
						currentWord = strings.ToLower(currentWord)
					}

					var score float64 = 0
					score, err = similarityFunc.Compare(ctx, currentWord, wordFromDbToCompare)
					if err != nil {
						err = fmt.Errorf("error compare string '%s': %s", currentWord, err.Error())
						words = []MappedString{}
						return
					}

					m := MappedString{
						original:   bufOriginal.String(),
						normalized: bufNormalized.String(),
					}

					// when score is higher or equal than minimum score in config, replace it
					if score >= replacerCfg.MinimumScore {
						m.replaced = wordFromDbReplacement
						m.isReplaced = true
					}

					words = append(words, m)
				}

				// this is for the separator chars, like space, newline or tab
				words = append(words, MappedString{
					original:   string(char),
					normalized: string(char),
				})

				// always reset after append to MappedString
				bufOriginal.Reset()
				bufNormalized.Reset()
				continue
			}

			if j+1 == textLength {
				if bufOriginal.Len() > 0 {
					var currentWord = bufOriginal.String()
					if replacerCfg.WordToCompare == WordNormalized {
						currentWord = bufNormalized.String()
					}

					if !replacerCfg.CaseSensitive {
						currentWord = strings.ToLower(currentWord)
					}

					var score float64 = 0
					score, err = similarityFunc.Compare(ctx, currentWord, wordFromDbToCompare)
					if err != nil {
						err = fmt.Errorf("error compare string '%s': %s", currentWord, err.Error())
						words = []MappedString{}
						return
					}

					m := MappedString{
						original:   bufOriginal.String(),
						normalized: bufNormalized.String(),
					}

					// when score is higher or equal than minimum score in config, replace it
					if score >= replacerCfg.MinimumScore {
						m.replaced = wordFromDbReplacement
						m.isReplaced = true
					}

					words = append(words, m)
				}

				// always reset after append to MappedString
				bufOriginal.Reset()
				bufNormalized.Reset()
				continue
			}
		}
	}

	strCollection.data = words
	return
}
