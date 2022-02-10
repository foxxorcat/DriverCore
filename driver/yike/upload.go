package yike

import (
	"context"
	"fmt"
	"math"
	"strings"

	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	"github.com/foxxorcat/DriverCore/tools"
	"github.com/guonaihong/gout"
	"github.com/guonaihong/gout/dataflow"
	"github.com/guonaihong/gout/filter"
)

func (b *YiKe) Upload(ctx context.Context, block []byte) (metaurl string, err error) {
	if len(block) > b.MaxSize() {
		err = drivercommon.ErrNoSuperBlockSize
		return
	}

	// 加密并编码块
	if block, err = b.option.Encoder.Encode(b.option.Crypto.Encrypt(block)); err != nil {
		return
	}

	// 处理上传需要的信息
	sliceList, src := strings.Builder{}, block

	sliceList.WriteByte('[')
	for len(src) > 0 {
		n := int(math.Min(float64(len(src)), float64(drivercommon.BlockSize4MIB)))
		sliceList.WriteByte('"')
		sliceList.WriteString(tools.Md5Hex(src[:n]))
		sliceList.WriteByte('"')
		src = src[n:]
		if len(src) > 0 {
			sliceList.WriteByte(',')
		}
	}
	sliceList.WriteByte(']')

	md5 := tools.Md5Hex(block)
	param := Param{
		Size:       fmt.Sprint(len(block)),
		ContentMd5: tools.Md5Hex(block[:int(math.Min(float64(len(block)), 256*1024))]),
		BlockMd5:   md5,
		SliceList:  sliceList.String(),
		Path:       fmt.Sprint("/", md5, b.option.Encoder.Type()),
	}

	// 判断文件状态
	p, _ := b.precreate(ctx, param)

	switch p.ReturnType {
	case 1:
		// 上传块
		var reqdata struct {
			YiKeError
			Md5, Partseq, Uploadid string
		}

		for src, i := block, 0; len(src) > 0; i++ {
			n := int(math.Min(float64(len(src)), float64(drivercommon.BlockSize4MIB))) //块最大4MIB
			b.client.POST("https://c3.pcs.baidu.com/rest/2.0/pcs/superfile2").
				WithContext(ctx).
				SetQuery(gout.H{
					"method":     "upload",
					"app_id":     "16051585",
					"channel":    "chunlei",
					"clienttype": "70",
					"web":        "1",
					"logid":      getLogID(),
					"path":       param.Path,
					"uploadid":   p.UploadID,
					"partseq":    fmt.Sprint(i),
				}).
				SetForm(gout.H{
					"file": gout.FormType{File: gout.FormMem(src[:n]), FileName: "blob", ContentType: "application/octet-stream"},
				}).
				Filter().Retry().Attempt(b.option.Attempt).
				Func(func(c *dataflow.Context) error {
					c.BindJSON(&reqdata).Do()
					if c.Error != nil || reqdata.Errno != 0 || reqdata.RequestID == 0 {
						err = drivercommon.ErrApiFailure
						return filter.ErrRetry
					}
					return nil
				}).MaxWaitTime(b.option.MaxWaitTime).WaitTime(b.option.WaitTime).
				Do()
			if err != nil {
				return
			}
			src = src[n:]
		}
		fallthrough
	case 2:
		// 创建文件
		precreateresp := YiKePrecreateResp{}
		b.client.POST("https://photo.baidu.com/youai/file/v1/create").
			WithContext(ctx).
			SetForm(gout.H{
				"block_list":  param.SliceList,
				"isdir":       "0",
				"rtype":       "1",
				"ctype":       "11",
				"path":        param.Path,
				"size":        param.Size,
				"slice-md5":   param.BlockMd5,
				"content-md5": param.ContentMd5,
				"uploadid":    p.UploadID,
			}).
			SetQuery(gout.H{
				"clienttype": "70",
				"bdstoken":   b.getbdstoken(),
			}).
			Filter().Retry().Attempt(b.option.Attempt).
			Func(func(c *dataflow.Context) error {
				c.BindJSON(&precreateresp).Do()
				if c.Error != nil || precreateresp.Errno != 0 || precreateresp.RequestID == 0 {
					err = drivercommon.ErrApiFailure
					return filter.ErrRetry
				}
				return nil
			}).MaxWaitTime(b.option.MaxWaitTime).WaitTime(b.option.WaitTime).
			Do()
		if err != nil {
			return
		}
		fallthrough
	case 3:
		// 返回metaurl
		metaurl = fmt.Sprintf("%s#%s#%s", param.BlockMd5, param.ContentMd5, param.Size)
		return
	default:
		err = drivercommon.ErrApiFailure
		return
	}
}
