package cryptocommon

/**
* 加密插件
 */
type CryptoPlugin interface {
	Encrypt(in []byte) (out []byte)
	Decrypt(in []byte) (out []byte)
}
