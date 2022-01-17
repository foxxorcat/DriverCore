package baijiahao

import (
	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	encodercommon "github.com/foxxorcat/DriverCore/common/encoder"
	"github.com/foxxorcat/DriverCore/crypto"
	"github.com/foxxorcat/DriverCore/encoder"
	encoderimage "github.com/foxxorcat/DriverCore/encoder/image"
	"github.com/guonaihong/gout"
)

const NAME = "baijiahao"

type BaiJiaHao struct {
	option drivercommon.DriverOption
	suffix string
	client *gout.Client
}

func New() *BaiJiaHao {
	e, _ := encoder.NewEncoder(encoderimage.PNGRGBA, encodercommon.EncoderOption{})
	driver := new(BaiJiaHao)
	driver.client = gout.NewWithOpt(gout.WithTimeout(driver.option.Timeout), gout.WithInsecureSkipVerify())
	driver.SetOption(drivercommon.WithCrypto(&crypto.None{}), drivercommon.WithEncoder(e))

	return driver
}
