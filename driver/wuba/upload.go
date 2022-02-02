package wuba

import (
	"context"
	"encoding/base64"
	"regexp"

	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	"github.com/foxxorcat/DriverCore/tools"
	"github.com/guonaihong/gout/dataflow"
	"github.com/guonaihong/gout/filter"
)

func (wb *WuBa) Upload(ctx context.Context, block []byte) (metaurl string, err error) {
	//判断块大小是否支持
	if len(block) > wb.MaxSize() {
		err = drivercommon.ErrNoSuperBlockSize
		return
	}

	// 加密并编码块
	if block, err = wb.option.Encoder.Encode(wb.option.Crypto.Encrypt(block)); err != nil {
		return
	}

	buf := make([]byte, base64.RawStdEncoding.EncodedLen(len(block))+93)
	copy(buf[:91], tools.Str2bytes(`{"Pic-Size": "0*0","Pic-Encoding": "base64","Pic-Path": "/nowater/webim/big/","Pic-Data": "`))
	base64.RawStdEncoding.Encode(buf[91:], block)
	copy(buf[len(buf)-2:], tools.Str2bytes(`"}`))

	var data string
	wb.client.POST("https://upload.58cdn.com.cn/json").
		SetHeader(map[string]string{
			"Referer": "https://ai.58.com/pc/",
		}).
		WithContext(ctx).
		SetBody(buf).
		Filter().Retry().
		Func(func(c *dataflow.Context) error {
			c.BindBody(&data).Do()
			metaurl = regexp.MustCompile(`\w{36}`).FindString(data)
			if c.Error != nil || metaurl == "" {
				err = drivercommon.ErrApiFailure
				return filter.ErrRetry
			}
			// 去除 n_v2
			metaurl = metaurl[4:]
			return nil
		}).
		Attempt(wb.option.Attempt).MaxWaitTime(wb.option.MaxWaitTime).WaitTime(wb.option.WaitTime).
		Do()
	return
}
