package main

import (
	"adian.com/advent_of_code_2022/utils"
	"testing"
)

var inputFromDescription = []string{
	"Valve AA has flow rate=0; tunnels lead to valves DD, II, BB",
	"Valve BB has flow rate=13; tunnels lead to valves CC, AA",
	"Valve CC has flow rate=2; tunnels lead to valves DD, BB",
	"Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE",
	"Valve EE has flow rate=3; tunnels lead to valves FF, DD",
	"Valve FF has flow rate=0; tunnels lead to valves EE, GG",
	"Valve GG has flow rate=0; tunnels lead to valves FF, HH",
	"Valve HH has flow rate=22; tunnel leads to valve GG",
	"Valve II has flow rate=0; tunnels lead to valves AA, JJ",
	"Valve JJ has flow rate=21; tunnel leads to valve II",
}

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
				lines: inputFromDescription,
			},
			want: 1651,
		},
		{
			name: "actual input",
			args: args{
				lines: utils.ReadFileLines("./input.txt"),
			},
			want: 1940,
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
		lines             []string
		maxGenerationSize int
	}
	tests := []struct {
		name       string
		args       args
		want       int
		skip       bool
		skipReason string
	}{
		{
			name: "input from description",
			args: args{
				lines:             inputFromDescription,
				maxGenerationSize: 30_000,
			},
			want: 1707,
		},
		{
			name: "actual input",
			args: args{
				lines:             utils.ReadFileLines("./input.txt"),
				maxGenerationSize: 120_000,
			},
			want:       2469,
			skip:       true,
			skipReason: "This case is to slow to run it everytime.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skip {
				t.Skipf(tt.skipReason)
			}
			if got := solvePart2(tt.args.lines, tt.args.maxGenerationSize); got != tt.want {
				t.Errorf("solvePart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
