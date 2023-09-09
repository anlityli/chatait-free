// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package request

type FinanceWalletChange struct {
	UserId     string `json:"user_id" v:"required"`
	WalletType string `json:"wallet_type" v:"required"`
	Amount     int    `json:"amount" v:"required"`
	Remark     string `json:"remark"`
}
