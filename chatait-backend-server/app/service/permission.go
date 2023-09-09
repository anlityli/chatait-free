// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package service

import (
	"github.com/anlityli/chatait-free/chatait-backend-server/app/model/request"
	"github.com/anlityli/chatait-free/chatait-backend-server/app/model/response"
	"github.com/anlityli/chatait-free/chatait-backend-server/app/service/column"
	"github.com/anlityli/chatait-free/chatait-backend-server/library/datalist"
	"github.com/anlityli/chatait-free/chatait-backend-server/library/menu"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/helper"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"reflect"
)

var Permission = &permissionService{}

type permissionService struct {
}

// RolePermission 获取角色全部的权限
func (s *permissionService) RolePermission(r *ghttp.Request) (re response.AdminAllPermissionList, err error) {
	requestModel := &request.AdminRolePermission{}
	if err := r.Parse(requestModel); err != nil {
		return nil, err
	}
	// 获取该角色的所有权限
	permissionData, err := dao.AdminRole.Where("id=?", requestModel.ID).One()
	if err != nil {
		return nil, err
	}
	permissionSlice := make([]string, 0)
	if !permissionData.IsEmpty() {
		permissionStr := gconv.String(permissionData["permission"])
		if permissionStr != "" {
			permissionJson, err := gjson.Decode(permissionStr)
			if err != nil {
				return nil, err
			}
			permissionSlice = gconv.SliceStr(permissionJson)
		}
	}
	re = make(response.AdminAllPermissionList, 0)
	// 获取所有的权限
	menuList := menu.Instance().GetList()
	for _, main := range menuList {
		permissionItem := &response.AdminAllPermissionItem{}
		permissionItem.ID = main.Key
		permissionItem.MainPermission = &response.AdminPermissionItem{
			Title: main.Title,
			Path:  main.RoutePath,
		}
		for _, child := range main.Children {
			childPermissionItem := &response.AdminPermissionItem{}
			childPermissionItem.Title = child.Title
			childPermissionItem.Path = child.RoutePath
			if helper.StrInArr(permissionSlice, child.RoutePath) {
				childPermissionItem.IsChecked = true
			}
			permissionItem.ChildPermission = append(permissionItem.ChildPermission, childPermissionItem)
		}
		re = append(re, permissionItem)
	}
	return re, nil
}

// RolePermissionEdit 编辑权限
func (s *permissionService) RolePermissionEdit(r *ghttp.Request) (err error) {
	requestModel := &request.AdminRolePermissionEdit{}
	if err := r.Parse(requestModel); err != nil {
		return err
	}
	permissionJson, err := gjson.Encode(requestModel.Permission)
	if err != nil {
		return err
	}
	if _, err := dao.AdminRole.Where("id=?", requestModel.ID).Data(g.Map{
		"permission": gconv.String(permissionJson),
	}).Update(); err != nil {
		return err
	}
	return nil
}

// RoleColumnPermission 管理员角色对应的所有表的全部字段权限
func (s *permissionService) RoleColumnPermission(r *ghttp.Request) (re response.AdminColumnPermissionList, err error) {
	requestModel := &request.AdminRolePermission{}
	if err := r.Parse(requestModel); err != nil {
		return nil, err
	}
	// 获取该角色的所有权限
	permissionData, err := dao.AdminRole.Where("id=?", requestModel.ID).One()
	if err != nil {
		return nil, err
	}
	permissionMap := make(map[string]interface{}, 0)
	if !permissionData.IsEmpty() {
		permissionStr := gconv.String(permissionData["column_permission"])
		if permissionStr != "" {
			permissionJson, err := gjson.Decode(permissionStr)
			if err != nil {
				return nil, err
			}
			permissionMap = gconv.Map(permissionJson)
		}
	}

	re = make(response.AdminColumnPermissionList, 0)
	columnObjs := column.AllListColumn()
	for _, obj := range columnObjs {
		permission := s.oneRoleColumnPermission(obj, permissionMap)
		re = append(re, permission...)
	}
	return re, nil
}

func (s *permissionService) oneRoleColumnPermission(columnObj interface{}, permissionMap map[string]interface{}) (re response.AdminColumnPermissionList) {
	re = make(response.AdminColumnPermissionList, 0)
	columnObjValue := reflect.ValueOf(columnObj)
	//columnObjType := columnObjValue.Type()
	for i := 0; i < columnObjValue.NumMethod(); i++ {
		//methodName := columnObjType.Method(i).Name
		inputs := make([]reflect.Value, 0)
		if columnObjValue.Method(i).Type().String() == "func() *datalist.Columns" {
			//glog.Line().Debug(columnObjType.Method(i).Name)
			callRe := columnObjValue.Method(i).Call(inputs)
			if len(callRe) > 0 {
				callReInterface := callRe[0].Interface()
				// if reflect.TypeOf(callReInterface).String() == "*datalist.Columns" {
				columnItem := &response.AdminColumnPermissionItem{}
				columnsRe := callReInterface.(*datalist.Columns)
				columnListRe := columnsRe.ColumnList
				columnItem.ListID = columnsRe.ListID
				columnItem.ListName = columnsRe.ListName
				columnItem.Columns = make([]*response.AdminColumnPermissionItemColumns, 0)

				columnPermission := make([]string, 0)
				if columnPermissionInterface, ok := permissionMap[columnItem.ListID]; ok {
					columnPermission = gconv.SliceStr(columnPermissionInterface)
				}

				for _, oneColumn := range columnListRe {
					tempColumn := &response.AdminColumnPermissionItemColumns{}
					if oneColumn.Hidden {
						continue
					}
					tempColumn.Field = oneColumn.Field
					tempColumn.FieldName = oneColumn.FieldName
					if helper.StrInArr(columnPermission, tempColumn.Field) {
						tempColumn.IsChecked = true
					}
					columnItem.Columns = append(columnItem.Columns, tempColumn)
				}
				re = append(re, columnItem)
				//}
			}
		}
	}

	return re
}

// RoleColumnPermissionEdit 编辑列表字段权限
func (s *permissionService) RoleColumnPermissionEdit(r *ghttp.Request) (err error) {
	requestModel := &request.AdminRoleColumnPermissionEdit{}
	if err := r.Parse(requestModel); err != nil {
		return err
	}
	permissionJson, err := gjson.Encode(requestModel.ColumnPermission)
	if err != nil {
		return err
	}
	if _, err := dao.AdminRole.Where("id=?", requestModel.ID).Data(g.Map{
		"column_permission": gconv.String(permissionJson),
	}).Update(); err != nil {
		return err
	}
	return nil
}
