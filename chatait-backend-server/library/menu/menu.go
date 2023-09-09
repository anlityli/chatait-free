// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package menu

import (
	"github.com/anlityli/chatait-free/chatait-public-lib/app/constant"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/helper"
	"sync"
)

// Menu 菜单
type Menu struct {
	listModel ListModel
}

var menuObj *Menu
var menuOnce sync.Once

func Instance() *Menu {
	menuOnce.Do(func() {
		menuObj = &Menu{}
		menuObj.list()
	})
	return menuObj
}

func (m *Menu) GetList() (re ListModel) {
	return m.listModel
}

// GetAllKeyData 获取所有子路由数据用map返回
func (m *Menu) GetAllKeyData() (re map[string]*ItemModel) {
	re = make(map[string]*ItemModel)
	for _, item := range m.listModel {
		for _, childItem := range item.Children {
			re[childItem.Key] = childItem
		}
	}
	return re
}

// GetRelevantRoutePathByKey 通过key获取关联的路由
func (m *Menu) GetRelevantRoutePathByKey(key string) (re []string) {
	re = make([]string, 0)
	allData := m.GetAllKeyData()
	item, ok := allData[key]
	if !ok {
		return re
	}
	return item.RelevantRoutePath
}

// list 一级主菜单
func (m *Menu) list() {
	list := ListModel{
		&ItemModel{Key: "dashboard", Title: "控制台", RoutePath: "dashboard", Show: true, Leaf: true, Children: m.dashboardChildren()},
		&ItemModel{Key: "user", Title: "会员管理", RoutePath: "user", Show: true, Children: m.userChildren()},
		&ItemModel{Key: "conversation", Title: "对话管理", RoutePath: "conversation", Show: true, Children: m.conversationChildren()},
		&ItemModel{Key: "shop", Title: "商城管理", RoutePath: "shop", Show: true, Children: m.shopChildren()},
		&ItemModel{Key: "finance", Title: "财务管理", RoutePath: "finance", Show: true, Children: m.financeChildren()},
		&ItemModel{Key: "admin", Title: "管理员管理", RoutePath: "admin", Show: true, Children: m.adminChildren()},
		&ItemModel{Key: "config", Title: "系统设置", RoutePath: "config", Show: true, Children: m.configChildren()},
	}
	menuObj.listModel = list
}

// dashboardChild 控制台
func (m *Menu) dashboardChildren() (re ListModel) {
	return ListModel{
		&ItemModel{Key: "dashboardIndex", Title: "控制台", RoutePath: "dashboard/index", Show: false, Leaf: true},
		&ItemModel{Key: "dashboardUserStatistic", Title: "用户数量统计", RoutePath: "dashboard/user-statistic", Show: false, Leaf: true},
		&ItemModel{Key: "dashboardAmountStatistic", Title: "收入统计", RoutePath: "dashboard/amount-statistic", Show: false, Leaf: true},
		&ItemModel{Key: "dashboardOrderStatistic", Title: "订单统计", RoutePath: "dashboard/order-statistic", Show: false, Leaf: true},
		&ItemModel{Key: "dashboardConversationStatistic", Title: "提问次数统计", RoutePath: "dashboard/conversation-statistic", Show: false, Leaf: true},
	}
}

func (m *Menu) userChildren() (re ListModel) {
	return ListModel{
		&ItemModel{Key: "userList", Title: "会员列表", RoutePath: "user/list", Show: true, Leaf: true},
		&ItemModel{Key: "userBan", Title: "禁用会员", RoutePath: "user/ban", Show: false, Leaf: true},
		&ItemModel{Key: "userChangeLevel", Title: "修改会员级别", RoutePath: "user/change-level", Show: false, Leaf: true},
		&ItemModel{Key: "userResetPassword", Title: "重置会员密码", RoutePath: "user/reset-password", Show: false, Leaf: true},
	}
}

func (m *Menu) conversationChildren() (re ListModel) {
	return ListModel{
		&ItemModel{Key: "conversationTopicList", Title: "话题列表", RoutePath: "conversation/topic-list", Show: true, Leaf: true},
		&ItemModel{Key: "conversationList", Title: "对话列表", RoutePath: "conversation/list", Show: true, Leaf: true},
	}
}

func (m *Menu) shopChildren() (re ListModel) {
	return ListModel{
		&ItemModel{Key: "shopGoodsList", Title: "商品列表", RoutePath: "shop/goods-list", Show: true, Leaf: true},
		&ItemModel{Key: "shopGoodsAdd", Title: "商品添加", RoutePath: "shop/goods-add", Show: false, Leaf: true},
		&ItemModel{Key: "shopGoodsEdit", Title: "商品编辑", RoutePath: "shop/goods-edit", RelevantRoutePath: []string{"shop/goods-one"}, Show: false, Leaf: true},
		&ItemModel{Key: "shopGoodsSort", Title: "商品排序", RoutePath: "shop/goods-sort", Show: false, Leaf: true},
		&ItemModel{Key: "shopGoodsDelete", Title: "商品删除", RoutePath: "shop/goods-delete", Show: false, Leaf: true},
		&ItemModel{Key: "shopOrderList", Title: "订单列表", RoutePath: "shop/order-list", Show: true, Leaf: true},
		&ItemModel{Key: "shopOrderStatus", Title: "修改订单状态", RoutePath: "shop/order-status", Show: false, Leaf: true},
	}
}

func (m *Menu) financeChildren() (re ListModel) {
	return ListModel{
		&ItemModel{Key: "financeWalletList", Title: "钱包列表", RoutePath: "finance/wallet-list", Show: true, Leaf: true},
		&ItemModel{Key: "financeWalletChange", Title: "钱包列表", RoutePath: "finance/wallet-change", Show: false, Leaf: true},
		&ItemModel{Key: "financeWalletFlowListBalance", Title: helper.GetWalletName(constant.WalletTypeBalance) + "流水", RoutePath: "finance/wallet-flow-list-balance", Show: true, Leaf: true},
		&ItemModel{Key: "financeWalletFlowListGpt3", Title: helper.GetWalletName(constant.WalletTypeGpt3) + "流水", RoutePath: "finance/wallet-flow-list-gpt3", Show: true, Leaf: true},
		&ItemModel{Key: "financeWalletFlowListGpt4", Title: helper.GetWalletName(constant.WalletTypeGpt4) + "流水", RoutePath: "finance/wallet-flow-list-gpt4", Show: true, Leaf: true},
		&ItemModel{Key: "financeWalletFlowListMidjourney", Title: helper.GetWalletName(constant.WalletTypeMidjourney) + "流水", RoutePath: "finance/wallet-flow-list-midjourney", Show: true, Leaf: true},
	}
}

// adminChild 管理员
func (m *Menu) adminChildren() (re ListModel) {
	return ListModel{
		&ItemModel{Key: "adminList", Title: "管理员列表", RoutePath: "admin/list", Show: true, Leaf: true},
		&ItemModel{Key: "adminAdd", Title: "添加管理员", RoutePath: "admin/add", Show: false, Leaf: true},
		&ItemModel{Key: "adminEdit", Title: "编辑管理员", RoutePath: "admin/edit", RelevantRoutePath: []string{"user/selector-list", "admin/info", "admin/one", "admin/all-role"}, Show: false, Leaf: true},
		&ItemModel{Key: "adminDelete", Title: "编辑管理员", RoutePath: "admin/delete", Show: false, Leaf: true},
		&ItemModel{Key: "adminRoleList", Title: "角色列表", RoutePath: "admin/role-list", Show: true, Leaf: true},
		&ItemModel{Key: "adminRoleAdd", Title: "添加角色", RoutePath: "admin/role-add", Show: false, Leaf: true},
		&ItemModel{Key: "adminRoleEdit", Title: "添加角色", RoutePath: "admin/role-edit", RelevantRoutePath: []string{"admin/role-one"}, Show: false, Leaf: true},
		&ItemModel{Key: "adminResetPassword", Title: "重置密码", RoutePath: "admin/reset-password", Show: false, Leaf: true},
		&ItemModel{Key: "adminResetOtherPassword", Title: "重置密码", RoutePath: "admin/reset-other-password", Show: false, Leaf: true},
	}
}

func (m *Menu) configChildren() (re ListModel) {
	return ListModel{
		&ItemModel{Key: "configAllOption", Title: "系统选项", RoutePath: "config/all-option", Show: true, Leaf: true},
		&ItemModel{Key: "configOptionEdit", RelevantRoutePath: []string{"config/all-option"}, Title: "编辑系统选项", RoutePath: "config/option-edit", Show: false, Leaf: true},
		&ItemModel{Key: "configWalletList", Title: "钱包配置列表", RoutePath: "config/wallet-list", Show: true, Leaf: true},
		&ItemModel{Key: "configWalletEdit", RelevantRoutePath: []string{"config/wallet-one"}, Title: "钱包编辑", RoutePath: "config/wallet-edit", Show: false, Leaf: true},
		&ItemModel{Key: "configPayList", Title: "支付方式列表", RoutePath: "config/pay-list", Show: true, Leaf: true},
		&ItemModel{Key: "configPayEdit", RelevantRoutePath: []string{"config/pay-one"}, Title: "编辑支付方式", RoutePath: "config/pay-edit", Show: false, Leaf: true},
		&ItemModel{Key: "configLevelList", Title: "级别列表", RoutePath: "config/level-list", Show: true, Leaf: true},
		&ItemModel{Key: "configLevelEdit", Title: "级别编辑", RoutePath: "config/level-edit", Show: false, Leaf: true},
		&ItemModel{Key: "configMidjourneyList", Title: "Midjourney配置列表", RoutePath: "config/midjourney-list", Show: true, Leaf: true},
		&ItemModel{Key: "configMidjourneyAdd", Title: "Midjourney配置添加", RoutePath: "config/midjourney-add", Show: false, Leaf: false},
		&ItemModel{Key: "configMidjourneyEdit", Title: "Midjourney配置编辑", RoutePath: "config/midjourney-edit", RelevantRoutePath: []string{"config/midjourney-one"}, Show: false, Leaf: false},
		&ItemModel{Key: "configMidjourneyDelete", Title: "Midjourney配置删除", RoutePath: "config/midjourney-delete", Show: false, Leaf: false},
		&ItemModel{Key: "configOpenaiList", Title: "Openai配置列表", RoutePath: "config/openai-list", Show: true, Leaf: true},
		&ItemModel{Key: "configOpenaiAdd", Title: "Openai配置添加", RoutePath: "config/openai-add", Show: false, Leaf: false},
		&ItemModel{Key: "configOpenaiEdit", Title: "Openai配置编辑", RoutePath: "config/openai-edit", RelevantRoutePath: []string{"config/openai-one"}, Show: false, Leaf: false},
		&ItemModel{Key: "configOpenaiDelete", Title: "Openai配置删除", RoutePath: "config/openai-delete", Show: false, Leaf: false},
	}
}
