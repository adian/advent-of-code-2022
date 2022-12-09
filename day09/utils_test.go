package main

import (
	"fmt"
	"testing"
)

func Test_absDiff(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		args args
		want uint
	}{
		{
			args: args{
				a: 5,
				b: 2,
			},
			want: 3,
		},
		{
			args: args{
				a: -5,
				b: 2,
			},
			want: 7,
		},
		{
			args: args{
				a: -5,
				b: -3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("a: %v, b: %v", tt.args.a, tt.args.b)
		t.Run(name, func(t *testing.T) {
			if got := absDiff(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("absDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
