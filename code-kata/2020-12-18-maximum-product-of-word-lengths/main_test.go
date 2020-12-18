package main

import "testing"

func Test_maxProduct(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test_case_1",
			args{[]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}},
			16,
		},
		{
			"test_case_2",
			args{[]string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}},
			4,
		},
		{
			"test_case_3",
			args{[]string{"a", "aa", "aaa", "aaaa"}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.input); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProductV2(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test_case_1",
			args{[]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}},
			16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProductV2(tt.args.input); got != tt.want {
				t.Errorf("maxProductV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
