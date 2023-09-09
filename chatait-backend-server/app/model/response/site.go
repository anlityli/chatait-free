// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package response

import "github.com/anlityli/chatait-free/chatait-backend-server/library/menu"

// SiteBaseInfo 站点基本信息
type SiteBaseInfo struct {
	Menu             menu.ListModel `json:"menu"`
	AdminRoles       interface{}    `json:"admin_roles"`
	SuperAdminRoleId string         `json:"super_admin_role_id"`
	FrontendUserId   string         `json:"frontendUserId"`
}
