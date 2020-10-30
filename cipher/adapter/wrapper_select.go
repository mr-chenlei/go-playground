package adapter

// CipherWrapperEnum 加密算法类型的枚举
type CipherWrapperEnum uint8

const (
	WrapperRSA = iota // RSA加密算法的封装
)

// NewWrapper ...
func NewWrapper(n CipherWrapperEnum) CiphersAdapter {
	switch n {
	case WrapperRSA:
		return NewRSAWrapper()
	default:
		return NewRSAWrapper()
	}
}
