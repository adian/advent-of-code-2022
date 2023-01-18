package main

import (
	"adian.com/advent_of_code_2022/utils/asserts"
	"reflect"
	"testing"
)

func Test_parseEdges(t *testing.T) {
	input := []string{"503,4 -> 502,4 -> 502,9 -> 494,9"}

	got := parseEdges(input)

	want := [][]point{
		{
			point{
				x: 503,
				y: 4,
			},
			point{
				x: 502,
				y: 4,
			},
			point{
				x: 502,
				y: 9,
			},
			point{
				x: 494,
				y: 9,
			},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}

}

func Test_parse(t *testing.T) {
	got := parse(inputFromDescription)

	asserts.Equal(t, len(got), 20)

	want := map[point]bool{
		point{494, 9}: true,
		point{495, 9}: true,
		point{496, 6}: true,
		point{496, 9}: true,
		point{497, 6}: true,
		point{497, 9}: true,
		point{498, 4}: true,
		point{498, 5}: true,
		point{498, 6}: true,
		point{498, 9}: true,
		point{499, 9}: true,
		point{500, 9}: true,
		point{501, 9}: true,
		point{502, 4}: true,
		point{502, 5}: true,
		point{502, 6}: true,
		point{502, 7}: true,
		point{502, 8}: true,
		point{502, 9}: true,
		point{503, 4}: true,
	}

	for rock := range got {
		if want[rock] == false {
			t.Errorf("unexpected %#v", rock)
		}
	}
}
