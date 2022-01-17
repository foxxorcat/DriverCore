package encodercommon

type EncoderOption struct {
	//MaxSize uint
	MinSize uint

	Mode EncoderMode
}

type EncoderMode uint8

// 图片模式
const (
	RGB      EncoderMode = iota //24位
	RGBA                        //32位
	Paletted                    //8位
	Gray                        //灰度图
	CUSTOM1
)
