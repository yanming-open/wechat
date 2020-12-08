// 公众号sdk
package mp

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/yanming-open/wechat/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"sort"
	"strings"
	"time"
)

var logger *zap.Logger

const (
	wxApiHost   = "https://api.weixin.qq.com/cgi-bin/" // 微信接口服务器地址
	wxSnsHost   = "https://api.weixin.qq.com/sns/"
	letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type mp struct {
	token          string
	encodingaeskey string
	appid          string
	appsecret      string
	aeskey         []byte
	accessToken    string
	jsTicket       string
}

// 实例化一个公众号接口实例，同一服务中只要实例化一个即可
// accesstoken 会自动刷新，初次调用需要注意先等待accesstoken请求完成
func NewMp(token, encodingaeskey, appid, appsecret string) *mp {
	var mp = mp{token: token, encodingaeskey: encodingaeskey, appid: appid, appsecret: appsecret}
	mp.initMp()
	return &mp
}

// 初始化日志
// 初始化accessToken调用,使用新线程异步调用
func (mp *mp) initMp() {
	hook := &lumberjack.Logger{
		Filename:   "./mp.log", // 日志文件路径
		MaxSize:    500,        // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 3,          // 日志文件最多保存多少个备份
		MaxAge:     28,         // 文件最多保存多少天
		Compress:   true,       // 是否压缩
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), // 编码器配置
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout),
			zapcore.AddSync(hook)), // 打印到控制台和文件
		zap.InfoLevel, // 日志级别
	)
	logger = zap.New(core)
	defer logger.Sync()
	mp.aeskey = encodingAESKey2AESKey(mp.encodingaeskey)
	go timerTicketToken(mp)
}

// 构建web端config时的参数
// 需要在前端验证用户权限
func (mp *mp) GetJsTicketSignature(url string) (resp JsTicketSignatureResponse) {
	noncestr := randString(16)
	timestamp := time.Now().Unix()
	sl := []string{fmt.Sprintf("noncestr=%s", noncestr),
		fmt.Sprintf("jsapi_ticket=%s", mp.jsTicket),
		fmt.Sprintf("timestamp=%v", timestamp),
		fmt.Sprintf("url=%s", url),
	}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))
	signature := fmt.Sprintf("%x", s.Sum(nil))
	resp.JsapiTicket = mp.jsTicket
	resp.NonceStr = noncestr
	resp.Timestamp = timestamp
	resp.Signature = signature
	return
}

// @title getAccessToken
// @description 获取access token
// @auth
// @return data TokenResponse "返回结果"
// @return error error "返回错误"
func requestAccessToken(mp *mp) (tokenResponse, error) {
	url := fmt.Sprintf("%stoken?grant_type=client_credential&appid=%s&secret=%s", wxApiHost, mp.appid, mp.appsecret)
	buf, err := utils.DoGet(url)
	var result tokenResponse
	if err != nil {
		return result, err
	} else {
		err = json.Unmarshal(buf, &result)
		return result, err
	}
}

func requestJsTicket(mp *mp) (result jsTicketResponse, err error) {
	url := fmt.Sprintf("%sticket/getticket?access_token=%s&type=jsapi", wxApiHost, mp.accessToken)
	var body []byte
	body, err = utils.DoGet(url)
	if err != nil {
		return
	} else {
		err = json.Unmarshal(body, &result)
		return
	}
}

func timerTicketToken(mp *mp) {
	var result tokenResponse
	var jsTicketResp jsTicketResponse
	var err error
	for {
		result, err = requestAccessToken(mp)
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		if result.ErrCode != 0 {
			logger.Error(result.ErrMsg)
			continue
		}
		mp.accessToken = result.AccessToken
		logger.Info("token　初始化成功，可以调用啦！")
		jsTicketResp, err = requestJsTicket(mp)
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		if jsTicketResp.ErrCode != 0 {
			logger.Error(jsTicketResp.ErrMsg)
			continue
		}
		mp.jsTicket = jsTicketResp.Ticket
		time.Sleep(time.Second * 7100)
	}
}
