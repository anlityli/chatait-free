// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package vmq

import "github.com/gogf/gf/net/ghttp"

type CreateOrderParams struct {
	R           *ghttp.Request
	PayFlowId   int64
	PayChannel  int
	OrderAmount int
	Param       string // 需要原样返回的参数
}

type CreateOrderResponse struct {
	PayId        string `json:"payId"`        // 商户订单号
	OrderId      string `json:"orderId"`      // 云端订单号，可用于查询订单是否支付成功
	PayType      int    `json:"payType"`      // 微信支付为1 支付宝支付为2
	Price        string `json:"price"`        // 订单金额
	ReallyPrice  string `json:"reallyPrice"`  // 实际需付金额
	PayUrl       string `json:"payUrl"`       // 支付二维码内容
	IsAuto       int    `json:"isAuto"`       // 1需要手动输入金额 0扫码后自动输入金额
	State        int    `json:"state"`        // 订单状态：-1|订单过期 0|等待支付 1|完成 2|支付完成但通知失败
	TimeOut      int    `json:"timeOut"`      // 订单有效时间（分钟）
	Date         string `json:"date"`         // 订单创建时间时间戳（10位）
	ResponseData string `json:"responseData"` // 支付接口返回原文
}

type GetStateParams struct {
	R *ghttp.Request
}

type GetStateResponse struct {
	LastPay      int64  `json:"Lastpay"`      // 最后一次监控到支付的时间戳（10位）
	LastHeart    int64  `json:"lastheart"`    // 最后一次监控端向服务器发送心跳的时间戳（10位）
	JKState      int    `json:"jkstate"`      // 监控端状态 1|在线 0|掉线 -1|还未绑定监控端
	ResponseData string `json:"responseData"` // 支付接口返回原文
}
