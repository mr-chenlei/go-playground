package main

import "testing"

func Test_untilRepeat(t *testing.T) {
	type args struct {
		s     string
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test_case_ababcdab",
			args{"ababcdab", 2, 8},
			4,
		},
		{
			"test_case_ababcdab",
			args{"ababcdab", 0, 8},
			2,
		},
		{
			"test_case_a",
			args{"a", 0, 1},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := untilRepeat(tt.args.s, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("untilRepeat() = %v, want %v", got, tt.want)
			}
		})
	}
}
