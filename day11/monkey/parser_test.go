package monkey

import (
	"adian.com/advent_of_code_2022/day11/util"
	"reflect"
	"testing"
)

func Test_parseMonkeys(t *testing.T) {
	lines := util.ReadFileLines("../input_test.txt")
	got := Parse(lines)

	want := Monkeys{
		{
			Id:    0,
			Items: []int{79, 98},
			operation: Operation{
				Num1:     "old",
				Num2:     "19",
				Operator: "*",
			},
			ThrowTest: ThrowTest{
				DivisibleBy: 23,
				ifTrue:      2,
				ifFalse:     3,
			},
		},
		{
			Id:    1,
			Items: []int{54, 65, 75, 74},
			operation: Operation{
				Num1:     "old",
				Num2:     "6",
				Operator: "+",
			},
			ThrowTest: ThrowTest{
				DivisibleBy: 19,
				ifTrue:      2,
				ifFalse:     0,
			},
		},
		{
			Id:    2,
			Items: []int{79, 60, 97},
			operation: Operation{
				Num1:     "old",
				Num2:     "old",
				Operator: "*",
			},
			ThrowTest: ThrowTest{
				DivisibleBy: 13,
				ifTrue:      1,
				ifFalse:     3,
			},
		},
		{
			Id:    3,
			Items: []int{74},
			operation: Operation{
				Num1:     "old",
				Num2:     "3",
				Operator: "+",
			},
			ThrowTest: ThrowTest{
				DivisibleBy: 17,
				ifTrue:      0,
				ifFalse:     1,
			},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}

}
