// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package service

import (
	"context"
	"errors"
	"github.com/anlityli/chatait-free/chatait-backend-server/app/model/request"
	"github.com/anlityli/chatait-free/chatait-backend-server/app/model/response"
	"github.com/anlityli/chatait-free/chatait-backend-server/app/service/column"
	"github.com/anlityli/chatait-free/chatait-backend-server/library/auth"
	"github.com/anlityli/chatait-free/chatait-backend-server/library/datalist"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/constant"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/libservice"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/model/entity"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/page"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var Finance = &financeService{}

type financeService struct {
}

func (s *financeService) WalletList(r *ghttp.Request) (re *datalist.Result, err error) {
	columnsModel := &column.Finance{}
	listColumns := columnsModel.WalletListColumns()
	// 筛选
	whereAndParams, err := datalist.FilterWhereAndParams(r, listColumns)
	if err != nil {
		return nil, err
	}
	listData := &response.FinanceWalletList{}
	// 获取会员数据
	data, err := page.Data(r, &page.Param{
		TableName:   dao.Wallet.Table + " w",
		Where:       whereAndParams.Where,
		WhereParams: whereAndParams.Params,
		Join: page.ParamJoin{
			&page.ParamJoinItem{
				JoinType:  "leftJoin",
				JoinTable: dao.User.Table + " u",
				On:        "u.id=w.user_id",
			},
			&page.ParamJoinItem{
				JoinType:  "leftJoin",
				JoinTable: dao.UserInfo.Table + " ui",
				On:        "ui.user_id=w.user_id",
			},
		},
		Field: "w.*, u.username,ui.avatar,ui.nickname",
	}, listData)
	if err != nil {
		return nil, err
	}
	return datalist.List(r, data, listColumns)
}

func (s *financeService) WalletChange(r *ghttp.Request) (err error) {
	requestModel := &request.FinanceWalletChange{}
	if err := r.Parse(requestModel); err != nil {
		return err
	}
	if requestModel.Amount == 0 {
		return errors.New("金额不得为0")
	}
	if requestModel.Remark == "" && requestModel.Amount > 0 {
		requestModel.Remark = "系统充入"
	} else if requestModel.Remark == "" && requestModel.Amount < 0 {
		requestModel.Remark = "系统扣除"
	}
	userData := &entity.User{}
	err = dao.User.Where("id=?", requestModel.UserId).Scan(userData)
	if err != nil {
		return err
	}
	if err = g.DB().Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) (err error) {
		if err = libservice.Wallet.ChangeWalletBalance(ctx, tx, &libservice.ChangeWalletParam{
			UserId:     gconv.Int64(requestModel.UserId),
			WalletType: requestModel.WalletType,
			Amount:     requestModel.Amount,
			Remark:     requestModel.Remark,
			TargetType: constant.WalletChangeTargetTypeBackend,
			AdminName:  auth.GetAdminName(r),
		}); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (s *financeService) WalletFlowList(r *ghttp.Request, walletType string) (re *datalist.Result, err error) {
	columnsModel := &column.Finance{}
	listColumns := columnsModel.WalletFlowListColumns(walletType)
	// 筛选
	whereAndParams, err := datalist.FilterWhereAndParams(r, listColumns)
	if err != nil {
		return nil, err
	}
	listData := &response.FinanceWalletFlowList{}
	table := ""
	if walletType == constant.WalletTypeBalance {
		table = dao.WalletFlowBalance.Table + " f"
	} else if walletType == constant.WalletTypeGpt3 {
		table = dao.WalletFlowGpt3.Table + " f"
	} else if walletType == constant.WalletTypeGpt4 {
		table = dao.WalletFlowGpt4.Table + " f"
	} else if walletType == constant.WalletTypeMidjourney {
		table = dao.WalletFlowMidjourney.Table + " f"
	} else {
		return nil, errors.New("钱包类型不存在")
	}
	// 获取会员数据
	data, err := page.Data(r, &page.Param{
		TableName:   table,
		Where:       whereAndParams.Where,
		WhereParams: whereAndParams.Params,
		Join: page.ParamJoin{
			&page.ParamJoinItem{
				JoinType:  "leftJoin",
				JoinTable: dao.User.Table + " u",
				On:        "u.id=f.user_id",
			},
			&page.ParamJoinItem{
				JoinType:  "leftJoin",
				JoinTable: dao.UserInfo.Table + " ui",
				On:        "ui.user_id=f.user_id",
			},
		},
		Field:   "f.*, u.username,ui.avatar,ui.nickname",
		OrderBy: "f.id DESC",
	}, listData)
	if err != nil {
		return nil, err
	}
	return datalist.List(r, data, listColumns)
}
