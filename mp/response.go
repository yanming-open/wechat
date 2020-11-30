package mp

import (
	"github.com/yanming-open/wechat/common"
)

// TokenResponse 获取token的返回结构体
type tokenResponse struct {
	common.BizResponse
	AccessToken string `json:"access_token"` // access_token
	ExpiresIn   int    `json:"expires_in"`   // 过期时间　默认7200秒
}

// jsapi jsticket输出
type jsTicketResponse struct {
	common.BizResponse
	Ticket    string `json:"ticket"`
	ExpiresIn int32  `json:"expires_in"`
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

// 获取用户列表
type UserListResponse struct {
	common.BizResponse
	Total int `json:"total"`
	Count int `json:"count"`
	Data  struct {
		OpenId []string `json:"openid"`
	} `json:"data"`
	NextOpenId string `json:"next_openid"`
}

// 上传临时素材时的结果返回
type MediaResponse struct {
	common.BizResponse
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt int    `json:"created_at"`
}

// 上传永久素材时的结果返回
type MaterialResponse struct {
	common.BizResponse
	MediaId string        `json:"media_id"`
	Url     string        `json:"url"`
	Item    []interface{} `json:"item,omitempty"`
}

// 获取素材总数结果返回
type MaterialCountResponse struct {
	common.BizResponse
	VoiceCount int `json:"voice_count"`
	VideoCount int `json:"video_count"`
	ImageCount int `json:"image_count"`
	NewsCount  int `json:"news_count"`
}
