package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TreeMap [][]int

func main() {
	raw := readInputFile("./day08/input.txt")
	treeMap := generateTreeMap(raw)

	result1 := solvePart1(treeMap)
	fmt.Printf("Part 1 answer: %d\n", result1)

	result2 := solvePart2(treeMap)
	fmt.Printf("Part 2 answer: %d\n", result2)
}

func readInputFile(name string) []byte {
	raw, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return raw
}

func solvePart2(treeMap TreeMap) int {
	highestScore := 0

	for y := 0; y < len(treeMap); y++ {
		for x := 0; x < len(treeMap); x++ {
			score := calculateScore(treeMap, y, x)
			if score > highestScore {
				highestScore = score
			}
		}
	}

	return highestScore
}

func calculateScore(treeMap TreeMap, y int, x int) int {
	if isOnEdge(treeMap, x) || isOnEdge(treeMap, y) {
		return 0
	}

	isHigherOrEqualToMaxHeight := func(iy, ix int) bool {
		maxHeight := treeMap[y][x]
		currHeight := treeMap[iy][ix]
		return currHeight >= maxHeight
	}

	score := 1

	// look up
	dist := 0
	for ty := y - 1; ty >= 0; ty-- {
		dist += 1
		if isHigherOrEqualToMaxHeight(ty, x) {
			break
		}
	}
	score *= dist

	// look down
	dist = 0
	for ty := y + 1; ty < len(treeMap); ty++ {
		dist += 1
		if isHigherOrEqualToMaxHeight(ty, x) {
			break
		}
	}
	score *= dist

	// look right
	dist = 0
	for tx := x + 1; tx < len(treeMap); tx++ {
		dist += 1
		if isHigherOrEqualToMaxHeight(y, tx) {
			break
		}
	}
	score *= dist

	// look left
	dist = 0
	for tx := x - 1; tx >= 0; tx-- {
		dist += 1
		if isHigherOrEqualToMaxHeight(y, tx) {
			break
		}
	}
	score *= dist

	return score
}

func solvePart1(treeMap TreeMap) int {
	size := len(treeMap)

	// trees around the edge
	sum := size*4 - 4

	// exclude first & last row
	for y := 1; y < size-1; y++ {
		// exclude first & last column
		for x := 1; x < size-1; x++ {
			if isVisible(treeMap, y, x) {
				sum += 1
			}
		}
	}

	return sum
}

func isVisible(tm TreeMap, y, x int) bool {
	if isOnEdge(tm, y) || isOnEdge(tm, x) {
		return true
	}

	tree := tm[y][x]
	biggestBefore := 0
	biggestAfter := 0
	for ty := 0; ty < len(tm); ty++ {
		curr := tm[ty][x]
		if ty < y && curr > biggestBefore {
			biggestBefore = curr
		}
		if ty > y && curr > biggestAfter {
			biggestAfter = curr
		}
	}

	if biggestBefore < tree || biggestAfter < tree {
		return true
	}

	biggestBefore = 0
	biggestAfter = 0
	for tx := 0; tx < len(tm); tx++ {
		curr := tm[y][tx]

		if tx < x && curr > biggestBefore {
			biggestBefore = curr
		}
		if tx > x && curr > biggestAfter {
			biggestAfter = curr
		}
	}

	return biggestBefore < tree || biggestAfter < tree
}

func isOnEdge(tm TreeMap, i int) bool {
	return i == 0 || i == len(tm)-1
}

func generateTreeMap(raw []byte) TreeMap {
	rows := strings.Split(string(raw), "\n")
	length := len(rows)

	tm := make([][]int, length)

	for rowIndex, row := range rows {
		tm[rowIndex] = make([]int, length)

		columns := strings.Split(row, "")
		for columnIndex, value := range columns {
			tm[rowIndex][columnIndex] = toInt(value)
		}
	}

	return tm
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
