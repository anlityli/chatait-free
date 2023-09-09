// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package server

import (
	"github.com/anlityli/chatait-free/chatait-backend-server/router/ser"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/helper"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/serv/http"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

// RunServer 运行程序
func RunServer() {
	InitConfig()
	http.Instance().Init("backend")
	http.Instance().Run(func(server *ghttp.Server) {
		ser.RouterInit(server)
	}, func() {

	})
	g.Wait()
}

func InitConfig() {
	// 初始化日志存放位置，日志存在文件中
	logPath := helper.FormatDirStr(g.Config().GetString("backendConf.logPath"))
	if err := glog.SetPath(logPath); err != nil {
		glog.Line().Fatalf("设置日志文件路径失败:%s\n", err.Error())
	}
	// 如果是debug模式的话，debug日志会保存，否则debug日志不会保存
	glog.SetDebug(g.Config().GetBool("backendConf.debugStatus"))
}
