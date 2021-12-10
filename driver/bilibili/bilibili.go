package bilibili

import (
	"DriverCore/common"
	"context"

	"github.com/go-resty/resty/v2"
)

const Name = "bilibili"

type BiLiBiLi struct {
	client  *resty.Client
	ctx     context.Context
	encoder common.EncoderPlugin
	crypto  common.CryptoPlugin
}
