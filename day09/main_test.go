package main

import (
	"testing"
)

func Test_solvePart1(t *testing.T) {
	type args struct {
		moves []Move
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "with puzzle description input",
			args: args{
				moves: parseMoves([]string{
					"R 4",
					"U 4",
					"L 3",
					"D 1",
					"R 4",
					"D 1",
					"L 5",
					"R 2",
				}),
			},
			want: 13,
		},
		{
			name: "with actual input",
			args: args{
				moves: parseMoves(readInputFile("./input.txt")),
			},
			want: 5930,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solvePart1(tt.args.moves); got != tt.want {
				t.Errorf("solvePart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solvePart2(t *testing.T) {
	type args struct {
		moves []Move
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "with puzzle description input",
			args: args{
				moves: parseMoves([]string{
					"R 5",
					"U 8",
					"L 8",
					"D 3",
					"R 17",
					"D 10",
					"L 25",
					"U 20",
				}),
			},
			want: 36,
		},
		//{
		//	name: "with actual input",
		//	args: args{
		//		moves: parseMoves(readInputFile("./input.txt")),
		//	},
		//	want: 5930,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solvePart2(tt.args.moves); got != tt.want {
				t.Errorf("solvePart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
