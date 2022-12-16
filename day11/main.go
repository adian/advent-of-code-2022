package main

import (
	"adian.com/advent_of_code_2022/day11/monkey"
	"adian.com/advent_of_code_2022/utils"
	"fmt"
	"sort"
)

func main() {

	lines := utils.ReadFileLines("./day11/input.txt")

	fmt.Printf("Part 1 answer: %v\n", solvePart1(lines))
	fmt.Printf("Part 2 answer: \n%v", solvePart2(lines))
}

func solvePart2(lines []string) int {
	monkeys := monkey.Parse(lines)

	lcm := calculateLeastCommonMultiplier(monkeys)
	for _, m := range monkeys {
		m.WorryLevelModifier = func(worryLevel int) int {
			return worryLevel % lcm
		}
	}

	numberOfRound := 10_000
	for i := 0; i < numberOfRound; i++ {
		runRound(monkeys)
	}

	return calculateMonkeyBusiness(monkeys)
}

func calculateLeastCommonMultiplier(monkeys monkey.Monkeys) int {
	lcm := 1
	for _, m := range monkeys {
		lcm *= m.ThrowTest.DivisibleBy
	}
	return lcm
}

func solvePart1(lines []string) int {
	monkeys := monkey.Parse(lines)
	for _, m := range monkeys {
		m.WorryLevelModifier = func(worryLevel int) int {
			return worryLevel / 3
		}
	}

	numberOfRound := 20
	for i := 0; i < numberOfRound; i++ {
		runRound(monkeys)
	}

	return calculateMonkeyBusiness(monkeys)
}

func calculateMonkeyBusiness(monkeys monkey.Monkeys) int {
	var inspections []int
	for _, m := range monkeys {
		inspections = append(inspections, m.NumberOfInspections)
	}
	sortReverse(inspections)
	return inspections[0] * inspections[1]
}

func sortReverse(inspections []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))
}

func runRound(monkeys monkey.Monkeys) {
	for i := 0; i < len(monkeys); i++ {
		currMonkey := monkeys[monkey.Id(i)]
		for len(currMonkey.Items) > 0 {
			to, what, ok := currMonkey.ThrowFirstItem()
			if ok {
				monkeys[to].Items = append(monkeys[to].Items, what)
			}
		}
	}
}
