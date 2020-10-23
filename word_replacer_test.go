package txtproc

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

const textWordReplacerTest = "this is a sentence"

// TestProfanityFilter1 test on default config
func TestWordReplacer(t *testing.T) {
	var mappedStrings, _ = WordSeparator(context.Background(), textWordReplacerTest)

	err := WordReplacer(context.Background(), mappedStrings, WordReplacerConfig{})
	if err != nil {
		t.Error("WordReplacer should return no error on default config")
		t.Fail()
	}
}

// TestWordReplacer2 test good word processing, repo Get return error
func TestWordReplacer2(t *testing.T) {
	var mappedStrings, _ = WordSeparator(context.Background(), textWordReplacerTest)

	goodWordData := RepoMocker()

	// Total and PerBatch is return 1
	goodWordData.On("Total", mock.Anything).Return(int64(1)).Once()
	goodWordData.On("PerBatch", mock.Anything).Return(int64(1)).Once()

	// repo Get is error
	wantGoodWordDataErr := fmt.Errorf("error getting data")
	goodWordData.On("Get", mock.Anything, int64(1)).Return([]ReplacerData{}, wantGoodWordDataErr).Once()

	err := WordReplacer(context.Background(), mappedStrings, WordReplacerConfig{
		WordToCompare:        WordNormalized,
		CaseSensitive:        true,
		ReplacerDataSeeder:   goodWordData,
		SimilarityFunc:       nil,
		ReplacerMinimumScore: 1,
	})

	if err != wantGoodWordDataErr {
		t.Errorf("want %v got %v", wantGoodWordDataErr, err)
		t.Fail()
		return
	}

}

// TestWordReplacer3 test good word processing, similarity function Compare return error
func TestWordReplacer3(t *testing.T) {
	var mappedStrings, _ = WordSeparator(context.Background(), textWordReplacerTest)

	goodWordData := RepoMocker()

	// Total and PerBatch is return 1
	goodWordData.On("Total", mock.Anything).Return(int64(1)).Once()
	goodWordData.On("PerBatch", mock.Anything).Return(int64(1)).Once()

	// repo Get will only return one
	wantGoodWordData := []ReplacerData{
		{
			StringToCompare:   "word",
			StringReplacement: "replacement",
		},
	}
	goodWordData.On("Get", mock.Anything, int64(1)).Return(wantGoodWordData, nil).Once()

	goodWordSim := SimilarityMocker()

	wantSimCompareErr := fmt.Errorf("compare method error")
	goodWordSim.On("Compare", mock.Anything, mock.Anything, mock.Anything).
		Return(0.0, wantSimCompareErr).Once()

	err := WordReplacer(context.Background(), mappedStrings, WordReplacerConfig{
		WordToCompare:        WordNormalized,
		CaseSensitive:        true,
		ReplacerDataSeeder:   goodWordData,
		SimilarityFunc:       goodWordSim,
		ReplacerMinimumScore: 0,
	})

	if err == nil {
		t.Errorf("should return error, but got nil")
		t.Fail()
		return
	}
}

// TestWordReplacer4 test good word processing, return success, even with nil values
func TestWordReplacer4(t *testing.T) {
	var mappedStrings, _ = WordSeparator(context.Background(), textWordReplacerTest)

	goodWordData := RepoMocker()

	// Total = 2 and PerBatch = 1, so it will do 2 iteration, so it will call w.IsReplaced() statement
	goodWordData.On("Total", mock.Anything).Return(int64(2)).Once()
	goodWordData.On("PerBatch", mock.Anything).Return(int64(1)).Once()

	// repo Get will only return one
	wantGoodWordData := []ReplacerData{
		{
			StringToCompare:   "this",
			StringReplacement: "****",
		},
	}
	goodWordData.On("Get", mock.Anything, mock.Anything).Return(wantGoodWordData, nil).Twice()

	goodWordSim := SimilarityMocker()
	goodWordSim.On("Compare", mock.Anything, mock.Anything, mock.Anything).
		Return(1.0, nil)

	err := WordReplacer(context.Background(), mappedStrings, WordReplacerConfig{
		WordToCompare:        WordNormalized,
		CaseSensitive:        false,
		ReplacerDataSeeder:   goodWordData,
		SimilarityFunc:       goodWordSim,
		ReplacerMinimumScore: 0,
	})

	if err != nil {
		t.Errorf("should return not error, but got '%v'", err)
		t.Fail()
		return
	}
}

// TestWordReplacer5 when mapped string is nil
func TestWordReplacer5(t *testing.T) {
	err := WordReplacer(context.Background(), nil, WordReplacerConfig{})
	if err != ErrMappedStringsNil {
		t.Errorf("WordReplacer want %v got %v", ErrMappedStringsNil, err)
		t.Fail()
	}
}

// TestWordReplacer6 test config replacer minimum score
func TestWordReplacer6(t *testing.T) {
	var mappedStrings, _ = WordSeparator(context.Background(), textWordReplacerTest)

	err := WordReplacer(context.Background(), mappedStrings, WordReplacerConfig{
		ReplacerMinimumScore: -0.1,
	})
	if err != ErrMinReplacerScore {
		t.Errorf("WordReplacer want '%v' got '%v'", ErrMinReplacerScore, err)
		t.Fail()
		return
	}

	err = WordReplacer(context.Background(), mappedStrings, WordReplacerConfig{
		ReplacerMinimumScore: 1.1,
	})
	if err != ErrMinReplacerScore {
		t.Errorf("WordReplacer want '%v' got '%v'", ErrMinReplacerScore, err)
		t.Fail()
	}
}
