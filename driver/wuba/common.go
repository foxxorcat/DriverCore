package wuba

import (
	"context"
	"fmt"

	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	encoderimage "github.com/foxxorcat/DriverCore/encoder/image"
	"github.com/foxxorcat/DriverCore/tools"
)

func (wb *WuBa) formatUrl(metaurl string) string {
	return fmt.Sprintf("https://pic%d.58cdn.com.cn/nowater/webim/big/n_v2%s%s", tools.RangeRand(1, 5), metaurl, wb.option.Encoder.Type())
}

func (wb *WuBa) Name() string {
	return NAME
}

func (wb *WuBa) MaxSize() int {
	switch wb.option.Encoder.(type) {
	case *encoderimage.Gif:
		return 14 * (2 << 19) // 16MIB
	case *encoderimage.Png:
		return 18 * (2 << 19) // 18MIB
	default:
		return 20 * (2 << 19) // 20MIB
	}
}

func (wb *WuBa) SuperEncoder() []string {
	return []string{
		encoderimage.GIFPALETTED,

		encoderimage.PNGRGB,
		encoderimage.PNGRGBA,
		encoderimage.PNGGRAY,
	}
}

func (wb *WuBa) DownloadUsable() bool {
	return true
}

func (wb *WuBa) UploadUsable() bool {
	return true
}

func (wb *WuBa) SpaceSize() drivercommon.SpaceSize {
	return drivercommon.SpaceSize{
		Total: -1,
		Usage: -1,
	}
}

func (wb *WuBa) CheckUrl(ctx context.Context, metaurl string) bool {
	var code int
	wb.client.HEAD(wb.formatUrl(metaurl)).Code(&code).WithContext(ctx).Do()
	return code == 200
}
