package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	intervals = []int{
		20,
		60,
		100,
		140,
		180,
		220,
	}
)

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	operations := parseLines(lines)
	total, crt := processOperations(operations)

	log.Printf("Result 1: %d\n", total)
	log.Println("Result 2:")
	printCrt(crt)
}

type Operation struct {
	Type  string
	Value int
}

func parseLines(lines []string) (ops []Operation) {
	for _, line := range lines {
		split := strings.Split(line, " ")
		op := split[0]
		value := 0
		if len(split) > 1 {
			value, _ = strconv.Atoi(split[1])
		}
		ops = append(ops, Operation{Type: op, Value: value})
	}
	return ops
}

func processOperations(ops []Operation) (int, []string) {
	cycle := 1
	register := 1
	total := 0
	crt := []string{}

	for _, op := range ops {
		if op.Type == "noop" {
			total += addRegisterValue(cycle, register)
			crt = append(crt, addCrtImage(cycle, register))
			cycle++
			continue
		}

		if op.Type == "addx" {
			for i := 0; i < 2; i++ {
				total += addRegisterValue(cycle, register)
				crt = append(crt, addCrtImage(cycle, register))
				if i == 1 {
					register += op.Value
				}
				cycle++
			}
		}
	}

	return total, crt
}

func printCrt(crt []string) {
	for index, val := range crt {
		fmt.Print(val)
		if ((index + 1) % 40) == 0 {
			fmt.Println("")
		}
	}
}
func addCrtImage(cycle, register int) string {
	position := cycle % 40
	if position == 0 {
		position = 40
	}
	lower := register
	upper := register + 2
	if position >= lower && position <= upper {
		return "#"
	}
	return "."
}

func addRegisterValue(cycle, register int) int {
	for _, val := range intervals {
		if cycle == val {
			log.Printf("Cycle: %d, Register: %d, Adding: %d", cycle, register, cycle*register)
			return cycle * register
		}
	}

	return 0
}
