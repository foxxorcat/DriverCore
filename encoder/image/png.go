package encoderimage

import (
	"bytes"
	"image/png"
	"sync"

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
	byteBuffer sync.Pool
}

// 编码
func (p *Png) Encode(in []byte) ([]byte, error) {
	c := p.byteBuffer.Get().(*bytes.Buffer)
	defer p.byteBuffer.Put(c)
	img, err := DataToImage(in, p.EncoderOption, c)
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
		byteBuffer: sync.Pool{
			New: func() interface{} {
				return &bytes.Buffer{}
			},
		},
	}
}
