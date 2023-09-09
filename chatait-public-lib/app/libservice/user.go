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
	"github.com/anlityli/chatait-free/chatait-public-lib/library/helper"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/snowflake"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/xtime"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

var User = &userService{}

type userService struct {
}

type ChangeLevelParams struct {
	UserId          int64
	NewLevelId      int
	LevelExpireDate *gtime.Time
	AdminName       string
	Remark          string
}

// ChangeLevel 改变会员级别
func (s *userService) ChangeLevel(ctx context.Context, tx *gdb.TX, params *ChangeLevelParams) (err error) {
	if params.NewLevelId != constant.ConfigLevelMember && params.LevelExpireDate == nil {
		return errors.New("高级别会员必须要传有效期")
	}
	if params.NewLevelId == constant.ConfigLevelMember {
		params.LevelExpireDate = nil
	}
	userData := &entity.User{}
	err = dao.User.Ctx(ctx).TX(tx).Where("id=?", params.UserId).Scan(userData)
	if err != nil {
		glog.Line().Debug(err)
		return err
	}
	updateData := g.Map{
		"level_id":          params.NewLevelId,
		"level_expire_date": params.LevelExpireDate,
	}
	if params.LevelExpireDate != nil {
		updateData["level_expire_year"] = params.LevelExpireDate.Format("Y")
		updateData["level_expire_month"] = params.LevelExpireDate.Format("m")
		updateData["level_expire_day"] = params.LevelExpireDate.Format("d")
	} else {
		updateData["level_expire_year"] = 0
		updateData["level_expire_month"] = 0
		updateData["level_expire_day"] = 0
	}
	if _, err = dao.User.Ctx(ctx).TX(tx).Data(updateData).Where("id=?", params.UserId).Update(); err != nil {
		glog.Line().Debug(err)
		return err
	}
	if _, err = dao.UserLevelFlow.Ctx(ctx).TX(tx).Data(g.Map{
		"id":              snowflake.GenerateID(),
		"user_id":         params.UserId,
		"old_level_id":    userData.LevelId,
		"new_level_id":    params.NewLevelId,
		"old_expire_date": userData.LevelExpireDate,
		"new_expire_date": params.LevelExpireDate,
		"admin_name":      params.AdminName,
		"remark":          params.Remark,
		"created_at":      xtime.GetNowTime(),
	}).Insert(); err != nil {
		glog.Line().Debug(err)
		return err
	}
	return nil
}

var HandOutExecTime = int64(0)

// DailyMemberHandler 每日会员级别及提问次数处理
func (s *userService) DailyMemberHandler() {
	today := xtime.GetNow().StartOfDay()
	s.clearExpireMembers(today, 0)
	s.handOutValidMembers(today)
}

// clearExpireMembers 清理过期会员
func (s *userService) clearExpireMembers(today *gtime.Time, offset int) {
	todayDate := today.Format("Y-m-d")
	userList := &[]*entity.User{}
	err := dao.User.Where("level_id>? AND level_expire_date<? AND created_at<?", constant.ConfigLevelMember, todayDate, HandOutExecTime).Order("id ASC").Offset(offset).Limit(1000).Scan(userList)
	if err != nil {
		s.handOutErrLog(0, 0, err.Error())
		return
	}
	if len(*userList) > 0 {
		for _, userData := range *userList {
			if err = g.DB().Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) (err error) {
				err = s.ChangeLevel(ctx, tx, &ChangeLevelParams{
					UserId:          userData.Id,
					NewLevelId:      constant.ConfigLevelMember,
					LevelExpireDate: nil,
					Remark:          "会员过期",
				})
				if err != nil {
					return err
				}
				balanceData := Wallet.GetAllBalance(userData.Id)

				if balanceData.Gpt3 > 0 {
					err = Wallet.ChangeWalletBalance(ctx, tx, &ChangeWalletParam{
						UserId:     userData.Id,
						WalletType: constant.WalletTypeGpt3,
						Amount:     -gconv.Int(balanceData.Gpt3),
						Remark:     "会员级别过期扣除",
						TargetType: constant.WalletChangeTargetTypeLevelMonth,
						TargetID:   userData.Id,
					})
					if err != nil {
						return err
					}
				}
				return nil
			}); err != nil {
				s.handOutErrLog(userData.Id, userData.LevelId, err.Error())
				continue
			}
		}
		s.clearExpireMembers(today, offset+1000)
	}
}

// handOutValidMembers 发放有效会员的月次数
func (s *userService) handOutValidMembers(today *gtime.Time) {
	todayDay := gconv.Int(today.Format("d"))
	todayDate := today.Format("Y-m-d")
	currentMonthEndDay := xtime.GetNow().EndOfMonth().Format("Y-m-d")
	isMonthEnd := todayDate == currentMonthEndDay
	selectDayArr := []int{todayDay}
	if isMonthEnd {
		if todayDay == 28 {
			selectDayArr = []int{28, 29, 30, 31}
		} else if todayDay == 29 {
			selectDayArr = []int{29, 30, 31}
		} else if todayDay == 30 {
			selectDayArr = []int{30, 31}
		}
	}
	for _, selectDay := range selectDayArr {
		s.handout(selectDay, 0)
	}
}

// handout 发放子函数
func (s *userService) handout(day, offset int) {
	userList := &[]*entity.User{}
	err := dao.User.Where("level_id>? AND level_expire_day=? AND created_at<?", constant.ConfigLevelMember, day, HandOutExecTime).Order("id ASC").Offset(offset).Limit(1000).Scan(userList)
	if err != nil {
		s.handOutErrLog(0, 0, err.Error())
		return
	}
	if len(*userList) > 0 {
		for _, userData := range *userList {
			if err = g.DB().Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) (err error) {
				leveData, err := helper.GetConfigLevel(userData.LevelId)
				if err != nil {
					return err
				}
				balanceData := Wallet.GetAllBalance(userData.Id)
				if balanceData.Gpt3 > 0 {
					err = Wallet.ChangeWalletBalance(ctx, tx, &ChangeWalletParam{
						UserId:     userData.Id,
						WalletType: constant.WalletTypeGpt3,
						Amount:     -gconv.Int(balanceData.Gpt3),
						Remark:     "每月会员级别扣除",
						TargetType: constant.WalletChangeTargetTypeLevelMonth,
						TargetID:   userData.Id,
					})
					if err != nil {
						return err
					}
				}
				if leveData.MonthGpt3 > 0 {
					err = Wallet.ChangeWalletBalance(ctx, tx, &ChangeWalletParam{
						UserId:     userData.Id,
						WalletType: constant.WalletTypeGpt3,
						Amount:     leveData.MonthGpt3,
						Remark:     "每月会员级别充值",
						TargetType: constant.WalletChangeTargetTypeLevelMonth,
						TargetID:   userData.Id,
					})
					if err != nil {
						return err
					}
				}
				return nil
			}); err != nil {
				s.handOutErrLog(userData.Id, userData.LevelId, err.Error())
				continue
			}
		}
		s.handout(day, offset+1000)
	}
}

func (s *userService) handOutErrLog(userId int64, levelId int, errData string) {
	_, _ = dao.HandOutErrorFlow.Data(g.Map{
		"id":         snowflake.GenerateID(),
		"user_id":    userId,
		"level_id":   levelId,
		"error_data": errData,
		"created_at": xtime.GetNowTime(),
	}).Insert()
}
