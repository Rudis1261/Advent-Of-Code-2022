package main

import (
	"encoding/json"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	directions = map[string]Position{
		"U": {1, 0},
		"D": {-1, 0},
		"R": {0, 1},
		"L": {0, -1},
	}
	// head    = Position{0, 0}
	// tail    = Position{0, 0}
	visited = []Position{}
	rope    = []Position{}
)

type Position struct {
	Y int `json:"y"`
	X int `json:"x"`
}

func main() {
	data, _ := os.ReadFile("input.medium.txt")
	lines := strings.Split(string(data), "\n")

	for i := 0; i < 10; i++ {
		rope = append(rope, Position{})
	}

	for _, line := range lines {
		parts := strings.Split(line, " ")
		direction := directions[parts[0]]
		amount, _ := strconv.Atoi(parts[1])

		for a := 0; a < amount; a++ {
			rope[0].X += direction.X
			rope[0].Y += direction.Y

			for i := 0; i < 9; i++ {
				head := &rope[i]
				tail := &rope[i+1]

				diff := Position{
					Y: head.Y - tail.Y,
					X: head.X - tail.X,
				}

				if abs(diff.X) > 1 {
					tail.X += floor(diff.X / 2)
					if diff.Y == 0 {

					} else if diff.Y > 0 {
						tail.Y += 1
					} else {
						tail.Y -= 1
					}
				}

				if abs(diff.Y) > 1 {
					tail.Y += floor(diff.Y / 2)
					if diff.X == 0 {

					} else if diff.X > 0 {
						tail.X += 1
					} else {
						tail.X -= 1
					}
				}
			}

			visited = append(visited, rope[9])
		}
	}

	log.Printf("Visited: %d, v: %v", len(unique(visited)), js(unique(visited)))
}

func js(input interface{}) (res string) {
	data, _ := json.MarshalIndent(input, "", "    ")
	return string(data)
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

func floor(a int) int {
	return int(math.Floor(float64(a)))
}

func unique(slice []Position) []Position {
	uniqMap := make(map[Position]struct{})
	for _, v := range slice {
		uniqMap[v] = struct{}{}
	}

	uniqSlice := make([]Position, 0, len(uniqMap))
	for v := range uniqMap {
		uniqSlice = append(uniqSlice, v)
	}
	return uniqSlice
}
