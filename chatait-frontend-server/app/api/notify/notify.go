// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package notify

import (
	"github.com/anlityli/chatait-free/chatait-frontend-server/app/service"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/notice"
	"github.com/gogf/gf/net/ghttp"
)

type Notify struct {
}

func (c *Notify) Vmq(r *ghttp.Request) {
	if err := service.Notify.Vmq(r); err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	} else {
		r.Response.Write("success")
		r.ExitAll()
	}
}
