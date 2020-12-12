package main

import "testing"

func Test_isValidParentheses(t *testing.T) {
	type args struct {
		s    string
		pos  int
		next byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test_case_()[]{}",
			args{"()[]{}", 0, 0},
			true,
		},
		{
			"test_case_(]",
			args{"(]", 0, 0},
			false,
		},
		{
			"test_case_([)]",
			args{"([)]", 0, 0},
			false,
		},
		{
			"test_case_{[]}",
			args{"{[]}", 0, 0},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidParentheses(tt.args.s, tt.args.pos, tt.args.next); got != tt.want {
				t.Errorf("isValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
