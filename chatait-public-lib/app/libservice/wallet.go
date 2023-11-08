// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package libservice

import (
	"context"
	"errors"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/constant"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/model/entity"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/snowflake"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/xtime"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

var Wallet = &walletService{}

type walletService struct {
}

// ChangeWalletParam 改变会员钱包余额的参数
type ChangeWalletParam struct {
	UserId     int64
	WalletType string
	Amount     int
	Remark     string
	TargetType string
	TargetID   int64
	AdminName  string // 如果不是管理员操作，留空
}

// ChangeWalletBalance 改变会员的钱包余额
func (s *walletService) ChangeWalletBalance(ctx context.Context, tx *gdb.TX, changeParam *ChangeWalletParam) (err error) {
	nowTime := xtime.GetNowTime()
	year := xtime.GetNowFormat("Y")
	month := xtime.GetNowFormat("m")
	day := xtime.GetNowFormat("d")
	// 校验会员余额是否充足
	data, err := dao.Wallet.TX(tx).Ctx(ctx).Where("user_id=?", changeParam.UserId).One()
	if err != nil {
		return err
	}
	balance := 0
	total := 0

	if data.IsEmpty() {
		// 新建一个会员数据
		if _, err := dao.Wallet.TX(tx).Ctx(ctx).Data(g.Map{
			"user_id": changeParam.UserId,
		}).Insert(); err != nil {
			return err
		}
	} else {
		if changeParam.WalletType == constant.WalletTypeBalance {
			balance = gconv.Int(data[constant.WalletTypeBalance])
		} else if changeParam.WalletType == constant.WalletTypeGpt3 {
			balance = gconv.Int(data[constant.WalletTypeGpt3])
		} else if changeParam.WalletType == constant.WalletTypeGpt4 {
			balance = gconv.Int(data[constant.WalletTypeGpt4])
		} else if changeParam.WalletType == constant.WalletTypeMidjourney {
			balance = gconv.Int(data[constant.WalletTypeMidjourney])
		} else {
			return errors.New("余额格式不正确")
		}
	}

	if changeParam.Amount < 0 && balance < (-changeParam.Amount) {
		return errors.New("余额不足")
	}

	total = balance + changeParam.Amount

	// 修改余额
	whereStr := "user_id=?"
	isIncr := 1
	if changeParam.Amount < 0 {
		whereStr += " AND " + changeParam.WalletType + ">=" + gconv.String(-changeParam.Amount)
		isIncr = 0
	}
	if _, err := dao.Wallet.TX(tx).Ctx(ctx).Where(whereStr, changeParam.UserId).Data(changeParam.WalletType + "=" + changeParam.WalletType + "+" + gconv.String(changeParam.Amount)).Update(); err != nil {
		return err
	}

	// 增加流水
	var flowModel *gdb.Model
	if changeParam.WalletType == constant.WalletTypeBalance {
		flowModel = dao.WalletFlowBalance.TX(tx).Ctx(ctx)
	} else if changeParam.WalletType == constant.WalletTypeGpt3 {
		flowModel = dao.WalletFlowGpt3.TX(tx).Ctx(ctx)
	} else if changeParam.WalletType == constant.WalletTypeGpt4 {
		flowModel = dao.WalletFlowGpt4.TX(tx).Ctx(ctx)
	} else if changeParam.WalletType == constant.WalletTypeMidjourney {
		flowModel = dao.WalletFlowMidjourney.TX(tx).Ctx(ctx)
	} else {
		return errors.New("余额格式不正确")
	}
	flowID := snowflake.GenerateID()
	if _, err := flowModel.Data(g.Map{
		"id":          flowID,
		"user_id":     changeParam.UserId,
		"amount":      changeParam.Amount,
		"total":       total,
		"is_incr":     isIncr,
		"target_type": changeParam.TargetType,
		"target_id":   changeParam.TargetID,
		"remark":      changeParam.Remark,
		"admin_name":  changeParam.AdminName,
		"year":        year,
		"month":       month,
		"day":         day,
		"created_at":  nowTime,
	}).Insert(); err != nil {
		return err
	}

	return nil
}

// GetBalanceParams 获取钱包余额的参数
type GetBalanceParams struct {
	UserId     int64
	WalletType string
}

// GetBalance 获取钱包余额
func (s *walletService) GetBalance(params *GetBalanceParams) (balance int, err error) {
	// 校验会员余额是否充足
	data, err := dao.Wallet.Where("user_id=?", params.UserId).One()
	if err != nil {
		return 0, err
	}

	if data.IsEmpty() {
		// 新建一个会员数据
		if _, err := dao.Wallet.Data(g.Map{
			"user_id": params.UserId,
		}).Insert(); err != nil {
			return 0, err
		}
		return 0, nil
	} else {
		if params.WalletType == constant.WalletTypeBalance {
			balance = gconv.Int(data[constant.WalletTypeBalance])
		} else if params.WalletType == constant.WalletTypeGpt3 {
			balance = gconv.Int(data[constant.WalletTypeGpt3])
		} else if params.WalletType == constant.WalletTypeGpt4 {
			balance = gconv.Int(data[constant.WalletTypeGpt4])
		} else if params.WalletType == constant.WalletTypeMidjourney {
			balance = gconv.Int(data[constant.WalletTypeMidjourney])
		} else {
			return 0, errors.New("钱包格式不正确")
		}
	}

	return balance, nil
}

// GetAllBalance 获取会员的全部余额
func (s *walletService) GetAllBalance(userId int64) (re *entity.Wallet) {
	re = &entity.Wallet{}
	_ = dao.Wallet.Where("user_id=?", userId).Scan(re)
	return re
}
