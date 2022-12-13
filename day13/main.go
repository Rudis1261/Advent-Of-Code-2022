package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	Parent *Pair
	Value  int
}

func main() {
	data, _ := os.ReadFile("input.small.txt")
	lines := strings.Split(string(data), "\n\n")
	values := processInput(lines)

	log.Printf("Result 1: %v\n", values)
	log.Printf("Result 2: %s\n", "lines")
}

func processInput(lines []string) (pairs []Pair) {
	for _, line := range lines {
		set := []int{}
		for _, value := range strings.Split(line, "\n") {
			val, _ := strconv.Atoi(value)
			set = append(set, val)
		}
		// pairs = append(pairs, set)
	}
	return pairs
}
