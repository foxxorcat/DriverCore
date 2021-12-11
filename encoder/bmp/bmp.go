package encoderbmp

import "github.com/foxxorcat/DriverCore/common"

const Name = "bmp"

func New(param common.EncoderParam) (common.EncoderPlugin, error) {
	switch param.Mod {
	case "2bit":
		return new(BMP2bit), nil
	default:
		return nil, common.ErrNoSuperMod
	}
}
