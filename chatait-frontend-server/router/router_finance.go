// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package router

import (
	"github.com/anlityli/chatait-free/chatait-frontend-server/app/api/finance"
	"github.com/anlityli/chatait-free/chatait-frontend-server/router/utils"
)

var financeApi = &finance.Finance{}

var financeRouter = []*utils.RouterItem{
	{Method: "GET", Pattern: "/wallet-flow-list", Object: financeApi.WalletFlowList},
	{Method: "GET", Pattern: "/wallet-info", Object: financeApi.WalletInfo},
}
