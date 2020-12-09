package main

import "testing"

func Test_isFullyContained(t *testing.T) {
	type args struct {
		n1 int
		n2 int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test_case_125874",
			args{125874, 251748},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isFullyContained(tt.args.n1, tt.args.n2); got != tt.want {
				t.Errorf("isFullyContained() = %v, want %v", got, tt.want)
			}
		})
	}
}
