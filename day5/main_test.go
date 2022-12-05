package main

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	// assert := assert.New(t)
}

func Test_BuildIndex(t *testing.T) {
	assert := assert.New(t)
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")
	assert.Equal(map[int]int{
		1: 1,
		2: 5,
		3: 9,
		4: 13,
		5: 17,
		6: 21,
		7: 25,
		8: 29,
		9: 33,
	}, BuildIndex(lines[8]).Columns)

	assert.Equal(map[int]int{
		1:  1,
		5:  2,
		9:  3,
		13: 4,
		17: 5,
		21: 6,
		25: 7,
		29: 8,
		33: 9,
	}, BuildIndex(lines[8]).InverseIndex)
}

func TestNewMatrix(t *testing.T) {
	assert := assert.New(t)
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")
	iDex := BuildIndex(lines[8])

	// "        [Q] [B]         [H]        "
	// " 1   2   3   4   5   6   7   8   9 "
	assert.Equal(map[int][]string{
		0: []string{},
		1: []string{},
		2: []string{},
		3: []string{"Q"},
		4: []string{"B"},
		5: []string{},
		6: []string{},
		7: []string{"H"},
		8: []string{},
		9: []string{},
	}, NewMatrix(lines[0:1], iDex).Values)

	// "    [R] [D] [L] [C] [N] [Q]     [R]"
	// " 1   2   3   4   5   6   7   8   9 "
	assert.Equal(map[int][]string{
		0: []string{},
		1: []string{},
		2: []string{"R"},
		3: []string{"D"},
		4: []string{"L"},
		5: []string{"C"},
		6: []string{"N"},
		7: []string{"Q"},
		8: []string{},
		9: []string{"R"},
	}, NewMatrix(lines[3:4], iDex).Values)
}

func Test_parseMove(t *testing.T) {
	assert := assert.New(t)
	type Test struct {
		Input     string
		Qty       int
		FromIndex int
		ToIndex   int
	}

	tests := []Test{
		{Input: "move 12 from 6 to 2", Qty: 12, FromIndex: 6, ToIndex: 2},
		{Input: "move 1 from 1 to 4", Qty: 1, FromIndex: 1, ToIndex: 4},
	}

	for _, test := range tests {
		result := parseMove(test.Input)
		assert.Equal(test.Qty, result.Qty)
		assert.Equal(test.FromIndex, result.FromIndex)
		assert.Equal(test.ToIndex, result.ToIndex)
	}
}

func TestMatrix_ParseMoves(t *testing.T) {
	assert := assert.New(t)
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")
	iDex := BuildIndex(lines[8])
	matrix := NewMatrix(lines[3:4], iDex)
	matrix.ParseMoves(lines[10:13])

	type Test struct {
		Index     int
		Qty       int
		FromIndex int
		ToIndex   int
	}

	tests := []Test{
		// move 1 from 4 to 1
		{Index: 0, Qty: 1, FromIndex: 4, ToIndex: 1},
		// move 2 from 4 to 8
		{Index: 1, Qty: 2, FromIndex: 4, ToIndex: 8},
		// move 5 from 9 to 6
		{Index: 2, Qty: 5, FromIndex: 9, ToIndex: 6},
	}

	for _, test := range tests {
		assert.Equal(test.Qty, matrix.Moves[test.Index].Qty)
		assert.Equal(test.FromIndex, matrix.Moves[test.Index].FromIndex)
		assert.Equal(test.ToIndex, matrix.Moves[test.Index].ToIndex)
	}
}

func TestMatrix_ApplyMoves(t *testing.T) {
	assert := assert.New(t)
	stack := []string{
		"    [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
	}
	index := " 1   2   3 "
	moves := []string{
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}
	iDex := BuildIndex(index)
	matrix := NewMatrix(stack, iDex)
	matrix.ParseMoves(moves)

	type Test struct {
		Index     int
		Qty       int
		FromIndex int
		ToIndex   int
	}

	tests := []Test{
		// "move 1 from 2 to 1",
		{Index: 0, Qty: 1, FromIndex: 2, ToIndex: 1},
		// "move 3 from 1 to 3",
		{Index: 1, Qty: 3, FromIndex: 1, ToIndex: 3},
		// "move 2 from 2 to 1",
		{Index: 2, Qty: 2, FromIndex: 2, ToIndex: 1},
		// "move 1 from 1 to 2",
		{Index: 3, Qty: 1, FromIndex: 1, ToIndex: 2},
	}

	for _, test := range tests {
		assert.Equal(test.Qty, matrix.Moves[test.Index].Qty)
		assert.Equal(test.FromIndex, matrix.Moves[test.Index].FromIndex)
		assert.Equal(test.ToIndex, matrix.Moves[test.Index].ToIndex)
	}

	matrix.ApplyMoves()
	assert.Equal([]string{"C"}, matrix.Values[1])
	assert.Equal([]string{"M"}, matrix.Values[2])
	assert.Equal([]string{"P", "D", "N", "Z"}, matrix.Values[3])
}

func TestMatrix_ResultMoves(t *testing.T) {
	assert := assert.New(t)
	stack := []string{
		"    [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
	}
	index := " 1   2   3 "
	moves := []string{
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}
	iDex := BuildIndex(index)
	matrix := NewMatrix(stack, iDex)
	matrix.ParseMoves(moves)
	matrix.ApplyMoves()
	assert.Equal([]string{"C"}, matrix.Values[1])
	assert.Equal([]string{"M"}, matrix.Values[2])
	assert.Equal([]string{"P", "D", "N", "Z"}, matrix.Values[3])
	assert.Equal("CMZ", matrix.Result(matrix.Values))
	assert.Equal("MCD", matrix.Result(matrix.ValuesGrouped))

	// Using the real data
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")
	matrix = NewMatrix(lines[0:8], BuildIndex(lines[8]))
	matrix.ParseMoves(lines[10:11])
	matrix.ApplyMoves()

	//         [Q] [B]         [H]
	//     [F] [W] [D] [Q]     [S]
	//     [D] [C] [N] [S] [G] [F]
	//     [R] [D] [L] [C] [N] [Q]     [R]
	// [V] [W] [L] [M] [P] [S] [M]     [M]
	// [J] [B] [F] [P] [B] [B] [P] [F] [F]
	// [B] [V] [G] [J] [N] [D] [B] [L] [V]
	// [D] [P] [R] [W] [H] [R] [Z] [W] [S]
	//  1   2   3   4   5   6   7   8   9
	// move 1 from 4 to 1
	assert.Equal("BFQDQGHFR", matrix.Result(matrix.Values))
}

func Test_reverseSlice(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]string{"c", "b", "a"}, reverseSlice([]string{"a", "b", "c"}))
	assert.Equal([]string{"3", "2", "1"}, reverseSlice([]string{"1", "2", "3"}))
	assert.Equal([]int{3, 2, 1}, reverseSlice([]int{1, 2, 3}))
}
