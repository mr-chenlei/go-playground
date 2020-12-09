package main

import "testing"

func Test_compressString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test_case_aabcccccaaa",
			args{"aabcccccaaa"},
			"a2b1c5a3",
		},
		{
			"test_case_abbccd",
			args{"abbccd"},
			"abbccd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compressString(tt.args.s); got != tt.want {
				t.Errorf("compressString() = %v, want %v", got, tt.want)
			}
		})
	}
}
