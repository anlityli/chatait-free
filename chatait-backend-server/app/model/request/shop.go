// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package request

type ShopId struct {
	Id string `json:"id" v:"required"`
}

type ShopIds struct {
	Selected []string `json:"selected" v:"required"`
}

type ShopGoodsAdd struct {
	Title             string `json:"title" v:"required"`
	Content           string `json:"content"`
	FeatItems         string `json:"feat_items"`
	BuyType           int    `json:"buy_type" v:"required|in:1,2,3,4,5"`
	ActiveLevelId     int    `json:"active_level_id"`
	ActiveExpireType  int    `json:"active_expire_type"`
	ActiveExpireValue int    `json:"active_expire_value"`
	BuyValue          int    `json:"buy_value"`
	MarketPrice       int    `json:"market_price"`
	RealPrice         int    `json:"real_price"`
	Status            int    `json:"status"`
	Sort              int    `json:"sort"`
}

type ShopGoodsEdit struct {
	Id string `json:"id" v:"required"`
	ShopGoodsAdd
}

type ShopGoodsSort struct {
	Sort     []string `json:"sort" v:"required"`
	Page     int      `json:"page" v:"required"`
	PageSize int      `json:"page_size" v:"required"`
}

type ShopOrderStatus struct {
	Id     string `json:"id" v:"required"`
	Status int    `json:"status" v:"required"`
}
