package mp

import (
	"encoding/json"
	"fmt"
	"github.com/yanming-open/wechat/common"
	"github.com/yanming-open/wechat/utils"
)

type MaterialNews struct {
	Title              string `json:"title"`                 // 标题 必填
	ThumbMediaId       string `json:"thumb_media_id"`        // 图文消息的封面图片素材id（必须是永久mediaID） 必填
	Author             string `json:"author"`                // 作者 非必填
	Digest             string `json:"digest"`                // 图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空。如果本字段为没有填写，则默认抓取正文前64个字。
	ShowCoverPic       int    `json:"show_cover_pic"`        // 是否显示封面，0为false，即不显示，1为true，即显示　必填
	Content            string `json:"content"`               // 图文消息的具体内容，支持HTML标签，必须少于2万字符，小于1M，且此处会去除JS,涉及图片url必须来源 "上传图文消息内的图片获取URL"接口获取。外部图片url将被过滤 必填
	ContentSourceUrl   string `json:"content_source_url"`    // 图文消息的原文地址，即点击“阅读原文”后的URL 必填
	NeedOpenComment    uint32 `json:"need_open_comment"`     // 是否打开评论，0不打开，1打开
	OnlyFansCanComment uint32 `json:"only_fans_can_comment"` // 是否粉丝才可评论，0所有人可评论，1粉丝才可评论
}

func (m *Mp)MaterialAdd(filePath,fileType string,vedioDesc utils.KV)(resp MaterialResponse,err error)  {
	switch fileType {
	case Video:
		break
	default:
		break
	}
	return
}

func (m *Mp) MaterialAddNews() {
	url := fmt.Sprintf("%smaterial/add_news?access_token=%s", wxApiHost, m.accessToken)
	params := utils.KV{}
	params["articles"] = []MaterialNews{{Title: "test",
		ThumbMediaId: "http://mmbiz.qpic.cn/sz_mmbiz_jpg/bOgRrYiakxTO6jzp97tBTskBk5SGmLrh49IJFrKQsz8lIA45Owlia2bliaB5nHXebib0JdibZw3ms3kmKvKacicFKxtg/0",
		Author:       "admin", ShowCoverPic: 1, Content: "内容", ContentSourceUrl: "http://baidu.com"}}
	body, err := utils.DoPost(url, params)
	if err != nil {
		logger.Error(err.Error())
	} else {
		logger.Info(string(body))
	}
}

// 上传图文消息内的图片获取URL
// 图片仅支持jpg/png格式，大小必须在1MB以下
func (m *Mp) MaterialUploadImg(filePath string) string {
	url := fmt.Sprintf("%smedia/uploadimg?access_token=%s", wxApiHost, m.accessToken)
	body, err := utils.DoUpload(url, filePath)
	if err != nil {
		logger.Error(err.Error())
		return ""
	} else {
		var resp = struct {
			common.BizResponse
			Url string `json:"url"`
		}{}
		json.Unmarshal(body, &resp)
		if resp.ErrCode != 0 {
			logger.Error(resp.ErrMsg)
			return ""
		} else {
			return resp.Url
		}
	}
}
