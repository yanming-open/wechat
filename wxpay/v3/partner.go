package v3

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/yanming-open/wechat/utils"
)

// 验证订单内容是否合法
// TODO: 级联下required字段处理
func (wepay *wxPay) validatePartnerOrder(order *PartnerOrder) (err error) {
	order.SpAppId = wepay.spAppId
	order.SpMchId = wepay.spMchId
	order.SubMchId = wepay.subMchId
	order.NotifyUrl = wepay.payNotifyUrl
	if wepay.subAppId != "" {
		order.SubAppId = wepay.subAppId
	}
	err = validate.Struct(order)
	return
}

// 服务商APP下单
func (wepay *wxPay) PartnerAppOrder(order PartnerOrder) (resp PrepayIdResponse, err error) {
	result, err := wepay.orderPartnerRequest(order, "app")
	if err != nil {
		return PrepayIdResponse{}, err
	} else {
		json.Unmarshal(result, &resp)
		return
	}
}

// 服务商Native下单
func (wepay *wxPay) PartnerNativeOrder(order PartnerOrder) (resp NativeOrderResponse, err error) {
	result, err := wepay.orderPartnerRequest(order, "native")
	if err != nil {
		return NativeOrderResponse{}, err
	} else {
		json.Unmarshal(result, &resp)
		return
	}
}

// 服务商H5下单
func (wepay *wxPay) PartnerH5Order(order PartnerOrder) (resp H5OrderResponse, err error) {
	result, err := wepay.orderPartnerRequest(order, "jsapi")
	if err != nil {
		return H5OrderResponse{}, err
	} else {
		json.Unmarshal(result, &resp)
		return
	}
}

// 服务商JsApi/小程序下单
func (wepay *wxPay) PartnerJsApiOrder(order PartnerOrder) (resp PrepayIdResponse, err error) {
	result, err := wepay.orderPartnerRequest(order, "jsapi")
	if err != nil {
		return PrepayIdResponse{}, err
	} else {
		json.Unmarshal(result, &resp)
		return
	}
}

// 执行订单相关请求
func (wepay *wxPay) orderPartnerRequest(order PartnerOrder, orderType string) (body []byte, err error) {
	err = wepay.validatePartnerOrder(&order)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	url := fmt.Sprintf("%stransactions/%s", urlPartnerPrefix, orderType)
	signBody, _ := json.Marshal(order)
	ts, nonceStr, _, signature := wepay.getSign("POST", url, string(signBody))
	body, err = wepay.doHttpRequest(fmt.Sprintf("%s%s", wePayHost, url), nonceStr, signature, string(signBody), ts)
	return
}

// 微信支付订单号查询
func (wepay *wxPay) PartnerQueryOrderTransactions(id string) (resp QueryPartnerOrderResponse, err error) {
	url := fmt.Sprintf("%stransactions/id/%s?sp_mchid=%s&sub_mchid=%s", urlPartnerPrefix, id, wepay.spMchId, wepay.subMchId)
	ts, nonceStr, _, signature := wepay.getSign("GET", url, "")
	var body []byte
	body, err = wepay.doHttpRequest(fmt.Sprintf("%s%s", wePayHost, url), nonceStr, signature, "", ts)
	if err != nil {
		return QueryPartnerOrderResponse{}, err
	}
	err = json.Unmarshal(body, &resp)
	return
}

// 商户订单号查询
func (wepay *wxPay) PartnerQueryOrderOutTradeNo(outTradeNo string) (resp QueryPartnerOrderResponse, err error) {
	url := fmt.Sprintf("%stransactions/out-trade-no/%s?sp_mchid=%s&sub_mchid=%s", urlPartnerPrefix, outTradeNo, wepay.spMchId, wepay.subMchId)
	ts, nonceStr, _, signature := wepay.getSign("GET", url, "")
	var body []byte
	body, err = wepay.doHttpRequest(fmt.Sprintf("%s%s", wePayHost, url), nonceStr, signature, "", ts)
	if err != nil {
		return QueryPartnerOrderResponse{}, err
	}
	err = json.Unmarshal(body, &resp)
	return
}

// 关闭订单
func (wepay *wxPay) PartnerCloseOrder(outTradeNo string) error {
	url := fmt.Sprintf("%stransactions/out-trade-no/%s/close", urlPartnerPrefix, outTradeNo)
	bodyParams := utils.KV{}
	bodyParams["sp_mchid"] = wepay.spMchId
	bodyParams["sub_mchid"] = wepay.subMchId
	bodyBuffer, _ := json.Marshal(bodyParams)
	ts, nonceStr, _, signature := wepay.getSign("POST", url, string(bodyBuffer))
	var body []byte
	var err error
	body, err = wepay.doHttpRequest(fmt.Sprintf("%s%s", wePayHost, url), nonceStr, signature, string(bodyBuffer), ts)
	if err != nil {
		return err
	}
	if len(body) > 0 {
		return errors.New(string(body))
	}
	return nil
}
