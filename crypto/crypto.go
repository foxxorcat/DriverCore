package crypto

import (
	"github.com/foxxorcat/DriverCore/common"
)

var CryptoList = []string{
	NONE,
	CHACHA20,
}

func NewCrypto(name string, param ...string) (common.CryptoPlugin, error) {
	switch name {
	case NONE:
		return new(None), nil
	case CHACHA20:
		return NewChacha20(param...)
	default:
		return nil, common.ErrNotFindCrypto
	}
}
