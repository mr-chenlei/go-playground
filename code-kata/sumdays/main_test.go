package main

import "testing"

func Test_isSunday(t *testing.T) {
	type args struct {
		y int
		m int
		d int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test_case_1989_10_1",
			args{
				1989,
				10,
				1,
			},
			true,
		},
		//{
		//	"test_case_1989_1_1",
		//	args{
		//		1989,
		//		1,
		//		1,
		//	},
		//	true,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSunday(tt.args.y, tt.args.m, tt.args.d); got != tt.want {
				t.Errorf("isSunday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_monthConvert(t *testing.T) {
	type args struct {
		y int
		m int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			"test_case_1999_2",
			args{
				1999,
				2,
			},
			1998,
			14,
		},
		{
			"test_case_1999_1",
			args{
				1999,
				1,
			},
			1998,
			13,
		},
		{
			"test_case_1999_3",
			args{
				1999,
				3,
			},
			1999,
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := monthConvert(tt.args.y, tt.args.m)
			if got != tt.want {
				t.Errorf("monthConvert() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("monthConvert() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
