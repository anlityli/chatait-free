// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package request

// OauthLogin 用户名密码登录
type OauthLogin struct {
	AdminName string `json:"admin_name" v:"required#用户名必填"`
	Password  string `json:"password" v:"required#密码必填"`
}
