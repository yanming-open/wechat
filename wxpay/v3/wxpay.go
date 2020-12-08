package v3

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type WeConfig struct {
	SpAppId         string // 服务商申请的公众号或移动应用appid
	SpMchId         string // 服务商户号
	SubMchId        string // 子商户号
	SubAppId        string // 子商户申请的公众号或移动应用appid
	SerialNo        string // 证书序列号　V3　版本api必传
	KeyPath         string // 证书私钥路径
	PayNotifyUrl    string // 支付通知地址
	RefundNotifyUrl string // 退款通知地址
	ApiV3Key        string // v3密钥
}

type wxPay struct {
	spAppId         string // 服务商申请的公众号或移动应用appid
	spMchId         string // 服务商户号
	subMchId        string // 子商户号
	subAppId        string // 子商户申请的公众号或移动应用appid
	serialNo        string // 证书序列号　V3　版本api必传
	keyPath         string // 证书私钥路径
	payNotifyUrl    string // 支付通知地址
	refundNotifyUrl string // 退款通知地址
	apiV3Key        string // v3密钥
	publicKey       string // 请求得到的公钥串　TODO:
}

func NewWxPay(c WeConfig) *wxPay {
	var wepay = wxPay{
		spAppId:         c.SpAppId,
		spMchId:         c.SpMchId,
		subAppId:        c.SubAppId,
		subMchId:        c.SubMchId,
		serialNo:        c.SerialNo,
		keyPath:         c.KeyPath,
		payNotifyUrl:    c.PayNotifyUrl,
		refundNotifyUrl: c.RefundNotifyUrl,
		apiV3Key:        c.ApiV3Key,
	}
	return &wepay
}

// 获取签名
func (wepay *wxPay) getSign(method, url, body string) (timeStamp int64, randomStr, signStr, cipherData string) {
	timeStamp = time.Now().Unix()
	randomStr = randString(16)
	signStr = fmt.Sprintf(signTemplate, method, url, timeStamp, randomStr, body)
	cipherData = rsa256Encrypt(signStr, wepay.keyPath)
	return
}

// 执行http操作
func (wepay *wxPay) doHttpRequest(url, nonceStr, signature, body string, timeStamp int64) ([]byte, error) {
	client := http.Client{Timeout: time.Second * 15}
	var request *http.Request
	if body != "" {
		request, _ = http.NewRequest("POST", url, strings.NewReader(body))
	} else {
		request, _ = http.NewRequest("GET", url, nil)
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Header.Set("User-Agent", userAgent)
	request.Header.Set("Accept", "application/json;charset=UTF-8")
	request.Header.Set("Authorization", fmt.Sprintf(authTemplate, wepay.spMchId, nonceStr, signature, timeStamp, wepay.serialNo))

	resp, err := client.Do(request)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	return buf, err
}
