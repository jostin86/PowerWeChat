package redpack

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	payment "github.com/ArtisanCloud/PowerWeChat/src/payment/kernel"
	"github.com/ArtisanCloud/PowerLibs/object"
)

type Client struct {
	*payment.BaseClient
}

func NewClient(app *payment.ApplicationPaymentInterface) *Client {
	return &Client{
		payment.NewBaseClient(app),
	}
}

// Send Miniprogram Normal redpack.
// https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=18_2&index=2
func (comp *Client) SendMiniProgramNormal(params *power.HashMap) (interface{}, error) {
	config := (*comp.App).GetConfig()

	externalRequest := (*comp.App).GetExternalRequest()
	clientIP := externalRequest.Host
	if (*params)["client_ip"] != nil && (*params)["client_ip"].(string) != "" {
		clientIP = (*params)["client_ip"].(string)
	}
	base := &object.HashMap{
		"total_num": 1,
		"client_ip": clientIP,
		"wxappid":   config.GetString("app_id", ""),
	}

	options := object.MergeHashMap(base, params.ToHashMap())

	endpoint := comp.Wrap("/mmpaymkttransfers/sendminiprogramhb")
	result, err := comp.SafeRequest(endpoint, nil, "POST", options, false, nil)

	return result, err
}
