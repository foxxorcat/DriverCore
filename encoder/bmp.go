package encoder

import (
	"math"
	"strings"

	"github.com/foxxorcat/DriverCore/tools"
)

const BMP2BIT = "bmp2bit"

type BMP2bit struct{}

func (*BMP2bit) Name() string {
	return BMP2BIT
}

func (*BMP2bit) Encoded(sile []byte) ([]byte, error) {
	buf := new(strings.Builder)
	buf.WriteString("BM")                                                   //2byte 头
	buf.Write(tools.Uint32ToByteUseLittle(uint32(14 + 40 + 8 + len(sile)))) //4byte 文件大小
	buf.Write([]byte{0, 0, 0, 0})                                           //4byte 保留项
	buf.Write([]byte{62, 0, 0, 0})                                          //4byte 图像数据偏移

	buf.Write([]byte{40, 0, 0, 0})                                                    //4byte 位图信息头的大小（固定）
	buf.Write(tools.Uint32ToByteUseLittle(uint32(len(sile))))                         //4byte 位图宽度
	buf.Write([]byte{1, 0, 0, 0})                                                     //4byte 位图高度
	buf.Write([]byte{1, 0})                                                           //2byte 位图的平面数
	buf.Write([]byte{1, 0})                                                           //2byte  颜色深度
	buf.Write([]byte{0, 0, 0, 0})                                                     //4byte 是否压缩
	buf.Write(tools.Uint32ToByteUseLittle(uint32(math.Ceil(float64(len(sile)) / 8)))) //4byte 图像数据部分大小
	buf.Write([]byte{0, 0, 0, 0})                                                     //4byte 水平分辨率
	buf.Write([]byte{0, 0, 0, 0})                                                     //4byte 垂直分辨率
	buf.Write([]byte{0, 0, 0, 0})                                                     //4byte 使用的颜色
	buf.Write([]byte{0, 0, 0, 0})                                                     //4byte 重要的颜色数
	buf.Write([]byte{0, 0, 0, 0})                                                     //调试板
	buf.Write([]byte{255, 255, 255, 0})                                               //调试板
	buf.Write(sile)

	return tools.Str2bytes(buf.String()), nil
}

func (*BMP2bit) Decode(in []byte) ([]byte, error) {
	return in[62:], nil
}
