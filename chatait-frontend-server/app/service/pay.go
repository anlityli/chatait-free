// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package service

import (
	"context"
	"errors"
	"github.com/anlityli/chatait-free/chatait-frontend-server/app/model/request"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/constant"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/libservice"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/model/entity"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/api/vmq"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/helper"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/snowflake"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/xtime"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtimer"
	"github.com/gogf/gf/util/gconv"
	"time"
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
	nowTime := xtime.GetNowTime()
	// 支付方式是否存在
	configPayData, err := libservice.Pay.OneConfigPay(params.ConfigPayId)
	if err != nil {
		glog.Line(true).Debug(err)
		return nil, err
	}
	if configPayData.Status != 1 {
		return nil, errors.New("支付方式不存在")
	}
	var payChannel *libservice.ConfigPayChannelItem
	// 支付渠道是否可用
	if len(configPayData.PayChannel) > 0 {
		for _, channelItem := range configPayData.PayChannel {
			if channelItem.Channel == params.PayChannel {
				payChannel = channelItem
			}
		}
	}
	if payChannel != nil && payChannel.Status != 1 {
		return nil, errors.New("支付渠道不可用")
	}
	// 生成支付记录
	payFlowInsertId := snowflake.GenerateID()
	if _, err = dao.PayFlow.Data(g.Map{
		"id":            payFlowInsertId,
		"flow_type":     params.FlowType,
		"target_id":     params.TargetId,
		"config_pay_id": params.ConfigPayId,
		"pay_channel":   payChannel.Channel,
		"order_amount":  params.OrderAmount,
		"status":        constant.PayStatusInit,
		"created_at":    nowTime,
	}).Insert(); err != nil {
		return nil, err
	}
	// 请求支付接口
	if params.ConfigPayId == constant.PayConfigVmq {
		paymentData, err := vmq.CreateOrder(&vmq.CreateOrderParams{
			R:           params.R,
			PayFlowId:   payFlowInsertId,
			PayChannel:  gconv.Int(payChannel.Channel),
			OrderAmount: params.OrderAmount,
			Param:       "",
		})
		if err != nil {
			glog.Line(true).Debug(err)
			s.SetPayFlowFailed(payFlowInsertId, err.Error())
			return nil, err
		}
		// 把接口的结果写入数据库
		if _, err := dao.PayFlow.Data(g.Map{
			"payment_response": paymentData.ResponseData,
			"pay_amount":       helper.YuanToCent(paymentData.ReallyPrice),
			"updated_at":       xtime.GetNowTime(),
		}).Where("id=?", payFlowInsertId).Update(); err != nil {
			glog.Line(true).Debug(err)
			s.SetPayFlowFailed(payFlowInsertId, err.Error())
			return nil, err
		}
		// 如果超时时间大于0，需要到期取消支付记录
		if paymentData.TimeOut > 0 {
			gtimer.AddOnce(time.Duration(paymentData.TimeOut)*time.Minute, func() {
				payFlowData := &entity.PayFlow{}
				err := dao.PayFlow.Where("id=?", payFlowInsertId).Scan(payFlowData)
				if err != nil {
					glog.Line(true).Println(err)
					return
				}
				if payFlowData.Status != constant.PayStatusInit {
					return
				}
				if _, err = dao.PayFlow.Data(g.Map{
					"status":     constant.PayStatusExpired,
					"expired_at": xtime.GetNowTime(),
				}).Where("id=?", payFlowInsertId).Update(); err != nil {
					glog.Line(true).Println(err)
					return
				}
			})
		}
		re = paymentData
	} else {
		s.SetPayFlowFailed(payFlowInsertId, "支付方式不存在")
		return nil, errors.New("支付方式不存在")
	}
	return re, nil
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
	flowData := &entity.PayFlow{}
	err = dao.PayFlow.Where("id=?", params.PayFlowId).Scan(flowData)
	if err != nil {
		return err
	}
	// 已经支付过，直接返回成功
	if flowData.Status == constant.PayStatusPaid {
		return nil
	}
	if flowData.Status != constant.PayStatusInit {
		return errors.New("当前状态已不能支付")
	}
	notifyResponse := ""
	payAmount := 0
	if flowData.ConfigPayId == constant.PayConfigVmq {
		notifyData := params.NotifyData.(*request.NotifyVmq)
		if flowData.OrderAmount != helper.YuanToCent(notifyData.Price) {
			return errors.New("金额不匹配")
		}
		notifyResponseBytes, err := gjson.Encode(notifyData)
		if err == nil {
			notifyResponse = gconv.String(notifyResponseBytes)
		}
		payAmount = helper.YuanToCent(notifyData.ReallyPrice)
	} else {
		return errors.New("支付方式不存在")
	}
	if err := g.DB().Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) (err error) {
		if _, err = dao.PayFlow.Ctx(ctx).TX(tx).Data(g.Map{
			"pay_amount":      payAmount,
			"status":          constant.PayStatusPaid,
			"notify_response": notifyResponse,
			"paid_at":         xtime.GetNowTime(),
		}).Where("id=?", params.PayFlowId).Update(); err != nil {
			return err
		}
		if flowData.FlowType == constant.PayFlowTypeShopOrder {
			// 把订单设置为已完成
			err = libservice.Shop.SetOrderPaid(ctx, tx, flowData.TargetId, payAmount)
			if err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}
