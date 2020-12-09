package main

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test_case_reverse_a-bC-dEf-ghIj",
			args{"a-bC-dEf-ghIj"},
			"j-Ih-gfE-dCba",
		},
		{
			"test_case_reverse_Test1ng-Leet=code-Q!",
			args{"Test1ng-Leet=code-Q!"},
			"Qedo1ct-eeLg=ntse-T!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.input); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
