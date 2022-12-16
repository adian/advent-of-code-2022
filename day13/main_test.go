package main

import (
	"adian.com/advent_of_code_2022/utils"
	"testing"
)

func Test_packetsAreInRightOrder(t *testing.T) {
	type args struct {
		leftJson  string
		rightJson string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "only integers -> right order",
			args: args{
				leftJson:  "[1,1,3,1,1]",
				rightJson: "[1,1,5,1,1]",
			},
			want: true,
		},
		{
			name: "only integers -> wrong order",
			args: args{
				leftJson:  "[1,1,6,1,1]",
				rightJson: "[1,1,5,1,1]",
			},
			want: false,
		},
		{
			name: "mixed types, converting one to list -> right order",
			args: args{
				leftJson:  "[[1],[2,3,4]]",
				rightJson: "[[1],4]",
			},
			want: true,
		},
		{
			name: "mixed types, converting one to list -> right order",
			args: args{
				leftJson:  "[[1],[5,3,4]]",
				rightJson: "[[1],4]",
			},
			want: false,
		},
		{
			name: "left run out of items -> right order",
			args: args{
				leftJson:  "[[4,4],4,4]",
				rightJson: "[[4,4],4,4,4]",
			},
			want: true,
		},
		{
			name: "right run out of items -> wrong order",
			args: args{
				leftJson:  "[7,7,7,7]",
				rightJson: "[7,7,7]",
			},
			want: false,
		},
		{
			name: "only lists, right run out of items -> wrong order",
			args: args{
				leftJson:  "[[[]]]",
				rightJson: "[[]]",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := packetsAreInRightOrder(
				[]byte(tt.args.leftJson),
				[]byte(tt.args.rightJson),
			)
			if got != tt.want {
				t.Errorf("packetsAreInRightOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solvePart1(t *testing.T) {
	type args struct {
		lines string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "with test input",
			args: args{
				lines: utils.ReadFile("./input_test.txt"),
			},
			want: 13,
		},
		{
			name: "with actual input",
			args: args{
				lines: utils.ReadFile("./input.txt"),
			},
			want: 5503,
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
		data string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "with test input",
			args: args{
				data: utils.ReadFile("./input_test.txt"),
			},
			want: 140,
		},
		{
			name: "with actual input",
			args: args{
				data: utils.ReadFile("./input.txt"),
			},
			want: 20952,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solvePart2(tt.args.data); got != tt.want {
				t.Errorf("solvePart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
