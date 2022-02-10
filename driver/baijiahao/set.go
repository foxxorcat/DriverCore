package baijiahao

import (
	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	encodercommon "github.com/foxxorcat/DriverCore/common/encoder"
)

func (b *BaiJiaHao) SetOption(options ...drivercommon.Option) error {
	if err := b.option.SetOption(options...); err != nil {
		return err
	}

	b.option.Encoder.SetOption(encodercommon.WithMinSize(-10))
	b.client.SetTimeout(b.option.Timeout)
	return nil
}
