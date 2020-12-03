package v3

import (
	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
)

const (
	wePayHost        = "https://api.mch.weixin.qq.com" // 支付接口主域名
	urlPartnerPrefix = "/v3/pay/partner/"  // 服务商模式URL前缀
	letterBytes      = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // 用于生成随机串
	signTemplate     = "%s\n%s\n%v\n%s\n%s\n" // 签名模板：HTTP请求方法/URL/请求时间戳/随机串/请求报文主体
	// 签名模板：认证类型/发起请求的商户/随机串/时间戳/签名值/证书序列号
	authTemplate     = `WECHATPAY2-SHA256-RSA2048 mchid="%s",nonce_str="%s",signature="%s",timestamp="%v",serial_no="%s"`
	userAgent        = "YanMing-Open WeChat1.0.1" // 请求头
)

func init() {
	validate = validator.New()
}

