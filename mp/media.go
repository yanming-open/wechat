package mp

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/yanming-open/wechat/utils"
)

// 临时素材media_id是可复用的。
// 媒体文件在微信后台保存时间为3天，即3天后media_id失效。
// 上传临时素材的格式、大小限制与公众平台官网一致。
// 图片（image）: 10M，支持PNG\JPEG\JPG\GIF格式
// 语音（voice）：2M，播放长度不超过60s，支持AMR\MP3格式
// 视频（video）：10MB，支持MP4格式
// 缩略图（thumb）：64KB，支持JPG格式
func (m *Mp) UploadMedia(filePath, fileType string) (resp *MediaResponse, err error) {
	url := fmt.Sprintf("%smedia/upload?access_token=%s&type=%s", wxApiHost, m.accessToken, fileType)
	body, err := utils.DoUpload(url, filePath)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return
	}
	if resp.ErrCode != 0 {
		return new(MediaResponse), errors.New(resp.ErrMsg)
	}
	return
}

// 获取临时素材，如果是视频素材　返回内容为
// {
//    "video_url":DOWN_URL
// }
// 需要根据自有业务场景处理返回
func (m *Mp) GetMeida(mediaid string) (body []byte, err error) {
	url := fmt.Sprintf("%smedia/get?access_token=%s&media_id=%s", wxApiHost, m.accessToken, mediaid)
	body, err = utils.DoGet(url)
	return
}
