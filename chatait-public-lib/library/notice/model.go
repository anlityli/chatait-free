// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package notice

import "github.com/gogf/gf/net/ghttp"

type ErrorCode int

type HttpModel struct {
	Error   ErrorCode   `json:"error"`
	Message interface{} `json:"message"`
}

type HttpShowDialogMessage struct {
	Data        string `json:"data"`
	ConfirmText string `json:"confirm_text"`
	ConfirmJump string `json:"confirm_jump"`
}

type WSMsgType string

type WSMsg struct {
	Type WSMsgType   `json:"type"`
	Time string      `json:"time"`
	Sign string      `json:"sign"`
	Data interface{} `json:"data"`
}

type wsHandleMapItem struct {
	wsHandle *ghttp.WebSocket
	msgChan  chan *wsMsgChanModel
}

type wsMsgChanModel struct {
	msgType    int
	msgContent []byte
}
