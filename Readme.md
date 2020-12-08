# 微信开发sdk

* mp 公众号接口
* wxpay/v3 微信支付APIV3版本


# 功能列表

各子模块都默认创建日志文件记录

* 公众号
    * 永久素材添加、统计、拉取
    * 临时素材上传、获取
    * 消息　解密　**持久化需要自行实现Save方法**
    * 生成临时/永久二维码;支持场景定义
    * GetJsTicketSignature 将返回web端开发时JsTicket、时间戳、随机串及签名数据
    * Tag管理
    * 用户管理
        * 黑名单管理
        * 用户详情获取(包含web端　`GetSnsUserInfo` 非静默授权)
        * 用户列表
        * web端code获取access token
        * web端刷新access token
    * 被动消息回复，构建好回复内容后调用 `PassiveReply`方法即可；需要配合 [gin](https://github.com/gin-gonic/gin) 框架
    * 基本access token 和 JsTicket 每`7200`秒自动刷新
    
* 微信支付V3
    * 服务商模式下单、关闭、查询
    * 退款、退款查询


# TODO

- [ ] 支付直连商户功能
- [ ] 支付数据提交加密
- [ ] 多级struct使用`required`验证时json序列化问题

# ChangeLog

## v1.1.0

* 增加微信支付api v3版本接口
* 支持服务商模式下单、订单关闭、查询
* 退款、退款查询
* 帐单下载

## v1.0.0

* 完成公众号消息、用户、标签、素材(永久/临时)管理功能
* 用户操作支持网页端access token获取/刷新，及用户详情获取(仅非静默授权)
* 消息支持加密模式，配置了 EncodingAesKey 后调用 DecryptMsg 可得到解密消息体
* 消息支持签名校验 ValidateMsg
