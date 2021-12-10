package main

import (
	"DriverCore/common"
	"DriverCore/encoder"
	"DriverCore/tools"
	"testing"
)

func BenchmarkEncoder_Encoded(b *testing.B) {
	for _, encodername := range encoder.EncoderList {
		b.Run(encodername, func(b *testing.B) {
			b.Run("BlockSize128KIB", func(b *testing.B) {
				benchmarkEncoded(b, common.BlockSize128KIB, encoder.PNGALPHA)
			})
			b.Run("BlockSize256KIB", func(b *testing.B) {
				benchmarkEncoded(b, common.BlockSize256KIB, encoder.PNGALPHA)
			})
			b.Run("BlockSize512KIB", func(b *testing.B) {
				benchmarkEncoded(b, common.BlockSize512KIB, encoder.PNGALPHA)
			})
			b.Run("BlockSize1MIB", func(b *testing.B) {
				benchmarkEncoded(b, common.BlockSize1MIB, encoder.PNGALPHA)
			})
			b.Run("BlockSize2MIB", func(b *testing.B) {
				benchmarkEncoded(b, common.BlockSize2MIB, encoder.PNGALPHA)
			})
			b.Run("BlockSize4MIB", func(b *testing.B) {
				benchmarkEncoded(b, common.BlockSize4MIB, encoder.PNGALPHA)
			})
			b.Run("BlockSize8MIB", func(b *testing.B) {
				benchmarkEncoded(b, common.BlockSize8MIB, encoder.PNGALPHA)
			})
			b.Run("BlockSize16MIB", func(b *testing.B) {
				benchmarkEncoded(b, common.BlockSize16MIB, encoder.PNGALPHA)
			})
			b.Run("BlockSize32MIB", func(b *testing.B) {
				benchmarkEncoded(b, common.BlockSize32MIB, encoder.PNGALPHA)
			})
		})
	}

}

func benchmarkEncoded(b *testing.B, i int, name string) {
	enc, err := encoder.NewEncoder(name)
	if err != nil {
		b.Error(err)
		b.FailNow()
	}
	rawblock := tools.RandomBytes(int64(i))
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_, err := enc.Encoded(rawblock)
		if err != nil {
			b.Error(err)
			b.Fail()
		}
	}
}
