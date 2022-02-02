package encoderimage

import (
	"encoding/binary"
	"image"
	"image/color/palette"
	"math"

	encodercommon "github.com/foxxorcat/DriverCore/common/encoder"
)

func DataToImage(in []byte, option encodercommon.EncoderOption) (image.Image, error) {
	var (
		img  image.Image
		buf  []byte
		size int
	)

	switch option.Mode {
	case encodercommon.RGB:
		size = int(math.Max(float64(option.MinSize), math.Ceil(math.Sqrt(math.Ceil(float64(len(in)+4)/3))))) // 计算图片长宽
		buf = make([]byte, 4*size*size)

		// 优先填充8bit
		binary.LittleEndian.PutUint32(buf[:4], uint32(len(in)))
		buf[4], buf[3], buf[7] = buf[3], 255, 255
		src, dct := in[copy(buf[5:7], in):], buf[8:]
		for len(dct) > 0 {
			if len(src) > 0 {
				src = src[copy(dct[:3], src):]
			}
			dct[3] = 255
			dct = dct[4:]
		}

	case encodercommon.RGBA:
		size = int(math.Max(float64(option.MinSize), math.Ceil(math.Sqrt(math.Ceil(float64(len(in)+4)/4))))) // 计算图片长宽
		buf = make([]byte, 4*size*size)
		binary.LittleEndian.PutUint32(buf[:4], uint32(len(in)))
		copy(buf[4:], in)

	case encodercommon.Paletted, encodercommon.Gray:
		size = int(math.Max(float64(option.MinSize), math.Ceil(math.Sqrt(float64(len(in)+4))))) // 计算图片长宽
		buf = make([]byte, 4*size*size)
		binary.LittleEndian.PutUint32(buf[:4], uint32(len(in)))
		copy(buf[4:], in)

	}
	switch option.Mode {
	case encodercommon.RGB:
		img = &image.RGBA{Pix: buf, Stride: 4 * size, Rect: image.Rect(0, 0, size, size)}
	case encodercommon.RGBA:
		img = &image.NRGBA{Pix: buf, Stride: 4 * size, Rect: image.Rect(0, 0, size, size)}
	case encodercommon.Paletted:
		img = &image.Paletted{Pix: buf, Stride: size, Rect: image.Rect(0, 0, size, size), Palette: palette.Plan9}
	case encodercommon.Gray:
		img = &image.Gray{Pix: buf, Stride: size, Rect: image.Rect(0, 0, size, size)}
	default:
		return nil, encodercommon.ErrNotSuperImageMod
	}
	return img, nil
}

func ImageToData(img image.Image) ([]byte, error) {
	var buf []byte // 缓存数据
	// 判断图片类型
	switch img := img.(type) {
	case *image.NRGBA: // 包含a通道
		buf = img.Pix
	case *image.RGBA: // 不包含a通道
		buf = make([]byte, len(img.Pix)/4*3)
		src, dct := img.Pix, buf
		for len(src) > 0 {
			copy(dct[:3], src[:3])

			src = src[4:]
			dct = dct[3:]
		}
	case *image.Paletted:
		buf = img.Pix
	case *image.Gray:
		buf = img.Pix
	default:
		return nil, encodercommon.ErrImageFormat
	}

	// 验证大小并去除头部
	pixsize := len(buf)
	var size int
	if pixsize >= 4 {
		size = int(binary.LittleEndian.Uint32(buf[:4]))
	}

	if pixsize < size+4 {
		return nil, encodercommon.ErrImageCorrupted
	}

	return buf[4 : size+4], nil
}
