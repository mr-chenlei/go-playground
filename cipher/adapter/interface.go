package adapter

// CiphersAdapter 与其他加密算法库的适配接口
type CiphersAdapter interface {
	// NewKeyPair 生成新的公/私密钥对
	NewKeyPair(bits int) (pub, pri []byte, err error)
	// EncryptWithPublicKey 使用公钥对数据加密
	EncryptWithPublicKey(plain, pub []byte) ([]byte, error)
	// DecryptWithPrivateKey 使用私钥解密
	DecryptWithPrivateKey(cipher, pri []byte) ([]byte, error)
}
