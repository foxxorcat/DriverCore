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
	EncoderImageOption
	size int
}

func (f *Gif) Encode(in []byte) ([]byte, error) {
	img, err := DataToImage(in, f.EncoderImageOption)
	if err != nil {
		return nil, err
	}

	w := new(bytes.Buffer)
	w.Grow(f.size)
	gif.Encode(w, img, nil)
	f.size = w.Cap()
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

func NewGif(mode EncoderMode, option encodercommon.EncoderOption) (encodercommon.EncoderPlugin, error) {
	switch mode {
	case Paletted:
		return &Png{EncoderImageOption: EncoderImageOption{EncoderOption: option, Mode: mode}}, nil
	default:
		return nil, encodercommon.ErrNotSuperImageMod
	}
}
