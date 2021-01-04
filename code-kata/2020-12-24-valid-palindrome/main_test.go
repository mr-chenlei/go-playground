package main

import "testing"

func Test_isValidPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test_case_A man, a plan, a canal: Panama",
			args{"A man, a plan, a canal: Panama"},
			true,
		},
		{
			"test_case_race a car",
			args{"race a car"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isValidPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSymbolOrNumber(t *testing.T) {
	type args struct {
		s int32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test_case_9",
			args{'9'},
			true,
		},
		{
			"test_case_a",
			args{'a'},
			true,
		},
		{
			"test_case_A",
			args{'A'},
			true,
		},
		{
			"test_case_,",
			args{','},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymbolOrNumber(tt.args.s); got != tt.want {
				t.Errorf("isSymbolOrNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
