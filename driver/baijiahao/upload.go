package baijiahao

import (
	"context"
	"fmt"
	"strings"

	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	"github.com/foxxorcat/DriverCore/tools"
	"github.com/guonaihong/gout"
	"github.com/guonaihong/gout/dataflow"
	"github.com/guonaihong/gout/filter"
)

func (bjh *BaiJiaHao) Upload(ctx context.Context, block []byte) (md5 string, err error) {
	//判断块大小是否支持
	if len(block) > bjh.MaxSize() {
		err = drivercommon.ErrNoSuperBlockSize
		return
	}

	// 加密并编码块
	if block, err = bjh.option.Encoder.Encode(bjh.option.Crypto.Encrypt(block)); err != nil {
		return
	}

	md5 = tools.Md5Hex(block) // 计算文件md5

	if !bjh.CheckUrl(ctx, md5) {
		var data string
		bjh.client.POST("https://baijiahao.baidu.com/builderinner/api/content/file/upload").
			WithContext(ctx).
			SetForm(gout.H{
				"no_compress": "1",
				"id":          "WU_FILE_0",
				"is_avatar":   "0",
				"media":       gout.FormType{File: gout.FormMem(block), FileName: fmt.Sprint(md5, bjh.option.Encoder.Type()), ContentType: bjh.option.Encoder.Mime()},
			}).
			Filter().Retry().
			Func(func(c *dataflow.Context) error {
				c.BindBody(&data).Do()
				if c.Error != nil || !strings.Contains(data, "success") {
					err = drivercommon.ErrApiFailure
					return filter.ErrRetry
				}
				return nil
			}).
			Attempt(bjh.option.Attempt).MaxWaitTime(bjh.option.MaxWaitTime).WaitTime(bjh.option.WaitTime).
			Do()
	}
	return
}
