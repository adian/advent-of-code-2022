package main

import (
	"adian.com/advent_of_code_2022/utils"
	"fmt"
	"math"
	"strconv"
)

func main() {
	lines := utils.ReadFileLines("./day25/input.txt")
	fmt.Printf("Part 1 answer: %v\n", solvePart1(lines))
}

func solvePart1(lines []string) string {
	var result int
	for _, line := range lines {
		result += snafuNumberToDecimal(line)
	}
	return toSnafu(result)
}

func snafuNumberToDecimal(input string) int {
	result := 0
	runes := []rune(input)
	for i, r := range runes {
		n := snafuDigitToDecimal(r)
		placePow := math.Pow(5, float64(len(runes)-i-1))
		if placePow != 0 {
			n *= int(placePow)
		}
		result += n
	}

	return result
}

func snafuDigitToDecimal(r rune) int {
	switch r {
	case '2':
		return 2
	case '1':
		return 1
	case '0':
		return 0
	case '-':
		return -1
	case '=':
		return -2
	default:
		msg := fmt.Sprintf("Incorrect SNAFU digit %v", r)
		panic(msg)
	}
}

func toSnafu(num int) string {
	fiveBaseNum := toBase5(num)

	toAdd := 0
	for i := len(fiveBaseNum) - 1; i >= 0; i-- {
		fiveBaseNum[i] += toAdd
		toAdd = 0

		if fiveBaseNum[i] >= 5 {
			toAdd += fiveBaseNum[i] / 5
			fiveBaseNum[i] = fiveBaseNum[i] % 5
		}

		if fiveBaseNum[i] >= 3 {
			toAdd += 1

			reminder := 5 - fiveBaseNum[i]
			fiveBaseNum[i] = 0 - reminder
		}
	}

	if toAdd != 0 {
		fiveBaseNum = append([]int{toAdd}, fiveBaseNum...)
	}

	result := ""
	for _, n := range fiveBaseNum {
		if n < 0 {
			if n == -1 {
				result += "-"
			} else if n == -2 {
				result += "="
			} else {
				msg := fmt.Sprintf("unexpected num %v", n)
				panic(msg)
			}
		} else {
			result += strconv.FormatInt(int64(n), 10)
		}
	}

	return result
}

func toBase5(num int) []int {
	formatInt := strconv.FormatInt(int64(num), 5)
	result := make([]int, len(formatInt))
	for i := range formatInt {
		a := formatInt[i] - '0'
		result[i] = int(a)
	}

	return result
}
