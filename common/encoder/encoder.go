package encodercommon

/**
* 编码插件
 */
type EncoderPlugin interface {
	Encode(in []byte) (out []byte, err error) //编码
	Decode(in []byte) (out []byte, err error) //解码
}
