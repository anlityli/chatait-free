// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package web

import (
	"strings"

	"github.com/gogf/gf/net/ghttp"
)

// GetHeaderToken 获取头信息中的token
func GetHeaderToken(r *ghttp.Request) string {
	headerAuthorization := r.Header.Get("Authorization")
	if headerAuthorization != "" {
		parts := strings.SplitN(headerAuthorization, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			return ""
		}
		return parts[1]
	}
	return ""
}

// GetHeaderSignInfo 获取请求头鉴权信息
func GetHeaderSignInfo(r *ghttp.Request) (timestamp, sign string) {
	headerTimestamp := r.Header.Get("x-site-time")
	headerSign := r.Header.Get("x-site-sign")
	return headerTimestamp, headerSign
}
