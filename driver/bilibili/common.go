package bilibili

import (
	"context"
	"fmt"

	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	encoderimage "github.com/foxxorcat/DriverCore/encoder/image"
	"github.com/foxxorcat/DriverCore/tools"
)

// 格式化链接
func (b *BiLiBiLi) formatUrl(metaurl string) string {
	return fmt.Sprintf("http://i%d.hdslb.com/bfs/album/%s%s", tools.RangeRand(0, 4), metaurl, b.option.Encoder.Type())
}

func (b *BiLiBiLi) Name() string {
	return NAME
}

func (b *BiLiBiLi) MaxSize() int {
	switch b.option.Encoder.(type) {
	case *encoderimage.Gif:
		return 14 * (2 << 19) // 16MIB
	case *encoderimage.Png:
		return 18 * (2 << 19) // 18MIB
	default:
		return 20 * (2 << 19) // 20MIB
	}
}

func (bjh *BiLiBiLi) SuperEncoder() []string {
	return []string{
		encoderimage.GIFPALETTED,
		encoderimage.PNGRGB,
		encoderimage.PNGRGBA,
		encoderimage.PNGPALETTED,
		encoderimage.PNGGRAY,
	}
}

func (b *BiLiBiLi) DownloadUsable() bool {
	return true
}

func (b *BiLiBiLi) UploadUsable() bool {
	return b.IsLogin()
}

func (b *BiLiBiLi) SpaceSize() drivercommon.SpaceSize {
	return drivercommon.SpaceSize{
		Total: -1,
		Usage: -1,
	}
}

// 检查链接是否有效
func (b *BiLiBiLi) CheckUrl(ctx context.Context, metaurl string) bool {
	var code int
	b.client.HEAD(b.formatUrl(metaurl)).Code(&code).WithContext(ctx).Do()
	return code == 200
}
