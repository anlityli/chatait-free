// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package column

import (
	"github.com/anlityli/chatait-free/chatait-backend-server/app/model/response"
	"github.com/anlityli/chatait-free/chatait-backend-server/library/datalist"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/constant"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/helper"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

type Shop struct {
}

func (c *Shop) GoodsListColumns() (re *datalist.Columns) {
	re = &datalist.Columns{}
	re.ListName = "商品列表"
	re.ListID = "shopGoodsList"
	re.ColumnList = datalist.ColumnList{
		&datalist.ColumnItem{
			Field:       "id",
			Hidden:      true,
			FilterField: "id",
		},
		&datalist.ColumnItem{
			Field:       "title",
			FieldName:   "标题",
			FieldAttr:   &datalist.FieldAttr{Width: 150},
			CanFilter:   true,
			FilterField: "title",
		},
		&datalist.ColumnItem{
			Field:       "buy_type",
			FieldName:   "购买类型",
			FieldAttr:   &datalist.FieldAttr{Width: 120},
			CanFilter:   true,
			FilterField: "buy_type",
			FilterType:  c.buyTypeFilterType(),
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.ShopGoods)
				if row.BuyType == constant.ShopGoodsBuyTypeLevel {
					return "购买级别"
				} else if row.BuyType == constant.ShopGoodsBuyTypeBalance {
					return "购买" + helper.GetWalletName(constant.WalletTypeBalance)
				} else if row.BuyType == constant.ShopGoodsBuyTypeGpt3 {
					return "购买" + helper.GetWalletName(constant.WalletTypeGpt3)
				} else if row.BuyType == constant.ShopGoodsBuyTypeGpt4 {
					return "购买" + helper.GetWalletName(constant.WalletTypeGpt4)
				}
				return ""
			},
		},
		&datalist.ColumnItem{
			Field:       "active_level_id",
			FieldName:   "开通级别",
			FieldAttr:   &datalist.FieldAttr{Width: 120},
			CanFilter:   true,
			FilterField: "active_level_id",
			FilterType:  c.levelFilterType(),
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.ShopGoods)
				if row.BuyType != constant.ShopGoodsBuyTypeLevel {
					return ""
				}
				levelData, err := helper.GetConfigLevel(row.ActiveLevelId)
				if err == nil {
					return levelData.LevelName
				} else {
					return ""
				}
			},
		},
		&datalist.ColumnItem{
			Field:       "active_expire_type",
			FieldName:   "开通有效期类型",
			FieldAttr:   &datalist.FieldAttr{Width: 120},
			CanFilter:   true,
			FilterField: "active_expire_type",
			FilterType:  c.expireTypeFilterType(),
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.ShopGoods)
				if row.ActiveExpireType == constant.ShopGoodsActiveExpireTypeNone {
					return "不开通级别"
				} else if row.ActiveExpireType == constant.ShopGoodsActiveExpireTypeDay {
					return "按天开通"
				} else if row.ActiveExpireType == constant.ShopGoodsActiveExpireTypeMonth {
					return "按月开通"
				} else if row.ActiveExpireType == constant.ShopGoodsActiveExpireTypeYear {
					return "按年开通"
				}
				return ""
			},
		},
		&datalist.ColumnItem{
			Field:       "active_expire_value",
			FieldName:   "开通时长",
			FieldAttr:   &datalist.FieldAttr{Width: 150},
			CanFilter:   true,
			FilterField: "active_expire_value",
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.ShopGoods)
				if row.ActiveExpireType == constant.ShopGoodsActiveExpireTypeNone {
					return ""
				} else if row.ActiveExpireType == constant.ShopGoodsActiveExpireTypeDay {
					return gconv.String(row.ActiveExpireValue) + "天"
				} else if row.ActiveExpireType == constant.ShopGoodsActiveExpireTypeMonth {
					return gconv.String(row.ActiveExpireValue) + "月"
				} else if row.ActiveExpireType == constant.ShopGoodsActiveExpireTypeYear {
					return gconv.String(row.ActiveExpireValue) + "日"
				}
				return ""
			},
		},
		&datalist.ColumnItem{
			Field:       "buy_value",
			FieldName:   "购买数量",
			FieldAttr:   &datalist.FieldAttr{Width: 120},
			CanFilter:   true,
			FilterField: "buy_value",
			FilterType: &datalist.FilterType{
				Attr: "amount",
			},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.ShopGoods)
				return helper.CentToYuan(row.BuyValue)
			},
		},
		&datalist.ColumnItem{
			Field:       "market_price",
			FieldName:   "市场价",
			FieldAttr:   &datalist.FieldAttr{Width: 120},
			CanFilter:   true,
			FilterField: "market_price",
			FilterType: &datalist.FilterType{
				Attr: "amount",
			},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.ShopGoods)
				return helper.CentToYuan(row.MarketPrice)
			},
		},
		&datalist.ColumnItem{
			Field:       "real_price",
			FieldName:   "实际价格",
			FieldAttr:   &datalist.FieldAttr{Width: 120},
			CanFilter:   true,
			FilterField: "real_price",
			FilterType: &datalist.FilterType{
				Attr: "amount",
			},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.ShopGoods)
				return helper.CentToYuan(row.RealPrice)
			},
		},
		&datalist.ColumnItem{
			Field:       "status",
			FieldName:   "是否上架",
			FieldAttr:   &datalist.FieldAttr{Width: 120},
			CanFilter:   true,
			FilterField: "status",
			FilterType:  datalist.YesNoFilterType(),
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.ShopGoods)
				return datalist.YesNoValue(row.Status)
			},
		},
		&datalist.ColumnItem{
			Field:       "sort",
			FieldName:   "排序",
			FieldAttr:   &datalist.FieldAttr{Width: 150},
			FilterField: "sort",
		},
		&datalist.ColumnItem{
			Field:     "created_at",
			FieldName: "创建时间",
			FieldAttr: &datalist.FieldAttr{Width: 200},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.ShopGoods)
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
			FieldName: "更新时间",
			FieldAttr: &datalist.FieldAttr{Width: 200},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.ShopGoods)
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

func (c *Shop) OrderListColumns() (re *datalist.Columns) {
	re = &datalist.Columns{}
	re.ListName = "订单列表"
	re.ListID = "shopOrderList"
	re.ColumnList = datalist.ColumnList{
		&datalist.ColumnItem{
			Field:       "id",
			Hidden:      true,
			FilterField: "o.id",
		},
		&datalist.ColumnItem{
			Field:       "order_sn",
			FieldName:   "订单编号",
			FieldAttr:   &datalist.FieldAttr{Width: 150},
			CanFilter:   true,
			FilterField: "o.order_sn",
		},
		&datalist.ColumnItem{
			Field:       "username",
			FieldName:   "会员名",
			FieldAttr:   &datalist.FieldAttr{Width: 120},
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
			Field:       "order_amount",
			FieldName:   "订单金额",
			FieldAttr:   &datalist.FieldAttr{Width: 120},
			CanFilter:   true,
			FilterField: "o.order_amount",
			FilterType: &datalist.FilterType{
				Attr: "amount",
			},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.ShopOrder)
				return helper.CentToYuan(row.OrderAmount)
			},
		},
		&datalist.ColumnItem{
			Field:       "pay_amount",
			FieldName:   "实付金额",
			FieldAttr:   &datalist.FieldAttr{Width: 120},
			CanFilter:   true,
			FilterField: "o.pay_amount",
			FilterType: &datalist.FilterType{
				Attr: "amount",
			},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.ShopOrder)
				return helper.CentToYuan(row.PayAmount)
			},
		},
		&datalist.ColumnItem{
			Field:       "status",
			FieldName:   "开通有效期类型",
			FieldAttr:   &datalist.FieldAttr{Width: 120},
			CanFilter:   true,
			FilterField: "o.status",
			FilterType:  c.orderStatusFilterType(),
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.ShopOrder)
				if row.Status == constant.ShopOrderStatusInit {
					return "待支付"
				} else if row.Status == constant.ShopOrderStatusPaid {
					return "已支付"
				} else if row.Status == constant.ShopOrderStatusFinish {
					return "已完成"
				} else if row.Status == constant.ShopOrderStatusCancel {
					return "已取消"
				}
				return ""
			},
		},
		&datalist.ColumnItem{
			Field:     "created_at",
			FieldName: "创建时间",
			FieldAttr: &datalist.FieldAttr{Width: 200},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.ShopOrder)
				if row.CreatedAt == 0 {
					return ""
				}
				return gtime.NewFromTimeStamp(gconv.Int64(row.CreatedAt)).Format("Y-m-d H:i:s")
			},
			CanFilter: true,
			FilterType: &datalist.FilterType{
				Attr: "date",
			},
			FilterField: "o.created_at",
		},
		&datalist.ColumnItem{
			Field:     "expire_at",
			FieldName: "过期时间",
			FieldAttr: &datalist.FieldAttr{Width: 200},
			ValueCallBack: func(rowData interface{}) string {
				row := rowData.(*response.ShopOrder)
				if row.ExpireAt == 0 {
					return ""
				}
				return gtime.NewFromTimeStamp(gconv.Int64(row.ExpireAt)).Format("Y-m-d H:i:s")
			},
			CanFilter: true,
			FilterType: &datalist.FilterType{
				Attr: "date",
			},
			FilterField: "o.expire_at",
		},
	}
	return
}

func (c *Shop) buyTypeFilterType() (re *datalist.FilterType) {
	// 获取全部会员级别
	selectData := g.Slice{
		g.Map{
			"label": "购买级别",
			"value": constant.ShopGoodsBuyTypeLevel,
		},
		g.Map{
			"label": "购买" + helper.GetWalletName(constant.WalletTypeBalance),
			"value": constant.ShopGoodsBuyTypeBalance,
		},
		g.Map{
			"label": "购买" + helper.GetWalletName(constant.WalletTypeGpt3),
			"value": constant.ShopGoodsBuyTypeGpt3,
		},
		g.Map{
			"label": "购买" + helper.GetWalletName(constant.WalletTypeGpt4),
			"value": constant.ShopGoodsBuyTypeGpt4,
		},
	}
	return &datalist.FilterType{
		Attr:       "select",
		SelectData: selectData,
	}
}

func (c *Shop) levelFilterType() (re *datalist.FilterType) {
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

func (c *Shop) expireTypeFilterType() (re *datalist.FilterType) {
	// 获取全部会员级别
	selectData := g.Slice{
		g.Map{
			"label": "不开通级别",
			"value": constant.ShopGoodsActiveExpireTypeNone,
		},
		g.Map{
			"label": "按天开通",
			"value": constant.ShopGoodsActiveExpireTypeDay,
		},
		g.Map{
			"label": "按月开通",
			"value": constant.ShopGoodsActiveExpireTypeMonth,
		},
		g.Map{
			"label": "按年开通",
			"value": constant.ShopGoodsActiveExpireTypeYear,
		},
	}
	return &datalist.FilterType{
		Attr:       "select",
		SelectData: selectData,
	}
}

func (c *Shop) orderStatusFilterType() (re *datalist.FilterType) {
	// 获取全部会员级别
	selectData := g.Slice{
		g.Map{
			"label": "待支付",
			"value": constant.ShopOrderStatusInit,
		},
		g.Map{
			"label": "已支付",
			"value": constant.ShopOrderStatusPaid,
		},
		g.Map{
			"label": "已完成",
			"value": constant.ShopOrderStatusFinish,
		},
		g.Map{
			"label": "已取消",
			"value": constant.ShopOrderStatusCancel,
		},
	}
	return &datalist.FilterType{
		Attr:       "select",
		SelectData: selectData,
	}
}
