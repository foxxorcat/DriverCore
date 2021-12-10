package tools

import (
	"encoding/hex"
	"net/http"
	"strings"
)

var fileTypeMap map[string]string = map[string]string{
	"ffd8ffe1":             "jpg",  //JPEG (jpg)
	"ffd8ffdb":             "jpg",  //JPEG (jpg)
	"ffd8ffe0":             "jpg",  //JPEG (jpg)
	"89504e470d0a1a0a0000": "png",  //PNG (png)
	"47494638396":          "gif",  //GIF (gif)
	"49492a00227105008037": "tif",  //TIFF (tif)
	"424d3e00010000000000": "bmp",  //2色位图(bmp)
	"424d228c010000000000": "bmp",  //16色位图(bmp)
	"424d8240090000000000": "bmp",  //24位位图(bmp)
	"424d8e1b030000000000": "bmp",  //256色位图(bmp)
	"41433130313500000000": "dwg",  //CAD (dwg)
	"3c21444f435459504520": "html", //HTML (html)   3c68746d6c3e0  3c68746d6c3e0
	"3c68746d6c3e0":        "html", //HTML (html)   3c68746d6c3e0  3c68746d6c3e0
	"3c21646f637479706520": "htm",  //HTM (htm)
	"48544d4c207b0d0a0942": "css",  //css
	"696b2e71623d696b2e71": "js",   //js
	"7b5c727466315c616e73": "rtf",  // 我（IBAS）猜想的rtf
	"38425053000100000000": "psd",  //Photoshop (psd)
	"46726f6d3a203d3f6762": "eml",  //Email [Outlook Express 6] (eml)
	"d0cf11e0a1b11ae10000": "vsd",  //Visio 绘图
	"5374616E64617264204A": "mdb",  //MS Access (mdb)
	"252150532D41646F6265": "ps",
	"255044462d312e350d0a": "pdf",  //Adobe Acrobat (pdf)
	"2e524d46000000120001": "rmvb", //rmvb/rm相同
	"464c5601050000000900": "flv",  //flv与f4v相同
	"00000020667479706d70": "mp4",
	"49443303000000002176": "mp3",
	"000001b3":             "mpg", //
	"000001ba":             "mpg", //
	"3026b275":             "wmv", //wmv与asf相同
	"52494646e27807005741": "wav",
	"57415645":             "wav", //Wave (wav)
	"52494646d07d60074156": "avi",
	"4d546864000000060001": "mid", //MIDI (mid)
	//"504b03041":"zip", // 我（IBAS）看到的zip
	"504b0304140000000800": "zip", // 我（IBAS）看到的zip
	"526172211a0700cf9073": "rar", // 我（IBAS）看到的rar
	"235468697320636f6e66": "ini",
	//"504b03040":"jar"      , // 我（IBAS）看到的jar
	"4d5a9000": "exe",        //可执行文件
	"3c254020": "jsp",        //jsp文件
	"4d616e69": "mf",         //MF文件
	"3c3f786d": "xml",        //xml文件
	"494e5345": "sql",        //xml文件
	"7061636b": "java",       //java文件
	"40656368": "bat",        //bat文件
	"1f8b0800": "gz",         //gz文件
	"6c6f6734": "properties", //bat文件
	"cafebabe": "class",      //bat文件
	"49545346": "chm",        //bat文件
	"04000000": "mxp",        //bat文件
	"d0cf11e0": "wps",        //WPS文字wps、表格et、演示dps都是一样的
	"6431303a": "torrent",
	"6D6F6F76": "mov", //Quicktime (mov)
	"FF575043": "wpd", //WordPerfect (wpd)
	"CFAD12FE": "dbx", //Outlook Express (dbx)
	"2142444E": "pst", //Outlook (pst)
	"AC9EBD8F": "qdf", //Quicken (qdf)
	"E3828596": "pwl", //Windows Password (pwl)
	"2E7261FD": "ram", //Real Audio (ram)
}

//var longFileTypeMap sync.Map

func GetFileType(src []byte) string {
	var fileType string = "unknown"

	if len(src) > 0 {
		if len(src) > 8 {
			src = src[:20]
		}
		fileCode := hex.EncodeToString(src)
		for k, v := range fileTypeMap {
			if strings.HasPrefix(strings.ToLower(fileCode), strings.ToLower(k)) || strings.HasPrefix(strings.ToLower(k), strings.ToLower(fileCode)) {
				fileType = v
				break
			}
		}
	}

	return fileType
}

//最少需要512字节
func GetContentType(src []byte) string {
	return http.DetectContentType(src)
}
