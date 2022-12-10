package main

import (
	"testing"

	"github.com/Rudis1261/Advent-Of-Code-2022/day9/board"
	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(0, head.Last.X)
	assert.Equal(0, head.Last.Y)

	head.AddMove(board.Position{X: 1, Y: 2})
	assert.Equal(1, head.Last.X)
	assert.Equal(2, head.Last.Y)
}
