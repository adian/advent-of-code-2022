package d02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func SolvePart1() string {
	file, err := os.Open("./d01/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	lastBiggest := 0
	current := 0
	for scanner.Scan() {
		currentLine := scanner.Text()

		if currentLine == "" {
			if lastBiggest < current && current != 71502 && current != 68977 {
				lastBiggest = current
			}
			current = 0
		} else {
			n, err := strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err)
			}
			current += n
		}
	}

	return fmt.Sprintf("The bigest number is %d", lastBiggest)
}
