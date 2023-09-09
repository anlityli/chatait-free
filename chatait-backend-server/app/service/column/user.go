// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package column

import (
	"github.com/anlityli/chatait-free/chatait-backend-server/app/model/response"
	"github.com/anlityli/chatait-free/chatait-backend-server/library/datalist"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/constant"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

type User struct {
}

func (c *User) ListColumns() (re *datalist.Columns) {
	re = &datalist.Columns{}
	re.ListName = "会员列表"
	re.ListID = "userList"
	re.ColumnList = datalist.ColumnList{
		&datalist.ColumnItem{
			Field:       "id",
			FieldName:   "会员ID",
			FieldAttr:   &datalist.FieldAttr{Width: 200},
			CanFilter:   true,
			FilterField: "u.id",
		},
		&datalist.ColumnItem{
			Field:       "avatar",
			FieldName:   "头像",
			FieldAttr:   &datalist.FieldAttr{Width: 120},
			FilterField: "ui.avatar",
		},
		&datalist.ColumnItem{
			Field:       "username",
			FieldName:   "会员名",
			FieldAttr:   &datalist.FieldAttr{Width: 150},
			CanFilter:   true,
			FilterField: "u.username",
		},
		&datalist.ColumnItem{
			Field:       "nickname",
			FieldName:   "昵称",
			FieldAttr:   &datalist.FieldAttr{Width: 120},
			CanFilter:   true,
			FilterField: "ui.nickname",
		},
		&datalist.ColumnItem{
			Field:       "level_id",
			FieldName:   "级别",
			FieldAttr:   &datalist.FieldAttr{Width: 120},
			CanFilter:   true,
			FilterField: "u.level_id",
			FilterType:  c.levelFilterType(),
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.UserListItem)
				if row.LevelId == constant.ConfigLevelMember {
					return "member"
				} else if row.LevelId == constant.ConfigLevelPlus {
					return "plus"
				} else {
					return ""
				}
			},
		},
		&datalist.ColumnItem{
			Field:       "level_expire_date",
			FieldName:   "级别到期日期",
			FieldAttr:   &datalist.FieldAttr{Width: 120},
			CanFilter:   true,
			FilterField: "u.level_expire_date",
		},
		&datalist.ColumnItem{
			Field:       "is_ban",
			FieldName:   "是否禁用",
			FieldAttr:   &datalist.FieldAttr{Width: 120},
			CanFilter:   true,
			FilterField: "u.is_ban",
			FilterType:  datalist.YesNoFilterType(),
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.UserListItem)
				return datalist.YesNoValue(row.IsBan)
			},
		},
		&datalist.ColumnItem{
			Field:     "last_login_at",
			FieldName: "最后登录时间",
			FieldAttr: &datalist.FieldAttr{Width: 200},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.UserListItem)
				if row.LastLoginAt == 0 {
					return ""
				}
				return gtime.NewFromTimeStamp(gconv.Int64(row.LastLoginAt)).Format("Y-m-d H:i:s")
			},
			CanFilter: true,
			FilterType: &datalist.FilterType{
				Attr: "date",
			},
			FilterField: "u.last_login_at",
		},
		&datalist.ColumnItem{
			Field:     "created_at",
			FieldName: "创建时间",
			FieldAttr: &datalist.FieldAttr{Width: 200},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.UserListItem)
				if row.CreatedAt == 0 {
					return ""
				}
				return gtime.NewFromTimeStamp(gconv.Int64(row.CreatedAt)).Format("Y-m-d H:i:s")
			},
			CanFilter: true,
			FilterType: &datalist.FilterType{
				Attr: "date",
			},
			FilterField: "u.created_at",
		},
		&datalist.ColumnItem{
			Field:     "updated_at",
			FieldName: "修改时间",
			FieldAttr: &datalist.FieldAttr{Width: 200},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.UserListItem)
				if row.UpdatedAt == 0 {
					return ""
				}
				return gtime.NewFromTimeStamp(gconv.Int64(row.UpdatedAt)).Format("Y-m-d H:i:s")
			},
			CanFilter: true,
			FilterType: &datalist.FilterType{
				Attr: "date",
			},
			FilterField: "u.updated_at",
		},
	}
	return
}

func (c *User) levelFilterType() (re *datalist.FilterType) {
	// 获取全部会员级别
	data, err := dao.ConfigLevel.Where("1=1").Order("id ASC").All()
	if err != nil {
		return nil
	}
	selectData := g.Slice{}
	for _, oneData := range data {
		selectData = append(selectData, g.Map{
			"value": gconv.String(oneData["id"]),
			"label": gconv.String(oneData["level_name"]),
		})
	}
	return &datalist.FilterType{
		Attr:       "select",
		SelectData: selectData,
	}
}
