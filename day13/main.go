package main

import (
	"adian.com/advent_of_code_2022/utils"
	"encoding/json"
	"fmt"
	"math"
	"sort"
	"strings"
)

func main() {
	data := utils.ReadFile("./day13/input.txt")
	fmt.Printf("Part 1 answer: %v\n", solvePart1(data))
	fmt.Printf("Part 2 answer: %v\n", solvePart2(data))
}

func solvePart2(data string) int {
	data = strings.ReplaceAll(data, "\n\n", "\n")
	lines := strings.Split(data, "\n")
	lines = append(lines, "[[2]]")
	lines = append(lines, "[[6]]")

	sort.Slice(lines, func(i, j int) bool {
		return packetsAreInRightOrder([]byte(lines[i]), []byte(lines[j]))
	})

	result := 1
	for i, line := range lines {
		if line == "[[2]]" || line == "[[6]]" {
			result *= i + 1
		}
	}

	return result
}

func solvePart1(lines string) int {
	pairs := strings.Split(lines, "\n\n")

	result := 0
	for i, pair := range pairs {
		split := strings.Split(pair, "\n")
		if packetsAreInRightOrder([]byte(split[0]), []byte(split[1])) {
			result += i + 1
		}
	}

	return result
}

func packetsAreInRightOrder(leftSignalJson, rightSignalJson []byte) bool {
	leftSignal := parseJson(leftSignalJson)
	rightSignal := parseJson(rightSignalJson)

	return compare(leftSignal, rightSignal) == rightOrder
}

type compareResult uint8

const (
	wrongOrder compareResult = 1 << iota
	rightOrder
	noConclusion
)

func compare(leftSignal, rightSignal []any) compareResult {
	for i := 0; i < minLength(leftSignal, rightSignal); i++ {
		left := leftSignal[i]
		right := rightSignal[i]

		var result compareResult
		switch {
		case isFloat(left) && isFloat(right):
			result = compareNumbers(toFloat(left), toFloat(right))

		case isSlice(left) && isFloat(right):
			result = compare(castToSlice(left), []any{right})

		case isFloat(left) && isSlice(right):
			result = compare([]any{left}, castToSlice(right))

		case isSlice(left) && isSlice(right):
			result = compare(castToSlice(left), castToSlice(right))

		default:
			panic("Branch without implementation!")
		}

		if result != noConclusion {
			return result
		}
	}

	if len(leftSignal) != len(rightSignal) {
		if len(leftSignal) < len(rightSignal) {
			return rightOrder
		}
		return wrongOrder
	}

	return noConclusion
}

func minLength[T any](a1, a2 []T) int {
	min := math.Min(
		float64(len(a1)),
		float64(len(a2)),
	)
	return int(min)
}

func castToSlice(left any) []any {
	return left.([]any)
}

func compareNumbers(fLeft, fRight float64) compareResult {
	if fLeft != fRight {
		if fLeft < fRight {
			return rightOrder
		} else {
			return wrongOrder
		}
	}
	return noConclusion
}

func toFloat(v any) float64 {
	return v.(float64)
}

func parseJson(data []byte) []any {
	var result []any
	err := json.Unmarshal(data, &result)
	if err != nil {
		panic(err)
	}
	return result
}

func isSlice(data any) bool {
	switch data.(type) {
	case []any:
		return true
	default:
		return false
	}
}

func isFloat(data any) bool {
	switch data.(type) {
	case float64:
		return true
	default:
		return false
	}
}
