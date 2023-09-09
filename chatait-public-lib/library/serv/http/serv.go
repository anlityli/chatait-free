// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package http

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"sync"
)

type HTTP struct {
	server *ghttp.Server
	site   string // frontend | backend
}

var httpObj *HTTP
var httpOnce sync.Once

// Instance 单例
func Instance() *HTTP {
	httpOnce.Do(func() {
		httpObj = &HTTP{}
	})
	return httpObj
}

// Init 的服务
func (h *HTTP) Init(site string) {
	appPath := g.Config().GetString(site + "Conf.appPath")
	sessionPath := appPath + "runtime/session"
	h.server = g.Server("httpServer")
	_ = h.server.SetConfigWithMap(g.Map{
		"SessionPath": sessionPath,
	})
	h.server.SetDumpRouterMap(false)
	h.server.SetPort(g.Config().GetInt(site + "Conf.serverPort"))

	h.server.SetClientMaxBodySize(int64(10 * 1024 * 1024))
	h.site = site
}

// GetHttpServer 获取server对象
func (h *HTTP) GetHttpServer() *ghttp.Server {
	return h.server
}

func (h *HTTP) GetSite() string {
	return h.site
}

// Run 运行
func (h *HTTP) Run(routerCallback func(server *ghttp.Server), eventCallback func()) {
	go eventCallback()
	routerCallback(h.server)
	_ = h.server.Start()
}
