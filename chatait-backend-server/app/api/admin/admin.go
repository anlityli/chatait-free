// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package admin

import (
	"github.com/anlityli/chatait-free/chatait-backend-server/app/service"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/notice"
	"github.com/gogf/gf/net/ghttp"
)

type Admin struct {
}

// List 列表
func (c *Admin) List(r *ghttp.Request) {
	re, err := service.Admin.List(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, re)
}

// One 信息
func (c *Admin) One(r *ghttp.Request) {
	re, err := service.Admin.One(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, re)
}

// AllRole 所有角色
func (c *Admin) AllRole(r *ghttp.Request) {
	re, err := service.Admin.AllRole(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, re)
}

// RoleList 角色列表
func (c *Admin) RoleList(r *ghttp.Request) {
	re, err := service.Admin.RoleList(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, re)
}

func (c *Admin) RoleOne(r *ghttp.Request) {
	re, err := service.Admin.RoleOne(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, re)
}

// RoleAdd 添加角色
func (c *Admin) RoleAdd(r *ghttp.Request) {
	re, err := service.Admin.RoleAdd(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, re)
}

func (c *Admin) RoleEdit(r *ghttp.Request) {
	err := service.Admin.RoleEdit(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, "操作成功")
}

func (c *Admin) RoleDelete(r *ghttp.Request) {
	err := service.Admin.RoleDelete(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, "操作成功")
}

// RolePermission 角色权限
func (c *Admin) RolePermission(r *ghttp.Request) {
	re, err := service.Permission.RolePermission(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, re)
}

// RolePermissionEdit 编辑权限
func (c *Admin) RolePermissionEdit(r *ghttp.Request) {
	err := service.Permission.RolePermissionEdit(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, "操作成功")
}

// RoleColumn 列表字段权限
func (c *Admin) RoleColumn(r *ghttp.Request) {
	re, err := service.Permission.RoleColumnPermission(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, re)
}

// RoleColumnPermissionEdit 编辑列表字段权限
func (c *Admin) RoleColumnPermissionEdit(r *ghttp.Request) {
	err := service.Permission.RoleColumnPermissionEdit(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, "操作成功")
}

// ResetPassword 重置密码
func (c *Admin) ResetPassword(r *ghttp.Request) {
	err := service.Admin.ResetPassword(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, "操作成功")
}

// ResetOtherPassword 重置其他管理员面密码
func (c *Admin) ResetOtherPassword(r *ghttp.Request) {
	err := service.Admin.ResetOtherPassword(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, "操作成功")
}

// Add 添加
func (c *Admin) Add(r *ghttp.Request) {
	err := service.Admin.Add(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, "操作成功")
}

// Edit 编辑
func (c *Admin) Edit(r *ghttp.Request) {
	err := service.Admin.Edit(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, "操作成功")
}

// Delete 删除
func (c *Admin) Delete(r *ghttp.Request) {
	err := service.Admin.Delete(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, "操作成功")
}

func (c *Admin) Info(r *ghttp.Request) {
	re, err := service.Admin.Info(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, re)
}
