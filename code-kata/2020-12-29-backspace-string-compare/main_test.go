package main

import "testing"

func Test_isEqual(t *testing.T) {
	type args struct {
		s1        string
		s2        string
		backspace int32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test_case_ab#c_ad#c",
			args{"ab#c", "ad#c", '#'},
			true,
		},
		{
			"test_case_ab##_c#d#",
			args{"ab##", "c#d#", '#'},
			true,
		},
		{
			"test_case_a##c_#a#c",
			args{"a##c", "#a#c", '#'},
			true,
		},
		{
			"test_cast_a#c_b",
			args{"a#c", "b", '#'},
			false,
		},
		{
			"test_case_empty_inputs",
			args{"", "", '#'},
			true,
		},
		{
			"test_case_#_#",
			args{"#", "#", '#'},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEqual(tt.args.s1, tt.args.s2, tt.args.backspace); got != tt.want {
				t.Errorf("isEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
