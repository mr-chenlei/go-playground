package main

import "testing"

func Test_doMathV1(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test_cast_8",
			args{8},
			3,
		},
		{
			"test_cast_5",
			args{5},
			2,
		},
		{
			"test_cast_10",
			args{10},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doMathV1(tt.args.n); got != tt.want {
				t.Errorf("doMathV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_doMathV2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test_cast_8",
			args{8},
			3,
		},
		{
			"test_cast_5",
			args{5},
			2,
		},
		{
			"test_cast_10",
			args{10},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doMathV2(tt.args.n); got != tt.want {
				t.Errorf("doMathV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
