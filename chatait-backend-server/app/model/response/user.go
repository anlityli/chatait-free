// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package response

type UserListItem struct {
	Id              string `json:"id"`
	Username        string `json:"username"`
	Nickname        string `json:"nickname"`
	Avatar          string `json:"avatar"`
	LevelId         int    `json:"level_id"`
	LevelName       string `json:"level_name"`
	LevelExpireDate string `json:"level_expire_date"`
	LastLoginAt     int    `json:"last_login_at"`
	IsBan           int    `json:"is_ban"`
	CreatedAt       int    `json:"created_at"`
	UpdatedAt       int    `json:"updated_at"`
}

type UserList []*UserListItem

type UserSensitiveWord struct {
	Id             string `json:"id"`
	Username       string `json:"username"`
	Nickname       string `json:"nickname"`
	Type           int    `json:"type"`
	TopicType      int    `json:"topic_type"`
	Content        string `json:"content"`
	ValidateResult string `json:"validate_result"`
	CreatedAt      int    `json:"created_at"`
}

type UserSensitiveWordList []*UserSensitiveWord
