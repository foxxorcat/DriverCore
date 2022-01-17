package tools

import (
	"bufio"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"io"

	"github.com/cespare/xxhash"
	"golang.org/x/crypto/sha3"
)

//[32]byte
func XXHash64ReaderHex(r io.Reader) string {
	return hex.EncodeToString(XXHash64Reader(r))
}

func XXHash64Reader(r io.Reader) []byte {
	x := xxhash.New()
	bufio.NewReader(r).WriteTo(x)
	return x.Sum(nil)
}

//[32]byte
func XXHash64Hex(sile []byte) string {
	return hex.EncodeToString(XXHash64(sile[:]))
}

//[8]byte
func XXHash64(sile []byte) []byte {
	x := xxhash.New()
	x.Write(sile[:])
	return x.Sum(nil)
}

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

func SHA3_512Hex(sile []byte) string {
	return hex.EncodeToString(SHA3_512(sile))
}

func SHA3_512(sile []byte) []byte {
	h := sha3.New512()
	h.Write(sile)
	return h.Sum(nil)
}
