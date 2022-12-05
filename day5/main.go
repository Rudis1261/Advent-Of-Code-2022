package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Index struct {
	Columns      map[int]int
	InverseIndex map[int]int
}

func BuildIndex(indexLine string) (iDex Index) {
	iDex.Columns = make(map[int]int)
	iDex.InverseIndex = make(map[int]int)

	for index, columnName := range indexLine {
		val, err := strconv.Atoi(string(columnName))
		if err == nil {
			iDex.Columns[val] = index
			iDex.InverseIndex[index] = val
		}
	}

	return iDex
}

type Matrix struct {
	Index
	Values        map[int][]string
	ValuesGrouped map[int][]string
	Moves         []Action
}

func NewMatrix(lines []string, iDex Index) (matrix Matrix) {
	lines = reverseSlice(lines)
	matrix.Index = iDex
	matrix.Values = make(map[int][]string)
	matrix.ValuesGrouped = make(map[int][]string)

	for _, line := range lines {
		for index, box := range line {
			val, ok := matrix.Index.InverseIndex[index]

			if len(matrix.Values[val]) == 0 {
				matrix.Values[val] = []string{}
				matrix.ValuesGrouped[val] = []string{}
			}

			if ok && box != ' ' {
				matrix.Values[val] = append(matrix.Values[val], string(box))
				matrix.ValuesGrouped[val] = append(matrix.ValuesGrouped[val], string(box))
			}
		}
	}

	return matrix
}

func (matrix *Matrix) ParseMoves(moves []string) Matrix {
	for _, move := range moves {
		matrix.Moves = append(matrix.Moves, parseMove(move))
	}
	return *matrix
}

func (matrix *Matrix) ApplyMoves() Matrix {
	for _, move := range matrix.Moves {
		// CrateMover 9000
		length := len(matrix.Values[move.FromIndex])
		subSlice := length - move.Qty
		toMove := matrix.Values[move.FromIndex][subSlice:length]
		matrix.Values[move.FromIndex] = matrix.Values[move.FromIndex][0:subSlice]
		toMove = reverseSlice(toMove)
		matrix.Values[move.ToIndex] = append(matrix.Values[move.ToIndex], toMove...)

		// CrateMover 9001
		length = len(matrix.ValuesGrouped[move.FromIndex])
		subSlice = length - move.Qty
		toMove = matrix.ValuesGrouped[move.FromIndex][subSlice:length]
		matrix.ValuesGrouped[move.FromIndex] = matrix.ValuesGrouped[move.FromIndex][0:subSlice]
		matrix.ValuesGrouped[move.ToIndex] = append(matrix.ValuesGrouped[move.ToIndex], toMove...)
	}
	return *matrix
}

func (matrix Matrix) Result(field map[int][]string) (result string) {
	for i := 1; i <= len(matrix.Index.Columns); i++ {
		index := len(field[i]) - 1
		result += field[i][index]
	}
	return result
}

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")
	matrix := NewMatrix(lines[0:8], BuildIndex(lines[8]))
	matrix.ParseMoves(lines[10:])
	matrix.ApplyMoves()

	log.Printf("Result from all moves: %s\n", matrix.Result(matrix.Values))
	log.Printf("Result from all moves CrateMover9001: %s\n", matrix.Result(matrix.ValuesGrouped))
}

type Action struct {
	Qty       int
	FromIndex int
	ToIndex   int
}

func parseMove(line string) (action Action) {
	re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	matches := re.FindAllStringSubmatch(line, -1)
	action.Qty, _ = strconv.Atoi(matches[0][1])
	action.FromIndex, _ = strconv.Atoi(matches[0][2])
	action.ToIndex, _ = strconv.Atoi(matches[0][3])
	return action
}

func reverseSlice[T comparable](s []T) (reversed []T) {
	for i := len(s) - 1; i >= 0; i-- {
		reversed = append(reversed, s[i])
	}
	return reversed
}
