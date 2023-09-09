// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package user

import (
	"github.com/anlityli/chatait-free/chatait-frontend-server/app/service"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/notice"
	"github.com/gogf/gf/net/ghttp"
)

type User struct {
}

func (c *User) Info(r *ghttp.Request) {
	if re, err := service.User.Info(r); err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	} else {
		notice.Write(r, notice.NoError, re)
	}
}

func (c *User) EditNickname(r *ghttp.Request) {
	if err := service.User.EditNickname(r); err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	} else {
		notice.Write(r, notice.NoError, "操作成功")
	}
}
