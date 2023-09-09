// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package ser

import (
	"github.com/anlityli/chatait-free/chatait-backend-server/library/auth"
	"github.com/anlityli/chatait-free/chatait-backend-server/router"
	"github.com/anlityli/chatait-free/chatait-backend-server/router/utils"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/notice"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/snowflake"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/xtime"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
)

func Middleware(r *ghttp.Request) {
	routerPath := r.GetRouterString("routerPath")
	routerPathArr := gstr.Explode("/", routerPath)
	if len(routerPathArr) < 2 {
		notice.Write(r, notice.OtherError, "路径错误")
		return
	}
	controller := routerPathArr[0]
	rule := router.Rules[controller]
	for _, item := range rule {
		if routerPath == controller+item.Pattern {
			// 验签
			if !auth.ValidateSign(r) {
				if !item.NoSign {
					notice.Write(r, notice.OtherError, "验签失败")
					return
				}
			}
			// token授权校验
			adminData, err := auth.ValidateAuth(r)
			if err != nil {
				// 如果路由权限里面没有不需要登录就能访问，说明该接口需要登录才能访问
				if !item.NoLogin {
					notice.Write(r, notice.NotAuth, err.Error())
					return
				}
			}

			if !item.NoLogin {
				if adminData == nil {
					notice.Write(r, notice.OtherError, "用户数据获取失败")
					return
				}
				if !item.NoPermission {
					// 对管理员的角色进行校验
					if !auth.ValidatePermission(r) {
						notice.Write(r, notice.AuthForbidden, "权限不足")
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

func MiddlewareBehind(r *ghttp.Request) {
	r.Middleware.Next()
	// 异步把操作的增删改的操作记录到日志表中
	go func() {
		// 记录操作日志
		operationLog(r)
	}()
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

func operationLog(r *ghttp.Request) {
	// 判断是不是POST请求
	if gstr.ToLower(r.Method) == "post" {
		routerPath := r.Router.Uri
		if routerPath == "/oauth/refresh-token" {
			return
		} else if routerPath == "/oauth/login" {
			return
		}
		paramsByte := r.GetBody()
		paramsStr := gconv.String(paramsByte)
		paramsJson, err := gjson.Decode(paramsStr)
		if err == nil {
			paramsMap := gconv.Map(paramsJson)
			if len(paramsMap) > 0 {
				// 是否存在密码字段，存在则加密处理
				if _, ok := paramsMap["password"]; ok {
					paramsMap["password"] = "******"
				}
				if _, ok := paramsMap["sure_password"]; ok {
					paramsMap["sure_password"] = "******"
				}
				paramsJsonNew, err := gjson.Encode(paramsMap)
				if err != nil {
					glog.Line(true).Println("记录日志发生错误", err.Error())
					return
				}
				paramsStr = gconv.String(paramsJsonNew)
			}
		}
		adminName := auth.GetAdminName(r)
		// 把操作内容写到日志表中
		id := snowflake.GenerateID()
		if _, err := dao.LogOperation.Data(g.Map{
			"id":          id,
			"status_code": gconv.String(r.GetParam("responseErrorCode")),
			"router":      r.URL,
			"content":     paramsStr,
			"admin_name":  adminName,
			"created_at":  xtime.GetNowTime(),
		}).Insert(); err != nil {
			glog.Line(true).Println("记录日志发生错误", err.Error())
			return
		}
	} else {
		//paramsByte := r.GetBody()
		//paramsStr := gconv.String(paramsByte)
		//glog.Line().Debug(r.Router.Uri)
		//glog.Line().Debug(r.URL)
	}
}
