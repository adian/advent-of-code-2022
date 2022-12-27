package main

import (
	"adian.com/advent_of_code_2022/utils"
	"strings"
	"testing"
)

var heightmapFromDescription = toHeightmap(splitLines(`Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`))

var heightmap = toHeightmap(utils.ReadFileLines("./input.txt"))

func splitLines(s string) []string {
	return strings.Split(s, "\n")
}

func Test_solvePart1(t *testing.T) {
	type args struct {
		heightmap [][]rune
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "input from description",
			args: args{
				heightmap: heightmapFromDescription,
			},
			want: 31,
		},
		{
			name: "actual input",
			args: args{
				heightmap: heightmap,
			},
			want: 330,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solvePart1(tt.args.heightmap); got != tt.want {
				t.Errorf("solvePart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solvePart2(t *testing.T) {
	type args struct {
		heightmap [][]rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "input from description",
			args: args{
				heightmap: heightmapFromDescription,
			},
			want: 29,
		},
		{
			name: "actual input",
			args: args{
				heightmap: heightmap,
			},
			want: 321,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solvePart2(tt.args.heightmap); got != tt.want {
				t.Errorf("solvePart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
