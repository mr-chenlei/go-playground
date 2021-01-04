package main

import (
	"reflect"
	"testing"
)

func Test_findCombinationSum1(t *testing.T) {
	type args struct {
		input []int
		t     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test_case_10_1_2_7_6_1_5",
			args{[]int{10, 1, 2, 7, 6, 1, 5}, 8},
			[][]int{
				[]int{1, 7},
				[]int{1, 2, 5},
				[]int{2, 6},
				[]int{1, 1, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCombinationSum(tt.args.input, tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findCombinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
