package main

import "testing"

func Test_isTriangleString(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test_case_SKY",
			args{"SKY"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTriangleString(tt.args.input); got != tt.want {
				t.Errorf("isTriangleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isTriangleNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test_cast_4",
			args{4},
			false,
		},
		{
			"test_cast_10",
			args{10},
			true,
		},
		{
			"test_cast_55",
			args{55},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTriangleNumber(tt.args.n); got != tt.want {
				t.Errorf("isTriangleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
