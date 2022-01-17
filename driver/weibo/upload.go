package weibo

import (
	"context"
	"encoding/base64"
	"regexp"

	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	"github.com/guonaihong/gout"
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

	var body []byte
	if err = b.client.POST("https://picupload.weibo.com/interface/pic_upload.php").
		WithContext(ctx).
		SetForm(gout.H{
			"b64_data": base64.StdEncoding.EncodeToString(block),
		}).
		SetQuery(gout.H{
			"ori": "1",
			//"mime": tools.GetContentType(block),
			"data": "base64",
		}).
		BindBody(&body).
		Filter().Retry().Attempt(b.option.Attempt).MaxWaitTime(b.option.MaxWaitTime).WaitTime(b.option.WaitTime).
		Do(); err != nil {
		return
	}

	metaurl = string(regexp.MustCompile(`\w{32}`).Find(body))
	if metaurl == "" {
		err = drivercommon.ErrApiFailure
	}
	return
}
