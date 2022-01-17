package weibo

import (
	"context"

	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
)

func (b *WeiBo) Download(ctx context.Context, metaurl string) ([]byte, error) {
	var (
		data []byte
		code int
		err  error
	)
	b.client.GET(b.formatUrl(metaurl)).WithContext(ctx).Code(&code).BindBody(&data).
		Filter().Retry().Attempt(b.option.Attempt).MaxWaitTime(b.option.MaxWaitTime).WaitTime(b.option.WaitTime).
		Do()

	if code != 200 {
		return nil, drivercommon.ErrDownload
	}

	if data, err = b.option.Encoder.Decode(data); err != nil {
		return nil, err
	}
	return b.option.Crypto.Decrypt(data), nil
}
