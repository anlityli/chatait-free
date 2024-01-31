// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package midjourney

import (
	"crypto/tls"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/constant"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/model/entity"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/snowflake"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/xtime"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtimer"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"github.com/gorilla/websocket"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type WsClient struct {
	Client          *ghttp.WebSocketClient
	Conn            *websocket.Conn
	Config          *entity.ConfigMidjourney
	UserId          int64
	SessionId       string
	heartbeatValue  int
	lastReceiveTime int64 // 最后一次收到消息时间
	onReadyChan     chan string
}

var WsClientMap sync.Map

func WsRun(config *entity.ConfigMidjourney, callback WsRunConnCallback) (err error) {
	_, ok := WsClientMap.Load(config.Id)
	if ok {
		glog.Line(true).Debug("已经有了ws对象直接返回")
		go callback()
		return
	}
	glog.Line(true).Debug("没有ws对象，开始创建")
	w := &WsClient{}
	w.Config = config
	w.Client = ghttp.NewWebSocketClient()
	w.Client.HandshakeTimeout = time.Second * 30 // 设置超时时间
	if config.Proxy != "" {
		glog.Line(true).Debug(config.Title + "WS使用代理" + config.Proxy)
		proxy, err := url.Parse(config.Proxy)
		if err != nil {
			glog.Line(true).Println(err)
			return err
		}
		w.Client.Proxy = http.ProxyURL(proxy)
	}
	w.Client.TLSClientConfig = &tls.Config{}
	WsClientMap.Store(config.Id, w)
	glog.Line(true).Debug("客户端创建完成")

	w.onReadyChan = make(chan string)
	for {
		err = w.connect()
		if err != nil {
			glog.Line(true).Println(err, "Retrying connection in 5 seconds...")
			time.Sleep(5 * time.Second)
		} else {
			break
		}
	}
	// 这里执行在连接之后要做的逻辑
	onReady := <-w.onReadyChan
	if onReady != "" {
		close(w.onReadyChan)
		go callback()
	}

	return nil
}

func (w *WsClient) connect() (err error) {
	conn, _, err := w.Client.Dial(WsUrl, nil)
	if err != nil {
		glog.Line(true).Println(err)
		return err
	}
	w.Conn = conn
	glog.Line(true).Debug("链接完成")
	err = w.auth()
	if err != nil {
		return err
	}
	glog.Line(true).Debug("授权完成")
	gtimer.SetInterval(time.Duration(40)*time.Second, func() {
		// 这里检测最后一次收到消息的时间和配置中闲置时长比较，如果闲置时间过长，则断开websocket不在监听，防止一直连接websocket有封号风险,当再次有任务时重新开启
		if xtime.GetNowTime()-gconv.Int64(w.Config.WsIdleTime) > w.lastReceiveTime {
			glog.Line(true).Debug("ws监听已闲置超过" + gconv.String(w.Config.WsIdleTime) + "秒，连接断开，删除连接对象Map元素")
			_ = w.Conn.Close()
			WsClientMap.Delete(w.Config.Id)
			gtimer.Exit()
			return
		}
		// 发送心跳
		err = w.heartBeat()
		if err != nil {
			glog.Line(true).Println("心跳发生错误", err)
			_ = w.Conn.Close()
			WsClientMap.Delete(w.Config.Id)
			gtimer.Exit()
		}
	})
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, message, err := w.Conn.ReadMessage()
			if err != nil {
				glog.Line(true).Println("read:", err)
				_ = w.Conn.Close()
				WsClientMap.Delete(w.Config.Id)
				return
			}
			w.parseMessage(string(message))
			//glog.Line(true).Debug("received:", string(message))
		}
	}()
	return nil
}

// heartBeat 心跳
func (w *WsClient) heartBeat() (err error) {
	//glog.Line(true).Debug("heartBeat start")
	w.heartbeatValue++
	err = w.Conn.WriteJSON(&WsMessage{
		Op: 1,
		D:  w.heartbeatValue,
	})
	if err != nil {
		glog.Line(true).Debug("heartBeat error", err)
		return err
	}
	//glog.Line(true).Debug("heartBeat end")
	return nil
}

func (w *WsClient) auth() (err error) {
	err = w.Conn.WriteJSON(&WsMessage{
		Op: 2,
		D: &WsMessageAuth{
			Token:        w.Config.UserToken,
			Capabilities: 8189,
			Properties: &WsMessageAuthProperties{
				Os:      "Mac OS X",
				Browser: "Chrome",
				Device:  "",
			},
			Compress: false,
		},
	})
	return nil
}

func (w *WsClient) parseMessage(data string) {
	dataJson, err := gjson.Decode(data)
	if err != nil {
		glog.Line(true).Debug("无法解析消息", data, err)
		return
	}
	dataObj := &WsReceiveMessage{}
	err = gconv.Scan(dataJson, dataObj)
	if err != nil {
		glog.Line(true).Debug("无法解析消息", data, err)
		return
	}
	switch dataObj.T {
	case "READY":
		w.onReadyChan <- dataObj.T
		w.lastReceiveTime = xtime.GetNowTime()
		message := &WsReceiveMessageDReady{}
		err = gconv.Scan(dataObj.D, message)
		if err != nil {
			glog.Line(true).Debug("无法解析消息", data, err)
			return
		}
		w.onReady(message)
	case "MESSAGE_CREATE":
		w.lastReceiveTime = xtime.GetNowTime()
		glog.Line(true).Debug("MESSAGE_CREATE", data)
		message := &WsReceiveMessageDCommon{}
		err = gconv.Scan(dataObj.D, message)
		if err != nil {
			glog.Line(true).Debug("无法解析消息", data, err)
			return
		}
		w.onMessageCreate(message)
	case "MESSAGE_UPDATE":
		w.lastReceiveTime = xtime.GetNowTime()
		glog.Line(true).Debug("MESSAGE_UPDATE", data)
		message := &WsReceiveMessageDCommon{}
		err = gconv.Scan(dataObj.D, message)
		if err != nil {
			glog.Line(true).Debug("无法解析消息", data, err)
			return
		}
		w.onMessageUpdate(message)
	case "MESSAGE_DELETE":
		w.lastReceiveTime = xtime.GetNowTime()
		glog.Line(true).Debug("MESSAGE_DELETE", data)
		message := &WsReceiveMessageDCommon{}
		err = gconv.Scan(dataObj.D, message)
		if err != nil {
			glog.Line(true).Debug("无法解析消息", data, err)
			return
		}
		w.onMessageDelete(message)
	case "INTERACTION_CREATE":
		w.lastReceiveTime = xtime.GetNowTime()
		glog.Line(true).Debug("INTERACTION_CREATE", data)
		message := &WsReceiveMessageDCommon{}
		err = gconv.Scan(dataObj.D, message)
		if err != nil {
			glog.Line(true).Debug("无法解析消息", data, err)
			return
		}
		w.onInteractionCreate(message)
	case "INTERACTION_SUCCESS":
		w.lastReceiveTime = xtime.GetNowTime()
		glog.Line(true).Debug("INTERACTION_SUCCESS", data)
		message := &WsReceiveMessageDCommon{}
		err = gconv.Scan(dataObj.D, message)
		if err != nil {
			glog.Line(true).Debug("无法解析消息", data, err)
			return
		}
		w.onInteractionSuccess(message)
	default:
		//glog.Line(true).Debug("其他消息", data)
	}
}

func (w *WsClient) onReady(message *WsReceiveMessageDReady) {
	w.UserId = gconv.Int64(message.User.Id)
	w.SessionId = message.SessionId
	glog.Line(true).Debug("ready", w.UserId, w.SessionId)
}

func (w *WsClient) onMessageCreate(message *WsReceiveMessageDCommon) {
	// 不是想要的消息直接过滤掉
	if !w.filterMessage(message) {
		return
	}
	if message.Nonce != "" && message.Nonce != "0" {
		// 拿到nonce 与队列匹配，写入content
		queueData := w.matchQueue(message)
		if queueData.Id > 0 {
			messageContent := msgContentHandler(message.Content)
			if messageContent != "" {
				queueData.MessageContent = messageContent
			}
			QueueInstance().changeTaskData(&changeTaskDataParams{
				eventType: constant.QueueMidjourneyEventWriteMessageContent,
				queueData: queueData,
				message:   message,
			})
		}
		if message.Embeds != nil && len(message.Embeds) > 0 {
			embed := message.Embeds[0]
			glog.Line(true).Debug("embed.Color", embed.Color)
			switch embed.Color {
			case 16711680:
				if embed.Title == "Action needed to continue" {
					// continue
					w.continueHandler(message, queueData)
					glog.Line(true).Debug("embed.continue", message)
					return
				} else if embed.Title == "Pending mod message" {
					// continue
					w.continueHandler(message, queueData)
					glog.Line(true).Debug("embed.continue", message)
					return
				}
				msgErr := embed.Description
				// createErrorEvent
				if queueData.Id > 0 {
					queueData.Status = constant.QueueMidjourneyStatusError
					queueData.ErrorAt = gconv.Int(xtime.GetNowTime())
					queueData.ErrorData = msgErr
					QueueInstance().changeTaskData(&changeTaskDataParams{
						eventType: constant.QueueMidjourneyEventError,
						queueData: queueData,
						message:   message,
					})
				}
				glog.Line(true).Debug("embed.Color error", msgErr)
			case 16776960:
				glog.Line(true).Debug("embed.Color warning", embed.Description)
			default:
				if gstr.Contains(embed.Title, "continue") && gstr.Contains(embed.Description, "verify you're human") {
					// verify human
					w.verifyHuman(message, queueData)
					glog.Line(true).Debug("verify you're human")
					return
				}
				if gstr.Contains(embed.Title, "Invalid") {
					msgErr := embed.Description
					// createErrorEvent
					if queueData.Id > 0 {
						queueData.Status = constant.QueueMidjourneyStatusError
						queueData.ErrorAt = gconv.Int(xtime.GetNowTime())
						queueData.ErrorData = msgErr
						QueueInstance().changeTaskData(&changeTaskDataParams{
							eventType: constant.QueueMidjourneyEventError,
							queueData: queueData,
							message:   message,
						})
					}
					glog.Line(true).Debug("Invalid", msgErr)
				}
			}
		}
	}

	if message.Nonce == "" && message.Attachments != nil && len(message.Attachments) > 0 && message.Components != nil {
		queueData := w.matchQueue(message)
		if queueData.Id > 0 {
			responseData, err := gjson.Encode(message)
			if err != nil {
				glog.Line(true).Println("响应数据写json错误", message, err)
				return
			}
			var referMessageId int64
			messageHash := ""
			if len(message.Attachments) > 0 {
				messageHash = getMessageHash(message.Attachments[0].Filename)
			}
			if message.ReferencedMessage != nil && message.ReferencedMessage.Id != "" {
				referMessageId = gconv.Int64(message.ReferencedMessage.Id)
			}
			queueData.MessageId = gconv.Int64(message.Id)
			queueData.ReferMessageId = referMessageId
			queueData.MessageHash = messageHash
			queueData.MessageType = message.Type
			queueData.Progress = 100
			queueData.Status = constant.QueueMidjourneyStatusEnded
			queueData.EndedAt = gconv.Int(xtime.GetNowTime())
			queueData.ResponseData = gconv.String(responseData)
			QueueInstance().changeTaskData(&changeTaskDataParams{
				eventType: constant.QueueMidjourneyEventEnded,
				queueData: queueData,
				message:   message,
			})
		}
		glog.Line(true).Debug("done", message.Attachments, message.Components)
		return
	}

	// messageUpdate
	w.onMessageUpdate(message)
}

func (w *WsClient) onMessageUpdate(message *WsReceiveMessageDCommon) {
	// 不是想要的消息直接过滤掉
	if !w.filterMessage(message) {
		return
	}
	if (message.Nonce == "" || message.Nonce == "0") && message.Interaction != nil && message.Interaction.Name != "" {
		switch message.Interaction.Name {
		//case "imagine":
		//	glog.Line(true).Debug("update Interaction imagine")
		//	// 更新进度
		//	progress := matchMsgContentProgress(message.Content)
		//	if progress > 0 {
		//		queueData := w.matchQueue(message)
		//		if queueData.Id > 0 {
		//			queueData.Progress = progress
		//			QueueInstance().changeTaskData(&changeTaskDataParams{
		//				eventType: constant.QueueMidjourneyEventProgress,
		//				queueData: queueData,
		//				message:   message,
		//			})
		//			return
		//		}
		//	}
		case "settings":
			// todo settingsEvent
			glog.Line(true).Debug("update settings")
			return
		case "describe":
			if message.Embeds != nil && len(message.Embeds) > 0 && message.Embeds[0].Image != nil {
				uri := message.Embeds[0].Image.Url
				glog.Line(true).Debug("update describe", uri)
			}
			// todo toMjEmit
		case "prefer remix":
			if message.Content != "" {
				//todo preferRemixEvent
				glog.Line(true).Debug("update prefer remix", message.Content)
			}
		case "shorten":
			glog.Line(true).Debug("update shorten", message.Embeds)
		case "info":
			// todo info
			glog.Line(true).Debug("update info", message.Embeds)
			return
		}
	}
	if message.Embeds != nil && len(message.Embeds) > 0 {
		embed := message.Embeds[0]
		if embed.Title == "Duplicate images detected" {
			msgErr := embed.Description
			// errorEvent
			queueData := w.matchQueue(message)
			if queueData.Id > 0 {
				queueData.Status = constant.QueueMidjourneyStatusError
				queueData.ErrorAt = gconv.Int(xtime.GetNowTime())
				queueData.ErrorData = msgErr
				QueueInstance().changeTaskData(&changeTaskDataParams{
					eventType: constant.QueueMidjourneyEventError,
					queueData: queueData,
					message:   message,
				})
			}
			glog.Line(true).Debug("update Duplicate images detected", msgErr)
		}
	}

	if message.Content != "" {
		// progressEvent
		progress := matchMsgContentProgress(message.Content)
		if progress > 0 {
			queueData := w.matchQueue(message)
			if queueData.Id > 0 {
				queueData.Progress = progress
				QueueInstance().changeTaskData(&changeTaskDataParams{
					eventType: constant.QueueMidjourneyEventProgress,
					queueData: queueData,
					message:   message,
				})
				return
			}
		}
		glog.Line(true).Debug("update progress", message)
	}
}

func (w *WsClient) onMessageDelete(message *WsReceiveMessageDCommon) {
	if message.ChannelId != w.Config.ChannelId {
		return
	}
}

func (w *WsClient) onInteractionCreate(message *WsReceiveMessageDCommon) {
	if message.Nonce != "" {
		// InteractionCreateEvent
		queueData := w.matchQueue(message)
		if queueData.Id > 0 {
			QueueInstance().changeTaskData(&changeTaskDataParams{
				eventType: constant.QueueMidjourneyEventInteractionCreate,
				queueData: queueData,
				message:   message,
			})
			return
		}
		glog.Line(true).Debug("InteractionCreate", message.Nonce, message.Id)
	}
}

func (w *WsClient) onInteractionSuccess(message *WsReceiveMessageDCommon) {
	if message.Nonce != "" && message.Id != "" {
		// 找到队列中的任务把interaction_id写入到队列数据中
		queueData := w.matchQueue(message)
		if queueData.Id > 0 {
			queueData.InteractionId = gconv.Int64(message.Id)
			QueueInstance().changeTaskData(&changeTaskDataParams{
				eventType: constant.QueueMidjourneyEventInteractionSuccess,
				queueData: queueData,
				message:   message,
			})
		}
		glog.Line(true).Debug("InteractionSuccess", message.Nonce, message.Id)
	}
}

func (w *WsClient) filterMessage(message *WsReceiveMessageDCommon) bool {
	if message.ChannelId != w.Config.ChannelId {
		return false
	}
	if message.Author != nil && message.Author.Id != MJApplicationId && message.Author.Id != NJApplicationId {
		return false
	}
	if message.Interaction != nil && gconv.Int64(message.Interaction.User.Id) != w.UserId {
		return false
	}
	return true
}

func (w *WsClient) continueHandler(message *WsReceiveMessageDCommon, queueData *entity.QueueMidjourney) {
	if message.Components != nil && len(message.Components) > 0 && message.Components[0].Components != nil && len(message.Components[0].Components) > 0 {
		appeal := message.Components[0].Components[0]
		glog.Line(true).Debug("continueHandler", appeal)
		if appeal != nil {
			// 修改当前队列信息为新的信息
			newNonce := snowflake.GenerateDiscordId()
			queueData.Nonce = newNonce
			QueueInstance().changeTaskData(&changeTaskDataParams{
				eventType: constant.QueueMidjourneyEventContinue,
				queueData: queueData,
				message:   message,
			})
			// 再次发起请求
			applicationId := MJApplicationId
			if queueData.ApplicationType == constant.QueueMidjourneyApplicationTypeNJ {
				applicationId = NJApplicationId
			}
			requestData := &ReqCustomIdDiscord{
				Type:          3,
				GuildId:       w.Config.GuildId,
				ChannelId:     w.Config.ChannelId,
				MessageFlags:  gconv.Int64(message.Flags),
				MessageId:     message.Id,
				ApplicationId: applicationId,
				SessionId:     w.Config.SessionId,
				Data: &CustomIdData{
					ComponentType: 2,
					CustomId:      appeal.CustomId,
				},
				Nonce: gconv.String(newNonce),
			}
			requestDataJson, err := gjson.Encode(requestData)
			if err != nil {
				glog.Line(true).Debug("continueHandler error", message, requestData, err)
				return
			}
			_, err = QueueInstance().request(w.Config, gconv.String(requestDataJson), ApiUrl+"interactions")
			if err != nil {
				glog.Line(true).Debug("continueHandler error", message, requestData, err)
				return
			}
		}
	}
}

func (w *WsClient) verifyHuman(message *WsReceiveMessageDCommon, queueData *entity.QueueMidjourney) {
	if w.Config.HuggingFaceToken == "" {
		return
	}
	imgUrl := message.Embeds[0].Image.Url
	if message.Components != nil && len(message.Components) > 0 {
		categories := message.Components[0].Components
		classify := make([]string, 0)
		for _, oneCategory := range categories {
			classify = append(classify, oneCategory.Label)
		}
		verifyClient, err := NewVerifyHuman(w.Config)
		if err != nil {
			glog.Line(true).Debug("verifyHuman new error", message, err)
			return
		}
		category, err := verifyClient.Verify(imgUrl, classify)
		if err != nil {
			glog.Line(true).Debug("verifyHuman Verify error", message, classify, err)
			return
		}
		if category != "" {
			customId := ""
			for _, oneCategory := range categories {
				if oneCategory.Label == category {
					customId = oneCategory.CustomId
				}
			}
			if customId == "" {
				glog.Line(true).Debug("verifyHuman customId error", message, category, categories, err)
				return
			}
			// 修改当前队列信息为新的信息
			newNonce := snowflake.GenerateDiscordId()
			queueData.Nonce = newNonce
			QueueInstance().changeTaskData(&changeTaskDataParams{
				eventType: constant.QueueMidjourneyEventVerifyHuman,
				queueData: queueData,
				message:   message,
			})
			// 再次发起请求
			applicationId := MJApplicationId
			if queueData.ApplicationType == constant.QueueMidjourneyApplicationTypeNJ {
				applicationId = NJApplicationId
			}
			requestData := &ReqCustomIdDiscord{
				Type:          3,
				GuildId:       w.Config.GuildId,
				ChannelId:     w.Config.ChannelId,
				MessageFlags:  gconv.Int64(message.Flags),
				MessageId:     message.Id,
				ApplicationId: applicationId,
				SessionId:     w.Config.SessionId,
				Data: &CustomIdData{
					ComponentType: 2,
					CustomId:      customId,
				},
				Nonce: gconv.String(newNonce),
			}
			// 构造请求数据，把请求内容写入到队列中
			requestDataJson, err := gjson.Encode(requestData)
			if err != nil {
				glog.Line(true).Debug("verifyHuman Verify error", message, requestData, err)
				return
			}
			_, err = QueueInstance().request(w.Config, gconv.String(requestDataJson), ApiUrl+"interactions")
			if err != nil {
				glog.Line(true).Debug("verifyHuman Verify error", message, requestData, err)
				return
			}
		}
	}

}

// matchQueue 匹配队列任务
func (w *WsClient) matchQueue(message *WsReceiveMessageDCommon) (queueData *entity.QueueMidjourney) {
	queueData = &entity.QueueMidjourney{}
	// 拿到工作池中所有正在执行的任务
	QueueInstance().workingPool.Range(func(key, value any) bool {
		task := value.(*QueueTask)
		if message.Nonce != "" && message.Nonce != "0" {
			if task.Data.Nonce == gconv.Int64(message.Nonce) {
				queueData = task.Data
				return false
			}
		} else if message.Interaction != nil && message.Interaction.Id != "" {
			if task.Data.InteractionId == gconv.Int64(message.Interaction.Id) {
				queueData = task.Data
				return false
			}
		} else if (message.ReferencedMessage != nil && message.ReferencedMessage.Id != "") || (message.MessageReference != nil && message.MessageReference.MessageID != "") {
			referMessageId := ""
			if message.ReferencedMessage != nil && message.ReferencedMessage.Id != "" {
				referMessageId = message.ReferencedMessage.Id
			} else if message.MessageReference != nil && message.MessageReference.MessageID != "" {
				referMessageId = message.MessageReference.MessageID
			} else {
				return true
			}
			// 根据提及ID匹配
			if isUpscaleString(message.Content) {
				referIndex := matchMsgContentIndex(message.Content)
				if task.Data.ConfigId == w.Config.Id && task.Data.RequestType == RequestTypeCustomId && task.Data.ReferMessageId == gconv.Int64(referMessageId) && task.Data.ActionType == constant.ActionTypeUpscale && task.Data.ReferIndex == referIndex {
					queueData = task.Data
					return false
				}
			} else if isVariationsString(message.Content) {
				if task.Data.ConfigId == w.Config.Id && task.Data.RequestType == RequestTypeCustomId && task.Data.ReferMessageId == gconv.Int64(referMessageId) && (task.Data.ActionType == constant.ActionTypeVariate || task.Data.ActionType == constant.ActionTypeVary) {
					queueData = task.Data
					return false
				}
			} else if isPanString(message.Content) {
				_, referIndex := matchMsgContentPan(message.Content)
				if task.Data.ConfigId == w.Config.Id && task.Data.RequestType == RequestTypeCustomId && task.Data.ReferMessageId == gconv.Int64(referMessageId) && task.Data.ActionType == constant.ActionTypePan && task.Data.ReferIndex == referIndex {
					queueData = task.Data
					return false
				}
			} else if isZoomOutString(message.Content) {
				if task.Data.ConfigId == w.Config.Id && task.Data.RequestType == RequestTypeCustomId && task.Data.ReferMessageId == gconv.Int64(referMessageId) && task.Data.ActionType == constant.ActionTypeZoomOut {
					queueData = task.Data
					return false
				}
			} else {
				queueData = task.Data
				glog.Line(true).Debug("其他情况匹配到的队列数据", queueData.Id)
				return false
			}
		}
		if message.Content != "" {
			matchContent := msgContentHandler(message.Content)
			if task.Data.ConfigId == w.Config.Id && task.Data.MessageContent == matchContent {
				queueData = task.Data
				return false
			}
		}
		return true
	})
	return queueData
}
