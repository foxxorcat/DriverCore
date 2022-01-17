package baijiahao

import (
	"context"
	"fmt"
)

func (bjh *BaiJiaHao) Download(ctx context.Context, metaurl string) ([]byte, error) {
	var (
		data []byte
		code int
	)
	err := bjh.client.GET(bjh.formatUrl(metaurl)).WithContext(ctx).Code(&code).BindBody(&data).
		Filter().Retry().Attempt(bjh.option.Attempt).MaxWaitTime(bjh.option.MaxWaitTime).WaitTime(bjh.option.WaitTime).
		Do()

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("下载失败,%s", string(data))
	}

	if data, err = bjh.option.Encoder.Decode(data); err != nil {
		return nil, err
	}
	return bjh.option.Crypto.Decrypt(data), nil
}
