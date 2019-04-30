package main

import "testing"

func BenchmarkGoAdd(b *testing.B) {
	GoAdd(b.N)
}

func BenchmarkCAdd(b *testing.B) {
	CAdd(b.N)
}
