package main

import "testing"

func Test_romeNumber2Arabic(t *testing.T) {
	type args struct {
		rome string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test_case_MCMXCIV",
			args{"MCMXCIV"},
			1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romeNumber2Arabic(tt.args.rome); got != tt.want {
				t.Errorf("romeNumber2Arabic() = %v, want %v", got, tt.want)
			}
		})
	}
}
