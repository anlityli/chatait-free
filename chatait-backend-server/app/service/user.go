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
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/libservice"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/model/entity"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/page"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/security"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
)

var User = &userService{}

type userService struct {
}

// List 会员列表
func (s *userService) List(r *ghttp.Request) (re *datalist.Result, err error) {
	columnsModel := &column.User{}
	listColumns := columnsModel.ListColumns()
	// 筛选
	whereAndParams, err := datalist.FilterWhereAndParams(r, listColumns)
	if err != nil {
		return nil, err
	}
	listData := &response.UserList{}
	// 获取会员数据
	data, err := page.Data(r, &page.Param{
		TableName:   dao.User.Table + " u",
		Where:       whereAndParams.Where,
		WhereParams: whereAndParams.Params,
		Join: page.ParamJoin{
			&page.ParamJoinItem{
				JoinType:  "leftJoin",
				JoinTable: dao.UserInfo.Table + " ui",
				On:        "u.id=ui.user_id",
			},
		},
		Field: "u.*, ui.avatar,ui.nickname",
	}, listData)
	if err != nil {
		return nil, err
	}
	return datalist.List(r, data, listColumns)
}

func (s *userService) Ban(r *ghttp.Request) error {
	requestModel := &request.UserBan{}
	if err := r.Parse(requestModel); err != nil {
		return err
	}
	for _, id := range requestModel.Selected {
		if _, err := dao.User.Data(g.Map{
			"is_ban": requestModel.IsBan,
		}).Where("id=?", id).Update(); err != nil {
			continue
		}
	}
	return nil
}

func (s *userService) ChangeLevel(r *ghttp.Request) (err error) {
	requestModel := &request.UserChangeLevel{}
	if err := r.Parse(requestModel); err != nil {
		return err
	}
	userData := &entity.User{}
	err = dao.User.Where("id=?", requestModel.UserId).Scan(userData)
	if err != nil {
		glog.Line().Debug(err)
		return err
	}
	if userData.LevelId == requestModel.LevelId && requestModel.LevelExpireDate == userData.LevelExpireDate.Format("Y-m-d") {
		return errors.New("级别和有效期没有发生变化")
	}
	var levelExpireDate *gtime.Time
	if requestModel.LevelExpireDate != "" {
		levelExpireDate = gtime.NewFromStr(requestModel.LevelExpireDate)
	}
	if err = g.DB().Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) (err error) {
		if err = libservice.User.ChangeLevel(ctx, tx, &libservice.ChangeLevelParams{
			UserId:          userData.Id,
			NewLevelId:      requestModel.LevelId,
			LevelExpireDate: levelExpireDate,
			AdminName:       auth.GetAdminName(r),
			Remark:          requestModel.Remark,
		}); err != nil {
			glog.Line().Debug(err)
			return err
		}
		return nil
	}); err != nil {
		glog.Line().Debug(err)
		return err
	}
	return nil
}

func (s *userService) ResetPassword(r *ghttp.Request) (err error) {
	requestModel := &request.UserResetPassword{}
	if err := r.Parse(requestModel); err != nil {
		return err
	}
	password, err := security.GeneratePassword(requestModel.Password)
	if err != nil {
		return err
	}
	if _, err = dao.User.Data(g.Map{
		"password": password,
	}).Where("id=?", requestModel.UserId).Update(); err != nil {
		glog.Line().Debug(err)
		return err
	}
	return nil
}

func (s *userService) SensitiveWordList(r *ghttp.Request) (re *datalist.Result, err error) {
	columnsModel := &column.User{}
	listColumns := columnsModel.SensitiveWordListColumns()
	// 筛选
	whereAndParams, err := datalist.FilterWhereAndParams(r, listColumns)
	if err != nil {
		return nil, err
	}
	listModel := &response.UserSensitiveWordList{}
	data, err := page.Data(r, &page.Param{
		TableName:   dao.UserSensitiveWord.Table + " usw",
		Where:       whereAndParams.Where,
		WhereParams: whereAndParams.Params,
		Join: page.ParamJoin{
			&page.ParamJoinItem{
				JoinType:  "leftJoin",
				JoinTable: dao.User.Table + " u",
				On:        "usw.user_id=u.id",
			},
			&page.ParamJoinItem{
				JoinType:  "leftJoin",
				JoinTable: dao.UserInfo.Table + " ui",
				On:        "usw.user_id=ui.user_id",
			},
		},
		Field:   "usw.*, u.username,ui.nickname",
		OrderBy: "usw.id Desc",
	}, listModel)
	if err != nil {
		return nil, err
	}
	return datalist.List(r, data, listColumns)
}

func (s *userService) SensitiveWordOne(r *ghttp.Request) (re *response.UserSensitiveWord, err error) {
	requestModel := &request.UserId{}
	if err := r.Parse(requestModel); err != nil {
		return nil, err
	}
	re = &response.UserSensitiveWord{}
	err = dao.UserSensitiveWord.As("usw").LeftJoin(dao.User.Table+" u", "usw.user_id=u.id").LeftJoin(dao.UserInfo.Table+" ui", "usw.user_id=ui.user_id").Where("usw.id=?", g.Slice{requestModel.Id}).Fields("usw.*, u.username,ui.nickname").Scan(re)
	if err != nil {
		return nil, err
	}
	return re, nil
}
