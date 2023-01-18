package main

import (
	"strings"
	"testing"

	"adian.com/advent_of_code_2022/utils"
	"adian.com/advent_of_code_2022/utils/asserts"
)

var inputFromDescription = strings.Split(
	`Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3`,
	"\n",
)

var exampleSensorsFromDescription = parse(inputFromDescription)

var inputSensors = parse(utils.ReadFileLines("input.txt"))

func Test_solvePart1(t *testing.T) {
	type args struct {
		sensors []pair
		y       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "input from description",
			args: args{
				sensors: exampleSensorsFromDescription,
				y:       10,
			},
			want: 26,
		},
		{
			name: "actual input",
			args: args{
				sensors: inputSensors,
				y:       2_000_000,
			},
			want: 5256611,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solvePart1(tt.args.sensors, tt.args.y); got != tt.want {
				t.Errorf("solvePart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_manhattanDistance(t *testing.T) {
	got := manhattanDistance(
		position{8, 7},
		position{2, 10},
	)

	want := 9

	asserts.Equal(t, got, want)
}

func Test_solvePart2(t *testing.T) {
	type args struct {
		sensors []pair
		size    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example from the description",
			args: args{
				sensors: exampleSensorsFromDescription,
				size:    20,
			},
			want: 56000011,
		},
		{
			name: "actual input",
			args: args{
				sensors: inputSensors,
				size:    4_000_000,
			},
			want: 13337919186981,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solvePart2(tt.args.sensors, tt.args.size); got != tt.want {
				t.Errorf("solvePart2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isReachable(t *testing.T) {
	pairs := []pair{
		{
			sensor: position{8, 7},
			beacon: position{2, 10},
		},
	}

	type args struct {
		pairs []pair
		p     position
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				pairs: pairs,
				p:     position{14, 3},
			},
			want: false,
		},
		{
			args: args{
				pairs: pairs,
				p:     position{13, 3},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isReachable(tt.args.pairs, tt.args.p); got != tt.want {
				t.Errorf("isReachable() = %v, want %v", got, tt.want)
			}
		})
	}
}
