// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package router

import (
	"github.com/anlityli/chatait-free/chatait-frontend-server/app/api/config"
	"github.com/anlityli/chatait-free/chatait-frontend-server/router/utils"
)

var configApi = &config.Config{}

var configRouter = []*utils.RouterItem{
	{Method: "GET", Pattern: "/wallet-list", Object: configApi.WalletList, NoLogin: true},
	{Method: "GET", Pattern: "/options", Object: configApi.Options, NoLogin: true},
	{Method: "GET", Pattern: "/pay-list", Object: configApi.PayList},
}
