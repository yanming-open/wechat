package mp

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/xml"
	"errors"
	"github.com/yanming-open/wechat/common"
)

const (
	Text       string = "text"       // 文本消息
	Image      string = "image"      // 图片消息
	Voice      string = "voice"      // 语音消息
	Video      string = "video"      // 视频消息
	ShortVideo string = "shortvideo" // 小视频消息
	Location   string = "location"   // 位置消息
	Link       string = "link"       // 连接消息
	Event      string = "event"      // 事件消息
	Music      string = "music"      // 音乐消息
	News       string = "news"       // 图文消息

)

//　EncryptMsg
type EncryptMsg struct {
	XMLName    xml.Name `xml:"xml"`
	ToUserName string   //　消息接收人，一般为appid
	Encrypt    string   // 消息加密内容
}

//　IBizMessage　接口
//　需要所有子类实现　Save 方法
type IBizMessage interface {
	DecryptMsg(encryptStr string) error // 解密消息体，会将结果保存到消息的　MessageContent　属性
	Save()                              // 用于实现消息的保存，持久化操作
}

// BizMessage 结构体
type BizMessage struct {
	IBizMessage
	XMLName        xml.Name     `xml:"xml",gorm:"-"`
	ToUserName     common.CDATA //　接收人，一般为appid
	FromUserName   common.CDATA // 消息来源，粉丝openid
	CreateTime     int          // 消息创建时间
	MsgType        common.CDATA // 消息类型
	MsgId          int64
	MessageContent []byte `gorm:"-"` // 触密后的内容
}

// 执行消息反序列化前，必需先调用解密操作
// 解密操作会校验　appid　是否一致，并截取真实消息体内容
// 其它各类消息继承自　BizMessage　即可，不用再实现解密操作
func (biz *BizMessage) DecryptMsg(encryptStr string, mp *Mp) error {
	cipherData, err := base64.StdEncoding.DecodeString(encryptStr)
	if err != nil {
		return err
	} else {
		plainData, err := aesDecrypt(cipherData, mp.aeskey)
		buf := bytes.NewBuffer(plainData[16:20])
		var length int32
		binary.Read(buf, binary.BigEndian, &length)
		appIDstart := 20 + length
		id := plainData[appIDstart : int(appIDstart)+len(mp.AppId)]
		if string(id) != mp.AppId {
			return errors.New("appid　is invalid")
		}
		biz.MessageContent = plainData[20 : 20+length]
		return err
	}
}

type TextMessage struct {
	IBizMessage
	BizMessage
	Content common.CDATA
}

type VoiceMessage struct {
	IBizMessage
	BizMessage
	MediaId     common.CDATA
	Format      common.CDATA
	Recognition common.CDATA
}

type VideoMessage struct {
	IBizMessage
	BizMessage
	MediaId      common.CDATA
	ThumbMediaId common.CDATA
}
type ShortVideoMessage struct {
	IBizMessage
	BizMessage
	MediaId      common.CDATA
	ThumbMediaId common.CDATA
}

type LocationMessage struct {
	IBizMessage
	BizMessage
	Location_X float64
	Location_Y float64
	Scale      int
	Label      common.CDATA
}

type LinkMessage struct {
	IBizMessage
	BizMessage
	Title       common.CDATA
	Description common.CDATA
	Url         common.CDATA
}
