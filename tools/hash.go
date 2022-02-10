package tools

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

//[32]byte
func Md5Hex(sile []byte) string {
	return hex.EncodeToString(Md5(sile[:]))
}

//[16]byte
func Md5(sile []byte) []byte {
	m := md5.New()
	m.Write(sile[:])
	return m.Sum(nil)
}

func SHA1Hex(sile []byte) string {
	return hex.EncodeToString(SHA1(sile[:]))
}

func SHA1(sile []byte) []byte {
	h := sha1.New()
	h.Write(sile[:])
	return h.Sum(nil)
}

func HmacSHA1(key string, data string) string {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(data))
	return hex.EncodeToString(mac.Sum(nil))
}
