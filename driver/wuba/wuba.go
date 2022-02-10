package wuba

import (
	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	encodercommon "github.com/foxxorcat/DriverCore/common/encoder"
	"github.com/foxxorcat/DriverCore/crypto"
	encoderimage "github.com/foxxorcat/DriverCore/encoder/image"
	"github.com/guonaihong/gout"
)

const NAME = "wuba"

type WuBa struct {
	option drivercommon.DriverOption
	client *gout.Client
}

func New() *WuBa {
	driver := new(WuBa)
	driver.client = gout.NewWithOpt(gout.WithTimeout(driver.option.Timeout), gout.WithInsecureSkipVerify())
	driver.SetOption(drivercommon.WithCrypto(&crypto.None{}), drivercommon.WithEncoder(encoderimage.NewPng(encodercommon.EncoderOption{Mode: encodercommon.RGBA})))
	return driver
}
