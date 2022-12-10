package main

import (
	"log"
	"os"
	"strings"

	"github.com/Rudis1261/Advent-Of-Code-2022/day9/board"
)

var (
	head = board.Tracking{}
	tail = board.Tracking{}
)

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")
	head.Init()
	tail.Init()

	log.Printf("Result 1: %s\n", lines)
	log.Printf("Result 2: %s\n", "lines")
}
