package main

import "testing"

func Test_decryptStringFromAlphabet2IntegerMapping(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test_case_10#11#12",
			args{"10#11#12"},
			"jkab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decryptStringFromAlphabet2IntegerMapping(tt.args.input); got != tt.want {
				t.Errorf("decryptStringFromAlphabet2IntegerMapping() = %v, want %v", got, tt.want)
			}
		})
	}
}
