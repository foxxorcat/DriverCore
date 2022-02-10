package weibo

import (
	"net/http"
	"net/http/cookiejar"

	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	encodercommon "github.com/foxxorcat/DriverCore/common/encoder"
	"github.com/foxxorcat/DriverCore/crypto"
	encoderimage "github.com/foxxorcat/DriverCore/encoder/image"
	"github.com/guonaihong/gout"
)

const NAME = "weibo"

type WeiBo struct {
	option    drivercommon.DriverOption
	client    *gout.Client
	cookiejar http.CookieJar
}

func New() *WeiBo {
	driver := new(WeiBo)
	driver.cookiejar, _ = cookiejar.New(nil)
	driver.client = gout.NewWithOpt(gout.WithClient(&http.Client{Jar: driver.cookiejar}), gout.WithTimeout(driver.option.Timeout), gout.WithInsecureSkipVerify())
	driver.SetOption(drivercommon.WithCrypto(&crypto.None{}), drivercommon.WithEncoder(encoderimage.NewGif(encodercommon.EncoderOption{Mode: encodercommon.Paletted})))
	return driver
}
