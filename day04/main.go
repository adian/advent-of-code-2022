package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./day04/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	result1 := solvePart1(lines)
	fmt.Printf("Part 1 answer: %d\n", result1)

	result2 := solvePart2(lines)
	fmt.Printf("Part 2 answer: %d\n", result2)
}

func solvePart2(lines []string) int {
	return solve(lines, pairAssignmentOverlapAtAll)
}

func pairAssignmentOverlapAtAll(pa pairAssignment) bool {
	return assignmentsOverlapAtAll(pa.a1, pa.a2) || assignmentsOverlapAtAll(pa.a2, pa.a1)
}

func assignmentsOverlapAtAll(f assignment, s assignment) bool {
	return (f.from <= s.from && s.from <= f.to) || (f.from <= s.to && s.to <= f.to)
}

func solvePart1(lines []string) int {
	return solve(lines, oneContainsOther)
}

type overlapComparator func(pairAssignment) bool

func solve(lines []string, oc overlapComparator) int {
	result := 0

	for _, line := range lines {
		pa := toPariAssignments(line)

		if oc(pa) {
			result += 1
		}
	}

	return result
}

func oneContainsOther(pa pairAssignment) bool {
	return firstContainsSecond(pa.a1, pa.a2) || firstContainsSecond(pa.a2, pa.a1)
}

func firstContainsSecond(f assignment, s assignment) bool {
	return f.from <= s.from && s.to <= f.to
}

type pairAssignment struct {
	a1 assignment
	a2 assignment
}

type assignment struct {
	from int
	to   int
}

func toPariAssignments(raw string) pairAssignment {
	split := strings.Split(raw, ",")
	return pairAssignment{
		a1: toAssignment(split[0]),
		a2: toAssignment(split[1]),
	}
}

func toAssignment(raw string) assignment {
	split := strings.Split(raw, "-")

	return assignment{
		from: parseToInt(split[0]),
		to:   parseToInt(split[1]),
	}
}

func parseToInt(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return n
}
