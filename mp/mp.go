package mp

import (
	"encoding/json"
	"fmt"
	"github.com/yanming-open/wechat/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var logger *zap.Logger

const wxApiHost = "https://api.weixin.qq.com/cgi-bin/" // 微信接口服务器地址

type Mp struct {
	Token          string
	EncodingAESKey string
	AppId          string
	AppSecret      string
	aeskey         []byte
	accessToken    string
}

// 实例化一个公众号接口实例，同一服务中只要实例化一个即可
// accesstoken 会自动刷新，初次调用需要注意先等待accesstoken请求完成
func NewMp(token, encodingaeskey, appid, appsecret string) *Mp {
	var mp = Mp{Token: token, EncodingAESKey: encodingaeskey, AppId: appid, AppSecret: appsecret}
	mp.initMp()
	return &mp
}

// 初始化日志
// 初始化accesstoken调用,使用新线程异步调用
func (mp *Mp) initMp() {
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
		zap.InfoLevel,              // 日志级别
	)
	logger = zap.New(core)
	defer logger.Sync()
	mp.aeskey = encodingAESKey2AESKey(mp.EncodingAESKey)
	go timerTicketToken(mp)
}

// @title getAccessToken
// @description 获取access token
// @auth
// @return data TokenResponse "返回结果"
// @return error error "返回错误"
func getAccessToken(mp *Mp) (tokenResponse, error) {
	url := fmt.Sprintf("%stoken?grant_type=client_credential&appid=%s&secret=%s", wxApiHost, mp.AppId, mp.AppSecret)
	buf, err := utils.DoGet(url)
	var result tokenResponse
	if err != nil {
		return result, err
	} else {
		err = json.Unmarshal(buf, &result)
		return result, err
	}
}

func timerTicketToken(mp *Mp) {
	var result tokenResponse
	var err error
	for {
		result, err = getAccessToken(mp)
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		if result.ErrCode != 0 {
			logger.Error(result.ErrMsg)
			continue
		}
		logger.Info(result.AccessToken)
		mp.accessToken = result.AccessToken
		logger.Info("token　初始化成功，可以调用啦！")
		time.Sleep(time.Second * 7100)
	}
}
