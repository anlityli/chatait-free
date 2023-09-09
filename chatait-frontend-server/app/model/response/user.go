// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package response

type UserInfo struct {
	Id              string `json:"id"`
	Username        string `json:"username"`
	Nickname        string `json:"nickname"`
	Avatar          string `json:"avatar"`
	LevelId         int    `json:"level_id"`
	LevelName       string `json:"level_name"`
	LevelExpireDate string `json:"level_expire_date"`
	CreatedAt       int    `json:"created_at"`
	LastLoginAt     int    `json:"last_login_at"`
}

type UserList []*UserInfo
