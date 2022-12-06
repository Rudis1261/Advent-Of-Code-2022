package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	// assert := assert.New(t)
}

func Test_findFirstUnique(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(5, findFirstUnique("bvwbjplbgvbhsrlpgdmjqwftvncz", 4))
	assert.Equal(6, findFirstUnique("nppdvjthqldpwncqszvftbrmjlhg", 4))
	assert.Equal(10, findFirstUnique("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4))
	assert.Equal(11, findFirstUnique("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4))
}
