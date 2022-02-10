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
	encodercommon.EncoderOption
}

func (b *Bmp) Encode(in []byte) ([]byte, error) {
	img, err := DataToImage(in, b.EncoderOption)
	if err != nil {
		return nil, err
	}

	w := new(bytes.Buffer)
	bmp.Encode(w, img)
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

func NewBmp(option encodercommon.EncoderOption) *Bmp {
	return &Bmp{
		EncoderOption: option,
	}
}
