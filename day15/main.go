package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	log.Printf("Result 1: %s\n", lines)
	log.Printf("Result 2: %s\n", lines)
}
