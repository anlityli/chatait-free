// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package response

type ShopGoods struct {
	Id               string               `json:"id"`
	Title            string               `json:"title"`
	Content          string               `json:"content"`
	FeatItems        string               `json:"feat_items"`
	FeatItemsSlice   []*ShopGoodsFeatItem `json:"feat_items_slice"`
	BuyType          int                  `json:"buy_type"`
	ActiveLevelId    int                  `json:"active_level_id"`
	ActiveExpireType int                  `json:"active_expire_type"`
	BuyValue         int                  `json:"buy_value"`
	MarketPrice      int                  `json:"market_price"`
	RealPrice        int                  `json:"real_price"`
	CreatedAt        int                  `json:"created_at"`
	UpdatedAt        int                  `json:"updated_at"`
}

type ShopGoodsFeatItem struct {
	Icon string `json:"icon"`
	Text string `json:"text"`
}

type ShopGoodsList []*ShopGoods

type ShopCalcOrderAmount struct {
	OrderAmount int `json:"order_amount"`
}

type ShopOrder struct {
	Id             string             `json:"id"`
	OrderSn        string             `json:"order_sn"`
	UserId         string             `json:"user_id"`
	OrderAmount    int                `json:"order_amount"`
	PayAmount      int                `json:"pay_amount"`
	Status         int                `json:"status"`
	CreatedAt      int                `json:"created_at"`
	UpdatedAt      int                `json:"updated_at"`
	DueExpireAt    int                `json:"due_expire_at"`
	ExpiredAt      int                `json:"expired_at"`
	OrderGoodsList ShopOrderGoodsList `json:"order_goods_list"`
}

type ShopOrderList []*ShopOrder

type ShopOrderGoods struct {
	Id            string     `json:"id"`
	OrderId       string     `json:"order_id"`
	UserId        string     `json:"user_id"`
	GoodsId       string     `json:"goods_id"`
	GoodsNum      int        `json:"goods_num"`
	GoodsSnapshot *ShopGoods `json:"goods_snapshot"`
}

type ShopOrderGoodsList []*ShopOrderGoods

type ShopPayOrder struct {
	PayAmount   int    `json:"pay_amount"`    // 实际付款金额
	PayUrl      string `json:"pay_url"`       // 支付二维码
	Timeout     int    `json:"timeout"`       // payFlow的超时时间(分钟)
	DueExpireAt int    `json:"due_expire_at"` // 应过期时间
}
