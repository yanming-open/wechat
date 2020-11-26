package mp

import "github.com/yanming-open/wechat/common"

// TokenResponse 获取token的返回结构体
type TokenResponse struct {
	common.BizResponse
	AccessToken string `json:"access_token"` // access_token
	ExpiresIn   int    `json:"expires_in"`   // 过期时间　默认7200秒
}

// 长短地址转换响应
type long2ShortResponse struct {
	ErrCode  int    `json:"errcode"`
	ErrMsg   string `json:"errmsg"`
	ShortUrl string `json:"short_url,omitempty"`
}