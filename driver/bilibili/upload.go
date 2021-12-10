package bilibili

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/foxxorcat/DriverCore/common"
	"github.com/foxxorcat/DriverCore/encoder"
	"github.com/foxxorcat/DriverCore/tools"
)

func (b *BiLiBiLi) Upload(block []byte) (string, error) {
	if len(block) > b.MaxSize() {
		return "", common.ErrNoSuperBlockSize
	}

	block, err := b.encoder.Encoded(b.crypto.Encrypt(block))
	if err != nil {
		return "", err
	}

	var suffix string
	switch b.encoder.Name() {
	case encoder.PNGALPHA, encoder.PNGNOTALPHA:
		suffix = "png"
	case encoder.BMP2BIT:
		suffix = "bmp"
	}

	url := fmt.Sprintf("%s.%s", tools.SHA1Hex(block), suffix)
	// 检查是否已经上传
	if !b.CheckUrl(url) {
		res, err := b.client.R().
			SetContext(b.ctx).
			SetFormData(map[string]string{
				"biz":      "draw",
				"category": "daily",
			}).
			SetFileReader("file_up", url, bytes.NewReader(block)).
			Post("https://api.vc.bilibili.com/api/v1/drawImage/upload")
		if err != nil {
			return "", err
		}

		var bilires struct {
			Code    int
			Message string
		}

		json.Unmarshal(res.Body(), &bilires)
		if bilires.Code != 0 {
			return "", fmt.Errorf(bilires.Message)
		}
	}
	return url, nil
}
