package crypto

import (
	"hash/crc32"
	"testing"

	cryptocommon "github.com/foxxorcat/DriverCore/common/crypto"
	"github.com/foxxorcat/DriverCore/tools"
)

func TestAes(t *testing.T) {
	param := cryptocommon.CryptoOption{
		K1: []byte("123456"),
		K2: []byte("654321"),
	}

	modes := []string{
		cryptocommon.CFB, cryptocommon.CTR, cryptocommon.OFB,
	}

	lengths := []int{
		16, 24, 32,
	}
	for _, length := range lengths {
		for _, mode := range modes {
			param := cryptocommon.CryptoOption{
				K1:     param.K1,
				K2:     param.K2,
				Mode:   mode,
				Length: length,
			}
			aes, err := NewAes(param)
			if err != nil {
				t.Errorf("参数%+v,错误信息%s", param, err)
				continue
			}
			t.Errorf("参数%+v,开始测试", param)
			for i := 0; i < 20; i++ {
				rawdata := tools.RandomBytes(tools.RangeRand(1024*1024, 1024*1024*4))
				newdata := aes.Decrypt(aes.Encrypt(rawdata))
				if crc32.ChecksumIEEE(rawdata) != crc32.ChecksumIEEE(newdata) {
					t.Errorf("块长度%d,参数%+v,错误信息%s", len(rawdata), param, "hash 验证失败")
				}
			}
			t.Errorf("参数%+v,测试完成", param)
		}
	}
}
