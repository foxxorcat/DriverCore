package crypto

const NONE = "none"

type None struct{}

func (c *None) Encrypt(in []byte) []byte {
	return in
}

func (c *None) Decrypt(in []byte) []byte {
	return in
}
