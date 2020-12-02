// 通用类型定义
package common

// BizResponse 基础返回类型，定义错误代码及错误消息
type BizResponse struct {
	ErrCode int    `json:"errcode,omitempty"` // 错误代码
	ErrMsg  string `json:"errmsg,omitempty"`  // 错误消息
}
