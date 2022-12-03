package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	assert := assert.New(t)
	// Some checking of correct charactorization of rune type
	assert.Equal(true, lowercase.Contains('a'))
	assert.Equal(true, lowercase.Contains('p'))
	assert.Equal(true, lowercase.Contains('z'))
	assert.Equal(false, lowercase.Contains('A'))
	assert.Equal(false, lowercase.Contains('P'))
	assert.Equal(false, lowercase.Contains('Z'))

	assert.Equal(false, lowercase.Contains('A'))
	assert.Equal(false, lowercase.Contains('Z'))
	assert.Equal(true, uppercase.Contains('A'))
	assert.Equal(true, uppercase.Contains('P'))
	assert.Equal(true, uppercase.Contains('Z'))

	assert.Equal(1, lowercase.NumericValue('a'))
	assert.Equal(16, lowercase.NumericValue('p'))
	assert.Equal(26, lowercase.NumericValue('z'))

	assert.Equal(27, uppercase.NumericValue('A'))
	assert.Equal(52, uppercase.NumericValue('Z'))
}

func Test_chunk(t *testing.T) {
	assert := assert.New(t)

	type test struct {
		result    [][]string
		input     []string
		chunkSize int
		err       error
	}
	tests := []test{
		{
			input: []string{"a", "b", "c", "d", "e", "f"},
			result: [][]string{
				{"a", "b", "c", "d"},
				{"e", "f"},
			},
			chunkSize: 4,
			err:       ErrNotDevisible,
		},
		{
			input: []string{"a", "b", "c", "d", "e", "f"},
			result: [][]string{
				{"a", "b"},
				{"c", "d"},
				{"e", "f"},
			},
			chunkSize: 2,
		},
		{
			input: []string{"a", "b", "c", "d", "e", "f"},
			result: [][]string{
				{"a"},
				{"b"},
				{"c"},
				{"d"},
				{"e"},
				{"f"},
			},
			chunkSize: 1,
		},
	}
	for _, testCase := range tests {
		result, err := chunk(testCase.input, testCase.chunkSize)
		if testCase.err != nil {
			assert.ErrorIs(err, testCase.err)
			continue
		}
		assert.Equal(testCase.result, result)
	}
}

func Test_findBadge(t *testing.T) {
	assert := assert.New(t)

	type test struct {
		result rune
		input  []string
		err    error
	}
	tests := []test{
		{
			input: []string{
				"abcd",
				"aZLKJ",
				"kljhkjahz",
			},
			result: 'a',
		},
		{
			input: []string{
				"fghi",
				"aZhLKJ",
				"kljhkjaz",
			},
			result: 'h',
		},
		{
			input: []string{
				"fghi",
				"aZLKJ",
				"kljhkjaz",
			},
			result: 0,
			err:    ErrNoMatchFound,
		},
	}
	for _, testCase := range tests {
		result, err := findBadge(testCase.input)
		if testCase.err != nil {
			assert.ErrorIs(err, testCase.err)
			continue
		}
		assert.Equal(testCase.result, result)
	}
}
