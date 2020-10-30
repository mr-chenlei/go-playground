package main

import (
	"reflect"
	"testing"
)

func Test_extract(t *testing.T) {
	type args struct {
		in    [][]int
		x     int
		y     int
		steps int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test_cast_5_x_5",
			args{
				[][]int{
					[]int{1, 2, 3, 4, 5},
					[]int{5, 4, 3, 2, 1},
					[]int{3, 5, 7, 9, 10},
					[]int{10, 9, 7, 5, 3},
					[]int{7, 2, 3, 1, 5},
				},
				2, 2, 3,
			},
			[][]int{
				[]int{7, 9, 10}, // right
				[]int{7, 5, 3},  // left
				[]int{7, 3, 3},  // up
				[]int{7, 7, 3},  // down
				[]int{7, 4, 1},  // upper left
				[]int{7, 2, 5},  // upper right
				[]int{7, 9, 7},  // down left
				[]int{7, 5, 5},  // down right
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extract(tt.args.in, tt.args.x, tt.args.y, tt.args.steps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extract() = %v, want %v", got, tt.want)
			}
		})
	}
}
