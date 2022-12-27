package main

import (
	"adian.com/advent_of_code_2022/utils"
	"fmt"
	"testing"
)

var snafuDecimalDict = []struct {
	snafu   string
	decimal int
}{
	{snafu: "0", decimal: 0},
	{snafu: "1", decimal: 1},
	{snafu: "2", decimal: 2},
	{snafu: "1=", decimal: 3},
	{snafu: "1-", decimal: 4},
	{snafu: "10", decimal: 5},
	{snafu: "11", decimal: 6},
	{snafu: "12", decimal: 7},
	{snafu: "2=", decimal: 8},
	{snafu: "2-", decimal: 9},
	{snafu: "20", decimal: 10},
	{snafu: "1=0", decimal: 15},
	{snafu: "1-0", decimal: 20},
	{snafu: "1-0", decimal: 20},
	{snafu: "2=01", decimal: 201},
	{snafu: "1-0---0", decimal: 12345},
	{snafu: "1121-1110-1=0", decimal: 314159265},
}

func Test_snafuToDecimal(t *testing.T) {
	for _, tt := range snafuDecimalDict {
		name := fmt.Sprintf("in SNAFU %v is equal to %v", tt.snafu, tt.decimal)
		t.Run(name, func(t *testing.T) {
			if got := snafuNumberToDecimal(tt.snafu); got != tt.decimal {
				t.Errorf("snafuNumberToDecimal() = %v, want %v", got, tt.decimal)
			}
		})
	}
}

func Test_toSnafu(t *testing.T) {
	for _, tt := range snafuDecimalDict {
		name := fmt.Sprintf("%v is equal to %v in SNAFU", tt.decimal, tt.snafu)
		t.Run(name, func(t *testing.T) {
			if got := toSnafu(tt.decimal); got != tt.snafu {
				t.Errorf("toSnafu() = %v, want %v", got, tt.snafu)
			}
		})
	}
}

func Test_solvePart1(t *testing.T) {
	lines := utils.ReadFileLines("input.txt")
	got := solvePart1(lines)

	want := "20=022=21--=2--12=-2"
	if got != want {
		t.Errorf("got = %v, want = %v", got, want)
	}
}
