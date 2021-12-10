package encoder

import (
	"DriverCore/tools"
	"bytes"
	"encoding/binary"
	"image"
	"image/png"
	"math"
	"strings"
)

const PNGNOTALPHA = "pngnotalpha"

type PNGNotAlpha struct{}

func (*PNGNotAlpha) Name() string {
	return PNGNOTALPHA
}

// 编码
func (*PNGNotAlpha) Encoded(in []byte) ([]byte, error) {
	size := int(math.Ceil(math.Sqrt(math.Ceil(float64(len(in)+4) / 3)))) // 计算图片长宽
	newData := make([]byte, size*size*4)

	binary.LittleEndian.PutUint32(newData[:4], uint32(len(in)))
	newData[4] = newData[3]

	for i, j := 0, 0; i < len(newData); i += 4 {
		newData[i+3] = 255
		if j < len(in) && i >= 8 {
			copy(newData[i:i+3], in[j:])
			j += 3
		}
	}

	rgba := &image.RGBA{Pix: newData, Stride: 4 * size, Rect: image.Rect(0, 0, size, size)}
	w := new(strings.Builder)
	if err := (&png.Encoder{CompressionLevel: png.NoCompression}).Encode(w, rgba); err != nil {
		return nil, err
	}
	return tools.Str2bytes(w.String()), nil
}

// 解码
func (*PNGNotAlpha) Decode(in []byte) ([]byte, error) {
	img, err := png.Decode(bytes.NewReader(in))
	if err != nil {
		return nil, err
	}

	var buf []byte // 缓存数据
	// 判断图片类型
	switch rgba := img.(type) {
	case *image.RGBA: // 包含a通道
		buf = make([]byte, int(math.Ceil(float64(len(rgba.Pix))/4*3)))
		for i, j := 0, 0; i < len(rgba.Pix); i += 4 {
			copy(buf[j:j+3], rgba.Pix[i:i+3])
			j += 3
		}
	default:
		return nil, ErrImageFormat
	}

	// 验证大小并去除头部
	pixsize := len(buf)
	var size int
	if pixsize >= 6 {
		size = int(binary.LittleEndian.Uint32(buf[:4]))
	}

	if pixsize < size+6 {
		return nil, ErrImageCorrupted
	}

	return buf[6 : size+6], nil
}
