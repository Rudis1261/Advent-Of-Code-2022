package main

import (
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	dirs      = make(map[string]int)
	files     = make(map[string]int)
	nodes     = make(map[string]int)
	pathRegex = regexp.MustCompile(`\$ cd (.+)`)
	sizeRegex = regexp.MustCompile(`(\d+) ([\w.]+)`)
)

func main() {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")
	buildDirTree(lines)
	total := 0

	canBeCleaned := []string{}
	for dir, size := range dirs {
		if dir == "/" {
			continue
		}
		if size < 100000 {
			total += size
		}
		minimumCleanupSize := 30000000 - (70000000 - dirs["/"])
		if size >= minimumCleanupSize {
			canBeCleaned = append(canBeCleaned, dir)
		}
	}

	// Sort cleanable directories by size
	sort.Slice(canBeCleaned, func(i, j int) bool {
		return nodes[canBeCleaned[i]] < nodes[canBeCleaned[j]]
	})

	// And amount of nodes
	sort.Slice(canBeCleaned, func(i, j int) bool {
		return dirs[canBeCleaned[i]] < dirs[canBeCleaned[j]]
	})

	log.Printf("Result 1: Total: %v\n", total)
	log.Printf("Result 2: Dir: %v, Size: %v, Nodes: %v\n", canBeCleaned[0], dirs[canBeCleaned[0]], nodes[canBeCleaned[0]])
}

func buildDirTree(lines []string) {
	path := []string{}
	for _, line := range lines {
		if strings.HasPrefix(line, "$ ls") || strings.HasPrefix(line, "dir") {
			continue
		}

		if strings.HasPrefix(line, "$ cd ..") {
			path = path[0 : len(path)-1]
			continue
		}

		if strings.HasPrefix(line, "$ cd") {
			matches := pathRegex.FindAllStringSubmatch(line, -1)
			path = append(path, matches[0][1])
			continue
		}

		key := strings.Join(path[:], "/")
		fileAndSize := sizeRegex.FindAllStringSubmatch(line, -1)
		size, _ := strconv.Atoi(fileAndSize[0][1])
		fileName := fileAndSize[0][2]
		fullpathName := strings.ReplaceAll(key+"/"+fileName, "//", "/")
		files[fullpathName] = size

		for i := 1; i <= len(path); i++ {
			dirKey := strings.ReplaceAll(strings.Join(path[0:i], "/"), "//", "/")
			if dirKey == "" {
				dirKey = "/"
			}
			dirs[dirKey] += size
			nodes[dirKey] += 1
		}
	}
}
