package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_encapsulated(t *testing.T) {
	assert := assert.New(t)
	// Base cases
	// For example, 2-8 fully contains 3-7,
	// and 6-6 is fully contained by 4-6.
	assert.Equal(true, encapsulated("2-8", "3-7"))
	assert.Equal(true, encapsulated("6-6", "4-6"))

	assert.Equal(true, encapsulated("2-4", "1-10"))
	assert.Equal(true, encapsulated("3-3", "1-10"))
	assert.Equal(true, encapsulated("3-3", "3-10"))
	assert.Equal(true, encapsulated("5-6", "4-6"))

	assert.Equal(false, encapsulated("2-4", "3-10"))
	assert.Equal(false, encapsulated("1-3", "2-10"))
	assert.Equal(false, encapsulated("11-14", "2-10"))
}

func Test_overlapping(t *testing.T) {
	assert := assert.New(t)
	// Base cases
	// 5-7,7-9 overlaps in a single section, 7.
	// 2-8,3-7 overlaps all of the sections 3 through 7.
	// 6-6,4-6 overlaps in a single section, 6.
	// 2-6,4-8 overlaps in sections 4, 5, and 6.
	assert.Equal(true, overlapping("5-7", "7-9"))
	assert.Equal(true, overlapping("2-8", "3-7"))
	assert.Equal(true, overlapping("6-6", "4-6"))
	assert.Equal(true, overlapping("2-6", "4-8"))

	assert.Equal(true, overlapping("2-4", "1-10"))
	assert.Equal(true, overlapping("3-3", "1-10"))
	assert.Equal(true, overlapping("3-3", "3-10"))
	assert.Equal(true, overlapping("5-6", "4-6"))
	assert.Equal(true, overlapping("2-4", "3-10"))
	assert.Equal(true, overlapping("1-3", "2-10"))

	assert.Equal(false, overlapping("1-2", "3-4"))
	assert.Equal(false, overlapping("11-14", "2-10"))
}

func Test_toRange(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]int{2, 8}, toRange("2-8"))
	assert.Equal([]int{6, 6}, toRange("6-6"))
	assert.Equal([]int{2, 4}, toRange("2-4"))
	assert.Equal([]int{3, 3}, toRange("3-3"))
	assert.Equal([]int{5, 6}, toRange("5-6"))
	assert.Equal([]int{1, 3}, toRange("1-3"))
	assert.Equal([]int{11, 14}, toRange("11-14"))
}

func Test_inRange(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(true, inRange(1, []int{0, 10}))
	assert.Equal(true, inRange(2, []int{0, 10}))
	assert.Equal(true, inRange(3, []int{3, 10}))
	assert.Equal(true, inRange(10, []int{9, 10}))
	assert.Equal(true, inRange(8, []int{6, 10}))

	assert.Equal(false, inRange(3, []int{1, 2}))
	assert.Equal(false, inRange(3, []int{6, 10}))
	assert.Equal(false, inRange(5, []int{6, 10}))
}
