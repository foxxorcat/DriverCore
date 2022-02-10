package yike

import (
	"net/http"
	"net/http/cookiejar"

	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	encodercommon "github.com/foxxorcat/DriverCore/common/encoder"
	"github.com/foxxorcat/DriverCore/crypto"
	encoderimage "github.com/foxxorcat/DriverCore/encoder/image"
	"github.com/guonaihong/gout"
)

const NAME = "yike"

type YiKe struct {
	option    drivercommon.DriverOption
	cookiejar http.CookieJar
	client    *gout.Client
	bdstoken  string
}

func New() *YiKe {
	driver := new(YiKe)
	driver.cookiejar, _ = cookiejar.New(nil)
	driver.client = gout.NewWithOpt(gout.WithClient(&http.Client{Jar: driver.cookiejar}), gout.WithTimeout(driver.option.Timeout), gout.WithInsecureSkipVerify())
	driver.SetOption(drivercommon.WithCrypto(&crypto.None{}), drivercommon.WithEncoder(encoderimage.NewPng(encodercommon.EncoderOption{Mode: encodercommon.RGBA})))
	return driver
}
