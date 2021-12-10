package tools

import (
	"net/http"
	"regexp"
	"strings"
)

// 字符串cookie拆分
func Str2Cookie(cookie string) (cookies []*http.Cookie) {
	for _, cookie := range regexp.MustCompile(`(\S+)=([^;]+)`).FindAllStringSubmatch(cookie, -1) {
		cookies = append(cookies, &http.Cookie{
			Name:  cookie[1],
			Value: cookie[2],
		})
	}
	return
}

func Cookies2Str(cookies []*http.Cookie) string {
	buf := new(strings.Builder)
	for _, cookie := range cookies {
		if buf.Len() > 0 {
			buf.WriteRune(';')
		}
		buf.WriteString(cookie.String())
	}
	return buf.String()
}
