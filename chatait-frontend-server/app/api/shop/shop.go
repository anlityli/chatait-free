// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package shop

import (
	"github.com/anlityli/chatait-free/chatait-frontend-server/app/service"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/notice"
	"github.com/gogf/gf/net/ghttp"
)

type Shop struct {
}

func (c *Shop) GoodsList(r *ghttp.Request) {
	if re, err := service.Shop.GoodsList(r); err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	} else {
		notice.Write(r, notice.NoError, re)
	}
}

func (c *Shop) GoodsDetail(r *ghttp.Request) {
	if re, err := service.Shop.GoodsDetail(r); err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	} else {
		notice.Write(r, notice.NoError, re)
	}
}

func (c *Shop) OrderList(r *ghttp.Request) {
	if re, err := service.Shop.OrderList(r); err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	} else {
		notice.Write(r, notice.NoError, re)
	}
}

func (c *Shop) OrderDetail(r *ghttp.Request) {
	if re, err := service.Shop.OrderDetail(r); err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	} else {
		notice.Write(r, notice.NoError, re)
	}
}

func (c *Shop) OrderCalcAmount(r *ghttp.Request) {
	if re, err := service.Shop.OrderCalcAmount(r); err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	} else {
		notice.Write(r, notice.NoError, re)
	}
}

func (c *Shop) GenerateOrder(r *ghttp.Request) {
	if re, err := service.Shop.GenerateOrder(r); err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	} else {
		notice.Write(r, notice.NoError, re)
	}
}

func (c *Shop) PayOrder(r *ghttp.Request) {
	if re, err := service.Shop.PayOrder(r); err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	} else {
		notice.Write(r, notice.NoError, re)
	}
}
