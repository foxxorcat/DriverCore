package baijiahao

import (
	"context"
	"hash/crc32"
	"testing"

	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	"github.com/foxxorcat/DriverCore/encoder"
	"github.com/foxxorcat/DriverCore/tools"
)

func Test(t *testing.T) {
	driver := New()
	for _, name := range driver.SuperEncoder() {
		encoder, _ := encoder.NewEncoder(name)
		driver.SetOption(drivercommon.WithEncoder(encoder))
		//rawdata := tools.RandomBytes(int64(driver.MaxSize()))
		rawdata := tools.RandomBytes(int64(1024 * 512))
		url, err := driver.Upload(context.Background(), rawdata)
		if err != nil {
			t.Errorf("%s 错误信息%s", name, err)
			continue
		}
		if !driver.CheckUrl(context.Background(), url) {
			t.Errorf("%s 错误信息%s", name, "checkerr")
			continue
		}
		downdata, err := driver.Download(context.Background(), url)
		if err != nil {
			t.Errorf("%s 错误信息%s", name, err)
			continue
		}

		if crc32.ChecksumIEEE(rawdata) != crc32.ChecksumIEEE(downdata) {
			t.Errorf("%s 错误信息%s", name, "hasherr")
			continue
		}
	}

}
