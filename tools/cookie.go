package tools

import (
	"net/http"
	"strings"
)

// 字符串cookie拆分
func Str2Cookie(cookie string) (cookies []*http.Cookie) {
	header := http.Header{}
	header.Add("Cookie", cookie)
	request := http.Request{Header: header}
	return request.Cookies()

}

func Cookies2Str(cookies []*http.Cookie) string {
	var buf strings.Builder
	for _, cookie := range cookies {
		buf.WriteString(cookie.String())
		buf.WriteRune(';')
	}
	return buf.String()
}
