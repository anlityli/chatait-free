// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package libservice

import (
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/constant"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/model/entity"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/helper"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"net/url"
	"regexp"
)

var PaySmsForwarder = &paySmsForwarderService{}

type paySmsForwarderService struct {
}

type PaySmsForwarderCreateParams struct {
	PayFlowId int64
}

type PaySmsForwarderCreateResponse struct {
	PayAmount int    `json:"pay_amount"` // 实际付款金额
	PayUrl    string `json:"pay_url"`    // 支付二维码
	Timeout   int    `json:"timeout"`    // payFlow的超时时间(分钟)
}

func (s *paySmsForwarderService) Create(params *PaySmsForwarderCreateParams) (re *PaySmsForwarderCreateResponse, err error) {
	payFlowData := &entity.PayFlow{}
	err = dao.PayFlow.Where("id=?", params.PayFlowId).Scan(payFlowData)
	if err != nil {
		return nil, err
	}
	// 确定数据库中可用的金额二维码
	payAmount, err := s.getPayAmount(params.PayFlowId, payFlowData.OrderAmount)
	if err != nil {
		return nil, err
	}
	payQrData := &entity.ConfigPayQr{}
	err = dao.ConfigPayQr.Where("amount=?", payAmount).Scan(payQrData)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if payQrData.Id <= 0 {
		return nil, errors.New("支付通道已满，请稍后再试")
	}
	timeoutStr, err := Pay.OneConfigPayParamValue(constant.PayConfigSmsForwarder, "timeout")
	timeout := gconv.Int(timeoutStr)
	re = &PaySmsForwarderCreateResponse{
		PayAmount: payAmount,
		PayUrl:    payQrData.PayUrl,
		Timeout:   timeout,
	}
	return re, nil
}

func (s *paySmsForwarderService) getPayAmount(currentId int64, orderAmount int) (re int, err error) {
	waitPayList := &[]*entity.PayFlow{}
	err = dao.PayFlow.Where("id<>? AND config_pay_id=? AND status=?", currentId, constant.PayConfigSmsForwarder, constant.PayStatusInit).Scan(waitPayList)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	if len(*waitPayList) <= 0 {
		return orderAmount, nil
	}
	return s.loopPayAmount(orderAmount, waitPayList, 0)
}

func (s *paySmsForwarderService) loopPayAmount(orderAmount int, waitPayList *[]*entity.PayFlow, loopTimes int) (re int, err error) {
	if loopTimes >= 10 {
		return 0, errors.New("当前支付通道已满，请稍后再试")
	}
	for _, item := range *waitPayList {
		if orderAmount == item.PayAmount {
			return s.loopPayAmount(orderAmount-1, waitPayList, loopTimes+1)
		}
	}
	return orderAmount, nil
}

type PaySmsForwarderNotifyParams struct {
	From      string
	Content   string
	Timestamp string
	Sign      string
}

type PaySmsForwarderNotifyResponse struct {
	PayFlowId int64
	PayAmount int
}

func (s *paySmsForwarderService) Notify(params *PaySmsForwarderNotifyParams) (re *PaySmsForwarderNotifyResponse, err error) {
	// 验签
	if !s.ValidateSign(params) {
		return nil, errors.New("验签失败")
	}
	// 查看金额与待付款到哪条订单一致，则确定该订单
	payAmount := 0
	payChannel := 0
	//if gstr.Pos(gstr.ToLower(params.From), "alipay") != -1 && gstr.PosIRune(params.Content, "成功收款") != -1 {
	if gstr.PosIRune(params.Content, "成功收款") != -1 {
		reg := regexp.MustCompile(`收款(\d+\.\d+)元`)
		match := reg.FindStringSubmatch(params.Content)
		glog.Line(true).Debug(match)
		if len(match) > 0 {
			payAmount = helper.YuanToCent(match[1])
			payChannel = 2
		}
	}
	if payAmount <= 0 || payChannel <= 0 {
		return nil, errors.New("金额校验错误")
	}
	payFlowData := &entity.PayFlow{}
	err = dao.PayFlow.Where("config_pay_id=? AND pay_channel=? AND pay_amount=? AND status=?", constant.PayConfigSmsForwarder, payChannel, payAmount, constant.PayStatusInit).Scan(payFlowData)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if payFlowData.Id <= 0 {
		return nil, errors.New("支付记录不存在")
	}
	re = &PaySmsForwarderNotifyResponse{
		PayFlowId: payFlowData.Id,
		PayAmount: payAmount,
	}
	return re, nil
}

func (s *paySmsForwarderService) ValidateSign(params *PaySmsForwarderNotifyParams) bool {
	// 获取支付方式密钥
	secret, err := Pay.OneConfigPayParamValue(constant.PayConfigSmsForwarder, "secret")
	if err != nil {
		glog.Line(true).Debug("验签过程中密钥报错", err)
		return false
	}
	secretEnc := []byte(secret)
	stringToSign := fmt.Sprintf("%s\n%s", params.Timestamp, secret)
	stringToSignEnc := []byte(stringToSign)
	hmacCode := hmac.New(sha256.New, secretEnc)
	hmacCode.Write(stringToSignEnc)
	encodedHmacCode := base64.StdEncoding.EncodeToString(hmacCode.Sum(nil))
	sign := url.QueryEscape(encodedHmacCode)
	//glog.Line().Debug(sign)
	//glog.Line().Debug(params.Sign)
	return params.Sign == sign
}
