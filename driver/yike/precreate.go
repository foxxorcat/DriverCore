package yike

import (
	"context"

	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	"github.com/guonaihong/gout"
	"github.com/guonaihong/gout/dataflow"
	"github.com/guonaihong/gout/filter"
)

type Param struct {
	Path       string
	Size       string
	SliceList  string
	BlockMd5   string
	ContentMd5 string
}

type (
	YiKePrecreateResp struct {
		YiKeError

		ReturnType int `json:"return_type"` //存在返回2 不存在返回1 已经保存3
		//存在返回
		Data YiKeFileData `json:"data"`

		//不存在返回
		Path      string  `json:"path"`
		UploadID  string  `json:"uploadid"`
		Blocklist []int64 `json:"block_list"`
	}
	YiKeFileData struct {
		FsID       int64  `json:"fs_id"`
		Size       int64  `json:"size"`
		Md5        string `json:"md5"`
		ServerName string `json:"server_filename"`
		Path       string `json:"path"`
		Ctime      int64  `json:"ctime"`
		Mtime      int64  `json:"mtime"`
		IsDir      int    `json:"isdir"`
		Category   int    `json:"category"`
		ServerMd5  string `json:"server_md5"`
		ShootTime  int64  `json:"shoot_time"`
	}
)

func (b *YiKe) precreate(ctx context.Context, param Param) (*YiKePrecreateResp, error) {
	precreateresp := YiKePrecreateResp{}
	b.client.POST("https://photo.baidu.com/youai/file/v1/precreate").
		WithContext(ctx).
		SetForm(gout.H{
			"autoinit":    "1",
			"isdir":       "0",
			"rtype":       "1",
			"ctype":       "11",
			"path":        param.Path,
			"size":        param.Size,
			"slice-md5":   param.BlockMd5,
			"content-md5": param.ContentMd5,
			"block_list":  param.SliceList,
		}).
		SetQuery(gout.H{
			"clienttype": "70",
			"bdstoken":   b.getbdstoken(),
		}).
		BindJSON(&precreateresp).
		Filter().Retry().Attempt(b.option.Attempt).MaxWaitTime(b.option.MaxWaitTime).WaitTime(b.option.WaitTime).
		Func(func(c *dataflow.Context) error {
			c.BindJSON(&precreateresp)
			if precreateresp.Errno != 0 {
				return filter.ErrRetry
			}
			return nil
		}).
		Do()
	if precreateresp.Errno != 0 || precreateresp.ReturnType == 0 {
		return nil, drivercommon.ErrApiFailure
	}
	return &precreateresp, nil
}
