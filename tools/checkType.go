package tools

import (
	"encoding/hex"
	"net/http"
	"strings"
)

var fileTypeList = map[string][]string{
	".dwg": {"41433130"}, //AC10

	".html": {
		"3c21444f4354595045", // <!DOCTYPE
		"3C68746D6C3E",       // <html>
	},

	".png": {"89504E470D0A1A0A"},
	".jpg": {"FFD8FF"},
	".gif": {"47494638"},
	".bmp": {"424D"},
	".tif": {"49492A00"},

	".rar": {"52617221"},
	".zip": {"504B0304"},
}

// 获取后缀
func GetFileType(src []byte) string {
	if len(src) >= 32 {
		src = src[:32]
	}
	fileCode := strings.ToUpper(hex.EncodeToString(src))
	for filetype, codes := range fileTypeList {
		for _, code := range codes {
			if strings.HasPrefix(fileCode, code) {
				return filetype
			}
		}
	}
	return ".unknown"
}

//获取mime
func GetContentType(src []byte) string {
	return http.DetectContentType(src)
}
