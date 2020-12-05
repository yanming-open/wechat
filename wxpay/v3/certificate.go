package v3

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
)

type EncryptCertificate struct {
	Algorithm      string `json:"algorithm"`
	AssociatedData string `json:"associated_data"`
	CipherText     string `json:"ciphertext"`
	Nonce          string `json:"nonce"`
}

type Certificate struct {
	EffectiveTime      string             `json:"effective_time"`
	EncryptCertificate EncryptCertificate `json:"encrypt_certificate"`
	ExpireTime         string             `json:"expire_time"`
	SerialNo           string             `json:"serial_no"`
}

// TODO:获取加密证书，加密证书序列号
func (wepay *WxPay) getCertficates() {
	url := "/v3/certificates"
	ts, nonceStr, _, signature := wepay.getSign("GET", url, "")
	buf, err := wepay.doHttpRequest(fmt.Sprintf("%s%s", wePayHost, url), nonceStr, signature, "", ts)
	if err != nil {
		logger.Error(err.Error())
	}
	var resp struct {
		Data []Certificate `json:"data"`
	}
	json.Unmarshal(buf, &resp)
	for _, cert := range resp.Data {
		certBuf, err := base64.StdEncoding.DecodeString(cert.EncryptCertificate.CipherText)
		decryptStr, err := certificateDecrypt(certBuf,
			wepay.apiV3Key,
			cert.EncryptCertificate.Nonce, cert.EncryptCertificate.AssociatedData)
		if err != nil {
			log.Fatal(err)
		}
		wepay.publicKey = decryptStr
	}
}
