package encodercommon

import "math"

type EncoderOption struct {
	//MaxSize uint
	MinSize int
}

func (b *EncoderOption) SetOption(options ...Option) error {
	var err error
	for _, option := range options {
		if err = option.Apply(b); err != nil {
			return err
		}
	}
	return nil
}

type Option interface {
	Apply(*EncoderOption) error
}

type minsize int

// 设置最小大小，如果为负值则设置为相对最小
func WithMinSize(t int) Option {
	return (*minsize)(&t)
}

func (t minsize) Apply(o *EncoderOption) error {
	if t < 0 {
		o.MinSize = int(math.Min(float64(o.MinSize), float64(-t)))
	}
	o.MinSize = int(t)
	return nil
}
