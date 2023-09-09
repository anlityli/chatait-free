// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package oauth

import (
	"github.com/anlityli/chatait-free/chatait-frontend-server/app/service"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/notice"
	"github.com/gogf/gf/net/ghttp"
)

// Oauth 授权控制器
type Oauth struct {
}

func (o *Oauth) SignupSendCode(r *ghttp.Request) {
	if re, err := service.Oauth.SignupSendCode(r); err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	} else {
		notice.Write(r, notice.NoError, re)
	}
}

func (o *Oauth) SignupValidate(r *ghttp.Request) {
	if re, err := service.Oauth.SignupValidate(r); err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	} else {
		notice.Write(r, notice.NoError, re)
	}
}

func (o *Oauth) SignupFinish(r *ghttp.Request) {
	if re, err := service.Oauth.SignupFinish(r); err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	} else {
		notice.Write(r, notice.NoError, re)
	}
}

func (o *Oauth) FindPasswordSendCode(r *ghttp.Request) {
	if re, err := service.Oauth.FindPasswordSendCode(r); err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	} else {
		notice.Write(r, notice.NoError, re)
	}
}

func (o *Oauth) FindPasswordValidate(r *ghttp.Request) {
	if re, err := service.Oauth.FindPasswordValidate(r); err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	} else {
		notice.Write(r, notice.NoError, re)
	}
}

func (o *Oauth) FindPasswordFinish(r *ghttp.Request) {
	if re, err := service.Oauth.FindPasswordFinish(r); err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	} else {
		notice.Write(r, notice.NoError, re)
	}
}

func (o *Oauth) Login(r *ghttp.Request) {
	if re, err := service.Oauth.Login(r); err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	} else {
		notice.Write(r, notice.NoError, re)
	}
}

func (o *Oauth) RefreshToken(r *ghttp.Request) {
	if re, err := service.Oauth.RefreshToken(r); err != nil {
		notice.Write(r, notice.NotAuth, "登录已失效，请重新登录")
	} else {
		notice.Write(r, notice.NoError, re)
	}
}
