package cryptocommon

import "github.com/foxxorcat/DriverCore/tools"

type CryptoOption struct {
	K1     []byte // 密码、密钥
	K2     []byte // 偏移、随机数、私钥密码
	Mode   string // 加密模式
	Length int    // 长度
}

// 加密模式

const (
	ECB = "ecb"
	CBC = "cbc"
	CTR = "ctr"
	CFB = "cfb"
	OFB = "ofb"
)

// 默认盐值
var Salt = tools.Str2bytes("3.1415926")
