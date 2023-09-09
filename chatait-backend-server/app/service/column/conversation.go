// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package column

import (
	"github.com/anlityli/chatait-free/chatait-backend-server/app/model/response"
	"github.com/anlityli/chatait-free/chatait-backend-server/library/datalist"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/constant"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
)

type Conversation struct {
}

func (c *Conversation) TopicListColumns() (re *datalist.Columns) {
	re = &datalist.Columns{}
	re.ListName = "话题列表"
	re.ListID = "conversationTopicList"
	re.ColumnList = datalist.ColumnList{
		&datalist.ColumnItem{
			Field:  "id",
			Hidden: true,
		},
		&datalist.ColumnItem{
			Field:       "title",
			FieldName:   "话题",
			FieldAttr:   &datalist.FieldAttr{Width: 200},
			CanFilter:   true,
			FilterField: "t.title",
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
			Field:       "type",
			FieldName:   "类型",
			FieldAttr:   &datalist.FieldAttr{Width: 120},
			CanFilter:   true,
			FilterField: "t.type",
			FilterType:  c.topicTypeFilterType(),
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.ConversationTopic)
				if row.Type == constant.TopicTypeOpenaiGPT3 {
					return "GPT3.5"
				}
				return ""
			},
		},
		&datalist.ColumnItem{
			Field:     "created_at",
			FieldName: "创建时间",
			FieldAttr: &datalist.FieldAttr{Width: 200},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.ConversationTopic)
				if row.CreatedAt == 0 {
					return ""
				}
				return gtime.NewFromTimeStamp(gconv.Int64(row.CreatedAt)).Format("Y-m-d H:i:s")
			},
			CanFilter: true,
			FilterType: &datalist.FilterType{
				Attr: "date",
			},
			FilterField: "t.created_at",
		},
		&datalist.ColumnItem{
			Field:     "updated_at",
			FieldName: "修改时间",
			FieldAttr: &datalist.FieldAttr{Width: 200},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.ConversationTopic)
				if row.UpdatedAt == 0 {
					return ""
				}
				return gtime.NewFromTimeStamp(gconv.Int64(row.UpdatedAt)).Format("Y-m-d H:i:s")
			},
			CanFilter: true,
			FilterType: &datalist.FilterType{
				Attr: "date",
			},
			FilterField: "t.updated_at",
		},
	}
	return
}

func (c *Conversation) ListColumns() (re *datalist.Columns) {
	re = &datalist.Columns{}
	re.ListName = "对话列表"
	re.ListID = "conversationList"
	re.ColumnList = datalist.ColumnList{
		&datalist.ColumnItem{
			Field:  "id",
			Hidden: true,
		},
		&datalist.ColumnItem{
			Field:  "mj_data",
			Hidden: true,
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
			Field:       "role",
			FieldName:   "昵称",
			FieldAttr:   &datalist.FieldAttr{Width: 150},
			CanFilter:   true,
			FilterField: "c.role",
		},
		&datalist.ColumnItem{
			Field:       "topic_type",
			FieldName:   "话题类型",
			FieldAttr:   &datalist.FieldAttr{Width: 120},
			CanFilter:   true,
			FilterField: "t.topic_type",
			FilterType:  c.topicTypeFilterType(),
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.Conversation)
				if row.TopicType == constant.TopicTypeOpenaiGPT3 {
					return "GPT3.5"
				}
				return ""
			},
		},
		&datalist.ColumnItem{
			Field:     "content",
			FieldName: "内容",
			FieldAttr: &datalist.FieldAttr{Width: 200},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.Conversation)
				if row.Content != "" {
					row.Content = gstr.SubStrRune(row.Content, 0, 20) + "..."
				}
				return row.Content
			},
		},
		&datalist.ColumnItem{
			Field:     "created_at",
			FieldName: "创建时间",
			FieldAttr: &datalist.FieldAttr{Width: 200},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.Conversation)
				if row.CreatedAt == 0 {
					return ""
				}
				return gtime.NewFromTimeStamp(gconv.Int64(row.CreatedAt)).Format("Y-m-d H:i:s")
			},
			CanFilter: true,
			FilterType: &datalist.FilterType{
				Attr: "date",
			},
			FilterField: "c.created_at",
		},
	}
	return
}

func (c *Conversation) topicTypeFilterType() (re *datalist.FilterType) {
	// 获取全部会员级别
	selectData := g.Slice{
		g.Map{
			"label": "GPT3.5",
			"value": constant.TopicTypeOpenaiGPT3,
		},
	}
	return &datalist.FilterType{
		Attr:       "select",
		SelectData: selectData,
	}
}
