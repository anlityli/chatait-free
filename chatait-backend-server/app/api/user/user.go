// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package user

import (
	"github.com/anlityli/chatait-free/chatait-backend-server/app/service"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/notice"
	"github.com/gogf/gf/net/ghttp"
)

type User struct {
}

// List 列表
func (c *User) List(r *ghttp.Request) {
	re, err := service.User.List(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, re)
}

func (c *User) Ban(r *ghttp.Request) {
	err := service.User.Ban(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, "操作成功")
}

func (c *User) ChangeLevel(r *ghttp.Request) {
	err := service.User.ChangeLevel(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, "操作成功")
}

func (c *User) ResetPassword(r *ghttp.Request) {
	err := service.User.ResetPassword(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, "操作成功")
}
