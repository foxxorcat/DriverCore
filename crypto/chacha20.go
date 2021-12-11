package crypto

import (
	"crypto/sha1"

	"github.com/foxxorcat/DriverCore/common"

	"github.com/aead/chacha20"
	"golang.org/x/crypto/pbkdf2"
)

const CHACHA20 = "chacha20"

type Chacha20 struct {
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

func NewChacha20(param ...string) (*Chacha20, error) {
	if len(param) < 2 {
		return nil, common.ErrParam
	}

	return &Chacha20{
		key:   pbkdf2.Key([]byte(param[0]), []byte(param[1]), 4096, 32, sha1.New),
		nonce: pbkdf2.Key([]byte(param[0]), []byte(param[1]), 4096, 24, sha1.New),
	}, nil
}
