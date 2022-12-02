package main

import (
	"log"
	"os"
	"strings"
)

type RockPaperScisors map[string]Gesture
type Gesture struct {
	Name  string
	Point int
	Beats string
	Loss  string
}

var (
	ROCK    = "rock"
	PAPER   = "paper"
	SCISSOR = "scissor"
	WIN     = "win"
	LOSS    = "loss"
	DRAW    = "draw"

	winLooseCondition = map[string]int{
		WIN:  6,
		DRAW: 3,
		LOSS: 0,
	}

	game = RockPaperScisors{
		ROCK: Gesture{
			Name:  ROCK,
			Point: 1,
			Beats: SCISSOR,
			Loss:  PAPER,
		},
		PAPER: Gesture{
			Name:  PAPER,
			Point: 2,
			Beats: ROCK,
			Loss:  SCISSOR,
		},
		SCISSOR: Gesture{
			Name:  SCISSOR,
			Point: 3,
			Beats: PAPER,
			Loss:  ROCK,
		},
	}

	gestures = map[string]string{
		"A": ROCK,
		"X": LOSS,
		"B": PAPER,
		"Y": DRAW,
		"C": SCISSOR,
		"Z": WIN,
	}
)

func main() {
	total := 0
	data, _ := os.ReadFile("input.txt")
	for _, line := range strings.Split(string(data), "\n") {
		parts := strings.Split(line, " ")
		opponent := parts[0]
		targetOutcome := parts[1]
		total += score(opponent, targetOutcome)
	}

	log.Printf("Total is: %d", total)
}

func score(opponent, targetOutcome string) int {
	if gestures[targetOutcome] == DRAW {
		return game[gestures[opponent]].Point + winLooseCondition[DRAW]
	}

	if gestures[targetOutcome] == WIN {
		idealGesture := game[gestures[opponent]].Loss
		return game[idealGesture].Point + winLooseCondition[WIN]
	}

	idealGesture := game[gestures[opponent]].Beats
	return game[idealGesture].Point + winLooseCondition[LOSS]
}
