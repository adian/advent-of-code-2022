package main

import (
	"adian.com/advent_of_code_2022/utils"
	"fmt"
)

func main() {
	lines := utils.ReadFileLines("./day12/input.txt")
	heightmap := toHeightmap(lines)
	fmt.Printf("Part 1 answer: %v\n", solvePart1(heightmap))
	fmt.Printf("Part 1 answer: %v\n", solvePart2(heightmap))
}

func solvePart2(heightmap [][]rune) int {
	startPositions := findAllWithHeight(heightmap, 'S', 'a')
	return findQuickestRoute(heightmap, startPositions)
}

type position struct {
	x int
	y int
}

func solvePart1(heightmap [][]rune) int {
	startPositions := findAllWithHeight(heightmap, 'S')
	return findQuickestRoute(heightmap, startPositions)
}

func findQuickestRoute(heightmap [][]rune, startPositions []position) int {
	possiblePositions := startPositions
	visited := make(map[position]bool)
	for _, p := range startPositions {
		visited[p] = true
	}

	steps := 0
	for {
		steps += 1
		var temp []position

		for _, pos := range possiblePositions {
			tryToClimb := func(p position) (isEnd bool) {
				heightDifference := calculateHeightDiff(heightmap, pos, p)
				if exist := visited[p]; !exist && heightDifference <= 1 {
					if heightmap[p.y][p.x] == 'E' {
						return true
					}

					visited[p] = true
					temp = append(temp, p)
				}
				return false
			}

			// try to go up
			if pos.y > 0 {
				np := position{pos.x, pos.y - 1}
				if tryToClimb(np) {
					return steps
				}
			}

			// try to go down
			if pos.y+1 < len(heightmap) {
				np := position{pos.x, pos.y + 1}
				if tryToClimb(np) {
					return steps
				}

			}

			// try to go left
			if pos.x > 0 {
				np := position{pos.x - 1, pos.y}
				if tryToClimb(np) {
					return steps
				}
			}

			// try to go right
			if pos.x+1 < len(heightmap[pos.y]) {
				np := position{pos.x + 1, pos.y}
				if tryToClimb(np) {
					return steps
				}
			}

		}
		possiblePositions = temp
		temp = []position{}
	}
}

func calculateHeightDiff(heightmap [][]rune, curr, next position) rune {
	currHeight := heightmap[curr.y][curr.x]
	nextHeight := heightmap[next.y][next.x]

	if currHeight == 'S' {
		currHeight = 'a'
	}
	if nextHeight == 'S' {
		currHeight = 'a'
	}

	if nextHeight == 'E' {
		nextHeight = 'z'
	}
	if nextHeight == 'E' {
		nextHeight = 'z'
	}

	return nextHeight - currHeight
}

func findAllWithHeight(heightmap [][]rune, heights ...rune) []position {
	var res []position

	for y := range heightmap {
		for x := range heightmap[y] {
			curr := heightmap[y][x]
			if contains(heights, curr) {
				res = append(res, position{x, y})
			}
		}
	}

	return res
}

func contains(runes []rune, r rune) bool {
	for _, curr := range runes {
		if curr == r {
			return true
		}
	}
	return false
}

func toHeightmap(lines []string) [][]rune {
	res := make([][]rune, len(lines))
	for i, line := range lines {
		res[i] = []rune(line)
	}
	return res
}
