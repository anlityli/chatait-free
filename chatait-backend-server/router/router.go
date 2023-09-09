// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package router

import "github.com/anlityli/chatait-free/chatait-backend-server/router/utils"

var Rules = map[string][]*utils.RouterItem{
	"admin":        adminRouter,
	"config":       configRouter,
	"conversation": conversationRouter,
	"dashboard":    dashboardRouter,
	"finance":      financeRouter,
	"oauth":        oauthRouter,
	"shop":         shopRouter,
	"site":         siteRouter,
	"user":         userRouter,
}
