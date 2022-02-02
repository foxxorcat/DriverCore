package driver

import (
	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	"github.com/foxxorcat/DriverCore/driver/baijiahao"
	"github.com/foxxorcat/DriverCore/driver/bilibili"
	"github.com/foxxorcat/DriverCore/driver/weibo"
	"github.com/foxxorcat/DriverCore/driver/wuba"
	"github.com/foxxorcat/DriverCore/driver/yike"
)

var DriverList = []string{
	baijiahao.NAME,
	weibo.NAME,
	bilibili.NAME,
	yike.NAME,
	wuba.NAME,
}

func NewDriver(name string) (drivercommon.DriverPlugin, error) {
	switch name {
	case baijiahao.NAME:
		return baijiahao.New(), nil
	case bilibili.NAME:
		return bilibili.New(), nil
	case weibo.NAME:
		return weibo.New(), nil
	case yike.NAME:
		return yike.New(), nil
	case wuba.NAME:
		return wuba.New(), nil
	default:
		return nil, drivercommon.ErrNoFindDriver
	}
}
