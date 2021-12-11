package common

/**
* 加密插件
 */
type CryptoPlugin interface {
	Encrypt(in []byte) (out []byte)
	Decrypt(in []byte) (out []byte)
}

/**
* 编码插件
 */
type EncoderPlugin interface {
	Encoded(in []byte) (out []byte, err error) //编码
	Decode(in []byte) (out []byte, err error)  //解码
	/* ContentType() string                       // Mime
	FileSuffix() string                        // 文件后缀 */
}

type EncoderParam struct {
	MinSize  int  // 最小数据大小(用于填充)
	MaxSize  int  // 最大数据大小(用于分块)
	Compress bool // 是否压缩
	Mod      string
}
