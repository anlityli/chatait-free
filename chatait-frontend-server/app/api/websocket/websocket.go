// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package websocket

import (
	"github.com/anlityli/chatait-free/chatait-frontend-server/app/service"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/notice"
	"github.com/gogf/gf/net/ghttp"
)

type Websocket struct {
}

func (w *Websocket) Index(r *ghttp.Request) {
	err := service.Websocket.Index(r)
	if err != nil {
		r.ExitAll()
	}
}

// NoticeToAll 开放一个接口，用于其他服务向前端发送websocket消息
// 注意该接口验签方式和前台请求服务端一样，不需要其他特殊验签
func (w *Websocket) NoticeToAll(r *ghttp.Request) {
	err := service.Websocket.NoticeToAll(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, "操作成功")
}
