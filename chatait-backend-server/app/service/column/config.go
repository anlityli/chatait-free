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

type Config struct {
}

func (c *Config) MidjourneyListColumns() (re *datalist.Columns) {
	re = &datalist.Columns{}
	re.ListName = "Midjourney配置列表"
	re.ListID = "configMidjourneyList"
	re.ColumnList = datalist.ColumnList{
		&datalist.ColumnItem{
			Field:  "id",
			Hidden: true,
		},
		&datalist.ColumnItem{
			Field:       "title",
			FieldName:   "配置标题",
			FieldAttr:   &datalist.FieldAttr{Width: 150},
			CanFilter:   true,
			FilterField: "title",
		},
		&datalist.ColumnItem{
			Field:       "guild_id",
			FieldName:   "服务ID(guild_id)",
			FieldAttr:   &datalist.FieldAttr{Width: 200},
			CanFilter:   true,
			FilterField: "guild_id",
		},
		&datalist.ColumnItem{
			Field:       "channel_id",
			FieldName:   "频道ID(channel_id)",
			FieldAttr:   &datalist.FieldAttr{Width: 200},
			CanFilter:   true,
			FilterField: "channel_id",
		},
		&datalist.ColumnItem{
			Field:       "user_token",
			FieldName:   "账户Token",
			FieldAttr:   &datalist.FieldAttr{Width: 200},
			CanFilter:   true,
			FilterField: "user_token",
		},
		&datalist.ColumnItem{
			Field:       "mj_bot_id",
			FieldName:   "MidjourneyBotId",
			FieldAttr:   &datalist.FieldAttr{Width: 200},
			CanFilter:   true,
			FilterField: "mj_bot_id",
		},
		&datalist.ColumnItem{
			Field:       "bot_token",
			FieldName:   "BotToken",
			FieldAttr:   &datalist.FieldAttr{Width: 200},
			CanFilter:   true,
			FilterField: "bot_token",
		},
		&datalist.ColumnItem{
			Field:       "session_id",
			FieldName:   "会话ID(session_id)",
			FieldAttr:   &datalist.FieldAttr{Width: 200},
			CanFilter:   true,
			FilterField: "session_id",
		},
		&datalist.ColumnItem{
			Field:       "proxy",
			FieldName:   "代理服务器",
			FieldAttr:   &datalist.FieldAttr{Width: 200},
			CanFilter:   true,
			FilterField: "proxy",
		},
		&datalist.ColumnItem{
			Field:     "status",
			FieldName: "是否启用",
			FieldAttr: &datalist.FieldAttr{Width: 120},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.ConfigMidjourney)
				return datalist.YesNoValue(row.Status)
			},
		},
		&datalist.ColumnItem{
			Field:     "listen_model",
			FieldName: "监听模式",
			FieldAttr: &datalist.FieldAttr{Width: 120},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.ConfigMidjourney)
				return datalist.YesNoValue(row.ListenModel, "UserWss", "Bot")
			},
		},
		&datalist.ColumnItem{
			Field:     "create_model",
			FieldName: "生图模式",
			FieldAttr: &datalist.FieldAttr{Width: 200},
		},
		&datalist.ColumnItem{
			Field:     "ws_idle_time",
			FieldName: "Websocket闲置时长(秒)",
			FieldAttr: &datalist.FieldAttr{Width: 200},
		},
		&datalist.ColumnItem{
			Field:     "call_num",
			FieldName: "调用次数",
			FieldAttr: &datalist.FieldAttr{Width: 150},
		},
		&datalist.ColumnItem{
			Field:     "created_at",
			FieldName: "创建时间",
			FieldAttr: &datalist.FieldAttr{Width: 200},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.ConfigMidjourney)
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
				row := rowData.(*response.ConfigMidjourney)
				if row.UpdatedAt == 0 {
					return ""
				}
				return gtime.NewFromTimeStamp(gconv.Int64(row.UpdatedAt)).Format("Y-m-d H:i:s")
			},
			CanFilter: true,
			FilterType: &datalist.FilterType{
				Attr: "date",
			},
			FilterField: "updated_at",
		},
	}
	return
}

func (c *Config) OpenaiListColumns() (re *datalist.Columns) {
	re = &datalist.Columns{}
	re.ListName = "Openai配置列表"
	re.ListID = "configOpenaiList"
	re.ColumnList = datalist.ColumnList{
		&datalist.ColumnItem{
			Field:  "id",
			Hidden: true,
		},
		&datalist.ColumnItem{
			Field:       "title",
			FieldName:   "配置标题",
			FieldAttr:   &datalist.FieldAttr{Width: 150},
			CanFilter:   true,
			FilterField: "title",
		},
		&datalist.ColumnItem{
			Field:       "api_url",
			FieldName:   "Api Url",
			FieldAttr:   &datalist.FieldAttr{Width: 200},
			CanFilter:   true,
			FilterField: "api_url",
		},
		&datalist.ColumnItem{
			Field:       "api_key",
			FieldName:   "Api Key",
			FieldAttr:   &datalist.FieldAttr{Width: 200},
			CanFilter:   true,
			FilterField: "api_key",
		},
		&datalist.ColumnItem{
			Field:       "proxy",
			FieldName:   "代理服务器",
			FieldAttr:   &datalist.FieldAttr{Width: 200},
			CanFilter:   true,
			FilterField: "proxy",
		},
		&datalist.ColumnItem{
			Field:       "max_tokens",
			FieldName:   "最大Token",
			FieldAttr:   &datalist.FieldAttr{Width: 150},
			CanFilter:   true,
			FilterField: "max_tokens",
		},
		&datalist.ColumnItem{
			Field:     "status",
			FieldName: "是否启用",
			FieldAttr: &datalist.FieldAttr{Width: 120},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.ConfigOpenai)
				if row.Status == 1 {
					return "是"
				}
				return "否"
			},
		},
		&datalist.ColumnItem{
			Field:     "call_num",
			FieldName: "调用次数",
			FieldAttr: &datalist.FieldAttr{Width: 150},
		},
		&datalist.ColumnItem{
			Field:     "created_at",
			FieldName: "创建时间",
			FieldAttr: &datalist.FieldAttr{Width: 200},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.ConfigOpenai)
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
				row := rowData.(*response.ConfigOpenai)
				if row.UpdatedAt == 0 {
					return ""
				}
				return gtime.NewFromTimeStamp(gconv.Int64(row.UpdatedAt)).Format("Y-m-d H:i:s")
			},
			CanFilter: true,
			FilterType: &datalist.FilterType{
				Attr: "date",
			},
			FilterField: "updated_at",
		},
	}
	return
}

func (c *Config) BaiduListColumns() (re *datalist.Columns) {
	re = &datalist.Columns{}
	re.ListName = "Baidu配置列表"
	re.ListID = "configBaiduList"
	re.ColumnList = datalist.ColumnList{
		&datalist.ColumnItem{
			Field:  "id",
			Hidden: true,
		},
		&datalist.ColumnItem{
			Field:       "title",
			FieldName:   "配置标题",
			FieldAttr:   &datalist.FieldAttr{Width: 150},
			CanFilter:   true,
			FilterField: "title",
		},
		&datalist.ColumnItem{
			Field:       "api_key",
			FieldName:   "Api Key",
			FieldAttr:   &datalist.FieldAttr{Width: 200},
			CanFilter:   true,
			FilterField: "api_key",
		},
		&datalist.ColumnItem{
			Field:       "secret_key",
			FieldName:   "Secret Key",
			FieldAttr:   &datalist.FieldAttr{Width: 200},
			CanFilter:   true,
			FilterField: "secret_key",
		},
		&datalist.ColumnItem{
			Field:     "status",
			FieldName: "是否启用",
			FieldAttr: &datalist.FieldAttr{Width: 120},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.ConfigBaidu)
				return datalist.YesNoValue(row.Status)
			},
		},
		&datalist.ColumnItem{
			Field:     "call_num",
			FieldName: "调用次数",
			FieldAttr: &datalist.FieldAttr{Width: 150},
		},
		&datalist.ColumnItem{
			Field:     "created_at",
			FieldName: "创建时间",
			FieldAttr: &datalist.FieldAttr{Width: 200},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.ConfigBaidu)
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
				row := rowData.(*response.ConfigBaidu)
				if row.UpdatedAt == 0 {
					return ""
				}
				return gtime.NewFromTimeStamp(gconv.Int64(row.UpdatedAt)).Format("Y-m-d H:i:s")
			},
			CanFilter: true,
			FilterType: &datalist.FilterType{
				Attr: "date",
			},
			FilterField: "updated_at",
		},
	}
	return
}

func (c *Config) SensitiveWordColumns() (re *datalist.Columns) {
	re = &datalist.Columns{}
	re.ListName = "敏感词列表"
	re.ListID = "configSensitiveWordList"
	re.ColumnList = datalist.ColumnList{
		&datalist.ColumnItem{
			Field:  "id",
			Hidden: true,
		},
		&datalist.ColumnItem{
			Field:       "content",
			FieldName:   "内容",
			FieldAttr:   &datalist.FieldAttr{Width: 150},
			CanFilter:   true,
			FilterField: "content",
		},
		&datalist.ColumnItem{
			Field:     "created_at",
			FieldName: "创建时间",
			FieldAttr: &datalist.FieldAttr{Width: 200},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.ConfigSensitiveWord)
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
	}
	return
}
