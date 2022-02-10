package encoderimage

import (
	"bytes"
	"image/gif"

	encodercommon "github.com/foxxorcat/DriverCore/common/encoder"
)

const (
	GIF         = "gif"
	GIFPALETTED = "gifpaletted"
)

type Gif struct {
	encodercommon.EncoderOption
}

func (f *Gif) Encode(in []byte) ([]byte, error) {
	if f.Mode != encodercommon.Paletted {
		return nil, encodercommon.ErrNotSuperImageMod
	}

	img, err := DataToImage(in, f.EncoderOption)
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

func (*Gif) Mime() string {
	return "image/gif"
}

func (*Gif) Type() string {
	return ".gif"
}

func NewGif(option encodercommon.EncoderOption) *Gif {
	return &Gif{
		EncoderOption: option,
	}
}
