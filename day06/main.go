package main

import (
	"fmt"
	"os"
)

func main() {
	raw, err := os.ReadFile("./day06/input.txt")
	if err != nil {
		panic(err)
	}

	input := string(raw)

	result1 := solvePart1(input)
	fmt.Printf("Part 1 answer: %d\n", result1)

	result2 := solvePart2(input)
	fmt.Printf("Part 2 answer: %d\n", result2)
}

func solvePart2(input string) int {
	return findUniqueSubSequence(input, 14)
}

func solvePart1(input string) int {
	return findUniqueSubSequence(input, 4)
}

func findUniqueSubSequence(input string, wantedSize int) int {
	seqStart := 0
	currentSize := 1

	fullSeq := []rune(input)

	for i := 1; i < len(fullSeq); i++ {
		c := fullSeq[i]
		li, found := find(fullSeq[seqStart:i], c)
		currentSize += 1

		if currentSize == wantedSize && !found {
			return i + 1
		}

		if found {
			seqStart += li + 1
			currentSize = i - seqStart + 1
		}
	}

	msg := fmt.Sprintf("input %v doesn't contains uniques sub sequence with lenght %v", input, wantedSize)
	panic(msg)
}

func find(runes []rune, c rune) (index int, found bool) {
	for i := range runes {
		if runes[i] == c {
			return i, true
		}
	}

	return -1, false
}
