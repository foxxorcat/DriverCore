package baijiahao

import (
	"context"
	"fmt"
	"strings"

	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	"github.com/foxxorcat/DriverCore/tools"
	"github.com/guonaihong/gout"
)

func (bjh *BaiJiaHao) Upload(ctx context.Context, block []byte) (string, error) {
	//判断块大小是否支持
	if len(block) > bjh.MaxSize() {
		return "", drivercommon.ErrNoSuperBlockSize
	}

	var err error
	// 加密并编码块
	if block, err = bjh.option.Encoder.Encode(bjh.option.Crypto.Encrypt(block)); err != nil {
		return "", err
	}

	md5 := tools.Md5Hex(block) // 计算文件md5

	if !bjh.CheckUrl(ctx, md5) {
		var body []byte
		if err = bjh.client.POST("https://baijiahao.baidu.com/builderinner/api/content/file/upload").
			WithContext(ctx).
			SetForm(gout.H{
				"no_compress": "1",
				"id":          "WU_FILE_0",
				"is_avatar":   "0",
				"media":       gout.FormMem(block),
			}).BindBody(&body).
			Filter().Retry().Attempt(bjh.option.Attempt).MaxWaitTime(bjh.option.MaxWaitTime).WaitTime(bjh.option.WaitTime).
			Do(); err != nil {
			return "", err
		}

		if !strings.Contains(string(body), "success") {
			return "", fmt.Errorf(string(body))
		}
	}
	return md5, nil
}
