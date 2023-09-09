// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package service

import (
	"errors"
	"github.com/anlityli/chatait-free/chatait-frontend-server/app/model/request"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/constant"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/notice"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var Websocket = &websocketService{}

type websocketService struct {
}

func (s *websocketService) Index(r *ghttp.Request) (err error) {
	wsHandle, err := r.WebSocket()
	if err != nil {
		return err
	}
	defer func() {
		notice.WsCloseMessageListener(wsHandle)
	}()

	// 发一个忽略的消息，以便开启监听通道的方法
	notice.WsSendJsonMessage(wsHandle, &notice.WSMsg{Type: constant.WSMsgResponseTypeIgnore, Data: nil})

	// 读取消息
FOR:
	for {
		_, data, err := wsHandle.ReadMessage()
		if err != nil {
			//glog.Line().Printf("disconnect:" + err.Error())
			break
		}
		msg := &notice.WSMsg{}
		if err := gjson.DecodeTo(data, msg); err != nil {
			dataStr := gconv.String(data)
			if dataStr != "ping" {
				return errors.New("消息内容解析失败:" + err.Error())
			}
			_ = wsHandle.WriteMessage(ghttp.WS_MSG_TEXT, []byte("pong"))
			continue
		}

		// 验签
		if !notice.WsValidateSign(msg) {
			return errors.New("验签失败")
		}

		// 根据消息类型分配不同的业务逻辑方法
		switch msg.Type {
		case constant.WSMsgRequestTypeExample:
			if err := s.requestExample(); err != nil {
				notice.WsSendJsonMessage(wsHandle, &notice.WSMsg{Type: constant.WSMsgResponseTypeError, Data: err.Error()})
				break FOR
			}
		}

	}
	return
}

// requestExample 一个请求的使用例子
func (s *websocketService) requestExample() (err error) {
	return
}

// NoticeToAll 开放一个接口，用于其他服务向前端发送websocket消息
func (s *websocketService) NoticeToAll(r *ghttp.Request) (err error) {
	reqModel := &request.WebsocketNoticeToAll{}
	if err := r.Parse(reqModel); err != nil {
		return err
	}
	notice.WsSendJsonMessageToAll(&notice.WSMsg{Type: reqModel.Type, Data: reqModel.Data})
	return nil
}
