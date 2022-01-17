package tools

import (
	"math/rand"
	"strings"
	"time"
)

var Random = rand.New(rand.NewSource(time.Now().Unix()))

//生成随机字节
func RandomBytes(n int64) []byte {
	buf := make([]byte, n)
	for i := int64(0); i < n; i++ {
		buf[i] = uint8(Random.Intn(255))
	}
	return buf
}

//随机生成字节读取器
func RandomBytesReader(n int64) *strings.Reader {
	return strings.NewReader(string(RandomBytes(n)))
}

// 随机生成 [m,n) 范围内随机数
func RangeRand(min, max int64) int64 {
	return Random.Int63n(max-min) + min
}

const char = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// 随机生成字符串
func RandChar(size int) string {
	var s strings.Builder
	for i := 0; i < size; i++ {
		s.WriteByte(char[Random.Int63()%int64(len(char))])
	}
	return s.String()
}
