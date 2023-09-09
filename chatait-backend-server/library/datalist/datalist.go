// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package datalist

import (
	"errors"
	"github.com/anlityli/chatait-free/chatait-backend-server/library/auth"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/helper"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/page"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// List 列表返回结果(主要是为了兼容之前的前端)
func List(r *ghttp.Request, pageData *page.Response, columns *Columns) (re *Result, err error) {
	columnList := columns.ColumnList
	// 获取管理员的角色
	roleID := auth.GetRoleID(r)
	// 获取管理员对应的列表字段权限
	columnPermission, err := auth.RoleColumnPermission(gconv.String(roleID))
	if err != nil {
		return nil, err
	}
	listID := columns.ListID
	_, columnOk := columnPermission[listID]
	if !columnOk {
		listID = columns.RelevantListID
		_, columnOk = columnPermission[listID]
	}
	if !columnOk {
		if roleID != 1 {
			return nil, errors.New("无此列表权限")
		}
	}
	columnPermissionSlice := gconv.SliceStr(columnPermission[listID])
	// 处理隐藏的column
	reColumns := make(ColumnList, 0)
	reFilterTypes := make([]*ResultFilterTypesItem, 0)
	for _, column := range columnList {
		// 判断字段权限
		if roleID != 1 && !helper.StrInArr(columnPermissionSlice, column.Field) {
			continue
		}
		// 可筛选内容
		if column.CanFilter {
			filterType := &ResultFilterTypesItem{}
			filterType.Field = column.Field
			filterType.FieldName = column.FieldName
			if column.FilterType != nil {
				filterType.Attr = column.FilterType.Attr
				filterType.AttrData = column.FilterType.SelectData
			}
			reFilterTypes = append(reFilterTypes, filterType)
		}
		if column.Hidden {
			continue
		}
		reColumns = append(reColumns, column)
	}
	re = &Result{}
	re.ListID = columns.ListID
	re.Columns = reColumns
	re.FilterTypes = reFilterTypes
	re.Page = pageData.Page
	re.PageSize = pageData.PageSize
	re.PageCount = pageData.PageCount
	re.TotalCount = pageData.TotalCount
	re.TotalPage = pageData.TotalPage
	re.List = make(ResultList, 0)
	if len(gconv.SliceAny(pageData.ListData)) == 0 {
		return
	}
	//listData := reflect.ValueOf(pageData.ListData).Elem()
	//listDataInterface := listData.Interface()
	listDataSlice := gconv.SliceAny(pageData.ListData)
	for _, rowStruct := range listDataSlice {
		rowMap := gconv.Map(rowStruct)
		tempMap := make(map[string]*ResultListValue)
		for rowKey, rowValue := range rowMap {
			tempMap[rowKey] = &ResultListValue{
				Value:    gconv.String(rowValue),
				OriValue: rowValue,
			}
		}
		reListValueMap := make(map[string]*ResultListValue)
		for _, column := range columnList {
			// 判断字段权限
			if roleID != 1 && !column.Hidden && !helper.StrInArr(columnPermissionSlice, column.Field) {
				continue
			}
			// 如果列里面有这个字段，就能付到结果集里
			if _, ok := tempMap[column.Field]; ok {
				reListValueMap[column.Field] = tempMap[column.Field]
			} else {
				reListValueMap[column.Field] = &ResultListValue{}
			}
			// 如果有值的回调，那么就用回调函数约束一下输出的值
			if column.ValueCallBack != nil {
				reListValueMap[column.Field].Value = column.ValueCallBack(rowStruct)
			}
			if column.OriValueCallBack != nil {
				reListValueMap[column.Field].OriValue = column.OriValueCallBack(rowStruct)
			}
			if column.ValueAttrCallBack != nil {
				reListValueMap[column.Field].ValueAttr = column.ValueAttrCallBack(rowStruct)
			}
		}
		re.List = append(re.List, reListValueMap)
	}
	return re, nil
}
