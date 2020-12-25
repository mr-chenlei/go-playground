package main

import "testing"

func Test_isMountainArray(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test_case_array_0321",
			args{[]int{0, 3, 2, 1}},
			true,
		},
		{
			"test_case_array_355",
			args{[]int{3, 5, 5}},
			false,
		},
		{
			"test_case_array_21",
			args{[]int{2, 1}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMountainArray(tt.args.input); got != tt.want {
				t.Errorf("isMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
