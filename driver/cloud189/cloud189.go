package cloud189

import (
	"net/http"

	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	"github.com/guonaihong/gout"
)

const NAME = "cloud189"

type Cloud189 struct {
	option    drivercommon.DriverOption
	cookiejar http.CookieJar
	client    *gout.Client
}
