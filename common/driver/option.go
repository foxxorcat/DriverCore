package drivercommon

import (
	"time"

	cryptocommon "github.com/foxxorcat/DriverCore/common/crypto"
	encodercommon "github.com/foxxorcat/DriverCore/common/encoder"
)

type DriverOption struct {
	Timeout time.Duration // 超时时间

	Attempt     int           // 重试次数
	WaitTime    time.Duration // 重试等待时间
	MaxWaitTime time.Duration // 重试最长等待时间

	Encoder encodercommon.EncoderPlugin // 编码器

	Crypto cryptocommon.CryptoPlugin // 加密器
}

func (b *DriverOption) SetOption(options ...Option) error {
	var err error
	for _, option := range options {
		if err = option.Apply(b); err != nil {
			return err
		}
	}
	return nil
}

type Option interface {
	Apply(*DriverOption) error
}

// 设置重试
type attempt int

func WithAttempt(t int) Option {
	return (*attempt)(&t)
}

func (t *attempt) Apply(o *DriverOption) error {
	o.Attempt = int(*t)
	return nil
}

// 设置超时
type timeout time.Duration

func WithTimeout(t time.Duration) Option {
	return (*timeout)(&t)
}

func (t *timeout) Apply(o *DriverOption) error {
	o.Timeout = time.Duration(*t)
	return nil
}

// 设置重试等待时间
type waitTime time.Duration

func WithWaitTime(t time.Duration) Option {
	return (*waitTime)(&t)
}

func (t waitTime) Apply(o *DriverOption) error {
	o.WaitTime = time.Duration(t)
	return nil
}

// 设置重试最大等待时间
type maxWaitTime time.Duration

func WithMaxWaitTime(t time.Duration) Option {
	return (*maxWaitTime)(&t)
}

func (t maxWaitTime) Apply(o *DriverOption) error {
	o.MaxWaitTime = time.Duration(t)
	return nil
}

// 设置encoder
type encoder struct {
	encodercommon.EncoderPlugin
}

func WithEncoder(t encodercommon.EncoderPlugin) Option {
	return &encoder{t}
}

func (t *encoder) Apply(o *DriverOption) error {
	if t.EncoderPlugin != nil {
		o.Encoder = t.EncoderPlugin
	}
	return nil
}

// 设置crypto
type crypto struct {
	cryptocommon.CryptoPlugin
}

func WithCrypto(t cryptocommon.CryptoPlugin) Option {
	return &crypto{t}
}

func (t *crypto) Apply(o *DriverOption) error {
	if t.CryptoPlugin != nil {
		o.Crypto = t.CryptoPlugin
	}
	return nil
}
