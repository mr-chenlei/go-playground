package adapter

import (
	"testing"
)

const (
	RSABits      = 4096
	RSAPlainText = string("www.lstaas.com")
)

func BenchmarkRSANewKeyPair_NewKeyPair(b *testing.B) {
	cipher := NewRSAWrapper()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cipher.NewKeyPair(RSABits)
	}
}

func BenchmarkRSAEncrypt_EncryptWithPublicKey(b *testing.B) {
	cipher := NewRSAWrapper()
	pub, _, _ := cipher.NewKeyPair(RSABits)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cipher.EncryptWithPublicKey([]byte(RSAPlainText), pub)
	}
}

func BenchmarkRSADecrypt_DecryptWithPrivateKey(b *testing.B) {
	cipher := NewRSAWrapper()
	pub, pri, _ := cipher.NewKeyPair(RSABits)
	cipherText, _ := cipher.EncryptWithPublicKey([]byte(RSAPlainText), pub)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cipher.DecryptWithPrivateKey(cipherText, pri)
	}
}
