// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/thinkeridea/go-extend/exnet"
)

// GetClientIP 获取用户的真是IP地址
func GetClientIP(r *ghttp.Request) string {
	return exnet.ClientIP(r.Request)
}
