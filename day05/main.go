package main

import (
	"adian.com/advent_of_code_2022/common/stack"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./day05/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	result1 := solvePart1(lines)
	fmt.Printf("Part 1 answer: %s\n", result1)

	result2 := solvePart2(lines)
	fmt.Printf("Part 2 answer: %s\n", result2)
}

func solvePart2(lines []string) string {
	sa, moves := parse(lines)

	for _, m := range moves {
		temp := sa[m.from].Take(m.number)
		sa[m.to].PutAll(temp...)
	}

	return getAnswer(sa)
}

func solvePart1(lines []string) string {
	sa, moves := parse(lines)

	for _, m := range moves {
		for i := 0; i < m.number; i++ {
			value, ok := sa[m.from].TakeOne()
			if !ok {
				panic("Peek unexpectedly returned empty value!")
			}
			sa[m.to].Put(value)
		}
	}

	return getAnswer(sa)
}

func getAnswer(sa stacksArrangement) string {
	result := ""

	for i := 1; i <= len(sa); i++ {
		currentStack := sa[i]
		value, ok := currentStack.Peek()
		if !ok {
			panic("Peek unexpectedly returned empty value!")
		}
		result += value
	}

	return result
}

func parse(lines []string) (stacksArrangement, []move) {
	var sa stacksArrangement
	var moves []move
	for i, line := range lines {
		if line == "" {
			sa = parseStacksArrangement(lines[:i])
			moves = parseMoves(lines[i+1:])
			return sa, moves
		}
	}

	panic(fmt.Sprintf("Couldn't parse input %#v", lines))
}

type stacksArrangement map[int]stack.Stack[string]

func parseStacksArrangement(lines []string) stacksArrangement {
	sa := stacksArrangement{}

	// skip last line, we know stack number implicit from stack position
	for lineIndex := len(lines) - 2; lineIndex >= 0; lineIndex-- {
		line := []rune(lines[lineIndex])
		stackNumber := 1

		for i := 1; i < len(line); i += 4 {
			currentStack := line[i]
			if currentStack != ' ' {
				if sa[stackNumber] == nil {
					sa[stackNumber] = stack.New[string]()
				}
				sa[stackNumber].Put(string(line[i]))
			}
			stackNumber += 1
		}
	}

	return sa
}

type move struct {
	number int
	from   int
	to     int
}

func parseMoves(lines []string) []move {
	moves := make([]move, len(lines))

	for i, line := range lines {
		moves[i] = parseMove(line)
	}

	return moves
}

func parseMove(line string) move {
	split := strings.Split(line, " ")

	m := move{}
	previous := ""
	for _, a := range split {
		if previous == "move" {
			m.number = toInt(a)
		}
		if previous == "from" {
			m.from = toInt(a)
		}
		if previous == "to" {
			m.to = toInt(a)
		}
		previous = a
	}

	return m
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
