package common

import (
	"context"
	"time"
)

/**
* 驱动的插件
 */
type DriverPlugin interface {
	DriverPluginInfo
	DriverPluginSet
	DriverPluginAction
	DriverPluginUsable
}

type DriverPluginInfo interface {
	Name() string           //名称
	MaxSize() int           //支持最大大小
	SuperEncoder() []string //支持的编码方式
}

type DriverPluginSet interface {
	SetContext(ctx context.Context) DriverPlugin
	SetTimeOut(time time.Duration) DriverPlugin   //设置超时
	SetAttempt(t uint) DriverPlugin               //设置重试次数
	SetEncoder(name string) error                 //设置数据编码
	SetCrypto(name string, param ...string) error // 设置加密
}

type DriverPluginAction interface {
	Upload(data []byte) (metaurl string, err error)   //上传数据
	Download(metaurl string) (data []byte, err error) //下载数据
	CheckUrl(metaurl string) bool                     // 检查链接是否失效
}

type DriverPluginUsable interface {
	UploadUsable() bool   //是否可以上传
	DownloadUsable() bool //是否可以下载
}

type DriverLoginPlugin interface {
	IsLogin() bool
	SetAuthorization(auto string) error // 设置授权
}
