package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LCM(t *testing.T) {
	assert := assert.New(t)
	type Test struct {
		A      int
		B      int
		C      []int
		Result int
	}
	tests := []Test{
		{A: 10, B: 15, Result: 30},
		{A: 3, B: 5, Result: 15},
		{A: 1, B: 9, Result: 9},
		{A: 199, B: 14, Result: 2786},
	}
	for _, tc := range tests {
		assert.Equal(tc.Result, LCM(tc.A, tc.B, tc.C...))
	}
}
