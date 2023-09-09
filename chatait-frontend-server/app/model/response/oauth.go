// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package response

// OauthSignUpSendCode 发送验证码的返回值
type OauthSignUpSendCode struct {
	Email          string `json:"email"`
	Code           string `json:"code"` // 只有体验测试时才能返回实际code值
	IntervalSecond int    `json:"interval_second"`
	ExpireIn       int    `json:"expire_in"`
}

type OauthSignUpValidateCode struct {
	Email       string `json:"email"`
	IsRight     int    `json:"is_right"`
	CodeExpired int    `json:"code_expired"`
}

type OauthUserToken struct {
	AccessToken          string `json:"access_token"`
	AccessTokenExpireIn  int    `json:"access_token_expire_in"`
	RefreshToken         string `json:"refresh_token"`
	RefreshTokenExpireIn int    `json:"refresh_token_expire_in"`
}
