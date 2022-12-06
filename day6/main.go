package main

import (
	"log"
	"os"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	chars := string(data)
	first := findFirstUnique(chars, 4)
	message := findFirstUnique(chars, 14)

	log.Printf("Result 1: %d\n", first)
	log.Printf("Result 2: %v\n", message)
}

func findFirstUnique(input string, size int) int {
	for i := size; i < len(input); i++ {
		isFound := unique(input[i-size : i])
		if isFound {
			return i
		}
	}
	return 0
}

func findAllUnique(input string, size int) map[int]string {
	uniques := make(map[int]string)
	for i := size; i < len(input); i++ {
		isFound := unique(input[i-size : i])
		if isFound {
			uniques[i] = input[i-size : i]
		}
	}
	return uniques
}

func unique(abc string) bool {
	values := make(map[rune]bool)
	for _, i := range abc {
		_, ok := values[i]
		if ok {
			return false
		}

		values[i] = true
	}
	return true
}
