package main

import (
	"fmt"
	"os"
	"strings"
)

type Item rune
type Rucksack []Item

func main() {
	input, err := os.ReadFile("./day03/input.txt")
	if err != nil {
		panic(err)
	}

	rucksacks := toRucksacks(input)

	result1 := solvePart1(rucksacks)
	fmt.Printf("Part 1 answer: %d\n", result1)

	result2 := solvePart2(rucksacks)
	fmt.Printf("Part 2 answer: %d\n", result2)

}

func toRucksacks(input []byte) []Rucksack {
	lines := strings.Split(string(input), "\n")
	var rucksacks []Rucksack
	for _, line := range lines {
		rucksacks = append(rucksacks, Rucksack(line))
	}
	return rucksacks
}

func solvePart2(rucksacks []Rucksack) uint {
	var result uint = 0
	var group []Rucksack
	for _, rucksack := range rucksacks {
		group = append(group, rucksack)

		if len(group) == 3 {
			var repeated []Item
			for _, r := range group {
				if repeated == nil {
					repeated = r
				} else {
					repeated = getRepeatedItems(repeated, r)
				}
			}

			for _, item := range repeated {
				result += toPriority(item)
			}

			group = nil
		}
	}

	return result
}

func solvePart1(rucksacks []Rucksack) uint {
	var result uint = 0
	for _, rucksack := range rucksacks {
		repeatedItems := getItemsInBothCompartments(rucksack)

		for _, item := range repeatedItems {
			result += toPriority(item)
		}
	}
	return result
}

func getItemsInBothCompartments(rucksack Rucksack) []Item {
	half := len(rucksack) / 2
	repeatedItems := getRepeatedItems(rucksack[:half], rucksack[half:])
	return repeatedItems
}

func getRepeatedItems(i1 []Item, i2 []Item) []Item {
	r1map := make(map[Item]bool, len(i1))
	for _, item := range i1 {
		r1map[item] = true
	}

	repeated := make(map[Item]bool)
	for _, item := range i2 {
		if r1map[item] {
			repeated[item] = true
		}
	}

	result := make([]Item, 0, len(repeated))
	for item := range repeated {
		result = append(result, item)
	}

	return result
}

func toPriority(item Item) uint {
	if item >= 'a' {
		return uint(item - 'a' + 1)
	} else {
		return uint(item - 'A' + 27)
	}
}
