// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package request

type ShopGoodsList struct {
}

type ShopGoodsDetail struct {
	GoodsId string `json:"goods_id" v:"required"`
}

type ShopOrderList struct {
	Status string `json:"status"`
}

type ShopOrderDetail struct {
	OrderId string `json:"order_id" v:"required"`
}

type ShopGenerateOrder struct {
	GoodsId  string `json:"goods_id" v:"required"`
	GoodsNum int    `json:"goods_num" v:"required"`
}

type ShopPayOrder struct {
	OrderId     string `json:"order_id" v:"required"`
	ConfigPayId int    `json:"config_pay_id" v:"required"`
	PayChannel  string `json:"pay_channel" v:"required"`
}
