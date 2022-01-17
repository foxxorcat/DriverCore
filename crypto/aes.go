package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"

	cryptocommon "github.com/foxxorcat/DriverCore/common/crypto"
	"github.com/foxxorcat/DriverCore/tools"
	"golang.org/x/crypto/pbkdf2"
)

var AES = "AES"

type Aes struct {
	cryptocommon.CryptoOption
	block cipher.Block
	iv    []byte
}

func (a *Aes) Encrypt(in []byte) []byte {
	var cip interface{}
	switch a.Mode {
	case cryptocommon.CBC:
		in = tools.PKCS7Padding(in, a.block.BlockSize())
		cip = cipher.NewCBCEncrypter(a.block, a.iv)
	case cryptocommon.ECB:
		in = tools.PKCS7Padding(in, a.block.BlockSize())
		cip = tools.NewECBEncrypter(a.block)
	case cryptocommon.CFB:
		cip = cipher.NewCFBEncrypter(a.block, a.iv)
	case cryptocommon.CTR:
		cip = cipher.NewCTR(a.block, a.iv)
	case cryptocommon.OFB:
		cip = cipher.NewOFB(a.block, a.iv)
	default:
		return nil
	}
	out := make([]byte, len(in))
	switch cip := cip.(type) {
	case cipher.Stream:
		cip.XORKeyStream(out, in)
	case cipher.BlockMode:
		cip.CryptBlocks(out, in)
	}

	return out
}

func (a *Aes) Decrypt(in []byte) []byte {
	var cip interface{}
	out := make([]byte, len(in))
	switch a.Mode {
	case cryptocommon.CBC:
		cip = cipher.NewCBCDecrypter(a.block, a.iv)
	case cryptocommon.ECB:
		cip = tools.NewECBDecrypter(a.block)
	case cryptocommon.CFB:
		cip = cipher.NewCFBDecrypter(a.block, a.iv)
	case cryptocommon.CTR:
		cip = cipher.NewCTR(a.block, a.iv)
	case cryptocommon.OFB:
		cip = cipher.NewOFB(a.block, a.iv)
	default:
		return nil
	}

	switch cip := cip.(type) {
	case cipher.Stream:
		cip.XORKeyStream(out, in)
		return out
	case cipher.BlockMode:
		cip.CryptBlocks(out, in)
	}
	return tools.PKCS7UnPadding(out)
}

func NewAes(option cryptocommon.CryptoOption) (*Aes, error) {
	switch option.Length {
	case 16, 24, 32:
	default:
		return nil, cryptocommon.ErrOption
	}

	switch option.Mode {
	case cryptocommon.CFB, cryptocommon.CTR, cryptocommon.OFB, cryptocommon.ECB, cryptocommon.CBC:
	default:
		return nil, cryptocommon.ErrOption
	}

	block, _ := aes.NewCipher(pbkdf2.Key(option.K1, cryptocommon.Salt, 4096, int(option.Length), sha1.New))
	return &Aes{
		block:        block,
		iv:           pbkdf2.Key(option.K2, cryptocommon.Salt, 4096, block.BlockSize(), sha1.New),
		CryptoOption: option,
	}, nil
}
