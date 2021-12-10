package bilibili

import (
	"DriverCore/common"
	"DriverCore/encoder"
	"DriverCore/tools"
	"fmt"
)

// 格式化链接
func (b *BiLiBiLi) formatUrl(sha1 string) string {
	return fmt.Sprintf("http://i%d.hdslb.com/bfs/album/%s", tools.RangeRand(0, 4), sha1)
}

func (b *BiLiBiLi) Name() string {
	return Name
}

func (b *BiLiBiLi) MaxSize() int {
	return common.BlockSize4MIB
}

func (b *BiLiBiLi) SuperEncoder() []string {
	return []string{encoder.BMP2BIT, encoder.PNGALPHA, encoder.PNGNOTALPHA}
}

/*DriverUsable*/
func (b *BiLiBiLi) DownloadUsable() bool {
	return true
}

func (b *BiLiBiLi) UploadUsable() bool {
	return b.IsLogin()
}
