package bilibili

import (
	"fmt"

	"github.com/foxxorcat/DriverCore/common"
	"github.com/foxxorcat/DriverCore/tools"
)

// 格式化链接
func (b *BiLiBiLi) formatUrl(metaurl string) string {
	return fmt.Sprintf("http://i%d.hdslb.com/bfs/album/%s", tools.RangeRand(0, 4), metaurl)
}

func (b *BiLiBiLi) Name() string {
	return "bilibili"
}

func (b *BiLiBiLi) MaxSize() int {
	return common.BlockSize4MIB
}

func (b *BiLiBiLi) SuperEncoder() []string {
	return []string{"pngrgb", "pngrgba", "bmp2bit"}
}

/*DriverUsable*/
func (b *BiLiBiLi) DownloadUsable() bool {
	return true
}

func (b *BiLiBiLi) UploadUsable() bool {
	return b.IsLogin()
}
