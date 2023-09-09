// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package router

import (
	"github.com/anlityli/chatait-free/chatait-frontend-server/app/api/site"
	"github.com/anlityli/chatait-free/chatait-frontend-server/router/utils"
)

var siteApi = &site.Site{}

var siteRouter = []*utils.RouterItem{
	{Method: "GET", Pattern: "/index", Object: siteApi.Index, NoLogin: true, NoSign: true},
	{Method: "GET", Pattern: "/datetime", Object: siteApi.Datetime, NoLogin: true, NoSign: true},
}
