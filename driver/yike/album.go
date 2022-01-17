package yike

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"

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
func (y *YiKe) GetAlbum(ctx context.Context, name ...string) (albumout []YiKeAlbum, err error) {
	albumlist := YiKeAlbumList{}
	err = y.client.GET("https://photo.baidu.com/youai/album/v1/list").BindJSON(&albumlist).WithContext(ctx).Do()
	if albumlist.Errno != 0 {
		err = drivercommon.ErrApiFailure
	}
	for _, album := range albumlist.List {
		if sort.SearchStrings(name, album.Title) < 0 {
			continue
		}
		albumout = append(albumout, album)
	}
	return
}

// 创建相册
func (y *YiKe) CreateAlbum(ctx context.Context, name string) (album YiKeAlbum, err error) {
	albumc := YiKeAlbumCreate{}
	y.client.GET("https://photo.baidu.com/youai/album/v1/create").BindJSON(&albumc).SetQuery(gout.H{
		//"clienttype": "70",
		"title":  name,
		"source": "0",
		"tid":    getTid(),
	}).Do()

	if albumc.Errno != 0 {
		err = drivercommon.ErrApiFailure
		return
	}
	album = albumc.Info
	return
}

// 删除相册
func (y *YiKe) DeleteAlbum(ctx context.Context, album YiKeAlbum, deletesource bool) error {
	ro := 0
	if deletesource {
		ro = 1
	}

	var err YiKeError

	y.client.POST("https://photo.baidu.com/youai/album/v1/delete").
		SetForm(gout.H{
			"album_id":            album.AlbumID,
			"delete_origin_image": ro,
			"tid":                 album.Tid,
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
	v, _ := json.Marshal(fileID)
	y.client.GET("https://photo.baidu.com/youai/file/v1/delete").
		WithContext(ctx).
		SetQuery(gout.H{
			"fsid_list": string(v),
		}).
		BindJSON(&err).
		Do()

	if err.Errno != 0 {
		return fmt.Errorf("删除错误")
	}
	return nil
}
