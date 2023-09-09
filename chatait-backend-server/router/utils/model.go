// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package utils

// RouterItem 路由参数
type RouterItem struct {
	Method       string      // get | post
	Pattern      string      // 路由路径 /index
	Object       interface{} // 路由对应的api方法
	NoLogin      bool        // 不需要登录就能访问 默认false 大部分接口都需要登录
	NoSign       bool        // 不需要验签就能访问 默认false 大部分接口都需要验签
	NoPermission bool        // 不校验角色权限，任何管理员都能访问
}
