// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package constant

const (
	ShopGoodsBuyTypeLevel      = 1 // 购买级别
	ShopGoodsBuyTypeBalance    = 2 // 购买balance
	ShopGoodsBuyTypeGpt3       = 3 // 购买gpt3
	ShopGoodsBuyTypeGpt4       = 4 // 购买gpt4
	ShopGoodsBuyTypeMidjourney = 5 // 购买midjourney
)

const (
	ShopOrderStatusInit     = 0 // 订单创建
	ShopOrderStatusPaid     = 1 // 已支付
	ShopOrderStatusShipped  = 2 // 已发货
	ShopOrderStatusReceived = 3 // 已收货
	ShopOrderStatusFinish   = 4 // 已完成
	ShopOrderStatusCancel   = 9 // 已取消
)

const (
	ShopGoodsActiveExpireTypeNone  = 0 // 激活类型无
	ShopGoodsActiveExpireTypeDay   = 1 // 激活类型天
	ShopGoodsActiveExpireTypeMonth = 2 // 激活类型月
	ShopGoodsActiveExpireTypeYear  = 3 // 激活类型年
)
