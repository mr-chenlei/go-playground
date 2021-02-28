package main

import "testing"

func Test_maskingEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test_case_1",
			args{"LeetCode@LeetCode.com"},
			"l*****e@leetcode.com",
		},
		{
			"AB@qq.com",
			args{"AB@qq.com"},
			"a*****b@qq.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maskingEmail(tt.args.email); got != tt.want {
				t.Errorf("maskEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maskingPhone(t *testing.T) {
	type args struct {
		phone string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		//{
		//	"test_case_1(234)567-890",
		//	args{"1(234)567-890"},
		//	"***-***-7890",
		//},
		//{
		//	"test_case_86-(10)12345678",
		//	args{"86-(10)12345678"},
		//	"+**-***-***-5678",
		//},
		{
			"test_case_(3906)2 07143 711",
			args{"(3906)2 07143 711"},
			"+***-***-***-3711",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maskingPhone(tt.args.phone); got != tt.want {
				t.Errorf("maskingPhone() = %v, want %v", got, tt.want)
			}
		})
	}
}
