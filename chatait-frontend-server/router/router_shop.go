// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package router

import (
	"github.com/anlityli/chatait-free/chatait-frontend-server/app/api/shop"
	"github.com/anlityli/chatait-free/chatait-frontend-server/router/utils"
)

var shopApi = &shop.Shop{}

var shopRouter = []*utils.RouterItem{
	{Method: "GET", Pattern: "/goods-list", Object: shopApi.GoodsList, NoLogin: true},
	{Method: "GET", Pattern: "/goods-detail", Object: shopApi.GoodsDetail, NoLogin: true},
	{Method: "GET", Pattern: "/order-list", Object: shopApi.OrderList},
	{Method: "GET", Pattern: "/order-detail", Object: shopApi.OrderDetail},
	{Method: "POST", Pattern: "/order-calc-amount", Object: shopApi.OrderCalcAmount},
	{Method: "POST", Pattern: "/generate-order", Object: shopApi.GenerateOrder},
	{Method: "POST", Pattern: "/pay-order", Object: shopApi.PayOrder},
}
