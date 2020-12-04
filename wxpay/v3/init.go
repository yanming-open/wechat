package v3

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var (
	validate *validator.Validate
	logger   *zap.Logger
)

const (
	wePayHost        = "https://api.mch.weixin.qq.com"                                  // 支付接口主域名
	urlPartnerPrefix = "/v3/pay/partner/"                                               // 服务商模式URL前缀
	letterBytes      = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // 用于生成随机串
	signTemplate     = "%s\n%s\n%v\n%s\n%s\n"                                           // 签名模板：HTTP请求方法/URL/请求时间戳/随机串/请求报文主体
	// 签名模板：认证类型/发起请求的商户/随机串/时间戳/签名值/证书序列号
	authTemplate = `WECHATPAY2-SHA256-RSA2048 mchid="%s",nonce_str="%s",signature="%s",timestamp="%v",serial_no="%s"`
	userAgent    = "YanMing-Open WeChat1.0.1" // 请求头
)

func init() {
	validate = validator.New()
	hook := &lumberjack.Logger{
		Filename:   "./wepayv3.log", // 日志文件路径
		MaxSize:    500,             // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 3,               // 日志文件最多保存多少个备份
		MaxAge:     28,              // 文件最多保存多少天
		Compress:   true,            // 是否压缩
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
}
