package crypto

import (
	"DriverCore/common"
	"errors"
)

var (
	ErrParam         = errors.New("参数错误")
	ErrNotFindCrypto = errors.New("无法找到加密器")
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
		return nil, ErrNotFindCrypto
	}
}
