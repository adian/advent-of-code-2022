package monkey

import (
	"fmt"
	"testing"
)

func TestMonkey_calculateWorryLevel(t *testing.T) {
	type args struct {
		worryLevel int
	}
	tests := []struct {
		name   string
		monkey Monkey
		args   args
		want   int
	}{
		{
			name: "without worry modifier",
			monkey: Monkey{
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
			args: args{
				79,
			},
			want: 1501,
		},
		{
			name: "with worry modifier",
			monkey: Monkey{
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
				WorryLevelModifier: func(worryLevel int) int {
					return worryLevel / 3
				},
			},
			args: args{
				79,
			},
			want: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.monkey.calculateWorryLevel(tt.args.worryLevel); got != tt.want {
				t.Errorf("calculateWorryLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperation_calculate(t *testing.T) {
	type args struct {
		item int
	}
	tests := []struct {
		operation Operation
		args      args
		want      int
	}{
		{
			operation: Operation{
				Num1:     "old",
				Num2:     "19",
				Operator: "*",
			},
			args: args{79},
			want: 1501,
		},
		{
			operation: Operation{
				Num1:     "old",
				Num2:     "old",
				Operator: "*",
			},
			args: args{79},
			want: 6241,
		},
		{
			operation: Operation{
				Num1:     "old",
				Num2:     "6",
				Operator: "+",
			},
			args: args{79},
			want: 85,
		},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%v %v %v", tt.operation.Num1, tt.operation.Operator, tt.operation.Num2)
		t.Run(name, func(t *testing.T) {
			if got := tt.operation.calculate(tt.args.item); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMonkey_ThrowFirstItem(t *testing.T) {
	m := Monkey{
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
		WorryLevelModifier: func(worryLevel int) int {
			return worryLevel / 3
		},
	}

	wantArr := []struct {
		to                  Id
		what                int
		ok                  bool
		numberOfInspections int
		itemsLength         int
	}{
		{
			to:                  3,
			what:                500,
			ok:                  true,
			numberOfInspections: 1,
			itemsLength:         1,
		},
		{
			to:                  3,
			what:                620,
			ok:                  true,
			numberOfInspections: 2,
			itemsLength:         0,
		},
		{
			to:                  0,
			what:                0,
			ok:                  false,
			numberOfInspections: 2,
			itemsLength:         0,
		},
	}

	for i := 0; i < 3; i++ {

		gotTo, gotWhat, gotOk := m.ThrowFirstItem()
		want := wantArr[i]
		if gotTo != want.to {
			t.Errorf("ThrowFirstItem() gotTo = %v, want %v", gotTo, want.to)
		}
		if gotWhat != want.what {
			t.Errorf("ThrowFirstItem() gotWhat = %v, want %v", gotWhat, want.what)
		}
		if gotOk != want.ok {
			t.Errorf("ThrowFirstItem() gotOk = %v, want %v", gotOk, want.ok)
		}
		if m.NumberOfInspections != want.numberOfInspections {
			t.Errorf("ThrowFirstItem() got = %v, want %v", m.NumberOfInspections, want.numberOfInspections)
		}
		if len(m.Items) != want.itemsLength {
			t.Errorf("ThrowFirstItem() got = %v, want %v", len(m.Items), want.itemsLength)
		}
	}

}
