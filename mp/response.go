package mp

import "github.com/yanming-open/wechat/common"

// TokenResponse 获取token的返回结构体
type tokenResponse struct {
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

// 标签列表返回数据
type tagsResponse struct {
	Tags []Tag `json:"tags"`
}

type tagsIdListResponse struct {
	common.BizResponse
	TagIdList []int `json:"tagid_list"`
}

// 获取标签下粉丝列表
type TagUsersResponse struct {
	common.BizResponse
	Count int `json:"count"`
	Data  struct {
		OpenId []string `json:"openid"`
	} `json:"data"`
	NextOpenId string `json:"next_openid"`
}
