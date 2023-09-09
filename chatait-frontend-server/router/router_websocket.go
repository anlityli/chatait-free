// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package router

import (
	"github.com/anlityli/chatait-free/chatait-frontend-server/app/api/websocket"
	"github.com/anlityli/chatait-free/chatait-frontend-server/router/utils"
)

var websocketApi = &websocket.Websocket{}

var websocketRouter = []*utils.RouterItem{
	{Method: "GET", Pattern: "/index", Object: websocketApi.Index, NoLogin: true, NoSign: true},
	{Method: "POST", Pattern: "/notice-to-all", Object: websocketApi.NoticeToAll, NoLogin: true},
}
