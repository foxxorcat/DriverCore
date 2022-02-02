package drivercommon

import "errors"

var (
	ErrNoFindDriver     = errors.New("不存在驱动")
	ErrNoSuperEncoder   = errors.New("不支持编码器")
	ErrLoginFail        = errors.New("登陆失败")
	ErrNoSuperBlockSize = errors.New("分块大小不支持")
)

var (
	ErrApiFailure = errors.New("api失效")

	ErrQRCodeGetFail = errors.New("二维码获取失败")
	ErrQRCodeFailure = errors.New("二维码失效")
)

var (
	ErrMetaUrlFailure = errors.New("metaUrl失效")
)
