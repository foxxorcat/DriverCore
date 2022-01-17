package cryptocommon

import "errors"

var (
	ErrOption        = errors.New("配置错误")
	ErrNotFindCrypto = errors.New("无法找到加密器")
)
