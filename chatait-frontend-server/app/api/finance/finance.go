// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package finance

import (
	"github.com/anlityli/chatait-free/chatait-frontend-server/app/service"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/notice"
	"github.com/gogf/gf/net/ghttp"
)

type Finance struct {
}

func (c *Finance) WalletFlowList(r *ghttp.Request) {
	if re, err := service.FinanceService.WalletFlowList(r); err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	} else {
		notice.Write(r, notice.NoError, re)
	}
}

func (c *Finance) WalletInfo(r *ghttp.Request) {
	if re, err := service.FinanceService.WalletInfo(r); err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	} else {
		notice.Write(r, notice.NoError, re)
	}
}
