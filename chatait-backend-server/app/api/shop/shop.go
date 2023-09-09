// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package shop

import (
	"github.com/anlityli/chatait-free/chatait-backend-server/app/service"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/notice"
	"github.com/gogf/gf/net/ghttp"
)

type Shop struct {
}

func (c *Shop) GoodsList(r *ghttp.Request) {
	re, err := service.Shop.GoodsList(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, re)
}

func (c *Shop) GoodsOne(r *ghttp.Request) {
	re, err := service.Shop.GoodsOne(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, re)
}

func (c *Shop) GoodsAdd(r *ghttp.Request) {
	err := service.Shop.GoodsAdd(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, "操作成功")
}

func (c *Shop) GoodsEdit(r *ghttp.Request) {
	err := service.Shop.GoodsEdit(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, "操作成功")
}

func (c *Shop) GoodsSort(r *ghttp.Request) {
	err := service.Shop.GoodsSort(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, "操作成功")
}

func (c *Shop) GoodsDelete(r *ghttp.Request) {
	err := service.Shop.GoodsDelete(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, "操作成功")
}

func (c *Shop) OrderList(r *ghttp.Request) {
	re, err := service.Shop.OrderList(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, re)
}

func (c *Shop) OrderStatus(r *ghttp.Request) {
	err := service.Shop.OrderStatus(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, "操作成功")
}
