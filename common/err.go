package common

import "errors"

var (
	ErrNoFindDriverErr  = errors.New("不存在驱动")
	ErrNoSuperEncoder   = errors.New("不支持编码器")
	ErrLoginFail        = errors.New("登陆失败")
	ErrNoSuperBlockSize = errors.New("分块大小不支持")
)
