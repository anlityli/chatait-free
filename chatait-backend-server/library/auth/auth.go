// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package auth

import (
	"database/sql"
	"errors"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/model/entity"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/helper"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/security"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/web"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// SetUserAuth 设置用户权限信息
func SetUserAuth(r *ghttp.Request, admin *entity.Admin) {
	r.SetParam("oauthAdminID", gconv.String(admin.Id))
	r.SetParam("oauthUserID", gconv.String(admin.UserId))
	r.SetParam("oauthAdminName", admin.AdminName)
	r.SetParam("oauthRoleID", gconv.String(admin.RoleId))
}

// GetAdminID 获取管理员ID
func GetAdminID(r *ghttp.Request) uint64 {
	return gconv.Uint64(gconv.String(r.GetParam("oauthAdminID")))
}

// GetUserID 获取前端用户ID
func GetUserID(r *ghttp.Request) uint64 {
	return gconv.Uint64(gconv.String(r.GetParam("oauthUserID")))
}

// GetAdminName 获取用户名
func GetAdminName(r *ghttp.Request) string {
	return gconv.String(r.GetParam("oauthAdminName"))
}

// GetRoleID 获取管理员角色
func GetRoleID(r *ghttp.Request) uint64 {
	return gconv.Uint64(gconv.String(r.GetParam("oauthRoleID")))
}

// ValidateAuth 鉴权
func ValidateAuth(r *ghttp.Request) (admin *entity.Admin, err error) {
	// 解析token
	parseRe, err := security.ParseAdminToken(web.GetHeaderToken(r))
	if err != nil {
		return nil, err
	}
	if parseRe.IsRefresh {
		err = errors.New("token 非法")
		return nil, err
	}
	// 从数据库中拿会员其他的信息
	if err = dao.Admin.Where("id=?", parseRe.Admin.Id).Scan(parseRe.Admin); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("会员不存在")
		}
		return nil, err
	}
	SetUserAuth(r, parseRe.Admin)
	return parseRe.Admin, nil
}

// ValidateRefresh 校验刷新Token
func ValidateRefresh(r *ghttp.Request) (admin *entity.Admin, err error) {
	// 获取refreshToken
	refreshToken := r.GetString("refresh_token")
	parseRe, err := security.ParseAdminToken(refreshToken)
	if err != nil {
		return
	}
	if !parseRe.IsRefresh {
		err = errors.New("token 非法")
		return
	}
	return parseRe.Admin, nil
}

// ValidateSign 校验验签
func ValidateSign(r *ghttp.Request) bool {
	headerTimestamp, headerSign := web.GetHeaderSignInfo(r)

	token := web.GetHeaderToken(r)
	dataMap := g.Map{
		"time":  headerTimestamp,
		"token": token,
		"sign":  headerSign,
	}
	return security.ValidateParamsSign(dataMap, g.Config().GetString("backendConf.privateKey"))
}

// ValidatePermission 校验权限
func ValidatePermission(r *ghttp.Request) bool {
	roleID := GetRoleID(r)
	if roleID == 1 {
		return true
	}
	routerPath := r.GetRouterString("routerPath")
	permissionData, err := RolePermission(gconv.String(roleID))
	if err != nil {
		return false
	}
	if !helper.StrInArr(permissionData, routerPath) {
		return false
	}
	return true
}
