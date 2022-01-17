package encodercommon

import "errors"

var (
	ErrNotSuperMod     = errors.New("不支持当前模式")
	ErrNotSuperEncoded = errors.New("不支持编码")
	ErrNotFindEncoder  = errors.New("无法找到编码器")

	// image
	ErrNotSuperImageMod = errors.New("不支持当前图片模式")
	ErrImageCorrupted   = errors.New("图片损坏")
	ErrImageFormat      = errors.New("图片格式错误")
	ErrImageDecode      = errors.New("图片解析错误")
)
