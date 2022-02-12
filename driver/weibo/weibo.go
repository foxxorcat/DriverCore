package weibo

import (
	"net/http"
	"net/http/cookiejar"

	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	"github.com/foxxorcat/DriverCore/crypto"
	"github.com/foxxorcat/DriverCore/encoder"
	encoderimage "github.com/foxxorcat/DriverCore/encoder/image"
	"github.com/guonaihong/gout"
)

const NAME = "weibo"

type WeiBo struct {
	option    drivercommon.DriverOption
	client    *gout.Client
	cookiejar http.CookieJar
}

func New(options ...drivercommon.Option) (drivercommon.DriverPlugin, error) {
	driver := new(WeiBo)
	if err := driver.option.SetOption(options...); err != nil {
		return nil, err
	}

	if driver.option.Encoder == nil {
		e, _ := encoder.NewEncoder(encoderimage.GIFPALETTED)
		driver.option.SetOption(drivercommon.WithEncoder(e))
	}

	if driver.option.Crypto == nil {
		driver.option.SetOption(drivercommon.WithCrypto(&crypto.None{}))
	}

	driver.cookiejar, _ = cookiejar.New(nil)
	driver.client = gout.NewWithOpt(gout.WithClient(&http.Client{Jar: driver.cookiejar}), gout.WithTimeout(driver.option.Timeout), gout.WithInsecureSkipVerify())
	return driver, nil
}
