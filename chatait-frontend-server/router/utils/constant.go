// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package utils

// 访问权限的级别
// 如果设置了不用登录就能访问 那么级别设置将不起作用，因为无需登录就能访问，意味着不校验会员级别
const (
	LevelAll    = 0 // 如果数组内没有规定级别，那么就相当于设置了all
	LevelNormal = 1 // 普通会员
	LevelPlus   = 2 //  VIP会员
)
