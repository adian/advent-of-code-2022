package main

import (
	"adian.com/advent_of_code_2022/utils"
	"fmt"
)

func main() {
	data := utils.ReadFileLines("./day14/input.txt")
	fmt.Printf("Part 1 answer: %v\n", solvePart1(data))
	fmt.Printf("Part 2 answer: %v\n", solvePart2(data))
}

func solvePart1(data []string) int {
	obstacles := parse(data)
	return solve(obstacles, withoutFloor)
}

func solvePart2(data []string) int {
	obstacles := parse(data)
	return solve(obstacles, withFloor)
}

type solveMode string

const (
	withFloor    solveMode = "withFloor"
	withoutFloor solveMode = "withoutFloor"
)

type point struct {
	x, y int
}

func solve(obstacles map[point]bool, mode solveMode) int {
	counter := 0
	lowestObstacleY := getLowestObstacleY(obstacles)
	if mode == withFloor {
		lowestObstacleY += 2
	}
	sandStart := point(createSandStartPoint())

	for {
		counter += 1
		sand, comeToRest := produceSand(obstacles, lowestObstacleY, withFloor == mode)
		obstacles[sand] = true

		if comeToRest == false {
			return counter - 1
		}
		if obstacles[sandStart] {
			return counter
		}
	}
}

func getLowestObstacleY(obstacles map[point]bool) int {
	var maxInt int
	for obstacle := range obstacles {
		if maxInt < obstacle.y {
			maxInt = obstacle.y
		}
	}
	return maxInt
}

func produceSand(obstacles map[point]bool, rockLowestLevel int, withFloor bool) (point, bool) {
	s := createSandStartPoint()

	for {
		switch {
		case s.y >= rockLowestLevel:
			// this sand will never come to rest
			return point{}, false
		case withFloor && isOnTheFloor(s, rockLowestLevel):
			return point(s), true
		case s.canGoDown(obstacles):
			s = s.goDown()
		case s.canGoDownLeft(obstacles):
			s = s.goDownLeft()
		case s.canGoDownRight(obstacles):
			s = s.goDownRight()
		default:
			return point(s), true
		}
	}
}

func isOnTheFloor(s sand, rockLowestLevel int) bool {
	return s.y+1 == rockLowestLevel
}
