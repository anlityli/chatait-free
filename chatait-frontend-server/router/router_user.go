// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package router

import (
	"github.com/anlityli/chatait-free/chatait-frontend-server/app/api/user"
	"github.com/anlityli/chatait-free/chatait-frontend-server/router/utils"
)

var userApi = &user.User{}

var userRouter = []*utils.RouterItem{
	{Method: "GET", Pattern: "/info", Object: userApi.Info},
	{Method: "POST", Pattern: "/edit-nickname", Object: userApi.EditNickname},
}
