// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package service

import (
	"github.com/anlityli/chatait-free/chatait-frontend-server/app/model/request"
	"github.com/anlityli/chatait-free/chatait-frontend-server/app/model/response"
	"github.com/anlityli/chatait-free/chatait-frontend-server/library/auth"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var User = &userService{}

type userService struct {
}

// Info 获取会员信息接口
func (s *userService) Info(r *ghttp.Request) (re *response.UserInfo, err error) {
	userId := auth.GetUserId(r)
	re = &response.UserInfo{}
	err = dao.User.As("u").LeftJoin(dao.UserInfo.Table+" ui", "u.id=ui.user_id").LeftJoin(dao.ConfigLevel.Table+" cl", "u.level_id=cl.id").Where("u.id=?", userId).Fields("u.*,ui.nickname,ui.avatar,cl.level_name").Scan(re)
	if err != nil {
		return nil, err
	}
	return re, nil
}

// EditNickname 编辑昵称
func (s *userService) EditNickname(r *ghttp.Request) (err error) {
	requestModel := &request.UserEditNickname{}
	if err = r.Parse(requestModel); err != nil {
		return err
	}
	userId := auth.GetUserId(r)
	if _, err = dao.UserInfo.Data(g.Map{
		"nickname": requestModel.Nickname,
	}).Where("user_id=?", userId).Update(); err != nil {
		return err
	}
	return nil
}
