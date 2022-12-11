package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readFileLines("./day10/input.txt")
	instructions := buildInstructions(lines)
	registerLog := buildRegisterLog(instructions)

	fmt.Printf("Part 1 answer: %v\n", solvePart1(registerLog))
	fmt.Printf("Part 2 answer: \n%v", solvePart2(registerLog))
}

const LineLength = 40

func solvePart2(registerLog RegisterLog) string {
	result := ""
	currentLine := ""

	for i, startOfSprite := range registerLog {
		currPixelIndex := i%LineLength + 1

		endOfSprite := startOfSprite + 2
		if startOfSprite <= currPixelIndex && currPixelIndex <= endOfSprite {
			currentLine += "#"
		} else {
			currentLine += "."
		}

		if len(currentLine)%LineLength == 0 {
			result += currentLine
			result += "\n"
			currentLine = ""
		}
	}

	return strings.TrimSuffix(result, "\n")
}

func solvePart1(log RegisterLog) int {
	result := 0

	for _, cycle := range []int{20, 60, 100, 140, 180, 220} {
		v := cycle * log[cycle-1]
		result += v
	}

	return result
}

func readFileLines(name string) []string {
	raw, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(raw), "\n")
}

type RegisterLog []int

func buildRegisterLog(firstInstruction *Instruction) RegisterLog {
	var registerLog RegisterLog
	currInstruction := firstInstruction
	lastElement := 1

	for currInstruction != nil {
		for cycle := 0; cycle < currInstruction.cycles; cycle++ {
			registerLog = append(registerLog, lastElement)
		}

		if currInstruction.value != 0 {
			lastElement += currInstruction.value
		}
		currInstruction = currInstruction.next
	}

	return append(registerLog, lastElement)
}

type Instruction struct {
	next   *Instruction
	cycles int
	value  int
}

func buildInstructions(lines []string) *Instruction {
	var head, tail *Instruction

	for _, line := range lines {
		curr := Instruction{}

		if line == "noop" {
			curr.cycles = 1
		} else {
			curr.cycles = 2
			curr.value = toInt(strings.Split(line, " ")[1])
		}

		if tail != nil {
			tail.next = &curr
		}
		tail = &curr

		if head == nil {
			head = tail
		}
	}

	return head
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
