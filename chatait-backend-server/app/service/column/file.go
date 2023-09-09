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

type File struct {
}

func (c *File) MidjourneyListColumns() (re *datalist.Columns) {
	re = &datalist.Columns{}
	re.ListName = "Midjourney文件列表"
	re.ListID = "fileMidjourneyList"
	re.ColumnList = datalist.ColumnList{
		&datalist.ColumnItem{
			Field:  "id",
			Hidden: true,
		},
		&datalist.ColumnItem{
			Field:  "prompt",
			Hidden: true,
		},
		&datalist.ColumnItem{
			Field:  "mj_file_name",
			Hidden: true,
		},
		&datalist.ColumnItem{
			Field:  "mj_url",
			Hidden: true,
		},
		&datalist.ColumnItem{
			Field:     "path",
			FieldName: "图片",
			FieldAttr: &datalist.FieldAttr{Width: 100},
		},
		&datalist.ColumnItem{
			Field:       "file_name",
			FieldName:   "文件名",
			FieldAttr:   &datalist.FieldAttr{Width: 200},
			CanFilter:   true,
			FilterField: "f.file_name",
		},
		&datalist.ColumnItem{
			Field:       "username",
			FieldName:   "会员",
			FieldAttr:   &datalist.FieldAttr{Width: 200},
			CanFilter:   true,
			FilterField: "u.username",
		},
		&datalist.ColumnItem{
			Field:       "nickname",
			FieldName:   "昵称",
			FieldAttr:   &datalist.FieldAttr{Width: 200},
			CanFilter:   true,
			FilterField: "ui.nickname",
		},
		&datalist.ColumnItem{
			Field:     "size",
			FieldName: "大小",
			FieldAttr: &datalist.FieldAttr{Width: 200},
		},
		&datalist.ColumnItem{
			Field:     "created_at",
			FieldName: "创建时间",
			FieldAttr: &datalist.FieldAttr{Width: 200},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.FileMidjourney)
				if row.CreatedAt == 0 {
					return ""
				}
				return gtime.NewFromTimeStamp(gconv.Int64(row.CreatedAt)).Format("Y-m-d H:i:s")
			},
			CanFilter: true,
			FilterType: &datalist.FilterType{
				Attr: "date",
			},
			FilterField: "f.created_at",
		},
	}
	return
}
