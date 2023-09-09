// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package dashboard

import (
	"github.com/anlityli/chatait-free/chatait-backend-server/app/service"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/notice"
	"github.com/gogf/gf/net/ghttp"
)

type Dashboard struct {
}

func (c *Dashboard) UserStatistic(r *ghttp.Request) {
	re, err := service.Dashboard.UserStatistic(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, re)
}

func (c *Dashboard) OrderStatistic(r *ghttp.Request) {
	re, err := service.Dashboard.OrderStatistic(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, re)
}

func (c *Dashboard) AmountStatistic(r *ghttp.Request) {
	re, err := service.Dashboard.AmountStatistic(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, re)
}

func (c *Dashboard) ConversationStatistic(r *ghttp.Request) {
	re, err := service.Dashboard.ConversationStatistic(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, re)
}
