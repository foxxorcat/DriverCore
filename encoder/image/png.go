package encoderimage

import (
	"bytes"
	"image/png"

	encodercommon "github.com/foxxorcat/DriverCore/common/encoder"
)

const (
	PNG         = "png"
	PNGRGB      = "pngrgb"
	PNGRGBA     = "pngrgba"
	PNGPALETTED = "pngpaletted"
	PNGGRAY     = "pnggray"
)

type Png struct {
	encodercommon.EncoderOption
}

// 编码
func (p *Png) Encode(in []byte) ([]byte, error) {
	img, err := DataToImage(in, p.EncoderOption)
	if err != nil {
		return nil, err
	}
	w := new(bytes.Buffer)
	(&png.Encoder{CompressionLevel: png.NoCompression}).Encode(w, img)
	return w.Bytes(), nil
}

// 解码
func (*Png) Decode(in []byte) ([]byte, error) {
	img, err := png.Decode(bytes.NewReader(in))
	if err != nil {
		return nil, err
	}

	return ImageToData(img)
}

func NewPng(option encodercommon.EncoderOption) *Png {
	return &Png{
		EncoderOption: option,
	}
}
