package main

import "testing"

// func TestCallPlusOne(t *testing.T) {
// 	CallAdd(3)
// }

func BenchmarkGoAdd(b *testing.B) {
	GoAdd(b.N)
}

func BenchmarkCAdd(b *testing.B) {
	CAdd(b.N)
}
