package v3

// 结算信息
type SettleInfo struct {
	ProfitSharing bool `json:"profit_sharing,omitempty"`
}

// 订单金额
type Amount struct {
	Total         int    `json:"total" validate:"required"` // 总金额
	Currency      string `json:"currency,omitempty"`        // 货币类型
	PayerTotal    int    `json:"payer_total,omitempty"`     // 用户支付金额
	PayerCurrency string `json:"payer_currency,omitempty"`  // 用户支付币种
}

// 商品详情
type Goods struct {
	MerchantGoodsId  string `json:"merchant_goods_id" validate:"required"`
	WechatpayGoodsId string `json:"wechatpay_goods_id,omitempty"`
	GoodsName        string `json:"goods_name,omitempty"`
	Quantity         int    `json:"quantity" validate:"required"`   // 商品数量
	UnitPrice        int    `json:"unit_price" validate:"required"` // 商品单价
}

// 优惠功能
type Discount struct {
	CostPrice   int     `json:"cost_price,omitempty"`
	InvoiceId   string  `json:"invoice_id,omitempty"`
	GoodsDetail []Goods `json:"goods_detail,omitempty"`
}

// 场景信息
type SceneInfo struct {
	PayerClientIp string    `json:"payer_client_ip" validate:"required"`
	DeviceId      string    `json:"device_id,omitempty"`
	StoreInfo     StoreInfo `json:"store_info,omitempty"`
}

// 商户门店信息
type StoreInfo struct {
	Id       string `json:"id" validate:"required"`
	Name     string `json:"name,omitempty"`
	AreaCode string `json:"area_code,omitempty"`
	Address  string `json:"address,omitempty"`
}

// 服务商下单
type PartnerOrder struct {
	SpAppId     string     `json:"sp_appid" validate:"required"`  // 服务商申请的公众号或移动应用appid
	SpMchId     string     `json:"sp_mchid" validate:"required"`  // 服务商户号，由微信支付生成并下发
	SubAppId    string     `json:"sub_appid,omitempty"`           // 子商户申请的公众号或移动应用appid
	SubMchId    string     `json:"sub_mchid" validate:"required"` // 子商户的商户号，有微信支付生成并下发
	Description string     `json:"description" validate:"required"`
	OutTradeNo  string     `json:"out_trade_no" validate:"required"`
	TimeExpire  string     `json:"time_expire,omitempty"`
	Attach      string     `json:"attach,omitempty"`
	NotifyUrl   string     `json:"notify_url" validate:"required"`
	GoodsTag    string     `json:"goods_tag,omitempty"`
	SettleInfo  SettleInfo `json:"settle_info,omitempty"`
	Amount      Amount     `json:"amount" validate:"required"`
	Detail      Discount   `json:"detail,omitempty"`
	SceneInfo   SceneInfo  `json:"scene_info",omitempty`
}

// 支付者信息
type Payer struct {
	SpOpenId  string `json:"sp_openid"`  // 用户服务标识
	SubOpenId string `json:"sub_openid"` // 用户子标识
}

// 订单查询时返回的场景信息
type QuerySceneInfo struct {
	DeviceId string `json:"device_id,omitempty"`
}

// 订单查询时返回的商品信息
type QueryGoods struct {
	GoodsId         string `json:"goods_id" validate:"required"`         // 商品编码
	Quantity        int    `json:"quantity" validate:"required"`         // 商品数量
	UnitPrice       int    `json:"unit_price" validate:"required"`       // 商品单价
	DiscountAmmount int    `json:"discount_ammount" validate:"required"` // 商品优惠金额
	GoodsRemark     string `json:"goods_remark"`                         // 商品备注
}

// 优惠功能 订单查询时返回
type PromotionDetail struct {
	CouponId            string       `json:"coupon_id"`            // 券ID
	Name                string       `json:"name"`                 // 优惠名称
	Scope               string       `json:"scope"`                // 优惠范围
	Type                string       `json:"type"`                 // 优惠类型
	Amount              int          `json:"amount"`               // 优惠券面额
	StockId             string       `json:"stock_id"`             // 活动ID
	WeChatPayContribute int          `json:"wechatpay_contribute"` // 微信出资
	MerchantContribute  int          `json:"merchant_contribute"`  // 商户出资
	OtherContribute     int          `json:"other_contribute"`     // 其他出资
	Currency            string       `json:"currency"`             // 优惠币种
	GoodsDetail         []QueryGoods `json:"goods_detail"`         // 单品列表
}

// 退款单－订单金额
type RefundAmount struct {
	Refund   int    `json:"refund" validate:"required"`   // 退款金额
	Total    int    `json:"total" validate:"required"`    // 总金额
	Currency string `json:"currency" validate:"required"` // 货币类型

}

// 订单退款申请
type RefundsOrder struct {
	SubMchId      string       `json:"sub_mchid" validate:"required"`
	SpAppId       string       `json:"sp_appid" validate:"required"`
	SubAppId      string       `json:"sub_appid,omitempty"`
	TransactionId string       `json:"transaction_id,omitempty"`
	OutTradeNo    string       `json:"out_trade_no,omitempty"`
	OutRefundNo   string       `json:"out_refund_no" validate:"required"`
	Reason        string       `json:"reason"`
	Amount        RefundAmount `json:"amount" validate:"required"`
	NotifyUrl     string       `json:"notify_url,omitempty"`
}
