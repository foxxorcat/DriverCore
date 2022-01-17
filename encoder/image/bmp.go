package encoderimage

import (
	"bytes"
	"sync"

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
	byteBuffer sync.Pool
}

func (b *Bmp) Encode(in []byte) ([]byte, error) {
	c := b.byteBuffer.Get().(*bytes.Buffer)
	defer b.byteBuffer.Put(c)
	img, err := DataToImage(in, b.EncoderOption, c)
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

func NewBmp(option encodercommon.EncoderOption) *Bmp {
	return &Bmp{
		EncoderOption: option,
		byteBuffer: sync.Pool{
			New: func() interface{} {
				return &bytes.Buffer{}
			},
		},
	}
}
