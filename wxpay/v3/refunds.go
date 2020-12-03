package v3

import (
	"encoding/json"
	"errors"
	"fmt"
)

// 订单退款
func (wepay *WxPay) Refund(refund RefundsOrder) (resp RefundOrderResponse, err error) {
	refund.SpAppId = wepay.spAppId
	refund.SubMchId = wepay.subMchId
	if wepay.subAppId != "" {
		refund.SubAppId = wepay.subAppId
	}
	err = validate.Struct(refund)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	url := "/v3/ecommerce/refunds/apply"
	signBody, _ := json.Marshal(refund)
	ts, nonceStr, _, signature := wepay.getSign("POST", url, string(signBody))
	var body []byte
	body, err = wepay.doHttpRequest(fmt.Sprintf("%s%s", wePayHost, url), nonceStr, signature, string(signBody), ts)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	if resp.Code != "" {
		return resp, errors.New(resp.Message)
	}
	return
}
