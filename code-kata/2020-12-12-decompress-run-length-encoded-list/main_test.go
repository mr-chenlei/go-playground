package main

import (
	"reflect"
	"testing"
)

func Test_decompress(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test_case_1234",
			args{[]int{1, 2, 3, 4}},
			[]int{2, 4, 4, 4},
		},
		//{
		//	"test_case_123",
		//	args{[]int{1, 2, 3}},
		//	nil,
		//},
		{
			"test_case_1123",
			args{[]int{1, 1, 2, 3}},
			[]int{1, 3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decompress(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decompress() = %v, want %v", got, tt.want)
			}
		})
	}
}
