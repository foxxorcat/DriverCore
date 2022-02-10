package weibo

import (
	"context"
	"encoding/base64"
	"regexp"
	"strings"

	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	"github.com/guonaihong/gout"
	"github.com/guonaihong/gout/dataflow"
	"github.com/guonaihong/gout/filter"
)

func (b *WeiBo) Upload(ctx context.Context, block []byte) (metaurl string, err error) {
	if len(block) > b.MaxSize() {
		err = drivercommon.ErrNoSuperBlockSize
		return
	}

	// 加密并编码块
	if block, err = b.option.Encoder.Encode(b.option.Crypto.Encrypt(block)); err != nil {
		return
	}

	var data string
	b.client.POST("https://picupload.weibo.com/interface/pic_upload.php").
		WithContext(ctx).
		SetForm(gout.H{
			"b64_data": base64.StdEncoding.EncodeToString(block),
		}).
		SetQuery(gout.H{
			"ori":  "1",
			"mime": b.option.Encoder.Mime(),
			"data": "base64",
		}).
		Filter().Retry().
		Func(func(c *dataflow.Context) error {
			c.BindBody(&data).Do()
			metaurl = strings.TrimFunc(regexp.MustCompile(`("|')\w{32}("|')`).FindString(data), func(r rune) bool { return r == '\'' || r == '"' })
			if c.Error != nil || metaurl == "" {
				err = drivercommon.ErrApiFailure
				return filter.ErrRetry
			}
			return nil
		}).
		Attempt(b.option.Attempt).MaxWaitTime(b.option.MaxWaitTime).WaitTime(b.option.WaitTime).
		Do()
	return
}
