// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package response

type FinanceWalletFLow struct {
	Id         string `json:"id"`
	UserId     string `json:"user_id"`
	Amount     int    `json:"amount"`
	Total      int    `json:"total"`
	IsIncr     int    `json:"is_incr"`
	TargetType int    `json:"target_type"`
	TargetId   int    `json:"target_id"`
	Remark     string `json:"remark"`
	Year       int    `json:"year"`
	Month      int    `json:"month"`
	Day        int    `json:"day"`
	CreatedAt  int    `json:"created_at"`
	UpdatedAt  int    `json:"updated_at"`
}

type FinanceWalletFLowList []*FinanceWalletFLow

type FinanceWalletInfo struct {
	UserId     string `json:"user_id"`
	Balance    int    `json:"balance"`
	Gpt3       int    `json:"gpt3"`
	Gpt4       int    `json:"gpt4"`
	Midjourney int    `json:"midjourney"`
}
