package weibo

import (
	"math"

	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	encoderimage "github.com/foxxorcat/DriverCore/encoder/image"
)

func (b *WeiBo) SetOption(options ...drivercommon.Option) error {
	if err := b.option.SetOption(options...); err != nil {
		return err
	}

	switch e := b.option.Encoder.(type) {
	case *encoderimage.Gif:
		e.MinSize = int(math.Max(float64(e.MinSize), 10))
	}
	b.client.SetTimeout(b.option.Timeout)
	return nil
}
