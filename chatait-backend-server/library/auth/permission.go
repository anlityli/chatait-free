// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package auth

import (
	"database/sql"
	"github.com/anlityli/chatait-free/chatait-backend-server/library/menu"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/model/entity"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/helper"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
)

// RolePermission 获取权限
func RolePermission(roleID string) (re []string, err error) {
	re = make([]string, 0)
	adminRoleData := &entity.AdminRole{}
	err = dao.AdminRole.Where("id=?", roleID).Scan(adminRoleData)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if adminRoleData.Id > 0 && adminRoleData.Permission != "" {
		permissionJson, err := gjson.Decode(adminRoleData.Permission)
		if err != nil {
			return nil, err
		}
		re = gconv.SliceStr(permissionJson)
	}

	// 循环全部权限，把菜单里关联到的路由，也加进去
	for _, route := range re {
		re = append(re, menu.Instance().GetRelevantRoutePathByKey(route)...)
	}

	return re, nil
}

// RoleColumnPermission 获取会员的列表字段权限
func RoleColumnPermission(roleID string) (re map[string]interface{}, err error) {
	adminRoleData := &entity.AdminRole{}
	err = dao.AdminRole.Where("id=?", roleID).Scan(adminRoleData)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if adminRoleData.Id > 0 && adminRoleData.ColumnPermission != "" {
		permissionJson, err := gjson.Decode(adminRoleData.ColumnPermission)
		if err != nil {
			return nil, err
		}
		re = gconv.Map(permissionJson)
	}
	return re, nil
}

// RoleMenu 角色对应的菜单
func RoleMenu(r *ghttp.Request) (re menu.ListModel) {
	roleID := GetRoleID(r)
	permissionData, err := RolePermission(gconv.String(roleID))
	if err != nil {
		return nil
	}
	permissionMainData := roleMainRoutePath(permissionData)
	menuList := menu.Instance().GetList()
	re = make(menu.ListModel, 0)
	for _, item := range menuList {
		// 菜单是否显示
		if !item.Show {
			continue
		}
		// 是否有该控制器的权限
		if roleID != 1 && !helper.StrInArr(permissionMainData, item.RoutePath) {
			continue
		}
		tmpItem := &menu.ItemModel{}
		tmpItem.Key = item.Key
		tmpItem.Title = item.Title
		tmpItem.RoutePath = item.RoutePath
		tmpItem.Show = item.Show
		tmpItem.Leaf = item.Leaf
		// 循环子菜单
		if item.Children != nil && len(item.Children) > 0 {
			tmpItem.Children = make(menu.ListModel, 0)
			for _, childItem := range item.Children {
				if !childItem.Show {
					continue
				}
				if roleID != 1 && !helper.StrInArr(permissionData, childItem.RoutePath) {
					continue
				}
				tmpItem.Children = append(tmpItem.Children, childItem)
			}
		}
		re = append(re, tmpItem)
	}
	return re
}

// roleMainRoutePath 角色权限中的一级注
func roleMainRoutePath(permissionData []string) (re []string) {
	re = make([]string, 0)
	if len(permissionData) > 0 {
		for _, path := range permissionData {
			pathSlice := gstr.Explode("/", path)
			if len(pathSlice) > 0 {
				if !helper.StrInArr(re, pathSlice[0]) {
					re = append(re, pathSlice[0])
				}
			}
		}
	}
	return re
}
