package drivercommon

import (
	"context"
	"image"
)

// 空间大小
type SpaceSize struct {
	Total int64
	Usage int64
}

/**
* 驱动的插件
 */
type DriverPlugin interface {
	DriverPluginInfo                   // driver信息
	SetOption(options ...Option) error // 设置配置
	DriverPluginAction
}

// 驱动信息
type DriverPluginInfo interface {
	Name() string           //名称
	MaxSize() int           //支持最大大小
	SuperEncoder() []string //支持的编码方式
	SpaceSize() SpaceSize   // [0]总空间 [1]使用空间
}

type DriverPluginAction interface {
	Upload(ctx context.Context, data []byte) (metaurl string, err error)   //上传数据
	Download(ctx context.Context, metaurl string) (data []byte, err error) //下载数据
	CheckUrl(ctx context.Context, metaurl string) bool                     // 检查链接是否失效
	UploadUsable() bool                                                    //是否可以上传
	DownloadUsable() bool                                                  //是否可以下载
}

type DriverPluginLogin interface {
	IsLogin() bool
	SetAuthorization(auto string) error // 设置授权
}

// 二维码登陆
type QRCodeLogin interface {
	DriverPluginLogin
	QrcodeLogin(ctx context.Context, show func(ctx context.Context, image image.Image) error) (auto string, err error)
}

//邮箱登陆
type EmailLogin interface {
	DriverPluginLogin
	EmailLogin(ctx context.Context, show func(ctx context.Context, image image.Image) (string, error)) (auto string, err error)
}
