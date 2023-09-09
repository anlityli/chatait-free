// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package request

type AdminIds struct {
	Selected []string `json:"selected" v:"required"`
}

// AdminId 获取一条管理员
type AdminId struct {
	Id string `json:"id" v:"required#id必填"`
}

// AdminRoleAdd 添加管理员
type AdminRoleAdd struct {
	RoleName string `json:"role_name" v:"required#role_name必填"`
	Remark   string `json:"remark"`
}

type AdminRoleEdit struct {
	Id string `json:"id" v:"required#必填"`
	AdminRoleAdd
}

// AdminRolePermission 管理员权限
type AdminRolePermission struct {
	ID string `json:"id" v:"required#id必填"`
}

// AdminRolePermissionEdit 管理员权限编辑
type AdminRolePermissionEdit struct {
	ID         string      `json:"id" v:"required#id必填"`
	Permission interface{} `json:"permission"`
}

// AdminRoleColumnPermissionEdit 列表字段权限编辑
type AdminRoleColumnPermissionEdit struct {
	ID               string      `json:"id" v:"required#id必填"`
	ColumnPermission interface{} `json:"column_permission"`
}

// AdminResetPassword 重置管理员密码
type AdminResetPassword struct {
	Password string `json:"password" v:"required"`
}

type AdminResetOtherPassword struct {
	AdminId  string `json:"admin_id" v:"required"`
	Password string `json:"password" v:"required"`
}

// AdminAdd 添加管理员
type AdminAdd struct {
	UserID    string   `json:"user_id"`
	AdminName string   `json:"admin_name"`
	RealName  string   `json:"real_name"`
	Remark    string   `json:"remark"`
	RoleID    string   `json:"role_id" v:"required"`
	IsEnable  int      `json:"is_enable" v:"required"`
	Password  string   `json:"password"`
	BindIP    []string `json:"bind_ip"`
}

// AdminEdit 编辑管理员
type AdminEdit struct {
	ID string `json:"id"`
	AdminAdd
}
