// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package request

type UserIds struct {
	Selected []string `json:"selected" v:"required"`
}

type UserChangeLevel struct {
	UserId          string `json:"user_id"`
	LevelId         int    `json:"level_id"`
	LevelExpireDate string `json:"level_expire_date"`
	Remark          string `json:"remark"`
}

type UserBan struct {
	Selected []string `json:"selected" v:"required"`
	IsBan    int      `json:"is_ban"`
}

type UserResetPassword struct {
	UserId   string `json:"user_id"`
	Password string `json:"password"`
}
