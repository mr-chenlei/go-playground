package main

import "testing"

func Test_parseString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		//{
		//	"test_case_G()(al)",
		//	args{"G()(al)"},
		//	"Goal",
		//},
		//{
		//	"test_case_G()()()()(al)",
		//	args{"G()()()()(al)"},
		//	"Gooooal",
		//},
		{
			"test_case_(al)G(al)()()G",
			args{"(al)G(al)()()G"},
			"alGalooG",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseString(tt.args.s); got != tt.want {
				t.Errorf("parseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
