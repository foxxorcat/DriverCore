package driver

import (
	"context"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"testing"

	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	"github.com/foxxorcat/DriverCore/encoder"
	"github.com/foxxorcat/DriverCore/tools"
)

func Test_Drivers(t *testing.T) {
	for _, name := range GetDrivers() {
		t.Run(name, func(t *testing.T) {
			t.Log(name, "开始测试")
			driver, _ := NewDriver(name)
			test_Driver(t, driver, 1024*512)
			t.Log(name, "测试通过")
		})
	}
}

func test_Driver(t *testing.T, driver drivercommon.DriverPlugin, datasize int64) {
	// 登录部分
	if ldriver, ok := driver.(drivercommon.QRCodeLogin); ok {
		path := filepath.Join(t.TempDir(), "qrcode.png")
		t.Log("扫描二维码", path)

		str, err := ldriver.QrcodeLogin(context.Background(), func(ctx context.Context, image image.Image) error {
			file, _ := os.OpenFile(path, os.O_CREATE|os.O_RDWR, os.ModePerm)
			defer file.Close()
			return png.Encode(file, image)
		})
		os.Remove(path)
		if err != nil {
			t.Error(driver.Name(), err)
			return
		}
		ldriver.SetAuthorization(str)
		t.Log("登录成功")
	}

	rawdata := tools.RandomBytes(datasize)
	for _, name := range driver.SuperEncoder() {
		encoder, _ := encoder.NewEncoder(name)
		driver.SetOption(drivercommon.WithEncoder(encoder))

		// 上传
		url, err := driver.Upload(context.Background(), rawdata)
		if err != nil {
			t.Errorf("Driver:%s Encoder:%s 错误信息:%s", driver.Name(), name, err)
			continue
		}
		// 检测链接
		if !driver.CheckUrl(context.Background(), url) {
			t.Errorf("Driver:%s Url:%s Encoder:%s 错误信息:%s", driver.Name(), url, name, "链接失效")
			continue
		}
		// 下载
		downdata, err := driver.Download(context.Background(), url)
		if err != nil {
			t.Errorf("Driver:%s Url:%s Encoder:%s 错误信息:%s", driver.Name(), url, name, err)
			continue
		}

		if tools.Bytes2str(rawdata) != tools.Bytes2str(downdata) {
			t.Errorf("Driver:%s Url:%s Encoder:%s 错误信息:%s", driver.Name(), url, name, "数据损坏")
			continue
		}
	}
}
