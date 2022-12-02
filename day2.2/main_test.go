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

	assert.Equal(gestures["X"], LOSS)
	assert.Equal(gestures["Y"], DRAW)
	assert.Equal(gestures["Z"], WIN)

	assert.Equal(game[ROCK].Point, 1)
	assert.Equal(game[PAPER].Point, 2)
	assert.Equal(game[SCISSOR].Point, 3)

	assert.Equal(game[ROCK].Loss, PAPER)
	assert.Equal(game[PAPER].Loss, SCISSOR)
	assert.Equal(game[SCISSOR].Loss, ROCK)
}

func Test_score(t *testing.T) {
	assert := assert.New(t)

	// Baseline
	assert.Equal(4, score("A", "Y"))
	assert.Equal(1, score("B", "X"))
	assert.Equal(7, score("C", "Z"))

	// Draw
	assert.Equal(4, score("A", "Y"))
	assert.Equal(5, score("B", "Y"))
	assert.Equal(6, score("C", "Y"))

	// Loss
	assert.Equal(1, score("B", "X"))
	assert.Equal(2, score("C", "X"))
	assert.Equal(3, score("A", "X"))

	// Win
	assert.Equal(7, score("C", "Z"))
	assert.Equal(8, score("A", "Z"))
	assert.Equal(9, score("B", "Z"))
}
