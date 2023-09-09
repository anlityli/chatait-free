// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package router

import (
	"github.com/anlityli/chatait-free/chatait-backend-server/app/api/admin"
	"github.com/anlityli/chatait-free/chatait-backend-server/router/utils"
)

var adminApi = &admin.Admin{}

var adminRouter = []*utils.RouterItem{
	{Method: "GET", Pattern: "/list", Object: adminApi.List},
	{Method: "GET", Pattern: "/one", Object: adminApi.One},
	{Method: "GET", Pattern: "/all-role", Object: adminApi.AllRole},
	{Method: "GET", Pattern: "/role-list", Object: adminApi.RoleList},
	{Method: "GET", Pattern: "/role-one", Object: adminApi.RoleOne},
	{Method: "POST", Pattern: "/role-add", Object: adminApi.RoleAdd},
	{Method: "POST", Pattern: "/role-edit", Object: adminApi.RoleEdit},
	{Method: "POST", Pattern: "/role-delete", Object: adminApi.RoleDelete},
	{Method: "GET", Pattern: "/role-permission", Object: adminApi.RolePermission},
	{Method: "POST", Pattern: "/role-permission-edit", Object: adminApi.RolePermissionEdit},
	{Method: "GET", Pattern: "/role-column", Object: adminApi.RoleColumn},
	{Method: "POST", Pattern: "/role-column-permission-edit", Object: adminApi.RoleColumnPermissionEdit},
	{Method: "POST", Pattern: "/reset-password", Object: adminApi.ResetPassword, NoPermission: true},
	{Method: "POST", Pattern: "/reset-other-password", Object: adminApi.ResetOtherPassword},
	{Method: "POST", Pattern: "/add", Object: adminApi.Add},
	{Method: "POST", Pattern: "/edit", Object: adminApi.Edit},
	{Method: "POST", Pattern: "/delete", Object: adminApi.Delete},
	{Method: "GET", Pattern: "/info", Object: adminApi.Info, NoPermission: true},
}
