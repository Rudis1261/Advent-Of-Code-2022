package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")
	count, scenicScore := getCount(lines)

	log.Printf("Result 1: %d\n", count)
	log.Printf("Result 2: %d\n", scenicScore)
}

func getCount(lines []string) (int, int) {
	canSee := 0
	maxScenic := 0
	for y := 1; y < len(lines)-1; y++ {
		for x := 1; x < len(lines[y])-1; x++ {
			if visible(lines, x, y) {
				canSee++
			}

			scenicScore := scenicScore(lines, x, y)
			if scenicScore > maxScenic {
				maxScenic = scenicScore
			}
		}
	}

	totalVisible := (len(lines[0]) * 2) + ((len(lines) - 2) * 2) // along the edge
	return totalVisible + canSee, maxScenic
}

func visible(lines []string, x int, y int) bool {
	horizontal := lines[y]
	currentValue, _ := strconv.Atoi(string(lines[y][x]))
	var vertical string

	for i := 0; i < len(lines); i++ {
		vertical += string(lines[i][x])
	}

	xMin := horizontal[0:x]
	xMax := horizontal[x+1:]
	yMin := vertical[0:y]
	yMax := vertical[y+1:]

	xMinVisible := true
	xMaxVisible := true
	yMinVisible := true
	yMaxVisible := true

	for _, val := range xMin {
		adjacentTree, _ := strconv.Atoi(string(val))
		if currentValue <= adjacentTree {
			xMinVisible = false
		}
	}

	for _, val := range xMax {
		adjacentTree, _ := strconv.Atoi(string(val))
		if currentValue <= adjacentTree {
			xMaxVisible = false
		}
	}

	for _, val := range yMin {
		adjacentTree, _ := strconv.Atoi(string(val))
		if currentValue <= adjacentTree {
			yMinVisible = false
		}
	}

	for _, val := range yMax {
		adjacentTree, _ := strconv.Atoi(string(val))
		if currentValue <= adjacentTree {
			yMaxVisible = false
		}
	}

	return xMaxVisible || xMinVisible || yMinVisible || yMaxVisible
}

func scenicScore(lines []string, x int, y int) int {
	horizontal := lines[y]
	currentValue, _ := strconv.Atoi(string(lines[y][x]))
	var vertical string
	for i := 0; i < len(lines); i++ {
		vertical += string(lines[i][x])
	}

	xMin := horizontal[0:x]
	xMax := horizontal[x+1:]
	yMin := vertical[0:y]
	yMax := vertical[y+1:]

	xMinScenicScore := 0
	xMaxScenicScore := 0
	yMinScenicScore := 0
	yMaxScenicScore := 0

	for i := len(xMin) - 1; i >= 0; i-- {
		adjacentTree, _ := strconv.Atoi(string(xMin[i]))
		xMinScenicScore++
		if currentValue <= adjacentTree {
			break
		}
	}

	for i := 0; i < len(xMax); i++ {
		adjacentTree, _ := strconv.Atoi(string(xMax[i]))
		xMaxScenicScore++
		if currentValue <= adjacentTree {
			break
		}
	}

	for i := len(yMin) - 1; i >= 0; i-- {
		adjacentTree, _ := strconv.Atoi(string(yMin[i]))
		yMinScenicScore++
		if currentValue <= adjacentTree {
			break
		}
	}

	for i := 0; i < len(yMax); i++ {
		adjacentTree, _ := strconv.Atoi(string(yMax[i]))
		yMaxScenicScore++
		if currentValue <= adjacentTree {
			break
		}
	}

	if xMinScenicScore == 0 {
		xMinScenicScore = 1
	}
	if xMaxScenicScore == 0 {
		xMaxScenicScore = 1
	}
	if yMinScenicScore == 0 {
		yMinScenicScore = 1
	}
	if yMaxScenicScore == 0 {
		yMaxScenicScore = 1
	}

	return yMinScenicScore * xMinScenicScore * xMaxScenicScore * yMaxScenicScore
}
