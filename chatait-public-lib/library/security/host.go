// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package security

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gmutex"
)

var hostUrl = ""
var hostUrlLock = gmutex.New()

func HostUrl(r ...*ghttp.Request) (re string, err error) {
	hostUrlLock.Lock()
	defer hostUrlLock.Unlock()
	if hostUrl == "" {
		if len(r) > 0 && r[0] != nil {
			protocol := "http"
			if r[0].TLS != nil || r[0].Header.Get("X-Forwarded-Proto") == "https" {
				protocol = "https"
			}
			host := r[0].Host
			hostUrl = protocol + "://" + host
		}
	}
	return hostUrl, nil
}
