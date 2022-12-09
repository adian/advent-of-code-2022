package main

import (
	"testing"
)

var inputFromPuzzleDescription = func() TreeMap {
	const input = `30373
25512
65332
33549
35390`
	return generateTreeMap([]byte(input))
}()

func Test_solvePart1(t *testing.T) {
	type args struct {
		treeMap TreeMap
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "with description input",
			args: args{
				treeMap: inputFromPuzzleDescription,
			},
			want: 21,
		},
		{
			name: "with real input",
			args: args{
				treeMap: generateTreeMap(readInputFileFromTest()),
			},
			want: 1546,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solvePart1(tt.args.treeMap); got != tt.want {
				t.Errorf("solvePart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func readInputFileFromTest() []byte {
	return readInputFile("./input.txt")
}

func Test_solvePart2(t *testing.T) {
	type args struct {
		treeMap TreeMap
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "with description input",
			args: args{
				treeMap: inputFromPuzzleDescription,
			},
			want: 8,
		},
		{
			name: "with real input",
			args: args{
				treeMap: generateTreeMap(readInputFileFromTest()),
			},
			want: 519064,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solvePart2(tt.args.treeMap); got != tt.want {
				t.Errorf("solvePart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
