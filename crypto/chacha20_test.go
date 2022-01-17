package crypto

import (
	"testing"

	cryptocommon "github.com/foxxorcat/DriverCore/common/crypto"
	"github.com/foxxorcat/DriverCore/tools"
)

func TestChacha20(t *testing.T) {
	param := cryptocommon.CryptoOption{
		K1: []byte("123456"),
		K2: []byte("654321"),
	}

	lengths := []int{
		8, 12, 24,
	}
	for _, length := range lengths {
		param := cryptocommon.CryptoOption{
			K1:     param.K1,
			K2:     param.K2,
			Length: length,
		}
		chacha20, err := NewChacha20(param)
		if err != nil {
			t.Errorf("参数%+v,错误信息%s", param, "hash 验证失败")
			continue
		}
		for i := 0; i < 20; i++ {
			rawdata := tools.RandomBytes(tools.RangeRand(1024*1024, 1024*1024*4))
			newdata := chacha20.Decrypt(chacha20.Encrypt(rawdata))
			if tools.XXHash64Hex(rawdata) != tools.XXHash64Hex(newdata) {
				t.Errorf("块长度%d,参数%+v,错误信息%s", len(rawdata), param, "hash 验证失败")
			}
		}
	}
}
