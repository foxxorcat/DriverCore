package encoder

import (
	"fmt"
	"sort"

	encodercommon "github.com/foxxorcat/DriverCore/common/encoder"
	encoderimage "github.com/foxxorcat/DriverCore/encoder/image"
)

type NewEnocerFunc func(options ...encodercommon.Option) (encodercommon.EncoderPlugin, error)

var encoderNameMapNew = map[string]NewEnocerFunc{
	encoderimage.BMPRGB:      encoderimage.WithEncoderNewFunc(encoderimage.RGB, encoderimage.NewBmp),
	encoderimage.BMPRGBA:     encoderimage.WithEncoderNewFunc(encoderimage.RGBA, encoderimage.NewBmp),
	encoderimage.BMPGRAY:     encoderimage.WithEncoderNewFunc(encoderimage.Gray, encoderimage.NewBmp),
	encoderimage.BMPPALETTED: encoderimage.WithEncoderNewFunc(encoderimage.Paletted, encoderimage.NewBmp),

	encoderimage.PNGRGB:      encoderimage.WithEncoderNewFunc(encoderimage.RGB, encoderimage.NewPng),
	encoderimage.PNGRGBA:     encoderimage.WithEncoderNewFunc(encoderimage.RGBA, encoderimage.NewPng),
	encoderimage.PNGGRAY:     encoderimage.WithEncoderNewFunc(encoderimage.Gray, encoderimage.NewPng),
	encoderimage.PNGPALETTED: encoderimage.WithEncoderNewFunc(encoderimage.Paletted, encoderimage.NewPng),

	encoderimage.GIFPALETTED: encoderimage.WithEncoderNewFunc(encoderimage.Paletted, encoderimage.NewGif),
}

func AddEncoder(name string, new NewEnocerFunc) error {
	if name == "" || new == nil {
		return fmt.Errorf("参数错误")
	}
	encoderNameMapNew[name] = new
	return nil
}

func GetEncoders() []string {
	list := make([]string, 0, len(encoderNameMapNew))
	for name := range encoderNameMapNew {
		list = append(list, name)
	}
	sort.Strings(list)
	return list
}

func NewEncoder(name string, options ...encodercommon.Option) (encodercommon.EncoderPlugin, error) {
	if new, ok := encoderNameMapNew[name]; ok {
		return new(options...)
	}
	return nil, encodercommon.ErrNotSuperEncoded
}
