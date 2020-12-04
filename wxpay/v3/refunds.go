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

// 通过微信支付退款单号查询退款
func (wepay *WxPay) RefundQueryByRefundId(refundId string) (resp RefundOrderQueryResponse, err error) {
	url := fmt.Sprintf("/v3/ecommerce/refunds/id/%s?sub_mchid=%s", refundId, wepay.subMchId)
	ts, nonceStr, _, signature := wepay.getSign("GET", url, "")
	var body []byte
	body, err = wepay.doHttpRequest(fmt.Sprintf("%s%s", wePayHost, url), nonceStr, signature, "", ts)
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

// 通过商户退款单号查询退款
func (wepay *WxPay) RefundQueryByOutRefundNo(outRefundNo string) (resp RefundOrderQueryResponse, err error) {
	url := fmt.Sprintf("/v3/ecommerce/refunds/out-refund-no/%s?sub_mchid=%s", outRefundNo, wepay.subMchId)
	ts, nonceStr, _, signature := wepay.getSign("GET", url, "")
	var body []byte
	body, err = wepay.doHttpRequest(fmt.Sprintf("%s%s", wePayHost, url), nonceStr, signature, "", ts)
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
