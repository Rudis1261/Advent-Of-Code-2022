package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getCount(t *testing.T) {
	assert := assert.New(t)
	sampleData := []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}
	count, scenicScore := getCount(sampleData)
	assert.Equal(21, count)
	assert.Equal(8, scenicScore)
}

func Test_tallestAmongsOthers(t *testing.T) {
	assert := assert.New(t)
	sampleData := []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}

	assert.Equal(false, visible(sampleData, 3, 1))
	assert.Equal(false, visible(sampleData, 1, 3))
	assert.Equal(false, visible(sampleData, 2, 2))
	assert.Equal(false, visible(sampleData, 3, 3))

	assert.Equal(true, visible(sampleData, 1, 1))
	assert.Equal(true, visible(sampleData, 1, 2))
	assert.Equal(true, visible(sampleData, 2, 1))
	assert.Equal(true, visible(sampleData, 2, 3))
	assert.Equal(true, visible(sampleData, 3, 2))
}

func Test_scenicScore(t *testing.T) {
	assert := assert.New(t)
	sampleData := []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}
	type test struct {
		X     int
		Y     int
		Score int
	}
	tests := []test{
		{X: 2, Y: 1, Score: 4},
		{X: 2, Y: 3, Score: 8},
	}

	for _, testCase := range tests {
		assert.Equal(testCase.Score, scenicScore(sampleData, testCase.X, testCase.Y))
	}
}
