package driver

import (
	"fmt"
	"sort"

	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	"github.com/foxxorcat/DriverCore/driver/baijiahao"
	"github.com/foxxorcat/DriverCore/driver/bilibili"
	"github.com/foxxorcat/DriverCore/driver/weibo"
	"github.com/foxxorcat/DriverCore/driver/wuba"
	"github.com/foxxorcat/DriverCore/driver/yike"
)

type NewDriverFunc func(...drivercommon.Option) (drivercommon.DriverPlugin, error)

var driverNameMapNew = map[string]NewDriverFunc{
	baijiahao.NAME: baijiahao.New,
	bilibili.NAME:  bilibili.New,
	weibo.NAME:     weibo.New,
	wuba.NAME:      wuba.New,
	yike.NAME:      yike.New,
}

func AddDriver(name string, new NewDriverFunc) error {
	if name == "" || new == nil {
		return fmt.Errorf("参数错误")
	}
	driverNameMapNew[name] = new
	return nil
}

func GetDrivers() []string {
	list := make([]string, 0, len(driverNameMapNew))
	for name := range driverNameMapNew {
		list = append(list, name)
	}
	sort.Strings(list)
	return list
}

func NewDriver(name string, options ...drivercommon.Option) (drivercommon.DriverPlugin, error) {
	if new, ok := driverNameMapNew[name]; ok {
		driver, err := new(options...)
		if err != nil {
			return nil, err
		}
		return driver, nil
	}
	return nil, drivercommon.ErrNoFindDriver
}
