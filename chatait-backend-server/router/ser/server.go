// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package ser

import (
	"github.com/anlityli/chatait-free/chatait-backend-server/router"
	"github.com/gogf/gf/net/ghttp"
)

// RouterInit 统一路由注册.
func RouterInit(server *ghttp.Server) {
	// http服务路由
	server.Group("/", func(g *ghttp.RouterGroup) {
		server.BindMiddleware("/*routerPath", Middleware, MiddlewareBehind)
		for controller, rule := range router.Rules {
			g.Group("/"+controller, func(group *ghttp.RouterGroup) {
				routerRuleHandle(group, rule)
			})
		}
	})
}
