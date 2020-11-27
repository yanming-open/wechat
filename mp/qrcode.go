package mp

import (
	"encoding/json"
	"fmt"
	"github.com/yanming-open/wechat/utils"
)

const (
	qrScene         = "QR_SCENE"           // 临时的整型参数值
	qrStrScene      = "QR_STR_SCENE"       // 临时的字符串参数值
	qrLimitScene    = "QR_LIMIT_SCENE"     // 永久的整型参数值
	qrLimitStrScene = "QR_LIMIT_STR_SCENE" // 永久的字符串参数值
)

// 二维码生成接口
type IQrCode interface {
	CreateQrCode(expire int, sceneinfo interface{})
	CreateLimitQrCode(sceneinfo interface{})
}

// 二维码请求参数
type qrCodeRequest struct {
	ActionName string `json:"action_name"` // 二维码类型
	ActionInfo struct {
		Scene struct {
			SceneId  int    `json:"scene_id,omitempty"`  // 场景值ID，临时二维码时为32位非0整型，永久二维码时最大值为100000（目前参数只支持1--100000）
			SceneStr string `json:"scene_str,omitempty"` // 场景值ID（字符串形式的ID），字符串类型，长度限制为1到64
		} `json:"scene"`
	} `json:"action_info"`                    // 二维码详细信息
	ExpireSeconds int `json:"expire_seconds"` // 该二维码有效时间，以秒为单位。 最大不超过2592000（即30天），此字段如果不填，则默认有效期为30秒。
}

type QrCode struct {
	IQrCode
	Ticket        string `json:"ticket"`
	ExpireSeconds int    `json:"expire_seconds"`
	Url           string `json:"url"`
	ShortUrl      string `json:"short_url"`
}

// 创建临时二维码
func (this *Mp) CreateQrCode(expire int, sceneinfo interface{}) (qc *QrCode) {
	url := fmt.Sprintf("%sqrcode/create?access_token=%s", wxApiHost, this.accessToken)
	var request = qrCodeRequest{}
	switch sceneinfo.(type) {
	case int:
		request.ActionName = qrScene
		request.ActionInfo.Scene.SceneId = sceneinfo.(int)
		break
	case string:
		request.ActionName = qrStrScene
		request.ActionInfo.Scene.SceneStr = sceneinfo.(string)
		break
	}
	request.ExpireSeconds = expire
	buf, err := json.Marshal(request)
	if err != nil {
		logger.Error(err.Error())
	}
	var params = utils.KV{}
	json.Unmarshal(buf, &params)
	body, err := utils.DoPost(url, params)
	if err != nil {
		logger.Error(err.Error())
	}
	json.Unmarshal(body, &qc)
	qc.Url = fmt.Sprintf("%sshowqrcode?ticket=%s", wxApiHost, qc.Ticket)
	qc.ShortUrl = long2short(qc.Url, this.accessToken)
	return qc
}

// 创建永久二维码
func (this *Mp) CreateLimitQrCode(sceneinfo interface{}) (qc *QrCode) {
	url := fmt.Sprintf("%sqrcode/create?access_token=%s", wxApiHost, this.accessToken)
	var request = qrCodeRequest{}
	switch sceneinfo.(type) {
	case int:
		request.ActionName = qrLimitScene
		request.ActionInfo.Scene.SceneId = sceneinfo.(int)
		break
	case string:
		request.ActionName = qrLimitStrScene
		request.ActionInfo.Scene.SceneStr = sceneinfo.(string)
		break
	}
	buf, err := json.Marshal(request)
	if err != nil {
		logger.Error(err.Error())
	}
	var params = utils.KV{}
	json.Unmarshal(buf, &params)
	body, err := utils.DoPost(url, params)
	if err != nil {
		logger.Error(err.Error())
	}
	json.Unmarshal(body, &qc)
	qc.Url = fmt.Sprintf("%sshowqrcode?ticket=%s", wxApiHost, qc.Ticket)
	qc.ShortUrl = long2short(qc.Url, this.accessToken)
	return qc
}
