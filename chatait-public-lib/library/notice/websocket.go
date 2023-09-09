// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package notice

import (
	"github.com/anlityli/chatait-free/chatait-public-lib/library/security"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/serv/http"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/xtime"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"sync"
)

var wsHandleMap sync.Map

func WsSendJsonMessageToAll(msgContent *WSMsg) {
	// 循环全部map内的wsHandle并发送消息
	wsHandleMap.Range(func(key, value interface{}) bool {
		wsHandle := key.(*ghttp.WebSocket)
		WsSendJsonMessage(wsHandle, msgContent)
		return true
	})
}

// WsSendJsonMessage 发送json消息
func WsSendJsonMessage(wsHandle *ghttp.WebSocket, msgContent *WSMsg) {
	// 从map中拿到对应的通道
	wsHandleMapItemObjOri := &wsHandleMapItem{wsHandle: wsHandle}
	wsHandleMapItemInterface, ok := wsHandleMap.LoadOrStore(wsHandle, wsHandleMapItemObjOri)
	if !ok {
		wsMessageListener(wsHandleMapItemObjOri)
	}
	wsHandleMapItemObj := wsHandleMapItemInterface.(*wsHandleMapItem)
	msgContent.Time = gconv.String(xtime.GetNowTime())
	// 生成sign
	dataMap := gconv.Map(msgContent)
	site := http.Instance().GetSite()
	msgContent.Sign = security.GenerateParamsSign(dataMap, g.Config().GetString(site+"Conf.privateKey"))
	msgJson, err := gjson.Encode(msgContent)
	if err == nil {
		msg := &wsMsgChanModel{}
		msg.msgType = ghttp.WS_MSG_TEXT
		msg.msgContent = msgJson
		wsHandleMapItemObj.msgChan <- msg
	}
}

// WsSendBinaryMessage 发送二进制消息
func WsSendBinaryMessage(wsHandle *ghttp.WebSocket, msgContent []byte) {
	// 从map中拿到对应的通道
	wsHandleMapItemObjOri := &wsHandleMapItem{wsHandle: wsHandle}
	wsHandleMapItemInterface, ok := wsHandleMap.LoadOrStore(wsHandle, wsHandleMapItemObjOri)
	if !ok {
		wsMessageListener(wsHandleMapItemObjOri)
	}
	wsHandleMapItemObj := wsHandleMapItemInterface.(*wsHandleMapItem)

	msg := &wsMsgChanModel{}
	msg.msgType = ghttp.WS_MSG_BINARY
	msg.msgContent = msgContent
	wsHandleMapItemObj.msgChan <- msg
}

// WsCloseMessageListener 停止一个wsHandle的消息监听
func WsCloseMessageListener(wsHandle *ghttp.WebSocket) {
	wsHandle.Close()
	wsHandleMapItemInterface, ok := wsHandleMap.Load(wsHandle)
	if ok {
		wsHandleMapItemObj := wsHandleMapItemInterface.(*wsHandleMapItem)
		msg := &wsMsgChanModel{}
		msg.msgType = ghttp.WS_MSG_CLOSE
		// 发送信号停止监听
		wsHandleMapItemObj.msgChan <- msg
		// 删除wsHandle键
		wsHandleMap.Delete(wsHandle)
	}
}

// wsMessageListener 监听wsHandle对应的通道内的消息
func wsMessageListener(wsHandleMapItemObj *wsHandleMapItem) {
	wsHandleMapItemObj.msgChan = make(chan *wsMsgChanModel)
	go func() {
		defer func() {
			close(wsHandleMapItemObj.msgChan)
		}()
		for {
			messageObj := <-wsHandleMapItemObj.msgChan
			if messageObj.msgType == ghttp.WS_MSG_CLOSE {
				break
			}
			wsHandle := wsHandleMapItemObj.wsHandle
			_ = wsHandle.WriteMessage(messageObj.msgType, messageObj.msgContent)
		}
	}()
}

// WsValidateSign 验签
func WsValidateSign(msg *WSMsg) bool {
	dataMap := g.Map{
		"time": msg.Time,
		"sign": msg.Sign,
		"type": msg.Type,
	}
	site := http.Instance().GetSite()
	return security.ValidateParamsSign(dataMap, g.Config().GetString(site+"Conf.privateKey"))
}
