package main

import (
	"testing"

	"adian.com/advent_of_code_2022/utils/asserts"
)

func Test_parse(t *testing.T) {
	got := parse(inputFromDescription[:2])

	want := []pair{
		{
			sensor: position{
				x: 2,
				y: 18,
			},
			beacon: position{
				x: -2,
				y: 15,
			},
		},
		{
			sensor: position{
				x: 9,
				y: 16,
			},
			beacon: position{
				x: 10,
				y: 16,
			},
		},
	}

	asserts.DeepEqual(t, got, want)
}
