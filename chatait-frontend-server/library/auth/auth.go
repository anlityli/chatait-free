// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package auth

import (
	"database/sql"
	"errors"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/model/entity"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/security"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/web"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// SetUserAuth 设置用户权限信息
func SetUserAuth(r *ghttp.Request, user *entity.User) {
	r.SetParam("oauthUserId", gconv.String(user.Id))
}

// GetUserId 获取前端用户ID
func GetUserId(r *ghttp.Request) int64 {
	return gconv.Int64(gconv.String(r.GetParam("oauthUserId")))
}

// ValidateAuth 鉴权
func ValidateAuth(r *ghttp.Request) (user *entity.User, err error) {
	// 解析token
	parseRe, err := security.ParseUserToken(web.GetHeaderToken(r))
	if err != nil {
		return nil, err
	}
	if parseRe.IsRefresh {
		err = errors.New("token 非法")
		return nil, err
	}
	// 从数据库中拿会员其他的信息
	if err = dao.User.Where("id=?", parseRe.User.Id).Scan(parseRe.User); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("会员不存在")
		}
		return nil, err
	}
	if parseRe.User.IsBan == 1 {
		return nil, errors.New("登录已失效")
	}
	SetUserAuth(r, parseRe.User)
	return parseRe.User, nil
}

// ValidateRefresh 校验刷新Token
func ValidateRefresh(r *ghttp.Request) (user *entity.User, err error) {
	// 获取refreshToken
	refreshToken := r.GetString("refresh_token")
	parseRe, err := security.ParseUserToken(refreshToken)
	if err != nil {
		return
	}
	if !parseRe.IsRefresh {
		err = errors.New("token 非法")
		return
	}
	return parseRe.User, nil
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
	return security.ValidateParamsSign(dataMap, g.Config().GetString("frontendConf.privateKey"))
}
