// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package service

import (
	"database/sql"
	"errors"
	"github.com/anlityli/chatait-free/chatait-backend-server/app/model/request"
	"github.com/anlityli/chatait-free/chatait-backend-server/app/model/response"
	"github.com/anlityli/chatait-free/chatait-backend-server/library/auth"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/model/entity"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/security"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/xtime"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var Oauth = &oauthService{}

type oauthService struct {
}

// LoginWithUsername 用户名密码登录
func (s *oauthService) LoginWithUsername(r *ghttp.Request) (tokenRe *response.OauthLogin, err error) {
	requestModel := &request.OauthLogin{}
	if err := r.Parse(requestModel); err != nil {
		return nil, err
	}
	adminModel := &entity.Admin{}
	if err := dao.Admin.Where(dao.Admin.Columns.AdminName+"=?", requestModel.AdminName).Scan(adminModel); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("用户名或密码错误")
		}
		return nil, err
	}
	// 查看管理员当前登录的IP是否在允许的IP段里
	if adminModel.BindIp != "" {
		//bindIPJson, err := gjson.Decode(adminModel.BindIp)
		//if err != nil {
		//	return nil, err
		//}
		// bindIPArr := gconv.SliceStr(bindIPJson)
		//if !helper.StrInArr(bindIPArr, web.GetClientIP(r)) {
		//	return nil, errors.New("您的IP禁止登录")
		//}
	}
	if !security.ValidatePassword(requestModel.Password, adminModel.PasswordHash) {
		return nil, errors.New("用户名或密码错误")
	}
	return s.generateToken(adminModel)
}

// generateToken 登录成功生成的token结果
func (s *oauthService) generateToken(adminModel *entity.Admin) (tokenRe *response.OauthLogin, err error) {
	var tokenStr string
	var exp int64
	var expIn int64
	var refreshToken string
	var refreshExp int64
	var refreshExpIn int64
	nowTime := xtime.GetNowTime()
	// 生成 accessToken
	tokenStr, exp, expIn, err = security.GenerateAdminAccessToken(adminModel)
	if err != nil {
		return nil, errors.New("token生成失败")
	}
	// 生成 refreshToken
	refreshToken, refreshExp, refreshExpIn, err = security.GenerateAdminRefreshToken(adminModel)
	if err != nil {
		return nil, errors.New("refresh_token生成失败")
	}
	updateData := g.Map{
		"last_login_at": nowTime,
	}
	if _, err := dao.Admin.Data(updateData).Where(dao.Admin.Columns.Id+"=?", adminModel.Id).Update(); err != nil {
		return nil, errors.New("更新会员失败")
	}
	tokenRe = &response.OauthLogin{}
	tokenRe.AccessToken = tokenStr
	tokenRe.AccessTokenExpire = gconv.Int(exp)
	tokenRe.AccessTokenExpireIn = gconv.Int(expIn)
	tokenRe.RefreshToken = refreshToken
	tokenRe.RefreshTokenExpire = gconv.Int(refreshExp)
	tokenRe.RefreshTokenExpireIn = gconv.Int(refreshExpIn)
	return tokenRe, nil
}

// RefreshToken 刷新access_token
func (s *oauthService) RefreshToken(r *ghttp.Request) (tokenRe *response.OauthLogin, err error) {
	userModel, err := auth.ValidateRefresh(r)
	if err != nil {
		return nil, err
	}
	// 生成 accessToken
	tokenStr, exp, expIn, err := security.GenerateAdminAccessToken(userModel)
	if err != nil {
		return nil, errors.New("token生成失败")
	}
	// 生成 refreshToken
	refreshToken, refreshExp, refreshExpIn, err := security.GenerateAdminRefreshToken(userModel)
	if err != nil {
		return nil, errors.New("refresh_token生成失败")
	}
	tokenRe = &response.OauthLogin{}
	tokenRe.AccessToken = tokenStr
	tokenRe.AccessTokenExpire = gconv.Int(exp)
	tokenRe.AccessTokenExpireIn = gconv.Int(expIn)
	tokenRe.RefreshToken = refreshToken
	tokenRe.RefreshTokenExpire = gconv.Int(refreshExp)
	tokenRe.RefreshTokenExpireIn = gconv.Int(refreshExpIn)
	return tokenRe, nil
}

// Info 获取管理员信息
func (s *oauthService) Info(r *ghttp.Request) (re *response.OauthInfo, err error) {
	re = &response.OauthInfo{}
	if err := dao.Admin.Where(dao.Admin.Columns.Id+"=?", g.Slice{auth.GetAdminID(r)}).Struct(re); err != nil {
		return nil, err
	}
	return
}
