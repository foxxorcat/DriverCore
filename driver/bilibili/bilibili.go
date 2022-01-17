package bilibili

import (
	"net/http"
	"net/http/cookiejar"

	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	encodercommon "github.com/foxxorcat/DriverCore/common/encoder"
	"github.com/foxxorcat/DriverCore/crypto"
	"github.com/foxxorcat/DriverCore/encoder"
	encoderimage "github.com/foxxorcat/DriverCore/encoder/image"
	"github.com/guonaihong/gout"
)

const NAME = "bilibili"

type BiLiBiLi struct {
	option    drivercommon.DriverOption
	suffix    string
	client    *gout.Client
	cookiejar http.CookieJar
}

func New() *BiLiBiLi {
	e, _ := encoder.NewEncoder(encoderimage.PNGRGBA, encodercommon.EncoderOption{})
	driver := new(BiLiBiLi)
	driver.cookiejar, _ = cookiejar.New(nil)
	driver.client = gout.NewWithOpt(gout.WithClient(&http.Client{Jar: driver.cookiejar}), gout.WithTimeout(driver.option.Timeout), gout.WithInsecureSkipVerify())
	driver.SetOption(drivercommon.WithCrypto(&crypto.None{}), drivercommon.WithEncoder(e))
	return driver
}
