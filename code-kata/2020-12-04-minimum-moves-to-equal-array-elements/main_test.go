package main

import "testing"

func Test_findMoves(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test_case_123",
			args{[]int{1, 2, 3}},
			3,
		},
		{
			"test_case_4789",
			args{[]int{4, 7, 8, 8}},
			11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMoves(tt.args.input); got != tt.want {
				t.Errorf("findMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}
