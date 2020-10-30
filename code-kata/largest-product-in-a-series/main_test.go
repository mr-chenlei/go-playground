package main

import "testing"

func Test_findLargest(t *testing.T) {
	type args struct {
		input string
		digit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test_case_out_of_boundry",
			args{
				"12345",
				6,
			},
			0,
		},
		{
			"test_case_digit_too_short",
			args{
				"12345",
				0,
			},
			0,
		},
		{
			"test_case_empty_input",
			args{
				"",
				3,
			},
			0,
		},
		{
			"test_case_5832",
			args{
				bigNumber,
				4,
			},
			5832,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLargest(tt.args.input, tt.args.digit); got != tt.want {
				t.Errorf("findLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
