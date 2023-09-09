// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package router

import (
	"github.com/anlityli/chatait-free/chatait-backend-server/app/api/shop"
	"github.com/anlityli/chatait-free/chatait-backend-server/router/utils"
)

var shopApi = &shop.Shop{}

var shopRouter = []*utils.RouterItem{
	{Method: "GET", Pattern: "/goods-list", Object: shopApi.GoodsList},
	{Method: "GET", Pattern: "/goods-one", Object: shopApi.GoodsOne},
	{Method: "POST", Pattern: "/goods-add", Object: shopApi.GoodsAdd},
	{Method: "POST", Pattern: "/goods-edit", Object: shopApi.GoodsEdit},
	{Method: "POST", Pattern: "/goods-sort", Object: shopApi.GoodsSort},
	{Method: "POST", Pattern: "/goods-delete", Object: shopApi.GoodsDelete},
	{Method: "GET", Pattern: "/order-list", Object: shopApi.OrderList},
	{Method: "POST", Pattern: "/order-status", Object: shopApi.OrderStatus},
}
