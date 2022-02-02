package wuba

import (
	"context"
	"fmt"
)

func (wb *WuBa) Download(ctx context.Context, metaurl string) ([]byte, error) {
	var (
		data []byte
		code int
		err  error
	)
	wb.client.GET(wb.formatUrl(metaurl)).WithContext(ctx).Code(&code).BindBody(&data).
		Filter().Retry().Attempt(wb.option.Attempt).MaxWaitTime(wb.option.MaxWaitTime).WaitTime(wb.option.WaitTime).
		Do()

	if code != 200 {
		return nil, fmt.Errorf("下载失败,%s", string(data))
	}

	if data, err = wb.option.Encoder.Decode(data); err != nil {
		return nil, err
	}
	return wb.option.Crypto.Decrypt(data), nil
}
