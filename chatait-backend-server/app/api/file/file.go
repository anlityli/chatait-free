// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package file

import (
	"github.com/anlityli/chatait-free/chatait-backend-server/app/service"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/notice"
	"github.com/gogf/gf/net/ghttp"
)

type File struct {
}

func (c *File) MidjourneyList(r *ghttp.Request) {
	re, err := service.File.MidjourneyList(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, re)
}

func (c *File) MidjourneyImage(r *ghttp.Request) {
	service.File.MidjourneyImage(r)
}
