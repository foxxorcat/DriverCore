package bilibili

import (
	"context"

	"github.com/foxxorcat/DriverCore/common"

	"github.com/go-resty/resty/v2"
)

var Info common.DriverPluginInfo = new(BiLiBiLi)

type BiLiBiLi struct {
	client  *resty.Client
	ctx     context.Context
	encoder common.EncoderPlugin
	suffix  string
	crypto  common.CryptoPlugin
}
