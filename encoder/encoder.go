package encoder

import (
	"github.com/foxxorcat/DriverCore/common"
	encoderbmp "github.com/foxxorcat/DriverCore/encoder/bmp"
	encoderpng "github.com/foxxorcat/DriverCore/encoder/png"
)

func NewEncoder(name string, param common.EncoderParam) (common.EncoderPlugin, error) {
	switch name {
	case encoderpng.Name:
		return encoderpng.New(param)
	case encoderbmp.Name:
		return encoderbmp.New(param)
	case "none":
		return new(None), nil
	default:
		return nil, common.ErrNotFindEncoder
	}
}
