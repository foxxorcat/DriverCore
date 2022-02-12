package baijiahao

import (
	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	"github.com/foxxorcat/DriverCore/crypto"
	"github.com/foxxorcat/DriverCore/encoder"
	encoderimage "github.com/foxxorcat/DriverCore/encoder/image"
	"github.com/guonaihong/gout"
)

const NAME = "baijiahao"

type BaiJiaHao struct {
	option drivercommon.DriverOption
	client *gout.Client
}

func New(options ...drivercommon.Option) (drivercommon.DriverPlugin, error) {
	driver := new(BaiJiaHao)
	if err := driver.option.SetOption(options...); err != nil {
		return nil, err
	}

	if driver.option.Encoder == nil {
		e, _ := encoder.NewEncoder(encoderimage.PNGRGBA)
		driver.option.SetOption(drivercommon.WithEncoder(e))
	}

	if driver.option.Crypto == nil {
		driver.option.SetOption(drivercommon.WithCrypto(&crypto.None{}))
	}

	driver.client = gout.NewWithOpt(gout.WithTimeout(driver.option.Timeout), gout.WithInsecureSkipVerify())
	return driver, nil
}
