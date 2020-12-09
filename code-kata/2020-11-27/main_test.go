package main

import "testing"

func Test_isPathCrossed(t *testing.T) {
	type args struct {
		path []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test_case_path_NES",
			args{[]string{"N", "E", "S"}},
			false,
		},
		{
			"test_case_path_NESWW",
			args{[]string{"N", "E", "S", "W", "W"}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPathCrossed(tt.args.path); got != tt.want {
				t.Errorf("isPathCrossed() = %v, want %v", got, tt.want)
			}
		})
	}
}
