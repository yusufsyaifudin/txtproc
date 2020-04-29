package txtproc

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

const textProfanityFilterTest = "this is a sentence"

func TestProfanityFilter(t *testing.T) {
	_, err := ProfanityFilter(context.Background(), "", ProfanityFilterConfig{})
	if err != ErrEmptyText {
		t.Errorf("want %v got %v", ErrEmptyText, err)
		t.Fail()
	}
}

// TestProfanityFilter2 test good word processing, repo Get return error
func TestProfanityFilter2(t *testing.T) {
	goodWordData := RepoMocker()

	// Total and PerBatch is return 1
	goodWordData.On("Total", mock.Anything).Return(int64(1)).Once()
	goodWordData.On("PerBatch", mock.Anything).Return(int64(1)).Once()

	// repo Get is error
	wantGoodWordDataErr := fmt.Errorf("error getting data")
	goodWordData.On("Get", mock.Anything, int64(1)).Return([]*ReplacerData{}, wantGoodWordDataErr).Once()

	_, err := ProfanityFilter(context.Background(), textWordReplacerTest, ProfanityFilterConfig{
		GoodWordsReplacerConfig: WordReplacerConfig{
			WordToCompare:        WordNormalized,
			CaseSensitive:        true,
			ReplacerDataSeeder:   goodWordData,
			SimilarityFunc:       nil,
			ReplacerMinimumScore: 0,
		},
		BadWordsReplacerConfig: WordReplacerConfig{},
	})

	if err != wantGoodWordDataErr {
		t.Errorf("want %v got %v", wantGoodWordDataErr, err)
		t.Fail()
		return
	}

}

// TestProfanityFilter3 test bad word processing, repo Get return error
func TestProfanityFilter3(t *testing.T) {
	badWordData := RepoMocker()

	// Total and PerBatch is return 1
	badWordData.On("Total", mock.Anything).Return(int64(1)).Once()
	badWordData.On("PerBatch", mock.Anything).Return(int64(1)).Once()

	// repo Get is error
	wantBadWordDataErr := fmt.Errorf("error getting data")
	badWordData.On("Get", mock.Anything, int64(1)).Return([]*ReplacerData{}, wantBadWordDataErr).Once()

	_, err := ProfanityFilter(context.Background(), textWordReplacerTest, ProfanityFilterConfig{
		GoodWordsReplacerConfig: WordReplacerConfig{},
		BadWordsReplacerConfig: WordReplacerConfig{
			WordToCompare:        WordNormalized,
			CaseSensitive:        true,
			ReplacerDataSeeder:   badWordData,
			SimilarityFunc:       nil,
			ReplacerMinimumScore: 0,
		},
	})

	if err != wantBadWordDataErr {
		t.Errorf("want %v got %v", wantBadWordDataErr, err)
		t.Fail()
		return
	}

}

// TestProfanityFilter4 success, without any mock
func TestProfanityFilter4(t *testing.T) {
	_, err := ProfanityFilter(context.Background(), textProfanityFilterTest, ProfanityFilterConfig{})
	if err != nil {
		t.Errorf("want nil error got '%v'", err)
		t.Fail()
	}
}
