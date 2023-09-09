// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package router

import (
	"github.com/anlityli/chatait-free/chatait-frontend-server/app/api/oauth"
	"github.com/anlityli/chatait-free/chatait-frontend-server/router/utils"
)

var oauthApi = &oauth.Oauth{}

var oauthRouter = []*utils.RouterItem{
	{Method: "POST", Pattern: "/signup-send-code", Object: oauthApi.SignupSendCode, NoLogin: true},
	{Method: "POST", Pattern: "/signup-validate", Object: oauthApi.SignupValidate, NoLogin: true},
	{Method: "POST", Pattern: "/signup-finish", Object: oauthApi.SignupFinish, NoLogin: true},
	{Method: "POST", Pattern: "/find-password-send-code", Object: oauthApi.FindPasswordSendCode, NoLogin: true},
	{Method: "POST", Pattern: "/find-password-validate", Object: oauthApi.FindPasswordValidate, NoLogin: true},
	{Method: "POST", Pattern: "/find-password-finish", Object: oauthApi.FindPasswordFinish, NoLogin: true},
	{Method: "POST", Pattern: "/login", Object: oauthApi.Login, NoLogin: true},
	{Method: "POST", Pattern: "/refresh-token", Object: oauthApi.RefreshToken, NoLogin: true},
}
