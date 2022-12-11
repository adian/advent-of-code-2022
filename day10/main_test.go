package main

import (
	"reflect"
	"strings"
	"testing"
)

var simpleTestInstructions = &Instruction{
	cycles: 1,
	value:  0,
	next: &Instruction{
		cycles: 2,
		value:  3,
		next: &Instruction{
			cycles: 2,
			value:  -5,
			next:   nil,
		},
	},
}

func Test_buildInstructions(t *testing.T) {
	type args struct {
		lines []string
	}

	tests := []struct {
		name string
		args args
		want *Instruction
	}{
		{
			name: "simple",
			args: args{
				lines: []string{
					"noop",
					"addx 3",
					"addx -5",
				},
			},
			want: simpleTestInstructions,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildInstructions(tt.args.lines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildInstructions() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func Test_buildRegisterLog(t *testing.T) {
	type args struct {
		firstInstruction *Instruction
	}
	tests := []struct {
		name string
		args args
		want RegisterLog
	}{
		{
			name: "simple",
			args: args{
				firstInstruction: simpleTestInstructions,
			},
			want: RegisterLog{
				1,
				1,
				1,
				4,
				4,
				-1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildRegisterLog(tt.args.firstInstruction); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildRegisterLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solvePart1(t *testing.T) {
	type args struct {
		registerLog RegisterLog
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "input from puzzle description",
			args: args{
				registerLog: buildRegisterLogFromFile("./test-input.txt"),
			},
			want: 13140,
		},
		{
			name: "actual puzzle input",
			args: args{
				registerLog: buildRegisterLogFromFile("./input.txt"),
			},
			want: 16880,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solvePart1(tt.args.registerLog); got != tt.want {
				t.Errorf("solvePart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func buildRegisterLogFromFile(name string) RegisterLog {
	return buildRegisterLog(buildInstructions(readFileLines(name)))
}

func Test_solveParty2(t *testing.T) {
	type args struct {
		registerLog RegisterLog
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "from puzzle description",
			args: args{
				registerLog: buildRegisterLogFromFile("./test-input.txt"),
			},
			want: `
##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....`,
		},
		{
			name: "with real input",
			args: args{
				registerLog: buildRegisterLogFromFile("./input.txt"),
			},
			want: `
###..#..#..##..####..##....##.###..###..
#..#.#.#..#..#....#.#..#....#.#..#.#..#.
#..#.##...#..#...#..#..#....#.###..#..#.
###..#.#..####..#...####....#.#..#.###..
#.#..#.#..#..#.#....#..#.#..#.#..#.#.#..
#..#.#..#.#..#.####.#..#..##..###..#..#.`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solvePart2(tt.args.registerLog)

			if trimNewLines(got) != trimNewLines(tt.want) {
				t.Errorf("solvePart2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func trimNewLines(got string) string {
	return strings.Trim(got, "\n")
}
