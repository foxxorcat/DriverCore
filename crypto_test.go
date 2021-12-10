package main

import (
	"DriverCore/common"
	"DriverCore/crypto"
	"testing"
)

func BenchmarkCryptoBlockSize128KIB(b *testing.B) {
	benchmarkCrypto(b, common.BlockSize128KIB, crypto.CHACHA20)
}
func BenchmarkCryptoBlockSize256KIB(b *testing.B) {
	benchmarkCrypto(b, common.BlockSize256KIB, crypto.CHACHA20)
}
func BenchmarkCryptoBlockSize512KIB(b *testing.B) {
	benchmarkCrypto(b, common.BlockSize512KIB, crypto.CHACHA20)
}
func BenchmarkCryptoBlockSize1MIB(b *testing.B) {
	benchmarkCrypto(b, common.BlockSize1MIB, crypto.CHACHA20)
}
func BenchmarkCryptoBlockSize2MIB(b *testing.B) {
	benchmarkCrypto(b, common.BlockSize2MIB, crypto.CHACHA20)
}
func BenchmarkCryptoBlockSize4MIB(b *testing.B) {
	benchmarkCrypto(b, common.BlockSize4MIB, crypto.CHACHA20)
}

func benchmarkCrypto(b *testing.B, i int, name string) {

}
