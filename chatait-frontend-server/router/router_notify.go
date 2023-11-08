// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package router

import (
	"github.com/anlityli/chatait-free/chatait-frontend-server/app/api/notify"
	"github.com/anlityli/chatait-free/chatait-frontend-server/router/utils"
)

var notifyApi = &notify.Notify{}

var notifyRouter = []*utils.RouterItem{
	{Method: "GET", Pattern: "/vmq", Object: notifyApi.Vmq, NoLogin: true, NoSign: true},
}
