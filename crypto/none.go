package crypto

import "io"

const NONE = "none"

type None struct{}

func (c *None) Encrypt(in []byte) []byte {
	return in
}

func (c *None) Decrypt(in []byte) []byte {
	return in
}

func (c *None) EncryptReader(r io.Reader) io.Reader {
	return r
}

func (c *None) DecryptReader(r io.Reader) io.Reader {
	return r
}
