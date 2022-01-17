package yike

import (
	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	"github.com/foxxorcat/DriverCore/tools"
)

func (b *YiKe) SetOption(options ...drivercommon.Option) error {
	var err error
	for _, option := range options {
		if err = option.Apply(&b.option); err != nil {
			return err
		}
	}

	v, _ := b.option.Encoder.Encode(tools.RandomBytes(512))
	b.suffix = tools.GetFileType(v)

	b.client.SetTimeout(b.option.Timeout)
	return nil
}
