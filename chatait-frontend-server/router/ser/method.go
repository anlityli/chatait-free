// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package ser

import (
	"github.com/anlityli/chatait-free/chatait-frontend-server/library/auth"
	"github.com/anlityli/chatait-free/chatait-frontend-server/router"
	"github.com/anlityli/chatait-free/chatait-frontend-server/router/utils"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/helper"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/notice"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
)

func Middleware(r *ghttp.Request) {
	routerPathArr := gstr.Explode("/", r.Router.Uri)
	if len(routerPathArr) < 2 {
		notice.Write(r, notice.OtherError, "路径错误")
		return
	}
	controller := routerPathArr[1]
	rule := router.Rules[controller]
	for _, item := range rule {
		if r.Router.Uri == "/"+controller+item.Pattern {

			// 验签
			if !auth.ValidateSign(r) {
				if !item.NoSign {
					notice.Write(r, notice.OtherError, "验签失败")
					return
				}
			}
			// token授权校验
			userData, err := auth.ValidateAuth(r)
			if err != nil {
				// 如果路由权限里面没有none，说明该接口需要登录才能访问
				if !item.NoLogin {
					notice.Write(r, notice.NotAuth, "尚未登录，请登录")
					return
				}
			}
			if !item.NoLogin {
				if userData == nil {
					notice.Write(r, notice.OtherError, "用户数据获取失败")
					return
				}

				// 对会员级别进行校验
				if len(item.PermissionLevel) > 0 && !helper.IntInArr(item.PermissionLevel, utils.LevelAll) {
					canRequest := false
					noticeStr := "您的级别权限不足"
					for _, routeLevelItem := range item.PermissionLevel {
						switch routeLevelItem {
						case utils.LevelNormal:
							noticeStr = "需要普通会员级别才能访问"
						case utils.LevelPlus:
							noticeStr = "需要VIP会员级别才能访问"
						}
					}

					// 普通会员
					if userData.LevelId == utils.LevelNormal {
						if helper.IntInArr(item.PermissionLevel, utils.LevelNormal) {
							canRequest = true
						}
					}
					// plus会员
					if userData.LevelId == utils.LevelPlus {
						if helper.IntInArr(item.PermissionLevel, utils.LevelPlus) {
							canRequest = true
						}
					}
					if !canRequest {
						notice.Write(r, notice.AuthForbidden, noticeStr)
						return
					}
				}
			}

			r.Middleware.Next()
			return
		}
	}
	notice.Write(r, notice.NotFind, "路径错误")
	return
}

func routerRuleHandle(g *ghttp.RouterGroup, routerRule []*utils.RouterItem) {
	for _, item := range routerRule {
		if gstr.ToLower(item.Method) == "get" {
			g.GET(item.Pattern, item.Object)
		} else if gstr.ToLower(item.Method) == "post" {
			g.POST(item.Pattern, item.Object)
		} else if gstr.ToLower(item.Method) == "all" {
			g.ALL(item.Pattern, item.Object)
		}
	}
}
