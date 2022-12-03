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

func Test_splitItems(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]string{"aaa", "bbb"}, splitItems("aaabbb"))
	assert.Equal([]string{"abc", "def"}, splitItems("abcdef"))
	assert.Equal([]string{"aAx", "adX"}, splitItems("aAxadX"))
}

func Test_findMatchingCharacter(t *testing.T) {
	assert := assert.New(t)
	type test struct {
		a     string
		b     string
		match rune
		err   error
	}

	tests := []test{
		{a: "abcd", b: "axyz", match: 'a', err: nil},
		{a: "ABCD", b: "BXYZ", match: 'B', err: nil},
		{a: "ABCD", b: "VXYZ", match: 0, err: ErrNoMatchFound},
	}

	for _, testCase := range tests {
		result, err := findMatchingCharacter(testCase.a, testCase.b)
		if testCase.err != nil {
			assert.ErrorIs(err, testCase.err)
		} else {
			assert.Equal(testCase.match, result)
		}
	}
}
