package driver

import (
	"context"

	"github.com/foxxorcat/DriverCore/common"
	"github.com/foxxorcat/DriverCore/driver/bilibili"
)

var DriverList = []string{
	bilibili.Name,
}

func NewDriver(name, encoder, crypto string, param []string) (driver common.DriverPlugin, err error) {
	switch name {
	case bilibili.Name:
		driver = new(bilibili.BiLiBiLi)
	default:
		return nil, common.ErrNoFindDriver
	}
	if err = driver.SetEncoder(encoder); err != nil {
		return
	}
	if err = driver.SetCrypto(crypto, param...); err != nil {
		return
	}
	driver.SetContext(context.Background())
	return
}
