package main

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test_case_1x1",
			args{
				[][]int{
					{3},
				},
			},
			[][]int{
				{3},
			},
		},
		{
			"test_case_3x3",
			args{
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			[][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			"test_case_5x5",
			args{
				[][]int{
					{1, 1, 1, 1, 1},
					{2, 2, 2, 2, 2},
					{3, 3, 3, 3, 3},
					{4, 4, 4, 4, 4},
					{5, 5, 5, 5, 5},
				},
			},
			[][]int{
				{5, 4, 3, 2, 1},
				{5, 4, 3, 2, 1},
				{5, 4, 3, 2, 1},
				{5, 4, 3, 2, 1},
				{5, 4, 3, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotate(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotate() = %v, want %v", got, tt.want)
			}
		})
	}
}
