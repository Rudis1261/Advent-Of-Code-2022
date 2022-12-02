package main

import (
	"log"
	"os"
	"strings"
)

type RockPaperScissors map[string]Gesture
type Gesture struct {
	Name  string
	Point int
	Beats string
}

var (
	ROCK    = "rock"
	PAPER   = "paper"
	SCISSOR = "scissor"
	WIN     = "win"
	LOSS    = "loss"
	DRAW    = "draw"

	winLossCondition = map[string]int{
		WIN:  6,
		DRAW: 3,
		LOSS: 0,
	}

	game = RockPaperScissors{
		ROCK: Gesture{
			Name:  ROCK,
			Point: 1,
			Beats: SCISSOR,
		},
		PAPER: Gesture{
			Name:  PAPER,
			Point: 2,
			Beats: ROCK,
		},
		SCISSOR: Gesture{
			Name:  SCISSOR,
			Point: 3,
			Beats: PAPER,
		},
	}

	gestures = map[string]string{
		"A": ROCK,
		"X": ROCK,
		"B": PAPER,
		"Y": PAPER,
		"C": SCISSOR,
		"Z": SCISSOR,
	}
)

func main() {
	total := 0
	data, _ := os.ReadFile("input.txt")
	for _, line := range strings.Split(string(data), "\n") {
		parts := strings.Split(line, " ")
		opponent := parts[0]
		you := parts[1]
		total += score(opponent, you)
	}

	log.Printf("Total is: %d", total)
}

func score(opponent, you string) int {
	if gestures[you] == gestures[opponent] {
		return game[gestures[you]].Point + winLossCondition[DRAW]
	}

	if game[gestures[you]].Beats == gestures[opponent] {
		return game[gestures[you]].Point + winLossCondition[WIN]
	}

	return game[gestures[you]].Point + winLossCondition[LOSS]
}
