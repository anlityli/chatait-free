// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package router

import (
	"github.com/anlityli/chatait-free/chatait-backend-server/app/api/user"
	"github.com/anlityli/chatait-free/chatait-backend-server/router/utils"
)

var userApi = &user.User{}

var userRouter = []*utils.RouterItem{
	{Method: "GET", Pattern: "/list", Object: userApi.List},
	{Method: "POST", Pattern: "/ban", Object: userApi.Ban},
	{Method: "POST", Pattern: "/change-level", Object: userApi.ChangeLevel},
	{Method: "POST", Pattern: "/reset-password", Object: userApi.ResetPassword},
	{Method: "GET", Pattern: "/sensitive-word-list", Object: userApi.SensitiveWordList},
	{Method: "GET", Pattern: "/sensitive-word-one", Object: userApi.SensitiveWordOne},
}
