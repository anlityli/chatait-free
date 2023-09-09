// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package router

import (
	"github.com/anlityli/chatait-free/chatait-backend-server/app/api/oauth"
	"github.com/anlityli/chatait-free/chatait-backend-server/router/utils"
)

var oauthApi = &oauth.Oauth{}

var oauthRouter = []*utils.RouterItem{
	{Method: "POST", Pattern: "/login", Object: oauthApi.Login, NoLogin: true, NoPermission: true},
	{Method: "POST", Pattern: "/refresh-token", Object: oauthApi.RefreshToken, NoLogin: true, NoPermission: true},
	{Method: "GET", Pattern: "/info", Object: oauthApi.Info, NoPermission: true},
}
