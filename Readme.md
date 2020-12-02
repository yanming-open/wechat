# 微信开发sdk

* mp 公众号接口


# ChangeLog

## v1.0.0

* 完成公众号消息、用户、标签、素材(永久/临时)管理功能
* 用户操作支持网页端access token获取/刷新，及用户详情获取(仅非静默授权)
* 消息支持加密模式，配置了 EncodingAesKey 后调用 DecryptMsg 可得到解密消息体
* 消息支持签名校验 ValidateMsg
