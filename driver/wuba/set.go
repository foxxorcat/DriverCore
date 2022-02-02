package wuba

import (
	"math"

	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	encoderimage "github.com/foxxorcat/DriverCore/encoder/image"
	"github.com/foxxorcat/DriverCore/tools"
)

func (wb *WuBa) SetOption(options ...drivercommon.Option) error {
	if err := wb.option.SetOption(options...); err != nil {
		return err
	}

	switch e := wb.option.Encoder.(type) {
	case *encoderimage.Png:
		e.MinSize = int(math.Max(float64(e.MinSize), 10))
		wb.suffix = "png"
	case *encoderimage.Gif:
		e.MinSize = int(math.Max(float64(e.MinSize), 10))
		wb.suffix = "gif"
	default:
		// 未知类型猜测
		v, _ := wb.option.Encoder.Encode(tools.RandomBytes(512))
		wb.suffix = tools.GetFileType(v)
	}
	wb.client.SetTimeout(wb.option.Timeout)
	return nil
}
