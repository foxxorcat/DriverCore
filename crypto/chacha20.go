package crypto

import (
	"crypto/sha1"

	"github.com/aead/chacha20"
	cryptocommon "github.com/foxxorcat/DriverCore/common/crypto"
	"golang.org/x/crypto/pbkdf2"
)

const CHACHA20 = "chacha20"

type Chacha20 struct {
	cryptocommon.CryptoOption
	key   []byte
	nonce []byte
}

func (c *Chacha20) Encrypt(in []byte) []byte {
	out := make([]byte, len(in))
	chacha20.XORKeyStream(out, in, c.nonce, c.key)
	return out
}

func (c *Chacha20) Decrypt(in []byte) []byte {
	out := make([]byte, len(in))
	chacha20.XORKeyStream(out, in, c.nonce, c.key)
	return out
}

func NewChacha20(option cryptocommon.CryptoOption) (*Chacha20, error) {
	switch option.Length {
	case 24, 12, 8:
		return &Chacha20{
			key:   pbkdf2.Key(option.K1, cryptocommon.Salt, 4096, 32, sha1.New),
			nonce: pbkdf2.Key(option.K2, cryptocommon.Salt, 4096, option.Length, sha1.New),
		}, nil
	default:
		return nil, cryptocommon.ErrOption
	}
}
