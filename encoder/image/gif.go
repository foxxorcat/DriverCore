package encoderimage

import (
	"bytes"
	"image/gif"
	"sync"

	encodercommon "github.com/foxxorcat/DriverCore/common/encoder"
)

const (
	GIF         = "gif"
	GIFPALETTED = "gifpaletted"
)

type Gif struct {
	encodercommon.EncoderOption
	byteBuffer sync.Pool
}

func (f *Gif) Encode(in []byte) ([]byte, error) {
	if f.Mode != encodercommon.Paletted {
		return nil, encodercommon.ErrNotSuperImageMod
	}

	c := f.byteBuffer.Get().(*bytes.Buffer)
	defer f.byteBuffer.Put(c)
	img, err := DataToImage(in, f.EncoderOption, c)
	if err != nil {
		return nil, err
	}

	w := new(bytes.Buffer)
	gif.Encode(w, img, nil)
	return w.Bytes(), nil
}

func (*Gif) Decode(in []byte) ([]byte, error) {
	img, err := gif.Decode(bytes.NewReader(in))
	if err != nil {
		return nil, err
	}

	return ImageToData(img)
}

func NewGif(option encodercommon.EncoderOption) *Gif {
	return &Gif{
		EncoderOption: option,
		byteBuffer: sync.Pool{
			New: func() interface{} {
				return &bytes.Buffer{}
			},
		},
	}
}
