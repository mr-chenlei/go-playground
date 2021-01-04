package main

import "testing"

func Test_isInterleavingString1(t *testing.T) {
	type args struct {
		s1 string
		s2 string
		s3 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test_case_s1_aabcc_s2_dbbca_s3_aadbbcbcac",
			args{
				"aabcc",
				"dbbca",
				"aadbbcbcac",
			},
			true,
		},
		{
			"test_case_s1_aabcc_s2_dbbca_s3_aadbbbaccc",
			args{
				"aabcc",
				"dbbca",
				"aadbbbaccc",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInterleavingString(tt.args.s1, tt.args.s2, tt.args.s3); got != tt.want {
				t.Errorf("isInterleavingString() = %v, want %v", got, tt.want)
			}
		})
	}
}
