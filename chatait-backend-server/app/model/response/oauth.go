// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package response

type OauthLogin struct {
	AccessToken          string `json:"access_token"`
	AccessTokenExpire    int    `json:"access_token_expire"`
	AccessTokenExpireIn  int    `json:"access_token_expire_in"`
	RefreshToken         string `json:"refresh_token"`
	RefreshTokenExpire   int    `json:"refresh_token_expire"`
	RefreshTokenExpireIn int    `json:"refresh_token_expire_in"`
}

// OauthInfo 会员登录后获取到的信息
type OauthInfo struct {
	ID        string `json:"id"`
	AdminName string `json:"admin_name"`
	RoleId    string `json:"role_id"`
}
