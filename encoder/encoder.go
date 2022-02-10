package encoder

import (
	encodercommon "github.com/foxxorcat/DriverCore/common/encoder"
	encoderimage "github.com/foxxorcat/DriverCore/encoder/image"
)

var EncoderList = []string{
	//encoderimage.BMP,
	encoderimage.BMPRGB,
	encoderimage.BMPRGBA,
	encoderimage.BMPPALETTED,
	encoderimage.BMPGRAY,

	//encoderimage.GIF,
	encoderimage.GIFPALETTED,

	//encoderimage.PNG,
	encoderimage.PNGRGB,
	encoderimage.PNGRGBA,
	encoderimage.PNGPALETTED,
	encoderimage.PNGGRAY,
}

func NewEncoder(name string, option ...encodercommon.Option) (encodercommon.EncoderPlugin, error) {
	var param encodercommon.EncoderOption
	if err := param.SetOption(option...); err != nil {
		return nil, err
	}

	switch name {
	//bmp
	case encoderimage.BMPRGB:
		param.Mode = encodercommon.RGB
		goto BMP
	case encoderimage.BMPRGBA:
		param.Mode = encodercommon.RGBA
		goto BMP
	case encoderimage.BMPPALETTED:
		param.Mode = encodercommon.Paletted
		goto BMP
	case encoderimage.BMPGRAY:
		param.Mode = encodercommon.Gray
		goto BMP

		// gif
	case encoderimage.GIFPALETTED:
		param.Mode = encodercommon.Paletted
		goto GIF

		//png
	case encoderimage.PNGRGB:
		param.Mode = encodercommon.RGB
		goto PNG
	case encoderimage.PNGRGBA:
		param.Mode = encodercommon.RGBA
		goto PNG
	case encoderimage.PNGPALETTED:
		param.Mode = encodercommon.Paletted
		goto PNG
	case encoderimage.PNGGRAY:
		param.Mode = encodercommon.Gray
		goto PNG
	default:
		return nil, encodercommon.ErrNotFindEncoder
	}

PNG:

	return encoderimage.NewPng(param), nil

GIF:
	return encoderimage.NewGif(param), nil

BMP:
	return encoderimage.NewBmp(param), nil

}
