package encoder

import (
	"sort"

	"github.com/foxxorcat/DriverCore/common"
)

var EncoderList = []string{
	PNGALPHA,
	PNGNOTALPHA,
	BMP2BIT,
	NONE,
}

func NewEncoder(name string, param ...string) (common.EncoderPlugin, error) {
	switch name {
	case PNGALPHA:
		return new(PNGAlpha), nil
	case PNGNOTALPHA:
		return new(PNGNotAlpha), nil
	case BMP2BIT:
		return new(BMP2bit), nil
	case NONE:
		return new(None), nil
	default:
		return nil, common.ErrNotFindEncoder
	}
}

func Exist(name string) bool {
	return sort.SearchStrings(EncoderList, name) > -1
}
