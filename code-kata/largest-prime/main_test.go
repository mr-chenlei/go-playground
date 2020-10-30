package main

import "testing"

func TestMaxPrime(t *testing.T) {
	tests := []struct {
		name string
		n    uint64
		want uint64
	}{
		// TODO: Add test cases.
		{
			"empty largest-prime",
			1,
			0,
		},
		{
			"happy case",
			15,
			5,
		},
		{
			"13195",
			13195,
			29,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LargestPrime(tt.n)
			if tt.want != got {
				t.Errorf("LargestPrime return %v, want %v", got, tt.want)
			}
		})
	}
}
