package d02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func SolvePart2() string {
	file, err := os.Open("./d01/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	threeBiggest := [3]int{0, 0, 0}
	current := 0
	for scanner.Scan() {
		currentLine := scanner.Text()

		if currentLine == "" {
			replaceIfSmaller(&threeBiggest, current)
			current = 0
		} else {
			n := toNumber(scanner)
			current += n
		}
	}

	// also check last number
	replaceIfSmaller(&threeBiggest, current)

	threeBiggestSum := 0
	for i := range threeBiggest {
		threeBiggestSum += threeBiggest[i]
	}

	return fmt.Sprintf("The sum of the Calories carried by these three elves is %d", threeBiggestSum)
}

func toNumber(scanner *bufio.Scanner) int {
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func replaceIfSmaller(threeBiggest *[3]int, current int) {
	indexOfSmallest := 0
	for i := range threeBiggest {
		if threeBiggest[i] < threeBiggest[indexOfSmallest] {
			indexOfSmallest = i
		}
	}

	if threeBiggest[indexOfSmallest] < current {
		threeBiggest[indexOfSmallest] = current
	}
}
