package main

import "testing"

func Test_alphabetLength(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want int
	}{
		{
			"test_case_twenty-one",
			"twenty-one",
			9,
		},
		{
			"test_case_one_hundred",
			"one-hundred",
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alphabetLength(tt.in); got != tt.want {
				t.Errorf("alphabetLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
