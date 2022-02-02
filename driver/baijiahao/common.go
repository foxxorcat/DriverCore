package baijiahao

import (
	"context"
	"fmt"

	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	encoderimage "github.com/foxxorcat/DriverCore/encoder/image"
)

func (bjh *BaiJiaHao) formatUrl(metaurl string) string {
	return fmt.Sprintf("https://pic.rmb.bdstatic.com/bjh/%s.%s", metaurl, bjh.suffix)
}

func (bjh *BaiJiaHao) Name() string {
	return NAME
}

func (bjh *BaiJiaHao) MaxSize() int {
	switch bjh.option.Encoder.(type) {
	case *encoderimage.Gif:
		return 14 * (2 << 19) // 16MIB
	case *encoderimage.Png:
		return 18 * (2 << 19) // 28MIB
	case *encoderimage.Bmp:
		return 19 * (2 << 19) // 29MIB
	default:
		return 20 * (2 << 19) // 30MIB
	}
}

func (bjh *BaiJiaHao) SuperEncoder() []string {
	return []string{
		encoderimage.BMPRGB,
		encoderimage.BMPRGBA,
		encoderimage.BMPPALETTED,
		encoderimage.BMPGRAY,

		encoderimage.GIFPALETTED,

		encoderimage.PNGRGB,
		encoderimage.PNGRGBA,
		encoderimage.PNGPALETTED,
		encoderimage.PNGGRAY,
	}
}

func (bjh *BaiJiaHao) DownloadUsable() bool {
	return true
}

func (bjh *BaiJiaHao) UploadUsable() bool {
	return true
}

func (bjh *BaiJiaHao) SpaceSize() drivercommon.SpaceSize {
	return drivercommon.SpaceSize{
		Total: -1,
		Usage: -1,
	}
}

func (bjh *BaiJiaHao) CheckUrl(ctx context.Context, metaurl string) bool {
	var code int
	bjh.client.HEAD(bjh.formatUrl(metaurl)).Code(&code).WithContext(ctx).Do()
	return code == 200
}
