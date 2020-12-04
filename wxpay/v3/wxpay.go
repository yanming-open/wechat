package v3

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type WeConfig struct {
	SpAppId         string
	SpMchId         string
	SubMchId        string
	SubAppId        string
	SerialNo        string
	KeyPath         string
	PayNotifyUrl    string
	RefundNotifyUrl string
}

type WxPay struct {
	spAppId         string
	spMchId         string
	subMchId        string
	subAppId        string
	serialNo        string
	keyPath         string
	payNotifyUrl    string
	refundNotifyUrl string
}

func NewWxPay(c WeConfig) *WxPay {
	var wepay = WxPay{
		spAppId:         c.SpAppId,
		spMchId:         c.SpMchId,
		subAppId:        c.SubAppId,
		subMchId:        c.SubMchId,
		serialNo:        c.SerialNo,
		keyPath:         c.KeyPath,
		payNotifyUrl:    c.PayNotifyUrl,
		refundNotifyUrl: c.RefundNotifyUrl,
	}
	return &wepay
}

// 获取签名
func (wepay *WxPay) getSign(method, url, body string) (timeStamp int64, randomStr, signStr, cipherData string) {
	timeStamp = time.Now().Unix()
	randomStr = randString(16)
	signStr = fmt.Sprintf(signTemplate, method, url, timeStamp, randomStr, body)
	cipherData = rsa256Encrypt(signStr, wepay.keyPath)
	return
}

// 执行http操作
func (wepay *WxPay) doHttpRequest(url, nonceStr, signature, body string, timeStamp int64) ([]byte, error) {
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

func (wepay *WxPay) getCertficates() {
	url := "/v3/certificates"
	ts, nonceStr, _, signature := wepay.getSign("GET", url, "")
	wepay.doHttpRequest(fmt.Sprintf("%s%s", wePayHost, url), nonceStr, signature, "", ts)
}
