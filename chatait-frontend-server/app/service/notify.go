// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package service

import (
	"github.com/anlityli/chatait-free/chatait-frontend-server/app/model/request"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/api/vmq"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
)

var Notify = &notifyService{}

type notifyService struct {
}

// Vmq 微免签回调
func (s *notifyService) Vmq(r *ghttp.Request) (err error) {
	reqData := r.GetBodyString()
	glog.Line(true).Debug("支付回调内容", reqData)
	requestModel := &request.NotifyVmq{}
	if err = r.Parse(requestModel); err != nil {
		return err
	}
	err = vmq.ValidateSign(requestModel)
	if err != nil {
		glog.Line(true).Println("验签失败", requestModel, err.Error())
		return err
	}
	return Pay.SetPayFlowSuccess(&SetPayFlowSuccessParams{
		PayFlowId:  gconv.Int64(requestModel.PayId),
		NotifyData: requestModel,
	})
}
