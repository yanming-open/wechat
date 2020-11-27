package mp

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"github.com/yanming-open/wechat/common"
)

type iReplyMessage interface {
	PassiveReply(c *gin.Context) // 被动消息回复，直接响应http请求
}

type BizReplyMessage struct {
	iReplyMessage
	XMLName      xml.Name     `xml:"xml",gorm:"-"`
	ToUserName   common.CDATA //　接收人，一般为appid
	FromUserName common.CDATA // 消息来源，粉丝openid
	CreateTime   int          // 消息创建时间
	MsgType      common.CDATA // 消息类型
}

type ReplyTextMessage struct {
	BizReplyMessage
	Content common.CDATA
}

// 被动消息回复，响应http请求
func (this *ReplyTextMessage) PassiveReply(c *gin.Context) {
	buf, err := xml.Marshal(this)
	if err != nil {
		logger.Error(err.Error())
		_, _ = c.Writer.WriteString("")
	} else {
		_, _ = c.Writer.WriteString(string(buf))
	}
}

type ReplyImageMessage struct {
	BizReplyMessage
	Image common.CDATA `xml:"Image>MediaId"`
}

func (this *ReplyImageMessage) PassiveReply(c *gin.Context) {
	buf, err := xml.Marshal(this)
	if err != nil {
		logger.Error(err.Error())
		_, _ = c.Writer.WriteString("")
	} else {
		_, _ = c.Writer.WriteString(string(buf))
	}
}

type ReplyVoiceMessage struct {
	BizReplyMessage
	Voice common.CDATA `xml:"Voice>MediaId"`
}

func (this *ReplyVoiceMessage) PassiveReply(c *gin.Context) {
	buf, err := xml.Marshal(this)
	if err != nil {
		logger.Error(err.Error())
		_, _ = c.Writer.WriteString("")
	} else {
		_, _ = c.Writer.WriteString(string(buf))
	}
}

type ReplyVideoMessage struct {
	BizReplyMessage
	Video struct {
		Media       common.CDATA
		Title       common.CDATA
		Description common.CDATA
	} `xml:"Video"`
}

func (this *ReplyVideoMessage) PassiveReply(c *gin.Context) {
	buf, err := xml.Marshal(this)
	if err != nil {
		logger.Error(err.Error())
		_, _ = c.Writer.WriteString("")
	} else {
		_, _ = c.Writer.WriteString(string(buf))
	}
}

type ReplyMusicMessage struct {
	BizReplyMessage
	Music struct {
		Title        common.CDATA
		Description  common.CDATA
		MusicUrl     common.CDATA
		HQMusicUrl   common.CDATA
		ThumbMediaId common.CDATA
	}
}

func (this *ReplyMusicMessage) PassiveReply(c *gin.Context) {
	buf, err := xml.Marshal(this)
	if err != nil {
		logger.Error(err.Error())
		_, _ = c.Writer.WriteString("")
	} else {
		_, _ = c.Writer.WriteString(string(buf))
	}
}

type Article struct {
	Title       common.CDATA // 图文消息标题
	Description common.CDATA // 图文消息描述
	PicUrl      common.CDATA // 图片链接，支持JPG、PNG格式，较好的效果为大图360*200，小图200*200
	Url         common.CDATA // 点击图文消息跳转链接
}

type ReplyNewsMessage struct {
	BizReplyMessage
	ArticleCount int
	Articles     []Article `xml:"Articles>item"`
}

func (this *ReplyNewsMessage) PassiveReply(c *gin.Context) {
	buf, err := xml.Marshal(this)
	if err != nil {
		logger.Error(err.Error())
		_, _ = c.Writer.WriteString("")
	} else {
		_, _ = c.Writer.WriteString(string(buf))
	}
}
