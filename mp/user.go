package mp

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/yanming-open/wechat/common"
	"github.com/yanming-open/wechat/utils"
)

// User　接口
// 定义获取用户信息方法
type IUser interface {
	GetUserInfo(openid string) error // 通过openid获取用户信息
}

// 用户定义
type User struct {
	IUser
	common.BizResponse
	SubScribe      int    `json:"subscribe"`       // 用户是否订阅该公众号标识，值为0时，代表此用户没有关注该公众号，拉取不到其余信息。
	OpenId         string `json:"openid"`          // 用户的标识，对当前公众号唯一
	NickName       string `json:"nickname"`        // 用户的昵称
	Sex            int    `json:"sex"`             // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	Language       string `json:"language"`        //用户的语言，简体中文为zh_CN
	City           string `json:"city"`            // 用户所在城市
	Province       string `json:"province"`        //用户所在省份
	Country        string `json:"country"`         //用户所在国家
	HeadImgUrl     string `json:"headimgurl"`      // 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
	SubScribeTime  int    `json:"subscribe_time"`  // 用户关注时间，为时间戳。如果用户曾多次关注，则取最后关注时间
	UnionId        string `json:"unionid"`         // 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
	ReMark         string `json:"remark"`          // 公众号运营者对粉丝的备注，公众号运营者可在微信公众平台用户管理界面对粉丝添加备注
	GroupId        int    `json:"groupid"`         // 用户所在的分组ID（兼容旧的用户分组接口）
	TagIdList      []int  `json:"tagid_list"`      // 用户被打上的标签ID列表
	SubScribeScene string `json:"subscribe_scene"` // 返回用户关注的渠道来源，ADD_SCENE_SEARCH 公众号搜索，ADD_SCENE_ACCOUNT_MIGRATION 公众号迁移，ADD_SCENE_PROFILE_CARD 名片分享，ADD_SCENE_QR_CODE 扫描二维码，ADD_SCENE_PROFILE_LINK 图文页内名称点击，ADD_SCENE_PROFILE_ITEM 图文页右上角菜单，ADD_SCENE_PAID 支付后关注，ADD_SCENE_WECHAT_ADVERTISEMENT 微信广告，ADD_SCENE_OTHERS 其他
	QrScene        int    `json:"qr_scene"`        // 二维码扫码场景（开发者自定义）
	QrSceneStr     string `json:"qr_scene_str"`    // 二维码扫码场景描述（开发者自定义）
}

// 获取用户信息
func (this *Mp) GetUserInfo(openid string) (u *User, err error) {
	url := fmt.Sprintf("%suser/info?access_token=%s&openid=%s&lang=zh_CN", wxApiHost, this.accessToken, openid)
	buf, err := utils.DoGet(url)
	if err != nil {
		return nil, err
	} else {
		err = json.Unmarshal(buf, &u)
		if err != nil {
			return nil, err
		} else {
			if u.ErrCode != 0 {
				return nil, errors.New(u.ErrMsg)
			} else {
				return u, nil
			}
		}
	}
}
