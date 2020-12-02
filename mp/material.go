package mp

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/yanming-open/wechat/common"
	"github.com/yanming-open/wechat/utils"
)

// 图文素材
type MaterialNews struct {
	Title              string `json:"title"`                 // 标题 必填
	ThumbMediaId       string `json:"thumb_media_id"`        // 图文消息的封面图片素材id（必须是永久mediaID） 必填
	ThumbUrl           string `json:"thumb_url,omitempty"`   // 获取素材时会返回，新建时不用赋值
	Author             string `json:"author"`                // 作者 非必填
	Digest             string `json:"digest"`                // 图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空。如果本字段为没有填写，则默认抓取正文前64个字。
	ShowCoverPic       int    `json:"show_cover_pic"`        // 是否显示封面，0为false，即不显示，1为true，即显示　必填
	Content            string `json:"content"`               // 图文消息的具体内容，支持HTML标签，必须少于2万字符，小于1M，且此处会去除JS,涉及图片url必须来源 "上传图文消息内的图片获取URL"接口获取。外部图片url将被过滤 必填
	ContentSourceUrl   string `json:"content_source_url"`    // 图文消息的原文地址，即点击“阅读原文”后的URL 必填
	NeedOpenComment    uint32 `json:"need_open_comment"`     // 是否打开评论，0不打开，1打开
	OnlyFansCanComment uint32 `json:"only_fans_can_comment"` // 是否粉丝才可评论，0所有人可评论，1粉丝才可评论
}

type VideoDesc struct {
	Title        string `json:"title"`
	Introduction string `json:"introduction"`
}

// 在上传视频素材时需要POST另一个表单，id为description，包含素材的描述信息，内容格式为JSON，格式如下：
// {
//    "title":VIDEO_TITLE,
//    "introduction":INTRODUCTION
//}
func (m *Mp) MaterialAdd(filePath, fileType string, videoDesc utils.KV) (resp MaterialResponse, err error) {
	url := fmt.Sprintf("%smaterial/add_material?access_token=%s&type=%s", wxApiHost, m.accessToken, fileType)
	if fileType == Video && videoDesc == nil {
		return resp, errors.New("视频描述不能为空")
	}
	var body []byte
	body, err = utils.DoUpload(url, filePath, videoDesc)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resp)
	return
}

// 新增永久图文素材
func (m *Mp) MaterialAddNews(news []MaterialNews) (resp MaterialResponse, err error) {
	url := fmt.Sprintf("%smaterial/add_news?access_token=%s", wxApiHost, m.accessToken)
	params := utils.KV{}
	params["articles"] = news
	var body []byte
	body, err = utils.DoPost(url, params)
	if err != nil {
		logger.Error(err.Error())
		return
	} else {
		err = json.Unmarshal(body, &resp)
	}
	return
}

// 上传图文消息内的图片获取URL
// 图片仅支持jpg/png格式，大小必须在1MB以下
func (m *Mp) MaterialUploadImg(filePath string) string {
	url := fmt.Sprintf("%smedia/uploadimg?access_token=%s", wxApiHost, m.accessToken)
	body, err := utils.DoUpload(url, filePath, nil)
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

// 返回结构请参考 https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Get_materials_list.html
func (m *Mp) BatchGetMaterial(Type string, offset, count int) (result interface{}, err error) {
	url := fmt.Sprintf("%smaterial/batchget_material?access_token=%s", wxApiHost, m.accessToken)
	params := utils.KV{}
	params["type"] = Type
	params["offset"] = offset
	params["count"] = count
	body, err := utils.DoPost(url, params)
	err = json.Unmarshal(body, &result)
	return
}

// 获取永久素材总数
func (m *Mp) GetMaterialCount() (resp MaterialCountResponse, err error) {
	url := fmt.Sprintf("%smaterial/get_materialcount?access_token=%s", wxApiHost, m.accessToken)
	var body []byte
	body, err = utils.DoGet(url)
	if err != nil {
		return resp, err
	} else {
		err = json.Unmarshal(body, &resp)
	}
	return
}
