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
	EncoderImageOption
	size int
}

// 编码
func (p *Png) Encode(in []byte) ([]byte, error) {
	img, err := DataToImage(in, p.EncoderImageOption)
	if err != nil {
		return nil, err
	}
	w := new(bytes.Buffer)
	w.Grow(p.size)
	(&png.Encoder{CompressionLevel: png.NoCompression}).Encode(w, img)
	p.size = w.Cap()
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

func (*Png) Mime() string {
	return "image/png"
}

func (*Png) Type() string {
	return ".png"
}

func NewPng(mode EncoderMode, option encodercommon.EncoderOption) (encodercommon.EncoderPlugin, error) {
	switch mode {
	case RGB, RGBA, Gray, Paletted:
		return &Png{EncoderImageOption: EncoderImageOption{EncoderOption: option, Mode: mode}}, nil
	default:
		return nil, encodercommon.ErrNotSuperImageMod
	}
}
