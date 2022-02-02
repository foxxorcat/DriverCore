package yike

import (
	"context"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"testing"

	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	encodercommon "github.com/foxxorcat/DriverCore/common/encoder"
	"github.com/foxxorcat/DriverCore/encoder"
	"github.com/foxxorcat/DriverCore/tools"
)

func Test(t *testing.T) {
	driver := New()

	str, err := driver.QrcodeLogin(context.Background(), func(ctx context.Context, image image.Image) error {
		file, _ := os.OpenFile(filepath.Join(tools.GetCurrentDir1(), "qrcode.png"), os.O_CREATE|os.O_RDWR, os.ModePerm)
		defer file.Close()
		return png.Encode(file, image)
	})
	os.Remove(filepath.Join(tools.GetCurrentDir1(), "qrcode.png"))
	if err != nil {
		t.Fatalf("错误信息%s", err)
		return
	}

	driver.SetAuthorization(str)
	for _, name := range encoder.EncoderList {
		encoder, _ := encoder.NewEncoder(name, encodercommon.EncoderOption{})
		driver.SetOption(drivercommon.WithEncoder(encoder))
		rawdata := tools.RandomBytes(int64(driver.MaxSize()))

		url, err := driver.Upload(context.Background(), rawdata)
		if err != nil {
			t.Errorf("%s 错误信息%s", name, err)
			continue
		}
		if !driver.CheckUrl(context.Background(), url) {
			t.Errorf("%s 链接%s 错误信息%s", name, url, "checkerr")
			continue
		}
		downdata, err := driver.Download(context.Background(), url)
		if err != nil {
			t.Errorf("%s 链接%s 错误信息%s", name, url, err)
			continue
		}

		if tools.XXHash64Hex(rawdata) != tools.XXHash64Hex(downdata) {
			t.Errorf("%s 错误信息%s", name, "hasherr")
			continue
		}
	}
}
