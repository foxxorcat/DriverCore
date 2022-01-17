package yike

import (
	"context"
	"encoding/json"
	"fmt"
	"image"
	"net/url"
	"regexp"
	"time"

	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	"github.com/foxxorcat/DriverCore/tools"
	"github.com/google/uuid"
	"github.com/guonaihong/gout"
	"github.com/skip2/go-qrcode"
)

//判断是否登录
func (b *YiKe) IsLogin() bool {
	var yiKeLoginRes struct {
		YiKeError
		YouaID   string `json:"youa_id"`
		Nickname string
	}
	b.client.GET("https://photo.baidu.com/youai/user/v1/getuinfo").
		SetHeader(gout.H{"Referer": "https://photo.baidu.com"}).
		BindJSON(&yiKeLoginRes).
		Do()
	return yiKeLoginRes.Errno == 0 && yiKeLoginRes.YouaID != ""
}

func (b *YiKe) SetAuthorization(auto string) error {
	cookies := tools.Str2Cookie(auto)
	for _, cookie := range cookies {
		if cookie.Domain == "" {
			cookie.Domain = ".baidu.com"
		}
	}
	u, _ := url.Parse("https://photo.baidu.com")
	b.cookiejar.SetCookies(u, cookies)

	if !b.IsLogin() {
		return drivercommon.ErrLoginFail
	}
	b.bdstoken = b.getbdstoken()
	return nil
}

func (b *YiKe) QrcodeLogin(ctx context.Context, show func(ctx context.Context, image image.Image) error) (cookie string, err error) {
	gid := uuid.New().String()
	tangramGuid := getTangramGuid()

	var body []byte
	b.client.GET("https://passport.baidu.com/v2/api/getqrcode").WithContext(ctx).
		BindBody(&body).
		SetQuery(gout.H{
			"lp":          "pc",
			"qrloginfrom": "pc",
			"gid":         gid,
			"callback":    tangramGuid,
			"tt":          getTimeStamp(),
			"tpl":         "yaxsp",
			"_":           getTimeStamp(),
		}).Do()
	var qrcodeinfo struct {
		Imgurl string
		YiKeError
		Sign   string
		Prompt string
	}
	unmarshal(body, &qrcodeinfo)
	if qrcodeinfo.Errno != 0 || qrcodeinfo.Sign == "" {
		err = drivercommon.ErrQRCodeGetFail
		return
	}

	// 显示二维码
	q, _ := qrcode.New(fmt.Sprintf("https://wappass.baidu.com/wp/?qrlogin&t=%s&error=0&sign=%s&cmd=login&lp=pc&tpl=yaxsp&adapter=3&qrloginfrom=pc", getTimeStamp(), qrcodeinfo.Sign), qrcode.Medium)
	if err = show(ctx, q.Image(256)); err != nil {
		return
	}

	for {
		b.client.GET("https://passport.baidu.com/channel/unicast").
			WithContext(ctx).
			BindBody(&body).SetQuery(gout.H{
			"channel_id": qrcodeinfo.Sign,
			"tpl":        "yaxsp",
			"gid":        gid,
			"callback":   tangramGuid,
			"apiver":     "v3",
			"tt":         getTimeStamp(),
			"_":          getTimeStamp(),
		}).Do()
		var (
			qrcode struct {
				YiKeError
				ChannelID string `json:"channel_id"`
				ChannelV  string `json:"channel_v"`
			}
			channelV struct {
				Status int
				V      string
			}
		)
		unmarshal(body, &qrcode)
		json.Unmarshal(tools.Str2bytes(qrcode.ChannelV), &channelV)

		switch channelV.Status {
		case 0:
			if channelV.V == "" {
				err = drivercommon.ErrQRCodeFailure
				return
			}
			// 获取cookie自动保存道cookiejar
			b.client.GET("https://passport.baidu.com/v3/login/main/qrbdusslogin").
				WithContext(ctx).
				SetQuery(gout.H{
					"v":            getTimeStamp(),
					"bduss":        channelV.V,
					"u":            "https://photo.baidu.com/photo/web/login/fromWebLogin%3Dtrue",
					"loginVersion": "v4",
					"qrcode":       "1",
					"tpl":          "yaxsp",
					"apiver":       "v3",
					"tt":           getTimeStamp(),
					"traceid":      "",
					"time":         fmt.Sprint(time.Now().Unix()),
					"alg":          "v3",
					"callback":     "bd__cbs__srtl8i",
				}).Response()
			b.client.GET("https://photo.baidu.com/photo/web/login?fromWebLogin=true").
				WithContext(ctx).
				Response()

			u, _ := url.Parse("https://photo.baidu.com")
			cookies := b.cookiejar.Cookies(u)
			cookie = tools.Cookies2Str(cookies)
			if len(cookies) == 0 {
				err = drivercommon.ErrApiFailure
			}
			return
		case 1:
		case 2:
			err = drivercommon.ErrQRCodeFailure
			return
		default:
			//未知错误
			err = drivercommon.ErrApiFailure
			return
		}
	}
}

func getTangramGuid() string {
	return fmt.Sprintf("tangram_guid_%s", getTimeStamp())
}

//获取时间戳
func getTimeStamp() string {
	return fmt.Sprint(time.Now().UnixNano() / 1e6)
}

func unmarshal(b []byte, v interface{}) error {
	return json.Unmarshal(regexp.MustCompile(`{.*}`).Find(b), v)
}
