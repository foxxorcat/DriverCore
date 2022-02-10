package wuba

import (
	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	encodercommon "github.com/foxxorcat/DriverCore/common/encoder"
)

func (wb *WuBa) SetOption(options ...drivercommon.Option) error {
	if err := wb.option.SetOption(options...); err != nil {
		return err
	}

	wb.option.Encoder.SetOption(encodercommon.WithMinSize(-10))
	wb.client.SetTimeout(wb.option.Timeout)
	return nil
}
