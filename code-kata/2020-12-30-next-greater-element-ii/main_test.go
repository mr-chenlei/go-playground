package main

import (
	"reflect"
	"testing"
)

func Test_sortBytNextGreaterElement(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test_case_1,2,1",
			args{[]int{1, 2, 1}},
			[]int{2, -1, 2},
		},
		{
			"test_case_3,2,1",
			args{[]int{3, 2, 1}},
			[]int{-1, 3, 3},
		},
		{
			"test_case_2,1",
			args{[]int{2, 1}},
			[]int{-1, 2},
		},
		{
			"test_case_[1,2,3,4,3]",
			args{[]int{1, 2, 3, 4, 3}},
			[]int{2, 3, 4, -1, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortBytNextGreaterElement(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortBytNextGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
