package bilibili

import (
	"encoding/json"

	"github.com/foxxorcat/DriverCore/common"
	"github.com/foxxorcat/DriverCore/tools"
)

// 是否登陆
func (b *BiLiBiLi) IsLogin() bool {
	res, err := b.client.R().Get("https://api.bilibili.com/x/web-interface/nav")
	if err != nil {
		return false
	}

	var bilibiliLoginRes struct {
		Code    int
		Message string
		Ttl     int
		Data    struct {
			IsLogin       bool //是否登陆
			EmailVerified int  `json:"email_verified"` //是否验证邮箱 1验证
		}
	}
	json.Unmarshal(res.Body(), &bilibiliLoginRes)
	return bilibiliLoginRes.Data.IsLogin
}

func (b *BiLiBiLi) SetAuthorization(auto string) error {
	b.client.SetCookies(tools.Str2Cookie(auto))
	if !b.IsLogin() {
		return common.ErrLoginFail
	}
	return nil
}
