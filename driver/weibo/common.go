package weibo

import (
	"context"
	"fmt"

	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	encoderimage "github.com/foxxorcat/DriverCore/encoder/image"
	"github.com/foxxorcat/DriverCore/tools"
)

// 格式化链接
func (b *WeiBo) formatUrl(metaurl string) string {
	return fmt.Sprintf("https://wx%d.sinaimg.cn/large/%s", tools.RangeRand(1, 5), metaurl)
}

func (*WeiBo) Name() string {
	return NAME
}

func (*WeiBo) MaxSize() int {
	return 16 * (2 << 19) // 16MIB
}

func (*WeiBo) SuperEncoder() []string {
	return []string{
		encoderimage.GIFPALETTED,
	}
}

func (*WeiBo) DownloadUsable() bool {
	return true
}

func (b *WeiBo) UploadUsable() bool {
	return b.IsLogin()
}

func (*WeiBo) SpaceSize() drivercommon.SpaceSize {
	return drivercommon.SpaceSize{
		Total: -1,
		Usage: -1,
	}
}

// 检查链接是否有效
func (b *WeiBo) CheckUrl(ctx context.Context, metaurl string) bool {
	var code int
	b.client.HEAD(b.formatUrl(metaurl)).Code(&code).WithContext(ctx).Do()
	return code == 200
}
