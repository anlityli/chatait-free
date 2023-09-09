// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package service

import (
	"errors"
	"github.com/anlityli/chatait-free/chatait-frontend-server/app/model/request"
	"github.com/anlityli/chatait-free/chatait-frontend-server/app/model/response"
	"github.com/anlityli/chatait-free/chatait-frontend-server/library/auth"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/constant"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/page"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/xtime"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var FinanceService = financeService{}

type financeService struct {
}

// WalletFlowList 钱包流水列表
func (s *financeService) WalletFlowList(r *ghttp.Request) (re *page.Response, err error) {
	requestModel := &request.FinanceWalletFlowList{}
	if err = r.Parse(requestModel); err != nil {
		return nil, err
	}
	var tableName string
	if requestModel.WalletType == constant.WalletTypeBalance {
		tableName = dao.WalletFlowBalance.Table
	} else if requestModel.WalletType == constant.WalletTypeGpt3 {
		tableName = dao.WalletFlowGpt3.Table
	} else {
		return nil, errors.New("钱包类型不正确")
	}
	userId := auth.GetUserId(r)
	// 6个月前的一天的开始时间
	lastTime := xtime.GetNow().AddDate(0, -6, 0).StartOfDay().Timestamp()
	listData := &response.FinanceWalletFLowList{}
	return page.Data(r, &page.Param{
		TableName:   tableName,
		Where:       "user_id=? AND created_at>?",
		WhereParams: g.Slice{userId, lastTime},
		OrderBy:     "id DESC",
	}, listData)
}

// WalletInfo 钱包信息
func (s *financeService) WalletInfo(r *ghttp.Request) (re *response.FinanceWalletInfo, err error) {
	userId := auth.GetUserId(r)
	re = &response.FinanceWalletInfo{}
	err = dao.Wallet.Where("user_id=?", userId).Scan(re)
	if err != nil {
		return nil, err
	}
	return re, nil
}
