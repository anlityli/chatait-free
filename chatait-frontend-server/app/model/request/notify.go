// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package request

type NotifyVmq struct {
	PayId       string `json:"payId"`
	Param       string `json:"param"`
	Type        int    `json:"type"`
	Price       string `json:"price"`
	ReallyPrice string `json:"reallyPrice"`
	Sign        string `json:"sign"`
}
