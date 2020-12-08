package v3

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type AccountType string

const (
	BasicAccount     AccountType = "BASIC"
	OperationAccount AccountType = "OPERATION"
	FeesAccount      AccountType = "FEES"
)

type BillType string

const (
	AllBill     BillType = "ALL"
	SuccessBill BillType = "SUCCESS"
	RefundBill  BillType = "REFUND"
)

// 申请交易账单API
// 对账单接口只能下载三个月以内的账单
//
// 小微商户不单独提供对账单下载，
// 如有需要，可在调取“下载对账单”API接口时不传sub_mch_id，
// 获取服务商下全量电商二级商户（包括小微商户和非小微商户）的对账单
// TODO: 校验直连商户是否成功下载帐单
func (wepay *wxPay) TradeBill(date, tar string, billtype BillType) (resp BillResponse, err error) {
	params := url.Values{}
	params.Add("bill_date", date)
	if wepay.subMchId != "" {
		// 如果是直连商户不传子商户号
		params.Add("sub_mchid", wepay.subMchId)
	}
	if billtype != "" {
		params.Add("bill_type", string(billtype))
	}
	if tar != "" {
		params.Add("tar_type", tar)
	}
	reqUrl := fmt.Sprintf("/v3/bill/tradebill?%s", params.Encode())
	ts, nonceStr, _, signature := wepay.getSign("GET", reqUrl, "")
	var body []byte
	body, err = wepay.doHttpRequest(fmt.Sprintf("%s%s", wePayHost, reqUrl), nonceStr, signature, "", ts)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resp)
	return
}

// 申请资金账单
func (wepay *wxPay) FundFlowBill(date, tar string, accountType AccountType) (resp BillResponse, err error) {
	params := url.Values{}
	params.Add("bill_date", date)
	if accountType != "" {
		params.Add("account_type", string(accountType))
	}
	if tar != "" {
		params.Add("tar_type", tar)
	}
	reqUrl := fmt.Sprintf("/v3/bill/fundflowbill?%s", params.Encode())
	ts, nonceStr, _, signature := wepay.getSign("GET", reqUrl, "")
	var body []byte
	body, err = wepay.doHttpRequest(fmt.Sprintf("%s%s", wePayHost, reqUrl), nonceStr, signature, "", ts)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resp)
	return
}

// 下载帐单；
// @return: 如果帐单申请时指定了为压缩包，则需要将body保存为.gzip的压缩文件
//			如示指定压缩格式　建议将body直接保存后缀为csv文件
//
func (wepay *wxPay) DownloadBill(bill BillResponse) (body []byte, err error) {
	ts, nonceStr, _, signature := wepay.getSign(
		"GET",
		strings.Replace(bill.DownloadUrl, wePayHost, "", 1),
		"")
	body, err = wepay.doHttpRequest(bill.DownloadUrl, nonceStr, signature, "", ts)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return
}
