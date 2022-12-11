package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	directions = map[string]Position{
		"U": {-1, 0},
		"D": {1, 0},
		"L": {0, -1},
		"R": {0, 1},
	}
	head    = Position{0, 0}
	tail    = Position{0, 0}
	visited = []Position{}
)

type Position struct {
	Y int
	X int
}

type Move struct {
	Direction Position
	Amount    int
}

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	parseMoves(parseLines(lines))

	log.Printf("Result 1: %d\n", len(visited))
	log.Printf("Result 2: %v\n", visited)
}

func parseLines(lines []string) (moves []Move) {
	for _, line := range lines {
		parts := strings.Split(line, " ")
		amount, _ := strconv.Atoi(parts[1])

		moves = append(moves, Move{
			Amount:    amount,
			Direction: directions[parts[0]],
		})
	}

	return moves
}

func parseMoves(moves []Move) {
	for _, move := range moves {
		for i := 0; i < move.Amount; i++ {
			head.Y += move.Direction.Y
			head.X += move.Direction.X

			diff := Position{}
			diffHeadTail := Position{
				Y: tail.Y - head.Y,
				X: tail.X - head.X,
			}

			switch true {
			case diffHeadTail.Y == 2 && diffHeadTail.X == 0:
				diff.Y = 1
				diff.X = 0
			case diffHeadTail.Y == -2 && diffHeadTail.X == 0:
				diff.Y = -1
				diff.X = 0
			case diffHeadTail.Y == 0 && diffHeadTail.X == 2:
				diff.Y = 0
				diff.X = 1
			case diffHeadTail.Y == 0 && diffHeadTail.X == -2:
				diff.Y = 0
				diff.X = -1
				// default:
				// 	diff.Y = diffHeadTail.Y
				// 	diff.X = diffHeadTail.X
			}

			newPos := Position{
				Y: head.Y + diff.Y,
				X: head.X + diff.X,
			}

			visited = append(visited, newPos)
		}
	}
}
