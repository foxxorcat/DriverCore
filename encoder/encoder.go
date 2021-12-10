package encoder

import (
	"DriverCore/common"
	"errors"
)

var (
	ErrImageCorrupted error = errors.New("图片损坏")
	ErrImageFormat          = errors.New("图片格式错误")
	ErrImageDecode          = errors.New("图片解析错误")
	ErrNotFindEncoder       = errors.New("无法找到编码器")
)

var EncoderList = []string{
	PNGALPHA,
	PNGNOTALPHA,
	BMP2BIT,
	NONE,
}

func NewEncoder(name string, param ...string) (common.EncoderPlugin, error) {
	switch name {
	case PNGALPHA:
		return new(PNGAlpha), nil
	case PNGNOTALPHA:
		return new(PNGNotAlpha), nil
	case BMP2BIT:
		return new(BMP2bit), nil
	case NONE:
		return new(None), nil
	default:
		return nil, ErrNotFindEncoder
	}
}
