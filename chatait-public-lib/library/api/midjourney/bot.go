// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package midjourney

import (
	"github.com/anlityli/chatait-free/chatait-public-lib/app/constant"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/model/entity"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/xtime"
	"github.com/bwmarrin/discordgo"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"net/http"
	"net/url"
	"sync"
)

type BotClient struct {
	Session *discordgo.Session
	Config  *entity.ConfigMidjourney
}

var BotClientMap sync.Map

func BotRun(config *entity.ConfigMidjourney) (err error) {
	_, ok := BotClientMap.Load(config.Id)
	if ok {
		return
	}
	b := &BotClient{}
	b.Config = config
	b.Session, err = discordgo.New("Bot " + config.BotToken)
	if err != nil {
		glog.Line(true).Debug(config.Title+"发生错误", err.Error())
		return err
	}
	BotClientMap.Store(config.Id, b)
	glog.Line(true).Debug("Bot客户端创建完成")
	if config.Proxy != "" {
		glog.Line(true).Debug(config.Title + "Bot使用代理" + config.Proxy)
		proxy, err := url.Parse(config.Proxy)
		if err != nil {
			glog.Line(true).Debug(err)
			return err
		}
		b.Session.Client.Transport = &http.Transport{Proxy: http.ProxyURL(proxy)}
		b.Session.Dialer.Proxy = http.ProxyURL(proxy)
	}
	err = b.Session.Open()
	if err != nil {
		glog.Line(true).Debug(config.Title+"发生错误", err)
		_ = b.Session.Close()
		BotClientMap.Delete(config.Id)
		return err
	}
	glog.Line().Debug(config.Title, "botMsgCreate")
	b.Session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		b.msgCreate(s, m)
	})
	glog.Line().Debug(config.Title, "botMsgUpdate")
	b.Session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageUpdate) {
		b.msgUpdate(s, m)
	})
	return nil
}

func (b *BotClient) msgCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// 过滤频道
	if m.ChannelID != b.Config.ChannelId {
		return
	}
	// 过滤掉自己发送的消息
	if m.Author.ID == s.State.User.ID {
		return
	}

	// 消息内容
	glog.Line(true).Debug(m)
	glog.Line(true).Debug(m.Content)
	glog.Line(true).Debug(m.Attachments)

	if gstr.Contains(m.Content, "(Waiting to start)") && !gstr.Contains(m.Content, "Rerolling **") {
		// 首次触发生成任务
		glog.Line(true).Debug("首次触发生成任务")
		queueData := b.matchQueue(m)
		if queueData.Id > 0 {
			queueData.MessageContent = msgContentHandler(m.Content)
			QueueInstance().changeTaskData(&changeTaskDataParams{
				eventType: constant.QueueMidjourneyEventInteractionSuccess,
				queueData: queueData,
				message:   m,
			})
		}
		return
	}
	for _, attachment := range m.Attachments {
		if attachment.Width > 0 && attachment.Height > 0 {
			queueData := b.matchQueue(m)
			if queueData.Id > 0 {
				responseData, err := gjson.Encode(m)
				if err != nil {
					glog.Line(true).Println("响应数据写json错误", m, err)
					return
				}
				var referMessageId int64
				messageHash := ""
				if len(m.Attachments) > 0 {
					messageHash = getMessageHash(m.Attachments[0].Filename)
				}
				if m.ReferencedMessage != nil && m.ReferencedMessage.ID != "" {
					referMessageId = gconv.Int64(m.ReferencedMessage.ID)
				}
				queueData.MessageId = gconv.Int64(m.ID)
				queueData.ReferMessageId = referMessageId
				queueData.MessageHash = messageHash
				queueData.MessageType = gconv.Int(m.Type)
				queueData.Progress = 100
				queueData.Status = constant.QueueMidjourneyStatusEnded
				queueData.EndedAt = gconv.Int(xtime.GetNowTime())
				queueData.ResponseData = gconv.String(responseData)
				QueueInstance().changeTaskData(&changeTaskDataParams{
					eventType: constant.QueueMidjourneyEventEnded,
					queueData: queueData,
					message:   m,
				})
			}
			return
		}
	}
}

func (b *BotClient) msgUpdate(s *discordgo.Session, m *discordgo.MessageUpdate) {
	// 过滤频道
	if m.ChannelID != b.Config.ChannelId {
		return
	}

	if m.Author == nil {
		return
	}

	// 过滤掉自己发送的消息
	if m.Author.ID == s.State.User.ID {
		return
	}
	// 消息内容
	glog.Line(true).Debug(m.ID)
	glog.Line(true).Debug(m.Content)

	if gstr.Contains(m.Content, "(Stopped)") {
		// 任务停止出错
		queueData := b.matchQueue(m)
		if queueData.Id > 0 {
			queueData.Status = constant.QueueMidjourneyStatusError
			queueData.ErrorAt = gconv.Int(xtime.GetNowTime())
			queueData.ErrorData = "任务停止出错，生成过程发生错误"
			QueueInstance().changeTaskData(&changeTaskDataParams{
				eventType: constant.QueueMidjourneyEventError,
				queueData: queueData,
				message:   m,
			})
		}
		return
	}

	progress := matchMsgContentProgress(m.Content)
	if progress > 0 {
		queueData := b.matchQueue(m)
		if queueData.Id > 0 {
			queueData.Progress = progress
			QueueInstance().changeTaskData(&changeTaskDataParams{
				eventType: constant.QueueMidjourneyEventProgress,
				queueData: queueData,
				message:   m,
			})
		}
	}

	if len(m.Embeds) > 0 {
		embed := m.Embeds[0]
		if embed.Title == "Duplicate images detected" {
			msgErr := embed.Description
			// errorEvent
			queueData := b.matchQueue(m)
			if queueData.Id > 0 {
				queueData.Status = constant.QueueMidjourneyStatusError
				queueData.ErrorAt = gconv.Int(xtime.GetNowTime())
				queueData.ErrorData = msgErr
				QueueInstance().changeTaskData(&changeTaskDataParams{
					eventType: constant.QueueMidjourneyEventError,
					queueData: queueData,
					message:   m,
				})
			}
		}
		// Embeds
		glog.Line(true).Debug("Embeds", m.Embeds)
		return
	}
}

// matchQueue 匹配队列任务
func (b *BotClient) matchQueue(m interface{}) (queueData *entity.QueueMidjourney) {
	queueData = &entity.QueueMidjourney{}
	// 无论是create还是update都转换成create格式
	messageMap := gconv.Map(m)
	message := &discordgo.MessageCreate{}
	err := gconv.Scan(messageMap, message)
	if err != nil {
		glog.Line(true).Println(" 消息转换失败", m)
		return
	}
	// 拿到工作池中所有正在执行的任务
	QueueInstance().workingPool.Range(func(key, value any) bool {
		task := value.(*QueueTask)
		if message.Interaction != nil && message.Interaction.ID != "" {
			if task.Data.InteractionId == gconv.Int64(message.Interaction.ID) {
				queueData = task.Data
				return false
			}
		} else if (message.ReferencedMessage != nil && message.ReferencedMessage.ID != "") || (message.MessageReference != nil && message.MessageReference.MessageID != "") {
			referMessageId := ""
			if message.ReferencedMessage != nil && message.ReferencedMessage.ID != "" {
				referMessageId = message.ReferencedMessage.ID
			} else if message.MessageReference != nil && message.MessageReference.MessageID != "" {
				referMessageId = message.MessageReference.MessageID
			} else {
				return true
			}
			// 根据提及ID匹配
			if isUpscaleString(message.Content) {
				referIndex := matchMsgContentIndex(message.Content)
				if task.Data.ConfigId == b.Config.Id && task.Data.RequestType == RequestTypeCustomId && task.Data.ReferMessageId == gconv.Int64(referMessageId) && task.Data.ActionType == constant.ActionTypeUpscale && task.Data.ReferIndex == referIndex {
					queueData = task.Data
					return false
				}
			} else if isVariationsString(message.Content) {
				if task.Data.ConfigId == b.Config.Id && task.Data.RequestType == RequestTypeCustomId && task.Data.ReferMessageId == gconv.Int64(referMessageId) && (task.Data.ActionType == constant.ActionTypeVariate || task.Data.ActionType == constant.ActionTypeVary) {
					queueData = task.Data
					return false
				}
			} else if isPanString(message.Content) {
				_, referIndex := matchMsgContentPan(message.Content)
				if task.Data.ConfigId == b.Config.Id && task.Data.RequestType == RequestTypeCustomId && task.Data.ReferMessageId == gconv.Int64(referMessageId) && task.Data.ActionType == constant.ActionTypePan && task.Data.ReferIndex == referIndex {
					queueData = task.Data
					return false
				}
			} else if isZoomOutString(message.Content) {
				if task.Data.ConfigId == b.Config.Id && task.Data.RequestType == RequestTypeCustomId && task.Data.ReferMessageId == gconv.Int64(referMessageId) && task.Data.ActionType == constant.ActionTypeZoomOut {
					queueData = task.Data
					return false
				}
			} else {
				queueData = task.Data
				glog.Line(true).Debug("其他情况匹配到的队列数据", queueData.Id)
				return false
			}
		} else if message.Content != "" {
			matchContent := msgContentHandler(message.Content)
			if task.Data.ConfigId == b.Config.Id && task.Data.MessageContent == matchContent {
				queueData = task.Data
				return false
			}
		}
		return true
	})
	return queueData
}
