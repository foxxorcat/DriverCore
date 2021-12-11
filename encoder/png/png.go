package encoderpng

import "github.com/foxxorcat/DriverCore/common"

const Name = "png"

func New(param common.EncoderParam) (common.EncoderPlugin, error) {
	switch param.Mod {
	case "rgb":
		return new(PNGAlpha), nil
	case "rgba":
		return new(PNGNotAlpha), nil
	default:
		return nil, common.ErrNoSuperMod
	}
}
