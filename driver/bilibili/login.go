package bilibili

import (
	"context"
	"encoding/json"
	"image"
	"net/http"
	"net/url"
	"regexp"
	"time"

	drivercommon "github.com/foxxorcat/DriverCore/common/driver"
	"github.com/foxxorcat/DriverCore/tools"
	"github.com/guonaihong/gout"
	"github.com/skip2/go-qrcode"
)

// 是否登陆
func (b *BiLiBiLi) IsLogin() bool {
	var bilibiliLoginRes struct {
		Code    int
		Message string
		Ttl     int
		Data    struct {
			IsLogin       bool //是否登陆
			EmailVerified int  `json:"email_verified"` //是否验证邮箱 1验证
		}
	}

	b.client.GET("https://api.bilibili.com/x/web-interface/nav").BindJSON(&bilibiliLoginRes).Do()
	return bilibiliLoginRes.Data.IsLogin
}

func (b *BiLiBiLi) SetAuthorization(auto string) error {
	cookies := tools.Str2Cookie(auto)
	for _, cookie := range cookies {
		if cookie.Domain == "" {
			cookie.Domain = ".bilibili.com"
		}
	}
	u, _ := url.Parse("https://bilibili.com")
	b.cookiejar.SetCookies(u, cookies)

	if !b.IsLogin() {
		return drivercommon.ErrLoginFail
	}
	return nil
}

//二维码登录
func (b *BiLiBiLi) QrcodeLogin(ctx context.Context, show func(ctx context.Context, image image.Image) error) (cookie string, err error) {

	// 获取二维码
	var qrres struct {
		Code   int
		Status bool
		Ts     int64
		Data   struct {
			Url      string
			OauthKey string
		}
	}
	b.client.GET("https://passport.bilibili.com/qrcode/getLoginUrl").WithContext(ctx).BindJSON(&qrres).Do()
	if qrres.Code != 0 || !qrres.Status {
		err = drivercommon.ErrQRCodeGetFail
		return
	}

	// 显示二维码
	q, _ := qrcode.New(qrres.Data.Url, qrcode.Medium) // 生成二维码
	if err = show(ctx, q.Image(256)); err != nil {
		return
	}

	// 检查二维码状态
	var body []byte
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		b.client.POST("https://passport.bilibili.com/qrcode/getLoginInfo").
			SetWWWForm(gout.H{
				"oauthKey": qrres.Data.OauthKey,
				"gourl":    "https://www.bilibili.com/",
			}).WithContext(ctx).BindBody(&body).Do()

		//未确认登陆结构
		var qrcoderes struct {
			Status  bool
			Data    int
			Message string
		}
		json.Unmarshal(body, &qrcoderes)

		switch qrcoderes.Data {
		case -2, -1:
			err = drivercommon.ErrQRCodeFailure
			return
		case -4:
			// 等待扫描二维码
			time.Sleep(time.Second * 2)
		case -5:
			// 手机确认登陆
			time.Sleep(time.Second * 2)
		default:
			if qrcoderes.Status {
				sessdata := regexp.MustCompile("SESSDATA=([^&]*)").FindStringSubmatch(string(body))
				if len(sessdata) >= 1 {
					return tools.Cookies2Str([]*http.Cookie{{Name: "SESSDATA", Value: sessdata[1]}}), nil
				}
			}
			err = drivercommon.ErrApiFailure
			return
		}
	}
}
