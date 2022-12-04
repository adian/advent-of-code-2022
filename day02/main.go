package main

import (
	"fmt"
	"os"
	"strings"
)

// https://adventofcode.com/2022/day/2
func main() {
	input, err := os.ReadFile("./day02/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	part1Answer := solvePart1(lines)
	fmt.Printf("Part one answer: %d\n", part1Answer)

	part2Answer := solvePart2(lines)
	fmt.Printf("Part two answer: %d\n", part2Answer)
}

type Hand uint8

const (
	Rock Hand = iota
	Paper
	Scissors
)

var handToPoints = map[Hand]int{
	Rock:     1,
	Paper:    2,
	Scissors: 3,
}

func solvePart1(lines []string) int {
	sum := 0
	for _, line := range lines {
		opponentHand, myHand := toHands(line)
		pointsForHand := handToPoints[myHand]
		pointsForOutcome := calculatePointsForOutcome(opponentHand, myHand)

		sum += pointsForHand + pointsForOutcome
	}

	return sum
}

func calculatePointsForOutcome(opponentHand Hand, myHand Hand) int {
	if myHand == opponentHand {
		return 3
	}

	if (myHand == Rock && opponentHand == Scissors) ||
		(myHand == Paper && opponentHand == Rock) ||
		(myHand == Scissors && opponentHand == Paper) {
		return 6
	}

	return 0
}

func toHands(text string) (Hand, Hand) {
	split := strings.Split(text, " ")
	return mapToHandConst(split[0]), mapToHandConst(split[1])
}

func mapToHandConst(v string) Hand {
	switch v {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissors
	}

	panic(fmt.Sprintf("Invalid input: %s!", v))
}

func solvePart2(lines []string) int {

	sum := 0
	for _, line := range lines {
		opponentHand, myHand := toHands2(line)
		pointsForHand := handToPoints[myHand]
		pointsForOutcome := calculatePointsForOutcome(opponentHand, myHand)

		sum += pointsForHand + pointsForOutcome
	}

	return sum
}

var strongerToWeak = map[Hand]Hand{
	Rock:     Scissors,
	Scissors: Paper,
	Paper:    Rock,
}

func toHands2(text string) (Hand, Hand) {
	split := strings.Split(text, " ")

	opponentHand := mapToHandConst(split[0])
	var myHand Hand

	// I need to lose
	if split[1] == "X" {
		myHand = strongerToWeak[opponentHand]
	}

	// I need to win
	if split[1] == "Z" {
		myHand = strongerToWeak[strongerToWeak[opponentHand]]
	}

	if split[1] == "Y" {
		myHand = opponentHand
	}

	return opponentHand, myHand
}
