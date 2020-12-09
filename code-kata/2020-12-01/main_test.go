package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test_case_aba",
			args{"aba"},
			true,
		},
		{
			"test_case_abca",
			args{"abca"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tryDeleteOne(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test_case_aba",
			args{"aba"},
			true,
		},
		{
			"test_case_abca",
			args{"abca"},
			true,
		},
		{
			"test_case_abcd",
			args{"abcd"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tryDeleteOne(tt.args.s); got != tt.want {
				t.Errorf("tryDeleteOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
