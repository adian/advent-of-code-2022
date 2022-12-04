package main

import (
	"testing"
)

const input = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func Test_solvePart1(t *testing.T) {
	got := solvePart1(getTestCaloriesLog())

	want := 24_000
	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}
}

func getTestCaloriesLog() []string {
	caloriesLog := splitLines([]byte(input))
	return caloriesLog
}

func Test_solvePart2(t *testing.T) {
	got := solvePart2(getTestCaloriesLog())

	want := 45_000
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}
