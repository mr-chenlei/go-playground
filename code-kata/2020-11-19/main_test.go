package main

import "testing"

func Test_multi(t *testing.T) {
	sum := 0
	type args struct {
		x   int
		y   int
		sum int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test_case_3*4",
			args{3, 4, sum},
			12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			multi(tt.args.x, tt.args.y, &sum)
			if tt.want != tt.args.x*tt.args.y {
				t.Fatalf("want: %v, got: %v", tt.want, sum)
			}
		})
	}
}
