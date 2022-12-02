package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(gestures["A"], ROCK)
	assert.Equal(gestures["B"], PAPER)
	assert.Equal(gestures["C"], SCISSOR)

	assert.Equal(gestures["X"], ROCK)
	assert.Equal(gestures["Y"], PAPER)
	assert.Equal(gestures["Z"], SCISSOR)

	assert.Equal(game[ROCK].Point, 1)
	assert.Equal(game[PAPER].Point, 2)
	assert.Equal(game[SCISSOR].Point, 3)
}

func Test_score(t *testing.T) {
	assert := assert.New(t)

	// Baseline
	assert.Equal(8, score("A", "Y"))

	// Draw
	assert.Equal(4, score("A", "X"))
	assert.Equal(5, score("B", "Y"))
	assert.Equal(6, score("C", "Z"))

	// Loss
	assert.Equal(1, score("B", "X"))
	assert.Equal(2, score("C", "Y"))
	assert.Equal(3, score("A", "Z"))

	// Win
	assert.Equal(7, score("C", "X"))
	assert.Equal(8, score("A", "Y"))
	assert.Equal(9, score("B", "Z"))
}
