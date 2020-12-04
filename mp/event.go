package mp

import (
	"encoding/xml"
	"github.com/yanming-open/wechat/common"
)

const (
	SubScribe   string = "subscribe"   // 订阅事件;用户未关注时　扫描带参数二维码事件　会带有　EventKey
	UnSubScribe string = "unsubscribe" //　取消订阅事件
	Scan        string = "SCAN"        // 用户已关注时的事件推送　扫描带参数二维码事件　会带有　EventKey
	LOCATION    string = "LOCATION"    // 上报地理位置事件
	Click       string = "CLICK"       // 自定义菜单事件--点击菜单拉取消息时的事件推送
	View        string = "VIEW"        // 自定义菜单事件--点击菜单跳转链接时的事件推送
)

type BizEvent struct {
	XMLName      xml.Name     `xml:"xml",gorm:"-"`
	ToUserName   common.CDATA //　接收人，一般为appid
	FromUserName common.CDATA // 消息来源，粉丝openid
	CreateTime   int          // 消息创建时间
	MsgType      common.CDATA // 消息类型
	Event        common.CDATA // 事件类型，subscribe(订阅)、unsubscribe(取消订阅)
}

// 用户扫描带场景值二维码时，可能推送以下两种事件：
//
//如果用户还未关注公众号，则用户可以关注公众号，关注后微信会将带场景值关注事件推送给开发者。
// event 值为 subscribe
//如果用户已经关注公众号，则微信会将带场景值扫描事件推送给开发者。
// event 值为 SCAN
//
type ScanSenceEvent struct {
	BizEvent
	EventKey common.CDATA // 事件KEY值，是一个32位无符号整数，即创建二维码时的二维码scene_id
	Ticket   common.CDATA // 二维码的ticket，可用来换取二维码图片
}

// 上报地理位置事件
type LocationEvent struct {
	BizEvent
	Latitude  float32 // 地理位置纬度
	Longitude float32 // 地理位置经度
	Precision float32 // 地理位置精度
}

// 自定义菜单事件
// 点击菜单跳转链接时的事件推送
type ClickEvent struct {
	BizEvent
	EventKey common.CDATA // 事件KEY值，与自定义菜单接口中KEY值对应/设置的跳转URL
}
