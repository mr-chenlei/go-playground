package main

import (
	"bytes"
	"crypto/sha512"
	"encoding/base64"
	"fmt"

	"github.com/google/uuid"
)

const (
	// AppKeyLength ....
	AppKeyLength = 4
)

// NewAppKey ...
func NewAppKey() []byte {
	var buffer bytes.Buffer
	for i := 0; i < AppKeyLength; i++ {
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
	hash.Write(keyID)
	hash.Write(key)
	return hash.Sum(nil)
}

func main() {
	appKey := NewAppKey()
	fmt.Println(appKey)

	encode := EncodeAppKey(appKey)
	fmt.Println(encode)
	decode, _ := DecodeAppKey(encode)
	fmt.Println(decode)
}
