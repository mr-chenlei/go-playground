package main

import (
	"reflect"
	"testing"
)

func Test_doAdd(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		//{
		//	"test_case_array_1234_k_is_23",
		//	args{
		//		[]int{1, 2, 3, 4},
		//		23,
		//	},
		//	[]int{1, 2, 5, 7},
		//},
		//{
		//	"test_case_array_1234_k_is_66",
		//	args{
		//		[]int{1, 2, 3, 4},
		//		66,
		//	},
		//	[]int{1, 3, 0, 0},
		//},
		{
			"test_case_array_1234_k_is_8766",
			args{
				[]int{1, 2, 3, 4},
				8766,
			},
			[]int{1, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doAdd(tt.args.arr, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doAdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
