package main

import (
	"errors"
	"log"
	"os"
	"sort"
	"strings"
)

type itemToASCII struct {
	Min    int
	Max    int
	Offset int
}

func (item itemToASCII) Contains(letter rune) bool {
	ascii := int(letter)
	return ascii >= item.Min && ascii <= item.Max
}

func (item itemToASCII) NumericValue(letter rune) int {
	ascii := int(letter)
	return ascii - item.Offset
}

var (
	// ASCII Values of lowercase chars 97 - 122
	uppercase = itemToASCII{Min: 65, Max: 90, Offset: 64 - 26}

	// ASCII Values of lowercase chars 65 - 90
	lowercase = itemToASCII{Min: 97, Max: 122, Offset: 96}

	ErrNoMatchFound = errors.New("failed to find matching string in comparison")
	ErrNotDevisible = errors.New("list of strings given isn't devisible by given chunk size")
)

func main() {
	total := 0
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")
	chunks, err := chunk(lines, 3)
	if err != nil {
		log.Panicf("could not chunk input")
	}

	for _, items := range chunks {
		match, err := findBadge(items)
		if err != nil {
			log.Panicf("could not find a badge between elf backpacks: '%v'", items)
		}

		if uppercase.Contains(match) {
			total += uppercase.NumericValue(match)
			continue
		}

		total += lowercase.NumericValue(match)
	}

	log.Printf("Total for matching badges is: %d", total)
}

func chunk(input []string, size int) (result [][]string, err error) {
	if len(input)%size != 0 {
		return result, ErrNotDevisible
	}

	group := []string{}
	for i, str := range input {
		group = append(group, str)
		if (i+1)%size == 0 {
			result = append(result, group)
			group = []string{}
		}
	}
	return result, nil
}

func findBadge(items []string) (rune, error) {
	sort.Slice(items, func(a, b int) bool {
		return len(items[a]) < len(items[b])
	})

	for _, letter := range items[0] {
		matchingFirst := strings.ContainsRune(items[1], letter)
		if !matchingFirst {
			continue
		}

		matchingSecond := strings.ContainsRune(items[2], letter)
		if matchingSecond {
			return letter, nil
		}
	}

	return rune(0), ErrNoMatchFound
}
