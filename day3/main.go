package main

import (
	"errors"
	"log"
	"os"
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
)

func main() {
	total := 0
	data, _ := os.ReadFile("input.txt")
	for _, line := range strings.Split(string(data), "\n") {
		items := splitItems(line)
		match, err := findMatchingCharacter(items[0], items[1])
		if err != nil {
			log.Panicf("could not find a matching character between strings: '%s' and '%s'", items[0], items[1])
		}

		if uppercase.Contains(match) {
			total += uppercase.NumericValue(match)
			continue
		}

		total += lowercase.NumericValue(match)
	}

	log.Printf("Total for matching chars is: %d", total)
}

func findMatchingCharacter(a, b string) (rune, error) {
	for _, letter := range a {
		if strings.ContainsRune(b, letter) {
			return letter, nil
		}
	}
	return rune(0), ErrNoMatchFound
}

func splitItems(input string) (items []string) {
	halfStringLenght := len(input) / 2
	items = append(items, input[0:halfStringLenght])
	items = append(items, input[halfStringLenght:])
	return items
}
