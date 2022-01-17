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

func NewEncoder(name string, option encodercommon.EncoderOption) (encodercommon.EncoderPlugin, error) {
	switch name {

	//bmp
	case encoderimage.BMPRGB:
		option.Mode = encodercommon.RGB
		goto BMP
	case encoderimage.BMPRGBA:
		option.Mode = encodercommon.RGBA
		goto BMP
	case encoderimage.BMPPALETTED:
		option.Mode = encodercommon.Paletted
		goto BMP
	case encoderimage.BMPGRAY:
		option.Mode = encodercommon.Gray
		goto BMP
	case encoderimage.BMP:
		goto BMP

		// gif
	case encoderimage.GIFPALETTED:
		option.Mode = encodercommon.Paletted
		goto GIF
	case encoderimage.GIF:
		goto GIF

		//png
	case encoderimage.PNGRGB:
		option.Mode = encodercommon.RGB
		goto PNG
	case encoderimage.PNGRGBA:
		option.Mode = encodercommon.RGBA
		goto PNG
	case encoderimage.PNGPALETTED:
		option.Mode = encodercommon.Paletted
		goto PNG
	case encoderimage.PNGGRAY:
		option.Mode = encodercommon.Gray
		goto PNG
	case encoderimage.PNG:
		goto PNG
	default:
		return nil, encodercommon.ErrNotFindEncoder
	}

PNG:

	return encoderimage.NewPng(option), nil

GIF:
	return encoderimage.NewGif(option), nil

BMP:
	return encoderimage.NewBmp(option), nil

}
