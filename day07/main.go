package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	raw, err := os.ReadFile("./day07/input.txt")
	if err != nil {
		panic(err)
	}

	input := splitLines(string(raw))

	dirTree := buildDirectoryTree(input)
	dirSizes := getDirSizes(dirTree)

	result1 := solvePart1(dirSizes)
	fmt.Printf("Part 1 answer: %d\n", result1)

	result2 := solvePart2(dirTree, dirSizes)
	fmt.Printf("Part 2 answer: %d\n", result2)
}

func solvePart1(dirSize []int) int {
	sum := 0
	for _, n := range dirSize {
		if n <= 100_000 {
			sum += n
		}
	}

	return sum
}

func solvePart2(dirTree Directory, dirSizes []int) int {
	unusedSpace := 70_000_000 - dirTree.getTotalSize()
	spaceToFree := 30_000_000 - unusedSpace

	result := math.MaxInt
	for _, size := range dirSizes {
		if size >= spaceToFree && size < result {
			result = size
		}
	}

	return result
}

type Directory struct {
	leafs     []*Directory
	parent    *Directory
	name      string
	filesSize int
}

func (n Directory) getTotalSize() int {
	var size int
	for _, n := range n.leafs {
		size += n.getTotalSize()
	}
	return size + n.filesSize

}

func getDirSizes(root Directory) []int {
	result := []int{root.getTotalSize()}

	for _, node := range root.leafs {
		result = append(result, getDirSizes(*node)...)
	}

	return result
}

func buildDirectoryTree(input []string) Directory {
	var (
		rootP    *Directory
		currentP *Directory
	)

	for _, line := range input {
		cmd := strings.Split(line, " ")

		isCommand := cmd[0] == "$"
		if isCommand {
			rootP, currentP = processCommand(cmd, currentP, rootP)
		} else {
			processLsOutput(cmd, currentP)
		}
	}

	return *rootP
}

func processCommand(cmd []string, currentP *Directory, rootP *Directory) (current, root *Directory) {
	if cmd[1] == "cd" {
		name := cmd[2]
		if name == ".." {
			currentP = currentP.parent
		} else {
			newCurrentP := &Directory{
				parent: currentP,
				name:   name,
			}

			if currentP != nil {
				currentP.leafs = append(currentP.leafs, newCurrentP)
			}

			currentP = newCurrentP
			if rootP == nil {
				rootP = currentP
			}
		}
	}

	return rootP, currentP
}

func processLsOutput(cmd []string, currentP *Directory) {
	isFile := cmd[0] != "dir"
	if isFile {
		currentP.filesSize += toInt(cmd[0])
	}
}
