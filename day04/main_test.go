package main

import (
	"reflect"
	"strings"
	"testing"
)

const input = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

func Test_solvePart1(t *testing.T) {
	got := solvePart1(getTestAssignments())

	want := 2
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func getTestAssignments() []string {
	return strings.Split(input, "\n")
}

func Test_toPariAssignments(t *testing.T) {
	tests := []struct {
		input string
		want  pairAssignment
		want1 assignment
	}{
		{
			input: "2-4,6-8",
			want: pairAssignment{
				a1: assignment{
					from: 2,
					to:   4,
				},
				a2: assignment{
					from: 6,
					to:   8,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := toPariAssignments(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toPariAssignments() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solvePart2(t *testing.T) {
	got := solvePart2(getTestAssignments())

	want := 4
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}
