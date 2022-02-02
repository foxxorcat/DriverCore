package baijiahao

import (
	"math"

	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	encoderimage "github.com/foxxorcat/DriverCore/encoder/image"
	"github.com/foxxorcat/DriverCore/tools"
)

func (b *BaiJiaHao) SetOption(options ...drivercommon.Option) error {
	if err := b.option.SetOption(options...); err != nil {
		return err
	}

	switch e := b.option.Encoder.(type) {
	case *encoderimage.Png:
		e.MinSize = int(math.Max(float64(e.MinSize), 10))
		b.suffix = "png"
	case *encoderimage.Bmp:
		e.MinSize = int(math.Max(float64(e.MinSize), 10))
		b.suffix = "bmp"
	case *encoderimage.Gif:
		e.MinSize = int(math.Max(float64(e.MinSize), 10))
		b.suffix = "gif"
	default:
		// 未知类型猜测
		v, _ := b.option.Encoder.Encode(tools.RandomBytes(512))
		b.suffix = tools.GetFileType(v)
	}
	b.client.SetTimeout(b.option.Timeout)
	return nil
}
