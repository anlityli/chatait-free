// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package notice

import (
	"github.com/gogf/gf/net/ghttp"
)

// Write 新版本以model的形式返回结果
func Write(r *ghttp.Request, errorCode ErrorCode, data interface{}) {
	noticeData := &HttpModel{}
	noticeData.Error = errorCode
	noticeData.Message = data
	_ = r.Response.WriteJson(noticeData)
	r.ExitAll()
}
