package main

import (
	"reflect"
	"testing"
)

func Test_sortInAlphabetical(t *testing.T) {
	tests := []struct {
		name   string
		intput []string
		want   []string
	}{
		{
			"hapyy_case",
			[]string{
				"DAB",
				"BBA",
				"ABD",
			},
			[]string{
				"ABD",
				"BBA",
				"DAB",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortInAlphabetical(tt.intput); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortInAlphabetical() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_charToNumber(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"string_to_number",
			"ABC",
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := charToNumber(tt.s); got != tt.want {
				t.Errorf("charToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
