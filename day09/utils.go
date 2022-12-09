package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInputFile(name string) []string {
	raw, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}

	str := string(raw)
	return strings.Split(str, "\n")
}

func toInt(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		msg := fmt.Sprintf("Parsing %v to int failed!", str)
		panic(msg)
	}
	return n
}

func absDiff(a, b int) uint {
	if a > b {
		return uint(a - b)
	}
	return uint(b - a)
}

func makeAndInitTails(numberOfTails int) []*Position {
	tails := make([]*Position, numberOfTails)
	for i := range tails {
		tails[i] = &Position{}
	}
	return tails
}

func moveToSecondIfNeeded(first *Position, second Position) {
	if isTooFarAway(*first, second) {
		first.moveCloserTo(second)
	}
}

func isTooFarAway(head Position, tail Position) bool {
	return absDiff(head.y, tail.y) > 1 || absDiff(head.x, tail.x) > 1
}
