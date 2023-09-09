// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package router

import (
	"github.com/anlityli/chatait-free/chatait-backend-server/app/api/finance"
	"github.com/anlityli/chatait-free/chatait-backend-server/router/utils"
)

var financeApi = &finance.Finance{}

var financeRouter = []*utils.RouterItem{
	{Method: "GET", Pattern: "/wallet-list", Object: financeApi.WalletList},
	{Method: "POST", Pattern: "/wallet-change", Object: financeApi.WalletChange},
	{Method: "GET", Pattern: "/wallet-flow-list-balance", Object: financeApi.WalletFlowListBalance},
	{Method: "GET", Pattern: "/wallet-flow-list-gpt3", Object: financeApi.WalletFlowListGpt3},
	{Method: "GET", Pattern: "/wallet-flow-list-gpt4", Object: financeApi.WalletFlowListGpt4},
	{Method: "GET", Pattern: "/wallet-flow-list-midjourney", Object: financeApi.WalletFlowListMidjourney},
}
