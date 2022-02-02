package wuba

import (
	"context"
	"testing"

	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	encodercommon "github.com/foxxorcat/DriverCore/common/encoder"
	"github.com/foxxorcat/DriverCore/encoder"
	"github.com/foxxorcat/DriverCore/tools"
)

func Test(t *testing.T) {
	driver := New()
	for _, name := range driver.SuperEncoder() {
		encoder, _ := encoder.NewEncoder(name, encodercommon.EncoderOption{})
		driver.SetOption(drivercommon.WithEncoder(encoder))
		rawdata := tools.RandomBytes(int64(driver.MaxSize()))
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

		if tools.XXHash64Hex(rawdata) != tools.XXHash64Hex(downdata) {
			t.Errorf("%s 错误信息%s", name, "hasherr")
			continue
		}
	}

}
