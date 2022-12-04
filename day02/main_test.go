package main

import (
	"strings"
	"testing"
)

const input = `A Y
B X
C Z`

func getTestInput() []string {
	return strings.Split(input, "\n")
}

func Test_solvePart1(t *testing.T) {
	got := solvePart1(getTestInput())

	want := 15
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func Test_solvePart2(t *testing.T) {
	got := solvePart2(getTestInput())

	want := 12
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}
