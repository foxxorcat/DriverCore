package yike

import (
	"context"
	"encoding/base64"
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strings"
	"time"

	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	encoderimage "github.com/foxxorcat/DriverCore/encoder/image"
	"github.com/foxxorcat/DriverCore/tools"
)

type YiKeError struct {
	Errno     int `json:"errno"`
	RequestID int `json:"request_id"`
}

//Tid生成
func getTid() string {
	return fmt.Sprintf("3%d%d", time.Now().Unix(), int64(math.Floor(9000000*rand.Float64()+1000000)))
}

//LogID生成
func getLogID() string {
	return base64.RawStdEncoding.EncodeToString(tools.Str2bytes(fmt.Sprintf("%v0.%v", time.Now().UnixNano()/1e6, rand.Uint64()/1e3)))
}

//获取百度bdstoken
func (y *YiKe) getbdstoken() string {
	if y.bdstoken != "" {
		return y.bdstoken
	}
	var data string
	y.client.GET("https://photo.baidu.com").BindBody(&data).Do()
	return strings.TrimFunc(regexp.MustCompile(`("|')\w{32}("|')`).FindString(data), func(r rune) bool { return r == '\'' || r == '"' })
}

func (*YiKe) Name() string {
	return NAME
}

func (y *YiKe) MaxSize() int {
	switch y.option.Encoder.(type) {
	case *encoderimage.Gif:
		return 16 * (2 << 19) // 16MIB
	case *encoderimage.Png:
		return 28 * (2 << 19) // 28MIB
	case *encoderimage.Bmp:
		return 29 * (2 << 19) // 29MIB
	default:
		return 30 * (2 << 19) // 30MIB
	}
}

func (*YiKe) SuperEncoder() []string {
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

func (b *YiKe) DownloadUsable() bool {
	return b.IsLogin()
}

func (b *YiKe) UploadUsable() bool {
	return b.IsLogin()
}

func (b *YiKe) SpaceSize() drivercommon.SpaceSize {
	var quotainfo struct {
		YiKeError
		Quota     int64
		Used      int64
		IsUnlimit int `json:"is_unlimit"`
	}
	b.client.GET("https://photo.baidu.com/youai/user/v1/quotainfo").BindJSON(&quotainfo).Do()
	return drivercommon.SpaceSize{
		Total: func() int64 {
			if quotainfo.IsUnlimit == 1 {
				return -1
			}
			return quotainfo.Quota
		}(),
		Usage: quotainfo.Used,
	}
}

// 检查链接是否有效
func (b *YiKe) CheckUrl(ctx context.Context, metaurl string) bool {
	p, _ := b.precreate(ctx, b.formatUrl(metaurl))
	if p.ReturnType == 2 {
		b.Delete(ctx, p.Data.FsID)
	}
	return p.ReturnType == 2 || p.ReturnType == 3
}

func (b *YiKe) formatUrl(metaurl string) (param Param) {
	info := strings.SplitN(metaurl, "#", 4)
	switch len(info) {
	case 4:
		param.Path = fmt.Sprint("/", info[3])
		fallthrough
	case 3:
		if param.Path == "" {
			param.Path = fmt.Sprint("/", param.BlockMd5, b.option.Encoder.Type())
		}
		param.Size = info[2]
		param.BlockMd5 = info[0]
		param.ContentMd5 = info[1]
	}
	return
}
