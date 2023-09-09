// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package response

import "github.com/anlityli/chatait-free/chatait-backend-server/library/menu"

type AdminListItem struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Username    string `json:"username"`
	AdminName   string `json:"admin_name"`
	RealName    string `json:"real_name"`
	Remark      string `json:"remark"`
	RoleID      string `json:"role_id"`
	RoleName    string `json:"role_name"`
	IsEnable    int    `json:"is_enable"`
	LoginNums   int    `json:"login_nums"`
	LastLoginIP string `json:"last_login_ip"`
	LastLoginAt int    `json:"last_login_at"`
	BindIP      string `json:"bind_ip"`
	CreateAdmin string `json:"create_admin"`
	UpdateAdmin string `json:"update_admin"`
	CreatedAt   int    `json:"created_at"`
	UpdatedAt   int    `json:"updated_at"`
	DontDel     int    `json:"dont_del"`
}

type AdminListItemList []*AdminListItem

type AdminRoleItem struct {
	ID               string `json:"id"`
	RoleName         string `json:"role_name"`
	Remark           string `json:"remark"`
	Permission       string `json:"permission"`
	ColumnPermission string `json:"column_permission"`
	DontDel          int    `json:"dont_del"`
	CreateAdmin      string `json:"create_admin"`
	UpdateAdmin      string `json:"update_admin"`
	CreatedAt        int    `json:"created_at"`
	UpdatedAt        int    `json:"updated_at"`
}

type AdminRoleItemList []*AdminRoleItem

type AdminAllPermissionItem struct {
	ID              string                 `json:"id"`
	MainPermission  *AdminPermissionItem   `json:"main_permission"`
	ChildPermission []*AdminPermissionItem `json:"child_permission"`
}

type AdminAllPermissionList []*AdminAllPermissionItem

type AdminPermissionItem struct {
	Title     string `json:"title"`
	Path      string `json:"path"`
	IsChecked bool   `json:"is_checked"`
}

type AdminColumnPermissionItem struct {
	ListName string                              `json:"list_name"`
	ListID   string                              `json:"list_id"`
	Columns  []*AdminColumnPermissionItemColumns `json:"columns"`
}

type AdminColumnPermissionList []*AdminColumnPermissionItem

type AdminColumnPermissionItemColumns struct {
	FieldName string `json:"header"`
	Field     string `json:"index"`
	IsChecked bool   `json:"is_checked"`
}

type AdminInfo struct {
	AdminListItem
	Menu            menu.ListModel `json:"menu"`
	AdminPermission []string       `json:"admin_permission"`
}
