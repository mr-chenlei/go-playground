package main

import "testing"

func Test_reverse2K(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse2K(tt.args.input); got != tt.want {
				t.Errorf("reverse2K() = %v, want %v", got, tt.want)
			}
		})
	}
}
