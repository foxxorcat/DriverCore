package bilibili

import (
	"context"
	"time"

	"github.com/foxxorcat/DriverCore/common"
	"github.com/foxxorcat/DriverCore/crypto"
	"github.com/foxxorcat/DriverCore/encoder"
)

func (b *BiLiBiLi) SetEncoder(name string) (err error) {
	switch name {
	case "pngrgb":
		b.encoder, err = encoder.NewEncoder("png", common.EncoderParam{MinSize: 400, Compress: false, Mod: "rgb"})
		b.suffix = "png"
	case "pngrgba":
		b.encoder, err = encoder.NewEncoder("png", common.EncoderParam{MinSize: 400, Compress: false, Mod: "rgba"})
		b.suffix = "png"
	case "bmp2bit":
		b.encoder, err = encoder.NewEncoder("bmp", common.EncoderParam{MinSize: 400, Mod: "2bit"})
		b.suffix = "bmp"
	default:
		return common.ErrNoSuperEncoder
	}
	return err
}

func (b *BiLiBiLi) SetCrypto(name string, param ...string) (err error) {
	b.crypto, err = crypto.NewCrypto(name, param...)
	return
}

func (b *BiLiBiLi) SetContext(ctx context.Context) common.DriverPlugin {
	b.ctx = ctx
	return b
}

func (b *BiLiBiLi) SetTimeOut(time time.Duration) common.DriverPlugin {
	b.client.SetTimeout(time)
	return b
}

func (b *BiLiBiLi) SetAttempt(t uint) common.DriverPlugin {
	b.client.SetRetryCount(int(t))
	return b
}
