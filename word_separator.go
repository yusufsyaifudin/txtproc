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

// wordSeparator will split text while holding its structure (spaces, punctuation, etc).
// It will replace the first match of data based on computed similarity score.
// When comparing string from your database and score return higher or equal
// than minimum score stated in `ReplacerConfig`, it will be replaced.
// It's up to you to order the string when `ReplacerDataSeeder.Get` is called.
func wordSeparator(ctx context.Context, text string, replacerCfg ReplacerConfig) (strCollection MappedStrings, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "wordSeparator")
	defer func() {
		ctx.Done()
		span.Finish()
	}()

	strCollection = MappedStrings{
		data:         []*MappedString{},
		originalText: text,
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

	// early return when not needed
	if !replacerCfg.enabled {
		return
	}

	var i = int64(0)
	for i < replacerCfg.ReplacerData.Total(ctx) {
		i++

		// call get data here to minimize db call if the data from db or external storage
		var dataReplacer = &ReplacerData{}
		dataReplacer, err = replacerCfg.ReplacerData.Get(ctx, i)
		if err != nil {
			return
		}

		for _, w := range words {
			var currentWord = w.GetOriginal()
			if replacerCfg.WordToCompare == WordNormalized {
				currentWord = w.GetNormalized()
			}

			if !replacerCfg.CaseSensitive {
				currentWord = strings.ToLower(currentWord)
			}

			var score float64 = 0
			score, err = similarityFunc.Compare(ctx, currentWord, dataReplacer.StringToCompare)
			if err != nil {
				err = fmt.Errorf(
					"error compare string '%s' vs '%s': %s",
					currentWord, dataReplacer.StringToCompare, err.Error(),
				)
				return
			}

			// when score is higher or equal than minimum score in config, replace it
			if score >= replacerCfg.MinimumScore {
				w.replaced = dataReplacer.StringReplacement
				w.isReplaced = true
			}
		}

	}

	return
}
