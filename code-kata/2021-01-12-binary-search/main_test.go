package main

import "testing"

func Test_binarySearch(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test_case_1",
			args{9, []int{-1, 0, 3, 5, 9, 12}},
			4,
		},
		{
			"test_case_2",
			args{2, []int{-1, 0, 3, 5, 9, 12}},
			-1,
		},
		{
			"test_case_3",
			args{-2, []int{-1, 0, 3, 5, 9, 12}},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
