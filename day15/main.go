package main

import (
	"fmt"

	"adian.com/advent_of_code_2022/utils"
)

type position struct {
	x, y int
}

type pair struct {
	sensor position
	beacon position
}

func main() {
	data := utils.ReadFileLines("./day15/input.txt")
	pairs := parse(data)
	fmt.Printf("Part 1 answer: %v\n", solvePart1(pairs, 2_000_000))
	fmt.Printf("Part 2 answer: %v\n", solvePart2(pairs, 4_000_000))
}

// Checking every point is too expensive.
// If only one beacon is not found, it should be one position away from the field that is already checked from sensor.
func solvePart2(pairs []pair, size int) int {

	sizePredicate := func(xOrY int) bool {
		return 0 <= xOrY && xOrY <= size
	}

	//    1
	//   2#2
	//  3#S#3
	//   4#4
	//    5
	for _, pair := range pairs {
		distPlusOne := manhattanDistance(pair.beacon, pair.sensor) + 1

		for r := -distPlusOne; r <= distPlusOne; r++ {
			y := pair.sensor.y + r

			if y < 0 {
				continue
			}

			if y > size {
				break
			}

			if pair.sensor.x == 14 && y == 11 {
				println("hahah")
			}

			offset := distPlusOne - abs(r)
			left := position{
				y: y,
				x: pair.sensor.x - offset,
			}
			if sizePredicate(left.x) && !isReachable(pairs, left) {
				return calculateFrequency(left)
			}

			right := position{
				y: y,
				x: pair.sensor.x + offset,
			}
			if sizePredicate(right.x) && !isReachable(pairs, right) {
				return calculateFrequency(right)
			}
		}

	}

	panic("unreachable")
}

func isReachable(pairs []pair, p position) bool {
	for _, pair := range pairs {
		pairDist := manhattanDistance(pair.beacon, pair.sensor)

		if manhattanDistance(pair.sensor, p) <= pairDist {
			return true
		}

	}
	return false
}

func calculateFrequency(p position) int {
	return p.x*4_000_000 + p.y
}

func solvePart1(pairs []pair, y int) int {
	beaconPossibleX := make(map[int]any)

	for _, s := range pairs {
		dist := manhattanDistance(s.sensor, s.beacon)
		deltaToY := abs(y - s.sensor.y)

		if deltaToY <= dist {
			leftDist := dist - deltaToY
			for i := 0; i <= leftDist; i++ {
				rightX := s.sensor.x + i
				beaconPossibleX[rightX] = struct{}{}

				leftX := s.sensor.x - i
				beaconPossibleX[leftX] = struct{}{}
			}
		}
	}

	// remove actual positions
	for _, pair := range pairs {
		if pair.beacon.y == y {
			delete(beaconPossibleX, pair.beacon.x)
		}
	}

	return len(beaconPossibleX)
}

func manhattanDistance(a, b position) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
