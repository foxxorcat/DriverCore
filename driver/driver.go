package driver

import (
	"context"

	"github.com/foxxorcat/DriverCore/common"
	"github.com/foxxorcat/DriverCore/driver/bilibili"
)

var DriverList = []common.DriverPluginInfo{
	bilibili.Info,
}

func NewDriver(name, encoder, crypto string, param ...string) (driver common.DriverPlugin, err error) {
	return NewDriverWithCtx(context.Background(), name, encoder, crypto, param...)
}

func NewDriverWithCtx(ctx context.Context, name, encoder, crypto string, param ...string) (driver common.DriverPlugin, err error) {
	switch name {
	case bilibili.Info.Name():
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
	return driver.SetContext(ctx), nil
}
