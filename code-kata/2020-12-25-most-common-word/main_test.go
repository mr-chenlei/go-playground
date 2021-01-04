package main

import "testing"

func Test_findMostCommonWord(t *testing.T) {
	type args struct {
		input  string
		banned string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test_cast_empty_string",
			args{"", "hahaha"},
			"",
		},
		{
			"test_case_1",
			args{"Bob hit a ball, the hit BALL flew far after it was hit.", "hit"},
			"ball",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMostCommonWord(tt.args.input, tt.args.banned); got != tt.want {
				t.Errorf("findMostCommonWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
