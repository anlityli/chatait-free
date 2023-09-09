// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package response

type FinanceWalletItem struct {
	UserId     string `json:"user_id"`
	Username   string `json:"username"`
	Nickname   string `json:"nickname"`
	Balance    int    `json:"balance"`
	Gpt3       int    `json:"gpt3"`
	Gpt4       int    `json:"gpt4"`
	Midjourney int    `json:"midjourney"`
}

type FinanceWalletList []*FinanceWalletItem

type FinanceWalletFlowItem struct {
	Id         string `json:"id"`
	UserId     string `json:"user_id"`
	Username   string `json:"username"`
	Nickname   string `json:"nickname"`
	Amount     int    `json:"amount"`
	Total      int    `json:"total"`
	IsIncr     int    `json:"is_incr"`
	TargetType string `json:"target_type"`
	TargetId   string `json:"target_id"`
	Remark     string `json:"remark"`
	AdminName  string `json:"admin_name"`
	CreatedAt  int    `json:"created_at"`
}

type FinanceWalletFlowList []*FinanceWalletFlowItem
