package common

import "errors"

var (
	ErrNoFindDriver     = errors.New("不存在驱动")
	ErrNoSuperEncoder   = errors.New("不支持编码器")
	ErrLoginFail        = errors.New("登陆失败")
	ErrNoSuperBlockSize = errors.New("分块大小不支持")
)

var (
	ErrImageCorrupted error = errors.New("图片损坏")
	ErrImageFormat          = errors.New("图片格式错误")
	ErrImageDecode          = errors.New("图片解析错误")
	ErrNotFindEncoder       = errors.New("无法找到编码器")
)

var (
	ErrParam         = errors.New("参数错误")
	ErrNotFindCrypto = errors.New("无法找到加密器")
)
