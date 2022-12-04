package main

import (
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	ErrNotDevisible = errors.New("list of strings given isn't devisible by given chunk size")
)

func main() {
	total := 0
	totalOverlapping := 0
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		sections := strings.Split(line, ",")
		if encapsulated(sections[0], sections[1]) {
			total += 1
		}
		if overlapping(sections[0], sections[1]) {
			totalOverlapping += 1
		}
	}

	log.Printf("Total for encasulating: %d\n", total)
	log.Printf("Total for overlapping: %d\n", totalOverlapping)
}

func encapsulated(a, b string) bool {
	aRange := toRange(a)
	bRange := toRange(b)
	aLower := inRange(aRange[0], bRange)
	aUpper := inRange(aRange[1], bRange)
	bLower := inRange(bRange[0], aRange)
	bUpper := inRange(bRange[1], aRange)

	return (aLower && aUpper) || (bLower && bUpper)
}

func overlapping(a, b string) bool {
	aRange := toRange(a)
	bRange := toRange(b)
	aLower := inRange(aRange[0], bRange)
	aUpper := inRange(aRange[1], bRange)
	bLower := inRange(bRange[0], aRange)
	bUpper := inRange(bRange[1], aRange)

	return (aLower || aUpper) || (bLower || bUpper)
}

func toRange(input string) (ints []int) {
	for _, letter := range strings.Split(input, "-") {
		asInt, _ := strconv.Atoi(letter)
		ints = append(ints, asInt)
	}
	return ints
}

func inRange(a int, inRange []int) bool {
	return a >= inRange[0] && a <= inRange[1]
}
