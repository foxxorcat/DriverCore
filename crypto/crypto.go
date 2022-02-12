package crypto

import (
	"fmt"

	cryptocommon "github.com/foxxorcat/DriverCore/common/crypto"
)

type NewCryptoFunc func(cryptocommon.CryptoOption) (cryptocommon.CryptoPlugin, error)

var cryptoNameMapNew = map[string]NewCryptoFunc{
	NONE:     func(cryptocommon.CryptoOption) (cryptocommon.CryptoPlugin, error) { return new(None), nil },
	CHACHA20: NewChacha20,
	AES:      NewAes,
}

func AddCrypto(name string, new NewCryptoFunc) error {
	if name == "" || new == nil {
		return fmt.Errorf("参数错误")
	}
	cryptoNameMapNew[name] = new
	return nil
}

func GetCryptos() []string {
	list := make([]string, 0, len(cryptoNameMapNew))
	for name := range cryptoNameMapNew {
		list = append(list, name)
	}
	return list
}

func NewCrypto(name string, option cryptocommon.CryptoOption) (cryptocommon.CryptoPlugin, error) {
	if new, ok := cryptoNameMapNew[name]; ok {
		return new(option)
	}
	return nil, cryptocommon.ErrNotFindCrypto
}
