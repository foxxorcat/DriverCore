package bilibili

import (
	"context"
	"fmt"

	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	"github.com/foxxorcat/DriverCore/tools"
	"github.com/guonaihong/gout"
	"github.com/guonaihong/gout/dataflow"
	"github.com/guonaihong/gout/filter"
)

func (b *BiLiBiLi) Upload(ctx context.Context, block []byte) (sha1 string, err error) {
	if len(block) > b.MaxSize() {
		err = drivercommon.ErrNoSuperBlockSize
		return
	}

	// 加密并编码块
	if block, err = b.option.Encoder.Encode(b.option.Crypto.Encrypt(block)); err != nil {
		return
	}

	sha1 = tools.SHA1Hex(block)

	// 检查是否已经上传
	if !b.CheckUrl(ctx, sha1) {
		var bilires struct {
			Code    int
			Message string
		}
		b.client.POST("https://api.vc.bilibili.com/api/v1/drawImage/upload").
			WithContext(ctx).
			SetForm(gout.H{
				"biz":      "draw",
				"category": "daily",
				"file_up":  gout.FormType{File: gout.FormMem(block), FileName: fmt.Sprint(sha1, b.option.Encoder.Type()), ContentType: b.option.Encoder.Mime()},
			}).
			Filter().Retry().
			Func(func(c *dataflow.Context) error {
				c.BindJSON(&bilires).Do()
				if c.Error != nil || bilires.Code != 0 || bilires.Message == "" {
					err = drivercommon.ErrApiFailure
					return filter.ErrRetry
				}
				return nil
			}).Attempt(b.option.Attempt).MaxWaitTime(b.option.MaxWaitTime).WaitTime(b.option.WaitTime).
			Do()
	}
	return
}
