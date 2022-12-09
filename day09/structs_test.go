package main

import (
	"reflect"
	"testing"
)

func Test_parseMoves(t *testing.T) {
	got := parseMoves([]string{
		"U 4",
		"D 3",
		"L 2",
		"R 1",
	})

	want := []Move{
		{
			direction: UP,
			steps:     4,
		},
		{
			direction: DOWN,
			steps:     3,
		},
		{
			direction: LEFT,
			steps:     2,
		},
		{
			direction: RIGHT,
			steps:     1,
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPosition_moveCloserTo(t *testing.T) {
	tests := []struct {
		name    string
		initial Position
		args    Position
		want    Position
	}{
		{
			name:    "Move right",
			initial: Position{1, 1},
			args:    Position{3, 1},
			want:    Position{2, 1},
		},
		{
			name:    "Move diagonally",
			initial: Position{1, 1},
			args:    Position{3, 2},
			want:    Position{2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.initial
			got.moveCloserTo(tt.args)
			if got != tt.want {
				t.Errorf("got %#v, want %#v", got, tt.want)
			}
		})
	}
}
