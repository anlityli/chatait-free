// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package column

import (
	"github.com/anlityli/chatait-free/chatait-backend-server/app/model/response"
	"github.com/anlityli/chatait-free/chatait-backend-server/library/datalist"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

type Admin struct {
}

func (a *Admin) ListColumns() (re *datalist.Columns) {
	re = &datalist.Columns{}
	re.ListName = "管理员列表"
	re.ListID = "adminList"
	re.ColumnList = datalist.ColumnList{
		&datalist.ColumnItem{
			Field:       "id",
			FieldName:   "管理员ID",
			FieldAttr:   &datalist.FieldAttr{Width: 200},
			CanFilter:   true,
			FilterField: "a.id",
		},
		&datalist.ColumnItem{
			Field:       "admin_name",
			FieldName:   "管理员名",
			FieldAttr:   &datalist.FieldAttr{Width: 150},
			CanFilter:   true,
			FilterField: "a.admin_name",
		},
		&datalist.ColumnItem{
			Field:       "real_name",
			FieldName:   "姓名",
			FieldAttr:   &datalist.FieldAttr{Width: 120},
			CanFilter:   true,
			FilterField: "a.real_name",
		},
		&datalist.ColumnItem{
			Field:     "remark",
			FieldName: "备注",
			FieldAttr: &datalist.FieldAttr{Width: 150},
		},
		&datalist.ColumnItem{
			Field:     "role_name",
			FieldName: "角色",
			FieldAttr: &datalist.FieldAttr{Width: 150},
		},
		&datalist.ColumnItem{
			Field:     "is_enable",
			FieldName: "是否启用",
			FieldAttr: &datalist.FieldAttr{Width: 120},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.AdminListItem)
				if row.IsEnable == 1 {
					return "是"
				}
				return "否"
			},
		},
		&datalist.ColumnItem{
			Field:     "login_nums",
			FieldName: "登录次数",
			FieldAttr: &datalist.FieldAttr{Width: 100},
		},
		&datalist.ColumnItem{
			Field:     "last_login_ip",
			FieldName: "最后登陆IP",
			FieldAttr: &datalist.FieldAttr{Width: 150},
		},
		&datalist.ColumnItem{
			Field:     "last_login_at",
			FieldName: "最后登录时间",
			FieldAttr: &datalist.FieldAttr{Width: 200},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.AdminListItem)
				if row.LastLoginAt == 0 {
					return ""
				}
				return gtime.NewFromTimeStamp(gconv.Int64(row.LastLoginAt)).Format("Y-m-d H:i:s")
			},
			CanFilter: true,
			FilterType: &datalist.FilterType{
				Attr: "date",
			},
			FilterField: "a.last_login_at",
		},
		&datalist.ColumnItem{
			Field:     "created_at",
			FieldName: "创建时间",
			FieldAttr: &datalist.FieldAttr{Width: 200},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.AdminListItem)
				if row.CreatedAt == 0 {
					return ""
				}
				return gtime.NewFromTimeStamp(gconv.Int64(row.CreatedAt)).Format("Y-m-d H:i:s")
			},
			CanFilter: true,
			FilterType: &datalist.FilterType{
				Attr: "date",
			},
			FilterField: "a.created_at",
		},
		&datalist.ColumnItem{
			Field:     "updated_at",
			FieldName: "修改时间",
			FieldAttr: &datalist.FieldAttr{Width: 200},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.AdminListItem)
				if row.UpdatedAt == 0 {
					return ""
				}
				return gtime.NewFromTimeStamp(gconv.Int64(row.UpdatedAt)).Format("Y-m-d H:i:s")
			},
			CanFilter: true,
			FilterType: &datalist.FilterType{
				Attr: "date",
			},
			FilterField: "a.updated_at",
		},
		&datalist.ColumnItem{
			Field:     "dont_del",
			FieldName: "是否可删除",
			FieldAttr: &datalist.FieldAttr{Width: 200},
			Hidden:    true,
		},
	}
	return
}

func (a *Admin) RoleListColumns() (re *datalist.Columns) {
	re = &datalist.Columns{}
	re.ListName = "管理员角色列表"
	re.ListID = "adminRoleList"
	re.ColumnList = datalist.ColumnList{
		&datalist.ColumnItem{
			Field:       "id",
			FieldName:   "角色ID",
			FieldAttr:   &datalist.FieldAttr{Width: 200},
			CanFilter:   true,
			FilterField: "u.id",
		},
		&datalist.ColumnItem{
			Field:     "role_name",
			FieldName: "角色名",
			FieldAttr: &datalist.FieldAttr{Width: 150},
		},
		&datalist.ColumnItem{
			Field:     "remark",
			FieldName: "备注",
			FieldAttr: &datalist.FieldAttr{Width: 150},
		},
		&datalist.ColumnItem{
			Field:     "create_admin",
			FieldName: "创建人",
			FieldAttr: &datalist.FieldAttr{Width: 150},
		},
		&datalist.ColumnItem{
			Field:     "update_admin",
			FieldName: "更新人",
			FieldAttr: &datalist.FieldAttr{Width: 150},
		},
		&datalist.ColumnItem{
			Field:     "created_at",
			FieldName: "创建时间",
			FieldAttr: &datalist.FieldAttr{Width: 200},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.AdminRoleItem)
				if row.CreatedAt == 0 {
					return ""
				}
				return gtime.NewFromTimeStamp(gconv.Int64(row.CreatedAt)).Format("Y-m-d H:i:s")
			},
			CanFilter: true,
			FilterType: &datalist.FilterType{
				Attr: "date",
			},
			FilterField: "created_at",
		},
		&datalist.ColumnItem{
			Field:     "updated_at",
			FieldName: "修改时间",
			FieldAttr: &datalist.FieldAttr{Width: 200},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.AdminRoleItem)
				if row.UpdatedAt == 0 {
					return ""
				}
				return gtime.NewFromTimeStamp(gconv.Int64(row.UpdatedAt)).Format("Y-m-d H:i:s")
			},
			CanFilter: true,
			FilterType: &datalist.FilterType{
				Attr: "date",
			},
			FilterField: "created_at",
		},
		&datalist.ColumnItem{
			Field:     "dont_del",
			FieldName: "是否可删除",
			FieldAttr: &datalist.FieldAttr{Width: 200},
			Hidden:    true,
		},
	}
	return
}
