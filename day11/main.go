package main

import (
	"encoding/json"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt.full")
	lines := strings.Split(string(data), "\n\n")

	monkies := processLines(lines)
	lcm := monkeyLCM(monkies)

	values := []int{}
	for _, monkey := range monkeyBusiness(monkies, 20, true, lcm) {
		values = append(values, monkey.Inspected)
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})

	monkies = processLines(lines)
	lcm = monkeyLCM(monkies)

	values2 := []int{}
	for _, monkey := range monkeyBusiness(monkies, 10000, false, lcm) {
		values2 = append(values2, monkey.Inspected)
	}
	sort.Slice(values2, func(i, j int) bool {
		return values2[i] > values2[j]
	})

	log.Printf("Result 1: %d\n", values[0]*values[1])
	log.Printf("Result 2: %d\n", values2[0]*values2[1])
}

type Operation struct {
	Subject string `json:"subject"`
	Ops     string `json:"ops"`
}

type Monkey struct {
	Number    int       `json:"number"`
	Items     []int     `json:"items"`
	Operation Operation `json:"operation"`
	Test      int       `json:"test"`
	OnTrue    int       `json:"on_true"`
	OnFalse   int       `json:"on_false"`
	Inspected int       `json:"inspected"`
	LCM       int       `json:"lowest_common_multiple"`
}

func monkeyLCM(monkies []Monkey) int {
	lcm := []int{}
	for _, monkey := range monkies {
		lcm = append(lcm, monkey.Test)
	}

	return LCM(lcm[0], lcm[1], lcm[2:]...)
}

func monkeyBusiness(monkies []Monkey, rounds int, devide bool, lcm int) []Monkey {
	for i := 1; i <= rounds; i++ {
		for x := 0; x < len(monkies); x++ {
			beforeItems := monkies[x].Items
			if len(monkies[x].Items) == 0 {
				continue
			}

			for _, item := range beforeItems {
				wl := worryLevel(monkies[x], item, devide)
				monkies[x].Inspected++
				index := monkies[x].OnFalse

				if wl%monkies[x].Test == 0 {
					index = monkies[x].OnTrue
				}

				monkies[index].Items = append(monkies[index].Items, wl%lcm)

				if len(monkies[x].Items) > 0 {
					monkies[x].Items = monkies[x].Items[1:]
				} else {
					monkies[x].Items = []int{}
				}
			}
		}
	}

	return monkies
}

func worryLevel(monkey Monkey, item int, devide bool) int {
	by := 0
	newValue := item

	if monkey.Operation.Subject == "old" {
		by = item
	} else {
		by, _ = strconv.Atoi(monkey.Operation.Subject)
	}

	if monkey.Operation.Ops == "*" {
		newValue = item * by
	}
	if monkey.Operation.Ops == "+" {
		newValue = item + by
	}

	if devide {
		return int(math.Floor(float64(newValue) / 3))
	}

	return newValue
}

func processLines(lines []string) (monkies []Monkey) {
	for _, line := range lines {
		monkey := Monkey{}
		data := strings.Split(line, "\n")

		// Monkey number
		m := regexp.MustCompile(`Monkey (\d):`).FindAllStringSubmatch(data[0], -1)
		number, _ := strconv.Atoi(m[0][1])
		monkey.Number = number

		// Starting items
		m = regexp.MustCompile(`Starting items: ([\d, ]+)`).FindAllStringSubmatch(data[1], -1)
		for _, i := range strings.Split(m[0][1], ", ") {
			number, _ := strconv.Atoi(i)
			monkey.Items = append(monkey.Items, number)
		}

		// Operation
		m = regexp.MustCompile(`Operation: new = old ([-+*/])\s(\w+)`).FindAllStringSubmatch(data[2], -1)
		monkey.Operation.Ops = m[0][1]
		monkey.Operation.Subject = m[0][2]

		// Test
		m = regexp.MustCompile(`Test: divisible by (.+)`).FindAllStringSubmatch(data[3], -1)
		devisible, _ := strconv.Atoi(m[0][1])
		monkey.Test = devisible

		// OnTrue
		m = regexp.MustCompile(`If true: throw to monkey (\d)`).FindAllStringSubmatch(data[4], -1)
		onTrue, _ := strconv.Atoi(m[0][1])
		monkey.OnTrue = onTrue

		// OnFalse
		m = regexp.MustCompile(`If false: throw to monkey (\d)`).FindAllStringSubmatch(data[5], -1)
		onFalse, _ := strconv.Atoi(m[0][1])
		monkey.OnFalse = onFalse

		monkies = append(monkies, monkey)
		monkeyAsString, _ := json.MarshalIndent(monkey, "", "  ")
		log.Printf("Monkey: %v", string(monkeyAsString))
	}

	return monkies
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
