package main

import (
	"reflect"
	"testing"
)

func Test_fractions(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test_cast_with_fraction_1",
			args{1},
			nil,
		},
		{
			"test_case_fraction_with_3",
			args{3},
			[]string{"1/2", "1/3", "2/3"},
		},
		{
			"test_case_fraction_with_4",
			args{4},
			[]string{"1/2", "1/3", "2/3", "1/4", "3/4"},
		},
		//{
		//	"test_case_fraction_with_9",
		//	args{9},
		//	[]string{"1/2", "1/3", "2/3", "1/4", "3/4"},
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fractions(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fractions() = %v, want %v", got, tt.want)
			}
		})
	}
}
