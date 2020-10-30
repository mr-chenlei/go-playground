package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		args int
		want bool
	}{
		// TODO: Add test cases.
		{
			"test_case_with_71",
			71,
			false,
		},
		{
			"test_case_with_101",
			101,
			true,
		},
		{
			"test_case_with_9009",
			9009,
			true,
		},
		//{
		//	"test_case_with_580085",
		//	58185,
		//	false,
		//},
		{
			"test_case_with_9900091",
			9900091,
			false,
		},
		{
			"test_case_with_906609",
			906609,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
