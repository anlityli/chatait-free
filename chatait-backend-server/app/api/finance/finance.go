// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package finance

import (
	"github.com/anlityli/chatait-free/chatait-backend-server/app/service"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/constant"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/notice"
	"github.com/gogf/gf/net/ghttp"
)

type Finance struct {
}

// WalletList 钱包列表
func (c *Finance) WalletList(r *ghttp.Request) {
	re, err := service.Finance.WalletList(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, re)
}

func (c *Finance) WalletChange(r *ghttp.Request) {
	err := service.Finance.WalletChange(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, "操作成功")
}

func (c *Finance) WalletFlowListBalance(r *ghttp.Request) {
	re, err := service.Finance.WalletFlowList(r, constant.WalletTypeBalance)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, re)
}

func (c *Finance) WalletFlowListGpt3(r *ghttp.Request) {
	re, err := service.Finance.WalletFlowList(r, constant.WalletTypeGpt3)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, re)
}

func (c *Finance) WalletFlowListGpt4(r *ghttp.Request) {
	re, err := service.Finance.WalletFlowList(r, constant.WalletTypeGpt4)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, re)
}

func (c *Finance) WalletFlowListMidjourney(r *ghttp.Request) {
	re, err := service.Finance.WalletFlowList(r, constant.WalletTypeMidjourney)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, re)
}
