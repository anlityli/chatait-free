// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package response

type ShopGoods struct {
	Id                string `json:"id"`
	Title             string `json:"title"`
	Content           string `json:"content"`
	FeatItems         string `json:"feat_items"`
	BuyType           int    `json:"buy_type"`
	ActiveLevelId     int    `json:"active_level_id"`
	ActiveExpireType  int    `json:"active_expire_type"`
	ActiveExpireValue int    `json:"active_expire_value"`
	BuyValue          int    `json:"buy_value"`
	MarketPrice       int    `json:"market_price"`
	RealPrice         int    `json:"real_price"`
	Status            int    `json:"status"`
	Sort              int    `json:"sort"`
	CreatedAt         int    `json:"created_at"`
	UpdatedAt         int    `json:"updated_at"`
}

type ShopGoodsList []*ShopGoods

type ShopOrder struct {
	Id          string `json:"id"`
	OrderSn     string `json:"order_sn"`
	UserId      string `json:"user_id"`
	Username    string `json:"username"`
	Nickname    string `json:"nickname"`
	OrderAmount int    `json:"order_amount"`
	PayAmount   int    `json:"pay_amount"`
	Status      int    `json:"status"`
	CreatedAt   int    `json:"created_at"`
	UpdatedAt   int    `json:"updated_at"`
	DueExpireAt int    `json:"due_expire_at"`
	ExpireAt    int    `json:"expire_at"`
}

type ShopOrderList []*ShopOrder
