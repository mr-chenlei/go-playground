package main

import (
	"bytes"
	"crypto/sha512"
	"encoding/base64"
	"fmt"

	"github.com/google/uuid"
)

const (
	// AppKeyLengthOfUUID AppKey所包含UUID的数量
	AppKeyLengthOfUUID = 2
)

// NewAppKey ...
func NewAppKey() []byte {
	var buffer bytes.Buffer
	for i := 0; i < AppKeyLengthOfUUID; i++ {
		k := uuid.New()
		buffer.Write(k[:])
	}
	return buffer.Bytes()
}

// EncodeAppKey ...
func EncodeAppKey(key []byte) string {
	return base64.StdEncoding.EncodeToString(key)
}

// DecodeAppKey ...
func DecodeAppKey(key string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(key)
}

// DigestAppKey ...
func DigestAppKey(keyID, key []byte) []byte {
	hash := sha512.New()
	digestedKeyID := hash.Sum(keyID)
	hash.Reset()
	hash.Write(key)
	hash.Write(digestedKeyID)
	return hash.Sum(nil)
}

func main() {
	key := NewAppKey()
	encode := EncodeAppKey(key)
	decode, err := DecodeAppKey(encode)
	if err != nil {
		fmt.Println(err)
		return
	}

	var b [64]byte
	copy(b[:], decode[:])
	fmt.Println("key:", encode, len(encode))
	fmt.Println("[64]byte:", b)
	fmt.Println("[]byte:", decode, len(decode))
}
