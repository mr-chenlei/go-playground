package main

import "testing"

func Test_findMaxPathSum(t *testing.T) {
	tests := []struct {
		name string
		in   [][]int
		want int
	}{
		{
			"test_case_23",
			[][]int{
				{3},
				{7, 4},
				{2, 4, 6},
				{8, 5, 9, 3},
			},
			23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxPathSum(tt.in); got != tt.want {
				t.Errorf("findMaxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
