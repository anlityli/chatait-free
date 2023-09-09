// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package request

type OauthSignupSendCode struct {
	Username string `json:"username" v:"required|email#邮箱必填｜邮箱格式不正确"`
}

type OauthSignUpValidateCode struct {
	Username string `json:"username" v:"required|email#邮箱必填｜邮箱格式不正确"`
	Code     string `json:"code" v:"required#code必填"`
}

type OauthSignupFinish struct {
	Username string `json:"username" v:"required|email#邮箱必填｜邮箱格式不正确"`
	Password string `json:"password" v:"required#密码必填"`
	Code     string `json:"code"`
	Nickname string `json:"nickname" v:"required#昵称必填"`
}

type OauthLogin struct {
	Username string `json:"mobile" v:"required"`
	Password string `json:"password" v:"required"`
}

type OauthFindPasswordFinish struct {
	Username string `json:"username" v:"required|email#邮箱必填｜邮箱格式不正确"`
	Password string `json:"password" v:"required#密码必填"`
	Code     string `json:"code" v:"required#验证码必填"`
}
