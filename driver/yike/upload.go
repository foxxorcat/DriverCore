package yike

import (
	"context"
	"crypto/md5"
	"encoding/hex"
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
	sliceList, md5r := strings.Builder{}, md5.New() //md5
	sliceList.Grow((len(block)/drivercommon.BlockSize4MIB+1)*35 + 1)

	sliceList.WriteByte('[')

	for src := block; len(src) > 0; {
		i := int(math.Min(float64(len(src)), float64(drivercommon.BlockSize4MIB)))
		sliceList.WriteByte('"')
		sliceList.WriteString(tools.Md5Hex(src[:i]))
		sliceList.WriteByte('"')
		md5r.Write(src[:i])
		src = src[i:]
		if len(src) > 0 {
			sliceList.WriteByte(',')
		}
	}
	sliceList.WriteByte(']')

	md5 := hex.EncodeToString(md5r.Sum(nil))
	param := Param{
		Size:       fmt.Sprint(len(block)),
		ContentMd5: tools.Md5Hex(block[:int(math.Min(float64(len(block)), 256*1024))]),
		BlockMd5:   md5,
		SliceList:  sliceList.String(),
		Path:       fmt.Sprintf("/%s.%s", md5, b.suffix),
	}

	// 判断状态
	p, err := b.precreate(ctx, param)
	if err != nil {
		return "", err
	}

	switch p.ReturnType {
	case 1:
		// 上传块
		var reqdata struct {
			YiKeError
			Md5, Partseq, Uploadid string
		}
		src, n := block, 0
		for len(src) > 0 {
			i := int(math.Min(float64(len(src)), float64(drivercommon.BlockSize4MIB)))
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
					"partseq":    fmt.Sprint(n),
				}).
				SetForm(gout.H{
					"file": gout.FormType{File: gout.FormMem(src[:i]), FileName: "blob", ContentType: "application/octet-stream"},
				}).
				BindJSON(&reqdata).
				Filter().Retry().Attempt(b.option.Attempt).Func(
				func(c *dataflow.Context) error {
					c.BindJSON(&reqdata)
					if reqdata.Errno != 0 {
						return filter.ErrRetry
					}
					return nil
				}).MaxWaitTime(b.option.MaxWaitTime).WaitTime(b.option.WaitTime).
				Do()

			if reqdata.RequestID == 0 {
				err = drivercommon.ErrApiFailure
				return
			}
			src = src[i:]
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
			BindJSON(&precreateresp).
			Filter().Retry().Attempt(b.option.Attempt).MaxWaitTime(b.option.MaxWaitTime).WaitTime(b.option.WaitTime).
			Func(func(c *dataflow.Context) error {
				c.BindJSON(&precreateresp)
				if precreateresp.Errno != 0 {
					return filter.ErrRetry
				}
				return nil
			}).
			Do()
		if precreateresp.Errno != 0 {
			err = drivercommon.ErrApiFailure
			return
		}
		fallthrough
	case 3:

		// 返回metaurl
		metaurl = fmt.Sprintf("%s#%s#%s#%s.%s", param.BlockMd5, param.ContentMd5, param.Size, param.BlockMd5, b.suffix)
		return
	default:
		err = drivercommon.ErrApiFailure
		return
	}
}
