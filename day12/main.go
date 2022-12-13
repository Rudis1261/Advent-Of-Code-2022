package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

var (
	costMap = map[rune]int{
		'a': 1, 'b': 2, 'c': 3,
		'd': 4, 'e': 5, 'f': 6,
		'g': 7, 'h': 8, 'i': 9,
		'j': 10, 'k': 11, 'l': 12,
		'm': 13, 'n': 14, 'o': 15,
		'p': 16, 'q': 17, 'r': 18,
		's': 19, 't': 20, 'u': 21,
		'v': 22, 'w': 23, 'x': 24,
		'y': 25, 'z': 26, 'S': 0,
		'E': 27,
	}
	vertexMap = map[int]map[int]Position{}
	visited   = map[int]map[int]bool{}
	board     = map[int]map[int]string{}
	start     = Position{}
	target    = Position{}
	path      = []Position{}

	ErrNoTarget = errors.New("no next targets")
)

type Position struct {
	Y         int    `json:"y"`
	X         int    `json:"x"`
	Char      rune   `json:"character"`
	Str       string `json:"string"`
	Value     int    `json:"value"`
	Direction string `json:"direction"`
}

type Distance struct {
	Y int `json:"y"`
	X int `json:"x"`
}

func (pos Position) Distance(another Position) (dist Distance) {
	dist.X = abs(another.X - pos.X)
	dist.Y = abs(another.Y - pos.Y)
	return dist
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

func main() {
	data, _ := os.ReadFile("input.small.txt")
	lines := strings.Split(string(data), "\n")
	point := start

	processInputData(lines)
	for {
		next := nextPoint(point, target)
		if next.X == target.X && next.Y == target.Y {
			break
		}

		path = append(path, next)
		visited[point.Y][point.X] = true
		board[point.Y][point.X] = next.Direction
		point = next
	}

	log.Printf("Result 1: %d\n", len(path))
	log.Printf("Result 2: %s\n", "lines")
}

func filter(from Position, positions []Position) (valid []Position) {
	for _, position := range positions {
		if abs(position.Value-from.Value) > 1 {
			continue
		}
		if done := visited[position.Y][position.X]; done {
			continue
		}
		valid = append(valid, position)
	}

	return valid
}

func neigbours(pos Position, target Position) (adjacent []Position) {
	if down, ok := vertexMap[pos.Y+1][pos.X]; ok {
		down.Direction = "v"
		adjacent = append(adjacent, down)
	}

	if right, ok := vertexMap[pos.Y][pos.X+1]; ok {
		right.Direction = ">"
		adjacent = append(adjacent, right)
	}

	if up, ok := vertexMap[pos.Y-1][pos.X]; ok {
		up.Direction = "^"
		adjacent = append(adjacent, up)
	}

	if left, ok := vertexMap[pos.Y][pos.X-1]; ok {
		left.Direction = "<"
		adjacent = append(adjacent, left)
	}

	return adjacent
}

func closest(points []Position, target Position) (next Position) {
	sort.Slice(points, func(i, j int) bool {
		offsetI := points[i].Distance(target)
		shortestI := offsetI.X
		if offsetI.X > offsetI.Y {
			shortestI = offsetI.Y
		}

		offsetJ := points[i].Distance(target)
		shortestJ := offsetJ.X
		if offsetJ.X > offsetJ.Y {
			shortestI = offsetJ.Y
		}

		iJson, _ := json.MarshalIndent(points[j], "", "  ")
		iOffset, _ := json.MarshalIndent(offsetI, "", "  ")
		jJson, _ := json.MarshalIndent(points[i], "", "  ")
		jOffset, _ := json.MarshalIndent(offsetJ, "", "  ")
		fmt.Printf("%s\n%s\n%s\n%s\n", iJson, iOffset, jJson, jOffset)

		return points[i].Value > points[j].Value && shortestI < shortestJ
	})

	next = points[0]
	return next
}

func nextPoint(pos Position, target Position) (next Position) {
	pointsToConsider := neigbours(pos, target)
	pointsToConsider = filter(pos, pointsToConsider)
	return closest(pointsToConsider, target)
}

func processInputData(lines []string) {
	for y, line := range lines {
		vertexLine := map[int]Position{}
		visitedLine := map[int]bool{}
		boardLine := map[int]string{}

		for x, char := range line {
			position := Position{
				X:     x,
				Y:     y,
				Char:  char,
				Str:   string(char),
				Value: costMap[char],
			}

			log.Printf("y: %d, x: %d, val: %s, pos: %v\n", y, x, string(char), position)

			if string(char) == "S" {
				start = position
			}

			if string(char) == "E" {
				target = position
			}

			vertexLine[x] = position
			visitedLine[x] = false
			boardLine[x] = "."
		}
		vertexMap[y] = vertexLine
		visited[y] = visitedLine
		board[y] = boardLine
	}
}
