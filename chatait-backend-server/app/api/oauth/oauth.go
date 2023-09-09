// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package oauth

import (
	"github.com/anlityli/chatait-free/chatait-backend-server/app/service"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/notice"
	"github.com/gogf/gf/net/ghttp"
)

// Oauth 授权控制器
type Oauth struct {
}

// Login 登陆方法
func (o *Oauth) Login(r *ghttp.Request) {
	if tokenRe, err := service.Oauth.LoginWithUsername(r); err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	} else {
		notice.Write(r, notice.NoError, tokenRe)
	}
}

// RefreshToken 刷新AccessToken
func (o *Oauth) RefreshToken(r *ghttp.Request) {
	if tokenRe, err := service.Oauth.RefreshToken(r); err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	} else {
		notice.Write(r, notice.NoError, tokenRe)
	}
}

// Info 管理员信息
func (o *Oauth) Info(r *ghttp.Request) {
	if tokenRe, err := service.Oauth.Info(r); err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	} else {
		notice.Write(r, notice.NoError, tokenRe)
	}
}
