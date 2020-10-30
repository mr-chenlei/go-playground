package main

import "testing"

func Test_isMultiple(t *testing.T) {
	tests := []struct {
		name       string
		n          int
		start      int
		maxDivisor int
		want       bool
	}{
		{
			"test_case_with_2520",
			2520,
			1,
			10,
			true,
		},
		{
			"test_case_with_252",
			252,
			1,
			10,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMultiple(tt.n, tt.start, tt.maxDivisor); got != tt.want {
				t.Errorf("isMultiple() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindMultipleWithDivisor(t *testing.T) {
	tests := []struct {
		name string
		d    int
		want int
	}{
		{
			"test_case_with_divisor_10",
			10,
			2520,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMultipleWithDivisor(tt.d); got != tt.want {
				t.Errorf("findMultipleWithDivisor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindMultipleWithDivisorV2(t *testing.T) {
	tests := []struct {
		name string
		d    int
		want int
	}{
		{
			"test_case_with_divisor_10",
			10,
			2520,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMultipleWithDivisorV2(tt.d); got != tt.want {
				t.Errorf("findMultipleWithDivisorV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
