package encodercommon

/**
* 编码插件
 */
type EncoderPlugin interface {
	Encode(in []byte) (out []byte, err error) //编码
	Decode(in []byte) (out []byte, err error) //解码
	SetOption(options ...Option) error
	Mime() string // 数据mime
	Type() string // 数据后缀
}
