package bilibili

import (
	"context"
	"fmt"
)

func (b *BiLiBiLi) Download(ctx context.Context, metaurl string) ([]byte, error) {
	var (
		data []byte
		code int
		err  error
	)
	b.client.GET(b.formatUrl(metaurl)).WithContext(ctx).Code(&code).BindBody(&data).
		Filter().Retry().Attempt(b.option.Attempt).MaxWaitTime(b.option.MaxWaitTime).WaitTime(b.option.WaitTime).
		Do()

	if code != 200 {
		return nil, fmt.Errorf("下载失败,%s", string(data))
	}

	if data, err = b.option.Encoder.Decode(data); err != nil {
		return nil, err
	}
	return b.option.Crypto.Decrypt(data), nil
}
