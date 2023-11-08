// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package router

import (
	"github.com/anlityli/chatait-free/chatait-backend-server/app/api/file"
	"github.com/anlityli/chatait-free/chatait-backend-server/router/utils"
)

var fileApi = &file.File{}

var fileRouter = []*utils.RouterItem{
	{Method: "GET", Pattern: "/midjourney-list", Object: fileApi.MidjourneyList},
	{Method: "GET", Pattern: "/midjourney-image", Object: fileApi.MidjourneyImage, NoLogin: true, NoSign: true, NoPermission: true},
}
