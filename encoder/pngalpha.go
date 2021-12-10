package encoder

import (
	"bytes"
	"encoding/binary"
	"image"
	"image/png"
	"math"
	"strings"

	"github.com/foxxorcat/DriverCore/common"

	"github.com/foxxorcat/DriverCore/tools"
)

const PNGALPHA = "pngalpha"

type PNGAlpha struct {
}

func (*PNGAlpha) Name() string {
	return PNGALPHA
}

// 编码
func (*PNGAlpha) Encoded(in []byte) ([]byte, error) {

	size := int(math.Max(math.Ceil(math.Sqrt(math.Ceil(float64(len(in)+4)/4))), 10)) // 计算图片长宽
	newData := make([]byte, 4*size*size)

	binary.LittleEndian.PutUint32(newData[:4], uint32(len(in)))
	copy(newData[4:], in)

	nrgba := &image.NRGBA{Pix: newData, Stride: 4 * size, Rect: image.Rect(0, 0, size, size)}
	w := new(strings.Builder)
	if err := (&png.Encoder{CompressionLevel: png.NoCompression}).Encode(w, nrgba); err != nil {
		return nil, err
	}
	return tools.Str2bytes(w.String()), nil
}

// 解码
func (*PNGAlpha) Decode(in []byte) ([]byte, error) {
	img, err := png.Decode(bytes.NewReader(in))
	if err != nil {
		return nil, err
	}

	var buf []byte // 缓存数据
	// 判断图片类型
	switch rgba := img.(type) {
	case *image.NRGBA: // 包含a通道
		buf = rgba.Pix
	case *image.RGBA: // 不包含a通道
		buf = make([]byte, (len(rgba.Pix)/4)*3)
		for i, j := 0, 0; i < len(rgba.Pix); i += 4 {
			copy(buf[j:], rgba.Pix[i:i+3])
			j++
		}
	default:
		return nil, common.ErrImageFormat
	}

	// 验证大小并去除头部
	pixsize := len(buf)
	var size int
	if pixsize >= 4 {
		size = int(binary.LittleEndian.Uint32(buf[:4]))
	}

	if pixsize < size+4 {
		return nil, common.ErrImageCorrupted
	}

	return buf[4 : size+4], nil
}
