package weibo

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
	"github.com/guonaihong/gout"
	"github.com/skip2/go-qrcode"
)

// 是否登陆
func (b *WeiBo) IsLogin() bool {
	var data struct {
		Code string
	}
	b.client.GET("https://weibo.com/aj/onoff/setstatus").
		BindJSON(&data).Do()
	return data.Code == "100000"
}

func (b *WeiBo) SetAuthorization(auto string) error {
	cookies := tools.Str2Cookie(auto)
	for _, cookie := range cookies {
		if cookie.Domain == "" {
			cookie.Domain = ".weibo.com"
		}
	}
	u, _ := url.Parse("https://weibo.com")
	b.cookiejar.SetCookies(u, cookies)

	if !b.IsLogin() {
		return drivercommon.ErrLoginFail
	}
	return nil
}

func (b *WeiBo) QrcodeLogin(ctx context.Context, show func(ctx context.Context, image image.Image) error) (cookie string, err error) {
	type WeiBoQrcodeResponse struct {
		Retcode int    `json:"retcode"`
		Msg     string `json:"msg"`
		Data    struct {
			Qrid  string `json:"qrid"`
			Image string `json:"image"`

			Alt string `json:"alt"`
		} `json:"data"`
	}

	// 获取二维码信息
	var (
		body        []byte
		qrcodeinfo  WeiBoQrcodeResponse
		checkqrcode WeiBoQrcodeResponse
	)
	b.client.GET("https://login.sina.com.cn/sso/qrcode/image").
		WithContext(ctx).
		SetHeader(gout.H{
			"Referer": "https://weibo.com/",
		}).
		BindBody(&body).SetQuery(gout.H{
		"service_id": "pc_protection",
		"entry":      "sso",
		"size":       "300",
		"callback":   getTimeStamp(),
	}).Do()

	unmarshal(body, &qrcodeinfo)
	if qrcodeinfo.Retcode != 20000000 || qrcodeinfo.Msg != "succ" {
		err = drivercommon.ErrQRCodeGetFail
		return
	}

	// 显示二维码
	q, _ := qrcode.New(fmt.Sprintf("https://passport.weibo.cn/signin/qrcode/scan?qr=%s&sinainternalbrowser=topnav&showmenu=0", qrcodeinfo.Data.Qrid), qrcode.Medium) // 生成二维码
	if err = show(ctx, q.Image(256)); err != nil {
		return
	}

	// 等待扫描
	for {
		b.client.GET("https://login.sina.com.cn/sso/qrcode/check").
			WithContext(ctx).
			BindBody(&body).SetQuery(gout.H{
			"entry":    "sso",
			"qrid":     qrcodeinfo.Data.Qrid,
			"callback": getTimeStamp(),
		}).Do()
		unmarshal(body, &checkqrcode)

		switch checkqrcode.Retcode {
		case 50114001:
			//等待扫码
			time.Sleep(time.Second)
		case 50114002:
			//等待确认
			time.Sleep(time.Second)
		case 50114004:
			//超时
			err = drivercommon.ErrQRCodeFailure
			return
		case 20000000:
			//成功
			b.client.GET("https://login.sina.com.cn/sso/login.php").
				WithContext(ctx).
				BindBody(&body).
				SetQuery(gout.H{
					"entry":       "weibo",
					"domain":      "weibo.com",
					"cdult":       "3",
					"crossdomain": "1",
					"returntype":  "TEXT",
					"alt":         checkqrcode.Data.Alt,
					"savestate":   "30",
					"callback":    getTimeStamp(),
				}).Do()

			// 解析json
			var reqinfo struct {
				Retcode            int
				Uid                string
				Nick               string
				CrossDomainUrlList []string
			}
			unmarshal(body, &reqinfo)

			for _, url := range reqinfo.CrossDomainUrlList {
				b.client.GET(url).
					WithContext(ctx).
					SetHeader(gout.H{
						"Referer": "https://weibo.com/",
					}).
					Do()
			}
			u, _ := url.Parse("https://weibo.com")
			cookies := b.cookiejar.Cookies(u)
			cookie = tools.Cookies2Str(cookies)
			if len(cookies) == 0 || reqinfo.Retcode != 0 {
				err = drivercommon.ErrApiFailure
			}
			return
		default:
			//未知错误
			err = drivercommon.ErrApiFailure
			return
		}
	}

}

func unmarshal(b []byte, ret interface{}) error {
	return json.Unmarshal(regexp.MustCompile(`{.*}`).Find(b), ret)
}

//获取时间戳
func getTimeStamp() string {
	return fmt.Sprintf("STK_%d", time.Now().UnixNano()/1e5)
}
