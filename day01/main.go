package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://adventofcode.com/2022/day/1
func main() {
	input, err := os.ReadFile("./day01/input.txt")
	if err != nil {
		panic(err)
	}

	caloriesLog := splitLines(input)

	part1Answer := solvePart1(caloriesLog)
	fmt.Printf("Part one answer: %d", part1Answer)

	part2Answer := solvePart2(caloriesLog)
	fmt.Printf("Part two answer: %d", part2Answer)
}

func solvePart2(caloriesLog []string) int {
	calories := toCalories(caloriesLog)
	sort.Ints(calories)

	sum := 0
	for _, c := range calories[len(calories)-3:] {
		sum += c
	}

	return sum
}

func splitLines(input []byte) []string {
	return strings.Split(string(input), "\n")
}

func solvePart1(caloriesLog []string) int {
	calories := toCalories(caloriesLog)
	sort.Ints(calories)

	return calories[len(calories)-1]
}

func toCalories(caloriesLog []string) []int {
	var calories []int
	var current = 0

	for _, line := range caloriesLog {
		if line == "" {
			calories = append(calories, current)
			current = 0
		} else {
			n, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			current += n
		}
	}

	calories = append(calories, current)
	return calories
}
