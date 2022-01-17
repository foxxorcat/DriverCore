package crypto

import (
	cryptocommon "github.com/foxxorcat/DriverCore/common/crypto"
)

var CryptoList = []string{
	NONE,
	CHACHA20,
	AES,
}

func NewCrypto(name string, option cryptocommon.CryptoOption) (cryptocommon.CryptoPlugin, error) {
	switch name {
	case NONE:
		return new(None), nil
	case CHACHA20:
		return NewChacha20(option)
	case AES:
		return NewAes(option)
	default:
		return nil, cryptocommon.ErrNotFindCrypto
	}
}
