package v3

type errorResponse struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Detail  struct {
		Field    string `json:"field,omitempty"`
		Value    string `json:"value,omitempty"`
		Issue    string `json:"issue,omitempty"`
		Location string `json:"location,omitempty"`
	} `json:"detail,omitempty"`
}

// Native下单响应
type NativeOrderResponse struct {
	errorResponse
	CodeUrl string `json:"code_url"`
}

// app/JsApi/小程序下单响应
type PrepayIdResponse struct {
	errorResponse
	PrepayId string `json:"prepay_id"`
}

// h5下单响应
type H5OrderResponse struct {
	errorResponse
	H5Url string `json:"h5_url"`
}

// 服务商模式订单查询响应
type QueryPartnerOrderResponse struct {
	errorResponse
	SpAppId        string            `json:"sp_appid" validate:"required"`     // 服务商申请的公众号或移动应用appid
	SpMchId        string            `json:"sp_mchid" validate:"required"`     // 服务商户号，由微信支付生成并下发
	SubAppId       string            `json:"sub_appid,omitempty"`              // 子商户申请的公众号或移动应用appid
	SubMchId       string            `json:"sub_mchid" validate:"required"`    // 子商户的商户号，有微信支付生成并下发
	OutTradeNo     string            `json:"out_trade_no" validate:"required"` // 商户订单号
	TransactionId  string            `json:"transaction_id"`                   // 微信支付订单号
	TradeType      string            `json:"trade_type"`                       // 交易类型
	TradeState     string            `json:"trade_state"`                      // 交易状态
	TradeStateDesc string            `json:"trade_state_desc"`                 // 交易状态描述
	BankType       string            `json:"bank_type"`                        // 付款银行
	Attach         string            `json:"attach"`                           // 附加数据
	SuccessTime    string            `json:"success_time"`                     // 支付完成时间
	Payer          Payer             `json:"payer"`                            // 支付者
	Amount         Amount            `json:"amount"`                           // 订单金额
	SceneInfo      QuerySceneInfo    `json:"scene_info,omitempty"`             // 场景信息
	Promotion      []PromotionDetail `json:"promotion_detail,omitempty"`       // 优惠功能
}

// 退款单－订单金额－响应
type RefundAmountReseponse struct {
	Refund         int    `json:"refund"`          // 退款金额
	PayerRefund    int    `json:"payer_refund"`    // 用户退款金额
	DiscountRefund int    `json:"discount_refund"` // 优惠退款金额
	Currency       string `json:"currency"`        // 退款币种
}

// 退款单-优惠退款详情-响应
type PromotionDetailResponse struct {
	PromotionId  string `json:"promotion_id"`  // 券ID
	Scope        string `json:"scope"`         // 优惠范围
	Type         string `json:"type"`          // 优惠类型
	Amount       int    `json:"amount"`        // 优惠券面额
	RefundAmount int    `json:"refund_amount"` // 优惠退款金额
}

// 退款单响应
type RefundOrderResponse struct {
	errorResponse
	RefundId        string                  `json:"refund_id"`        // 微信退款单号
	OutRefundNo     string                  `json:"out_refund_no"`    // 商户退款单号
	CreateTime      string                  `json:"create_time"`      // 退款创建时间
	Amount          RefundAmountReseponse   `json:"amount"`           // 订单金额
	PromotionDetail PromotionDetailResponse `json:"promotion_detail"` // 优惠退款详情
}

type RefundOrderQueryResponse struct {
	errorResponse
	RefundId            string                    `json:"refund_id"`                        // 微信退款单号
	OutRefundNo         string                    `json:"out_refund_no"`                    // 商户退款单号
	OutTradeNo          string                    `json:"out_trade_no" validate:"required"` // 商户订单号
	TransactionId       string                    `json:"transaction_id"`                   // 微信支付订单号
	Channel             string                    `json:"channel"`                          // 退款渠道
	UserReceivedAccount string                    `json:"user_received_account"`            // 退款入账账户
	SuccessTime         string                    `json:"success_time"`                     // 退款成功时间
	CreateTime          string                    `json:"create_time"`                      // 退款创建时间
	Status              string                    `json:"status"`                           // 退款状态
	Amount              RefundAmountReseponse     `json:"amount"`                           // 订单金额
	PromotionDetail     []PromotionDetailResponse `json:"promotion_detail"`                 // 优惠退款详情
}
