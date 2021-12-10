package bilibili

import (
	"DriverCore/common"
	"DriverCore/crypto"
	"DriverCore/encoder"
	"context"
	"sort"
	"time"
)

func (b *BiLiBiLi) SetEncoder(name string) (err error) {
	if sort.SearchStrings(b.SuperEncoder(), name) >= 0 {
		b.encoder, err = encoder.NewEncoder(name, "10")
		return
	}
	return common.ErrNoSuperEncoder
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
