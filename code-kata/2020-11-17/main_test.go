package main

import (
	"reflect"
	"testing"
)

func Test_sortArray(t *testing.T) {
	type args struct {
		input [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test_caset_sort_success",
			args{
				[][]int{{3, 7}, {1, 4}, {8, 11}},
			},
			[][]int{{1, 4}, {3, 7}, {8, 11}},
		},
		{
			"test_caset_no_need_sort",
			args{
				[][]int{{1, 4}, {3, 7}, {8, 11}},
			},
			[][]int{{1, 4}, {3, 7}, {8, 11}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArray(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_combine(t *testing.T) {
	type args struct {
		input [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test_case_example_1",
			args{
				[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			},
			[][]int{{1, 6}, {8, 10}, {15, 18}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combine(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combine() = %v, want %v", got, tt.want)
			}
		})
	}
}
