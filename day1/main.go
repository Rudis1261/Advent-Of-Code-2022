package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	elves := map[int]int{}
	sorted := map[int]int{}
	values := []int{}
	elf := 1
	data, _ := os.ReadFile("input.txt")

	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			elf++
			continue
		}

		calories, _ := strconv.Atoi(line)
		elves[elf] = elves[elf] + calories
	}

	// Flip the index, which automatically will sort it low to high, indexed by the value (if we need the elf)
	// And append the values to a flat index array
	for elf, calories := range elves {
		values = append(values, calories)
		sorted[calories] = elf
	}

	sort.Ints(values)
	max := values[len(values)-1]
	maxThree := sum(values[len(values)-3:])
	fmt.Printf("Single elf max: %d\nTop three total: %v\n", max, maxThree)
}

func sum(values []int) int {
	sum := 0
	for _, i := range values {
		sum += i
	}
	return sum
}
