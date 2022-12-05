package main

import (
	"adian.com/advent_of_code_2022/common/stack"
	"reflect"
	"strings"
	"testing"
)

const input = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func Test_solvePart1(t *testing.T) {
	got := solvePart1(getTestAssignments())

	want := "CMZ"
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func getTestAssignments() []string {
	return toLines(input)
}

func toLines(str string) []string {
	return strings.Split(str, "\n")
}

func Test_parseStacksArrangement(t *testing.T) {
	p := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 `
	got := parseStacksArrangement(toLines(p))

	want := stacksArrangement{
		1: stack.NewWith("Z", "N"),
		2: stack.NewWith("M", "C", "D"),
		3: stack.NewWith("P"),
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test_parseMoves(t *testing.T) {
	p := `move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
	got := parseMoves(toLines(p))

	want := []move{
		{
			number: 1,
			from:   2,
			to:     1,
		},
		{
			number: 3,
			from:   1,
			to:     3,
		},
		{
			number: 2,
			from:   2,
			to:     1,
		},
		{
			number: 1,
			from:   1,
			to:     2,
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}

func Test_parseMove(t *testing.T) {
	tests := []struct {
		input string
		want  move
	}{
		{
			input: "move 1 from 2 to 1",
			want: move{
				number: 1,
				from:   2,
				to:     1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := parseMove(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseMove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solvePart2(t *testing.T) {
	got := solvePart2(getTestAssignments())

	want := "MCD"
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}
