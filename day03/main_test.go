package main

import (
	"fmt"
	"testing"
)

func Test_getItemTypeInBothCompartment(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  Item
	}{
		{
			name:  "first rucksack",
			input: "vJrwpWtwJgWrhcsFMMfFFhFp",
			want:  'p',
		},
		{
			name:  "second rucksack",
			input: "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			want:  'L',
		},
		{
			name:  "third rucksack",
			input: "PmmdzqPrVvPwwTWBwg",
			want:  'P',
		},
		{
			name:  "fourth rucksack",
			input: "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
			want:  'v',
		},
		{
			name:  "fifth rucksack",
			input: "ttgJtRGJQctTZtZT",
			want:  't',
		},
		{
			name:  "sixth rucksack",
			input: "CrZsJsPPZsGzwwsLwLmpwMDw",
			want:  's',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rucksack := Rucksack(tt.input)

			got := getItemsInBothCompartments(rucksack)

			if len(got) != 1 {
				t.Errorf("want length 1 but is %d", len(got))
			}

			if !contains(got, tt.want) {
				t.Errorf("wantet Item %c but non found in %v", tt.want, got)
			}
		})
	}
}

func contains[T comparable](slice []T, item T) bool {
	for _, i := range slice {
		if i == item {
			return true
		}
	}

	return false
}

func Test_countPriority(t *testing.T) {
	tests := []struct {
		item Item
		want uint
	}{
		{
			item: 'p',
			want: 16,
		},
		{
			item: 'L',
			want: 38,
		},
		{
			item: 'P',
			want: 42,
		},
		{
			item: 'v',
			want: 22,
		},
		{
			item: 't',
			want: 20,
		},
		{
			item: 's',
			want: 19,
		},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%c has prority %d", tt.item, tt.want)
		t.Run(name, func(t *testing.T) {
			if got := toPriority(tt.item); got != tt.want {
				t.Errorf("countPriority() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solvePart1(t *testing.T) {
	rucksacks := getTestRucksacks()
	result := solvePart1(rucksacks)

	if result != 157 {
		t.Errorf("wanted 157, got %v", result)
	}
}

const testInput = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

func getTestRucksacks() []Rucksack {
	return toRucksacks([]byte(testInput))
}

func Test_solvePart2(t *testing.T) {
	result := solvePart2(getTestRucksacks())

	if result != 70 {
		t.Errorf("wanted 70, got %v", result)
	}
}
