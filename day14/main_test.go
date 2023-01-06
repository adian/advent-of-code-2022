package main

import (
	"adian.com/advent_of_code_2022/utils"
	"strings"
	"testing"
)

var inputFromDescription = strings.Split(
	`498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`,
	"\n",
)

func Test_solvePart1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name:  "for input from description",
			input: inputFromDescription,
			want:  24,
		},
		{
			name:  "actual input",
			input: utils.ReadFileLines("./input.txt"),
			want:  774,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solvePart1(tt.input); got != tt.want {
				t.Errorf("solvePart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solvePart2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name:  "for input from description",
			input: inputFromDescription,
			want:  93,
		},
		{
			name:  "actual input",
			input: utils.ReadFileLines("./input.txt"),
			want:  22499,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solvePart2(tt.input); got != tt.want {
				t.Errorf("solvePart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
