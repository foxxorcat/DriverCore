package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"io"

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

func (a *Aes) EncryptReader(r io.Reader) io.Reader {
	var crypto interface{}
	switch a.Mode {
	case cryptocommon.CBC:
		crypto = cipher.NewCBCEncrypter(a.block, a.iv)
	case cryptocommon.ECB:
		crypto = tools.NewECBEncrypter(a.block)
	case cryptocommon.CFB:
		crypto = cipher.NewCFBEncrypter(a.block, a.iv)
	case cryptocommon.CTR:
		crypto = cipher.NewCTR(a.block, a.iv)
	case cryptocommon.OFB:
		crypto = cipher.NewOFB(a.block, a.iv)
	default:
		return nil
	}

	switch v := crypto.(type) {
	case cipher.BlockMode:
		return &cryptoBlockEncrypt{
			cryptoReader: cryptoReader{
				crypto: v.CryptBlocks,
				reader: r,
			},
			blockSize: v.BlockSize(),
		}
	case cipher.Stream:
		return &cryptoStream{
			cryptoReader: cryptoReader{
				crypto: v.XORKeyStream,
				reader: r,
			},
			bufSize: a.block.BlockSize(),
		}
	}
	return nil
}

func (a *Aes) DecryptReader(r io.Reader) io.Reader {
	var crypto interface{}
	switch a.Mode {
	case cryptocommon.CBC:
		crypto = cipher.NewCBCDecrypter(a.block, a.iv)
	case cryptocommon.ECB:
		crypto = tools.NewECBDecrypter(a.block)
	case cryptocommon.CFB:
		crypto = cipher.NewCFBDecrypter(a.block, a.iv)
	case cryptocommon.CTR:
		crypto = cipher.NewCTR(a.block, a.iv)
	case cryptocommon.OFB:
		crypto = cipher.NewOFB(a.block, a.iv)
	default:
		return nil
	}

	switch v := crypto.(type) {
	case cipher.BlockMode:
		return &cryptoBlockDecrypt{
			cryptoReader: cryptoReader{
				crypto: v.CryptBlocks,
				reader: r,
			},
			blockSize: v.BlockSize(),
		}
	case cipher.Stream:
		return &cryptoStream{
			cryptoReader: cryptoReader{
				crypto: v.XORKeyStream,
				reader: r,
			},
			bufSize: a.block.BlockSize(),
		}
	}
	return nil
}

func (a *Aes) Encrypt(in []byte) []byte {
	reader := a.EncryptReader(bytes.NewReader(in))
	out := make([]byte, len(in)+(a.block.BlockSize()-len(in)%a.block.BlockSize()))
	n, _ := io.ReadFull(reader, out)
	return out[:n]
}

func (a *Aes) Decrypt(in []byte) []byte {
	reader := a.DecryptReader(bytes.NewReader(in))
	out := make([]byte, len(in))
	n, _ := io.ReadFull(reader, out)
	return out[:n]
}

func NewAes(option cryptocommon.CryptoOption) (cryptocommon.CryptoPlugin, error) {
	switch option.Length {
	case 16, 24, 32:
	default:
		return nil, cryptocommon.ErrOption
	}

	switch option.Mode {
	case cryptocommon.CFB, cryptocommon.CTR, cryptocommon.OFB, cryptocommon.CBC, cryptocommon.ECB:
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
