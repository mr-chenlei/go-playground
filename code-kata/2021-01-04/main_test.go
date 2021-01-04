package main

import "testing"

func Test_pop(t *testing.T) {
	type args struct {
		s   string
		pos int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test_case_hello_1",
			args{"hello", 1},
			"hllo",
		},
		{
			"test_case_hello_4",
			args{"hello", 4},
			"hell",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pop(tt.args.s, tt.args.pos); got != tt.want {
				t.Errorf("pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_increasingDecreasingString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test_case_aaaabbbbcccc",
			args{"aaaabbbbcccc"},
			"abccbaabccba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := increasingDecreasingString(tt.args.s); got != tt.want {
				t.Errorf("increasingDecreasingString() = %v, want %v", got, tt.want)
			}
		})
	}
}
