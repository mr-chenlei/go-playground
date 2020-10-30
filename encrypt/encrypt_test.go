package main

import "testing"

func BenchmarkNewAppKey(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewAppKey()
	}
}

func BenchmarkEncodeAppKey(b *testing.B) {
	appKey := NewAppKey()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EncodeAppKey(appKey)
	}
}

func BenchmarkDecodeAppKey(b *testing.B) {
	appKey := NewAppKey()
	encode := EncodeAppKey(appKey)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DecodeAppKey(encode)
	}
}

func BenchmarkDigestAppKey(b *testing.B) {
	keyID := []byte("0102030405060")
	appKey := NewAppKey()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DigestAppKey(keyID, appKey)
	}
}
