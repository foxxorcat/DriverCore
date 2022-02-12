package encoderimage

import (
	"bytes"

	encodercommon "github.com/foxxorcat/DriverCore/common/encoder"
	"golang.org/x/image/bmp"
)

const (
	BMP         = "bmp"
	BMPRGB      = "bmprgb"
	BMPRGBA     = "bmprgba"
	BMPPALETTED = "bmppaletted"
	BMPGRAY     = "bmpgray"
)

type Bmp struct {
	EncoderImageOption
	size int
}

func (b *Bmp) Encode(in []byte) ([]byte, error) {
	img, err := DataToImage(in, b.EncoderImageOption)
	if err != nil {
		return nil, err
	}
	w := new(bytes.Buffer)
	w.Grow(b.size)
	bmp.Encode(w, img)
	b.size = w.Cap()
	return w.Bytes(), nil
}

func (*Bmp) Decode(in []byte) ([]byte, error) {
	img, err := bmp.Decode(bytes.NewReader(in))
	if err != nil {
		return nil, err
	}

	return ImageToData(img)
}

func (*Bmp) Mime() string {
	return "image/x-ms-bmp"
}

func (*Bmp) Type() string {
	return ".bmp"
}

func NewBmp(mode EncoderMode, option encodercommon.EncoderOption) (encodercommon.EncoderPlugin, error) {
	switch mode {
	case RGB, RGBA, Gray, Paletted:
		return &Bmp{EncoderImageOption: EncoderImageOption{EncoderOption: option, Mode: mode}}, nil
	default:
		return nil, encodercommon.ErrNotSuperImageMod
	}
}
