package bilibili

import "fmt"

func (b *BiLiBiLi) Download(metaurl string) ([]byte, error) {
	res, err := b.client.R().SetContext(b.ctx).Get(b.formatUrl(metaurl))
	if err != nil {
		return nil, fmt.Errorf("下载失败,错误信息：%s", err)
	}

	if res.StatusCode() != 200 {
		return nil, fmt.Errorf("下载失败,%s", res.String())
	}

	data, err := b.encoder.Decode(res.Body())
	if err != nil {
		return nil, err
	}
	return b.crypto.Decrypt(data), nil
}
