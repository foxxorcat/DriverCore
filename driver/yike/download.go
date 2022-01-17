package yike

import (
	"context"
	"fmt"
	"strings"

	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	"github.com/guonaihong/gout"
	"github.com/guonaihong/gout/dataflow"
	"github.com/guonaihong/gout/filter"
)

func (b *YiKe) Download(ctx context.Context, metaurl string) ([]byte, error) {
	p, err := b.precreate(ctx, b.formatUrl(metaurl))
	if err != nil {
		return nil, err
	}
	if (p.ReturnType != 2 && p.ReturnType != 3) || p.Data.FsID == 0 {
		return nil, drivercommon.ErrMetaUrlFailure
	}

	var url struct {
		YiKeError
		Dlink string
	}
	b.client.POST("https://photo.baidu.com/youai/file/v2/download").
		WithContext(ctx).
		SetQuery(gout.H{"fsid": fmt.Sprint(p.Data.FsID)}).
		BindJSON(&url).
		Filter().Retry().Attempt(b.option.Attempt).MaxWaitTime(b.option.MaxWaitTime).WaitTime(b.option.WaitTime).
		Func(func(c *dataflow.Context) error {
			c.BindJSON(&url)
			if url.Errno != 0 {
				return filter.ErrRetry
			}
			return nil
		}).
		Do()

	if url.Errno != 0 || url.Dlink == "" {
		return nil, drivercommon.ErrDownload
	}

	if !strings.Contains(url.Dlink, "baidupcs.com/file") {
		return nil, fmt.Errorf("文件被屏蔽")
	}

	var (
		data []byte
		code int
	)
	b.client.GET(url.Dlink).WithContext(ctx).Code(&code).BindBody(&data).
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
