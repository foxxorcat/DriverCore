package encoder

import (
	"testing"

	encodercommon "github.com/foxxorcat/DriverCore/common/encoder"
	"github.com/foxxorcat/DriverCore/tools"
)

func TestEncoder(t *testing.T) {
	for _, encoderName := range EncoderList {
		encoder, err := NewEncoder(encoderName, encodercommon.EncoderOption{})
		if err != nil {
			t.Errorf("%s 创建失败，错误信息%s", encoderName, err)
			continue
		}
		for i := 0; i < 1; i++ {
			rawdata := tools.RandomBytes(tools.RangeRand(1024*1024, 1024*1024*4))
			newdata, err := encoder.Encode(rawdata)
			if err != nil {
				t.Errorf("块长度%d,参数%s ,错误信息%s", len(rawdata), encoderName, err)
				continue
			}
			newdata, err = encoder.Decode(newdata)
			if err != nil {
				t.Errorf("块长度%d,参数%s,错误信息%s", len(rawdata), encoderName, err)
				continue
			}
			if tools.XXHash64Hex(rawdata) != tools.XXHash64Hex(newdata) {
				t.Errorf("块长度%d,参数%s,错误信息%s", len(rawdata), encoderName, "hash 验证失败")
			}
		}
	}
}

func BenchmarkEncode(b *testing.B) {
	b.StopTimer()
	rawdata := tools.RandomBytes(1024 * 1024 * 4)
	b.ReportAllocs()
	b.StartTimer()
	for _, encoderName := range EncoderList {
		b.Run(encoderName, func(b *testing.B) {
			encoder, err := NewEncoder(encoderName, encodercommon.EncoderOption{})
			if err != nil {
				b.Errorf("%s 创建失败，错误信息%s", encoderName, err)
				return
			}
			for i := 0; i < b.N; i++ {
				data, err := encoder.Encode(rawdata)
				if err != nil {
					b.Errorf("块长度%d,参数%s,错误信息%s", len(rawdata), encoderName, err)
					continue
				}
				newdata, err := encoder.Decode(data)
				if err != nil {
					b.Errorf("块长度%d,参数%s,错误信息%s", len(rawdata), encoderName, err)
					continue
				}
				if tools.XXHash64Hex(rawdata) != tools.XXHash64Hex(newdata) {
					b.Errorf("块长度%d,参数%s,错误信息%s", len(rawdata), encoderName, "hash 验证失败")
				}
			}
		})
	}
}
