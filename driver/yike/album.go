package yike

import (
	"context"
	"encoding/json"

	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	"github.com/guonaihong/gout"
)

type (
	YiKeAlbumList struct {
		List []YiKeAlbum `json:"list"`
		YiKeError
	}

	YiKeAlbumCreate struct {
		AlbumID string    `json:"album_id"`
		Info    YiKeAlbum `json:"info"`
		YiKeError
	}

	YiKeAlbum struct {
		AppID   int    `json:"app_id"`
		AlbumID string `json:"album_id"`
		Tid     int64  `json:"tid"`
		Title   string `json:"title"`
	}
)

// 获取相册
func (y *YiKe) GetAlbum(ctx context.Context, names ...string) (albumout []YiKeAlbum, err error) {
	var albumlist YiKeAlbumList
	y.client.GET("https://photo.baidu.com/youai/album/v1/list").BindJSON(&albumlist).WithContext(ctx).Do()
	if albumlist.Errno != 0 || albumlist.RequestID == 0 {
		err = drivercommon.ErrApiFailure
		return
	}
	for _, album := range albumlist.List {
		for _, name := range names {
			if name == album.Title {
				albumout = append(albumout, album)
				break
			}
		}
	}
	return
}

// 创建相册
func (y *YiKe) CreateAlbum(ctx context.Context, name string) (*YiKeAlbum, error) {
	var albumc YiKeAlbumCreate
	y.client.GET("https://photo.baidu.com/youai/album/v1/create").BindJSON(&albumc).SetQuery(gout.H{
		"clienttype": "70",
		"title":      name,
		"source":     "0",
		"tid":        getTid(),
	}).Do()

	if albumc.Errno != 0 || albumc.RequestID == 0 {
		return nil, drivercommon.ErrApiFailure
	}
	return &albumc.Info, nil
}

// 删除相册
func (y *YiKe) DeleteAlbum(ctx context.Context, album *YiKeAlbum, deletesource bool) error {
	var err YiKeError
	y.client.POST("https://photo.baidu.com/youai/album/v1/delete").
		SetForm(gout.H{
			"album_id": album.AlbumID,
			"delete_origin_image": func() int {
				if deletesource {
					return 1
				}
				return 0
			}(),
			"tid": album.Tid,
		}).
		BindJSON(&err).
		Do()
	if err.Errno != 0 || err.RequestID == 0 {
		return drivercommon.ErrApiFailure
	}
	return nil
}

//删除文件
func (y *YiKe) Delete(ctx context.Context, fileID ...int64) error {
	var err YiKeError
	data, _ := json.Marshal(fileID)
	y.client.GET("https://photo.baidu.com/youai/file/v1/delete").
		WithContext(ctx).
		SetQuery(gout.H{
			"fsid_list": string(data),
		}).
		BindJSON(&err).
		Do()

	if err.Errno != 0 || err.RequestID == 0 {
		return drivercommon.ErrApiFailure
	}
	return nil
}
