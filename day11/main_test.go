package main

import (
	"adian.com/advent_of_code_2022/day11/util"
	"testing"
)

var inputFromPuzzleDescription = util.ReadFileLines("input_test.txt")
var actualInput = util.ReadFileLines("input.txt")

func Test_solvePart1(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "input from description",
			args: args{
				lines: inputFromPuzzleDescription,
			},
			want: 10605,
		},
		{
			name: "actual input",
			args: args{
				lines: actualInput,
			},
			want: 102399,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solvePart1(tt.args.lines); got != tt.want {
				t.Errorf("solvePart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solvePart2(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "input from description",
			args: args{
				lines: inputFromPuzzleDescription,
			},
			want: 2713310158,
		},
		{
			name: "actual input",
			args: args{
				lines: actualInput,
			},
			want: 23641658401,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solvePart2(tt.args.lines); got != tt.want {
				t.Errorf("solvePart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
