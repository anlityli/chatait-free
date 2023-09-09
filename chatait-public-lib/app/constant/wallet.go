// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package constant

// 钱包类型
const (
	WalletTypeBalance    = "balance"    // WalletTypeBalance 余额
	WalletTypeGpt3       = "gpt3"       // WalletTypeGpt3 GPT3提问次数
	WalletTypeGpt4       = "gpt4"       // WalletTypeGpt4 GPT4提问次数
	WalletTypeMidjourney = "midjourney" // WalletTypeMidjourney Midjourney提问次数
)

const (
	WalletChangeTargetTypeAddUser                = "addUser"        // WalletChangeTargetTypeAddUser 会员新增
	WalletChangeTargetTypeShopOrderGoods         = "shopOrderGoods" // WalletChangeTargetTypeShopOrderGoods 商商品购买
	WalletChangeTargetTypeConversationGpt3       = "gpt3"           // WalletChangeTargetTypeConversationGpt3 对话类型
	WalletChangeTargetTypeConversationGpt4       = "gpt4"           // WalletChangeTargetTypeConversationGpt4 对话类型
	WalletChangeTargetTypeConversationMidjourney = "midjourney"     // WalletChangeTargetTypeConversationMidjourney Midjourney类型
	WalletChangeTargetTypeBackend                = "backend"        // WalletChangeTargetTypeBackend 后台处理
	WalletChangeTargetTypeLevelMonth             = "levelMonth"     // WalletChangeTargetTypeLevelMonth 每月会员级别
)
