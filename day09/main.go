package main

import (
	"fmt"
)

func main() {
	lines := readInputFile("./day09/input.txt")
	moves := parseMoves(lines)

	result1 := solvePart1(moves)
	fmt.Printf("Part 1 answer: %d\n", result1)

	result2 := solvePart2(moves)
	fmt.Printf("Part 2 answer: %d\n", result2)
}

func solvePart2(headMoves []Move) int {
	return solve(headMoves, 9)
}

func solvePart1(headMoves []Move) int {
	return solve(headMoves, 1)
}

func solve(headMoves []Move, numberOfTails int) int {
	var head Position
	tails := makeAndInitTails(numberOfTails)

	lastTailPositions := make(map[string]bool)
	// save starting position
	lastTailPositions[head.String()] = true

	for _, hMove := range headMoves {
		for i := 0; i < hMove.steps; i++ {
			head.move(hMove.direction)

			// move first tail after head
			lastTail := tails[0]
			moveToSecondIfNeeded(lastTail, head)

			// move the rest of tails after first one
			for _, currTail := range tails {
				moveToSecondIfNeeded(currTail, *lastTail)
				lastTail = currTail
			}
			lastTailPositions[lastTail.String()] = true
		}
	}

	return len(lastTailPositions)
}
