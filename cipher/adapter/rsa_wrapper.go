package adapter

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log"
)

type RSAWrapper struct {
}

func NewRSAWrapper() *RSAWrapper {
	return &RSAWrapper{}
}

func (r *RSAWrapper) NewKeyPair(bits int) (pub, pri []byte, err error) {
	priKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	pubASN1, err := x509.MarshalPKIXPublicKey(priKey.Public())
	if err != nil {
		return nil, nil, err
	}
	pri = pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(priKey),
		},
	)
	pub = pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	})
	return pub, pri, nil
}

func (r *RSAWrapper) EncryptWithPublicKey(plain, pub []byte) ([]byte, error) {
	pubKey, err := r.bytesToPublicKey(pub)
	if err != nil {
		return nil, err
	}
	hash := sha512.New()
	label := []byte("")
	cipherText, err := rsa.EncryptOAEP(hash, rand.Reader, pubKey, plain, label)
	if err != nil {
		return nil, err
	}
	return cipherText, nil
}

func (r *RSAWrapper) DecryptWithPrivateKey(cipher, pri []byte) ([]byte, error) {
	priKey, err := r.bytesToPrivateKey(pri)
	if err != nil {
		return nil, err
	}
	hash := sha512.New()
	label := []byte("")
	plainText, err := rsa.DecryptOAEP(hash, rand.Reader, priKey, cipher, label)
	if err != nil {
		return nil, err
	}
	return plainText, nil
}

// privateKeyToBytes private key to bytes
func (r *RSAWrapper) privateKeyToBytes(priv *rsa.PrivateKey) []byte {
	privBytes := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(priv),
		},
	)

	return privBytes
}

// publicKeyToBytes public key to bytes
func (r *RSAWrapper) publicKeyToBytes(pub *rsa.PublicKey) ([]byte, error) {
	pubASN1, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return nil, err
	}

	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	})

	return pubBytes, nil
}

// bytesToPrivateKey bytes to private key
func (r *RSAWrapper) bytesToPrivateKey(priv []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(priv)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		log.Println("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			return nil, err
		}
	}
	key, err := x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// bytesToPublicKey bytes to public key
func (r *RSAWrapper) bytesToPublicKey(pub []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pub)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		log.Println("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			return nil, err
		}
	}
	ifc, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		return nil, err
	}
	key, ok := ifc.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not a valid public key")
	}
	return key, nil
}
