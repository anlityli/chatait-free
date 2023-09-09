// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package service

import (
	"errors"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/constant"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/xtime"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Pay = &payService{}

type payService struct {
}

type AddPayFlowParams struct {
	R           *ghttp.Request
	FlowType    int
	TargetId    int64
	ConfigPayId int
	OrderAmount int
	PayChannel  string
}

// AddPayFlow 添加支付记录
func (s *payService) AddPayFlow(params *AddPayFlowParams) (re interface{}, err error) {
	// todo 不同的支付方式不同的记录
	return nil, errors.New("支付方式不存在")
}

// SetPayFlowFailed 设置为失败
func (s *payService) SetPayFlowFailed(payFlowId int64, paymentFailReason string) {
	_, _ = dao.PayFlow.Data(g.Map{
		"status":           constant.PayStatusFailed,
		"payment_response": paymentFailReason,
		"updated_at":       xtime.GetNowTime(),
	}).Where("id=?", payFlowId).Update()
}

type SetPayFlowSuccessParams struct {
	PayFlowId  int64
	NotifyData interface{}
}

// SetPayFlowSuccess 支付记录设置为成功
func (s *payService) SetPayFlowSuccess(params *SetPayFlowSuccessParams) (err error) {
	err = s.setPayFlowSuccessHandler(params)
	if err != nil {
		// 记录失败日志
		_, _ = dao.PayFlow.Data(g.Map{
			"notify_fail_reason": err.Error(),
		}).Where("id=?", params.PayFlowId).Update()
		return err
	}
	return nil
}

// setPayFlowSuccessHandler 支付记录设置为成功的逻辑
func (s *payService) setPayFlowSuccessHandler(params *SetPayFlowSuccessParams) (err error) {
	// todo 不同的支付方式不同的记录
	return nil
}
