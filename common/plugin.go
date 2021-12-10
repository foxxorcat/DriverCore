package common

/**
* 加密插件
 */
type CryptoPlugin interface {
	Name() string
	Encrypt(in []byte) (out []byte)
	Decrypt(in []byte) (out []byte)
}

/**
* 编码插件
 */
type EncoderPlugin interface {
	Name() string
	Encoded(in []byte) (out []byte, err error) //编码
	Decode(in []byte) (out []byte, err error)  //解码
	/* ContentType() string                       // Mime
	FileSuffix() string                        // 文件后缀 */
}
