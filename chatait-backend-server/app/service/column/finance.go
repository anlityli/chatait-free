// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package column

import (
	"github.com/anlityli/chatait-free/chatait-backend-server/app/model/response"
	"github.com/anlityli/chatait-free/chatait-backend-server/library/datalist"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/helper"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

type Finance struct {
}

func (c *Finance) WalletListColumns() (re *datalist.Columns) {
	re = &datalist.Columns{}
	re.ListName = "钱包列表"
	re.ListID = "financeWalletList"
	re.ColumnList = datalist.ColumnList{
		&datalist.ColumnItem{
			Field:       "user_id",
			Hidden:      true,
			FilterField: "w.user_id",
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
			Field:       "balance",
			FieldName:   "充值余额",
			FieldAttr:   &datalist.FieldAttr{Width: 120},
			CanFilter:   true,
			FilterField: "w.balance",
			FilterType: &datalist.FilterType{
				Attr: "amount",
			},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.FinanceWalletItem)
				return helper.CentToYuan(row.Balance)
			},
		},
		&datalist.ColumnItem{
			Field:       "gpt3",
			FieldName:   "gpt3次数",
			FieldAttr:   &datalist.FieldAttr{Width: 120},
			CanFilter:   true,
			FilterField: "w.gpt3",
			FilterType: &datalist.FilterType{
				Attr: "amount",
			},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.FinanceWalletItem)
				return helper.CentToYuan(row.Gpt3)
			},
		},
		&datalist.ColumnItem{
			Field:       "gpt4",
			FieldName:   "gpt4次数",
			FieldAttr:   &datalist.FieldAttr{Width: 120},
			CanFilter:   true,
			FilterField: "w.gpt4",
			FilterType: &datalist.FilterType{
				Attr: "amount",
			},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.FinanceWalletItem)
				return helper.CentToYuan(row.Gpt4)
			},
		},
		&datalist.ColumnItem{
			Field:       "midjourney",
			FieldName:   "midjourney次数",
			FieldAttr:   &datalist.FieldAttr{Width: 120},
			CanFilter:   true,
			FilterField: "w.midjourney",
			FilterType: &datalist.FilterType{
				Attr: "amount",
			},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.FinanceWalletItem)
				return helper.CentToYuan(row.Midjourney)
			},
		},
	}
	return
}

func (c *Finance) WalletFlowListColumns(walletType string) (re *datalist.Columns) {
	re = &datalist.Columns{}
	re.ListName = "钱包流水"
	re.ListID = "financeWalletFlowList" + "_" + walletType
	re.ColumnList = datalist.ColumnList{
		&datalist.ColumnItem{
			Field:       "user_id",
			Hidden:      true,
			FilterField: "f.user_id",
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
			Field:       "amount",
			FieldName:   "充值余额",
			FieldAttr:   &datalist.FieldAttr{Width: 120},
			CanFilter:   true,
			FilterField: "f.amount",
			FilterType: &datalist.FilterType{
				Attr: "amount",
			},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.FinanceWalletFlowItem)
				return helper.CentToYuan(row.Amount)
			},
		},
		&datalist.ColumnItem{
			Field:     "total",
			FieldName: "变动后的余额",
			FieldAttr: &datalist.FieldAttr{Width: 120},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.FinanceWalletFlowItem)
				return helper.CentToYuan(row.Amount)
			},
		},
		&datalist.ColumnItem{
			Field:       "remark",
			FieldName:   "备注",
			FieldAttr:   &datalist.FieldAttr{Width: 120},
			CanFilter:   true,
			FilterField: "f.remark",
		},
		&datalist.ColumnItem{
			Field:       "admin_name",
			FieldName:   "操作管理员",
			FieldAttr:   &datalist.FieldAttr{Width: 120},
			CanFilter:   true,
			FilterField: "f.admin_name",
		},
		&datalist.ColumnItem{
			Field:     "created_at",
			FieldName: "创建时间",
			FieldAttr: &datalist.FieldAttr{Width: 200},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.FinanceWalletFlowItem)
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
