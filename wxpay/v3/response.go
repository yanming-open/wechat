package v3

type errorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Detail  struct {
		Field    string `json:"field"`
		Value    string `json:"value"`
		Issue    string `json:"issue"`
		Location string `json:"location"`
	} `json:"detail"`
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
