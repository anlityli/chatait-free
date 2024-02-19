// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/anlityli/chatait-free/chatait-frontend-server/app/model/request"
	"github.com/anlityli/chatait-free/chatait-frontend-server/app/model/response"
	"github.com/anlityli/chatait-free/chatait-frontend-server/library/auth"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/constant"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/libservice"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/model/entity"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/api/baidu"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/api/baidu/trans"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/api/midjourney"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/file"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/helper"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/notice"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/security"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/snowflake"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/xtime"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"regexp"
	"strings"
)

var ConversationMidjourney = &conversationMidjourneyService{}

type conversationMidjourneyService struct {
}

func (s *conversationMidjourneyService) Speak(r *ghttp.Request) (re *response.ConversationMidjourneySpeak, err error) {
	requestModel := &request.ConversationMidjourneySpeak{}
	if err = r.Parse(requestModel); err != nil {
		return nil, err
	}
	userId := auth.GetUserId(r)
	walletType := constant.WalletTypeMidjourney
	amount := 100
	// 如果用户次数不足直接报错
	walletData := libservice.Wallet.GetAllBalance(userId)
	if gconv.Int(walletData.Midjourney) < amount {
		midjourneyUseBalance, err := helper.GetConfig("midjourneyUseBalance")
		if err != nil {
			return nil, err
		}
		walletType = constant.WalletTypeBalance
		amount = gconv.Int(midjourneyUseBalance)
		if gconv.Int(walletData.Balance) < gconv.Int(midjourneyUseBalance) {
			notice.Write(r, notice.ShowDialog, &notice.HttpShowDialogMessage{
				Data:        "您的" + helper.GetWalletName(constant.WalletTypeMidjourney) + "或" + helper.GetWalletName(constant.WalletTypeBalance) + "不足请充值",
				ConfirmText: "去购买",
				ConfirmJump: "/purchase/goods-list",
			})
			return nil, nil
		}
	}
	// 如果达到每日最高上限，则提示报错
	midjourneyDailyLimit, err := helper.GetConfig("midjourneyDailyLimit")
	if err != nil {
		return nil, errors.New("获取系统配置参数失败")
	}
	todayStart := xtime.GetTodayBegin()
	todayEnd := xtime.GetTodayEnd()
	todayTimes, err := dao.Conversation.As("c").LeftJoin(dao.Topic.Table+" t", "c.topic_id=t.id").Where("t.type=? AND c.created_at>=? AND c.created_at<=?", constant.TopicTypeMidjourney, todayStart, todayEnd).Count()
	if todayTimes/2 >= gconv.Int(midjourneyDailyLimit) {
		return nil, errors.New("今日Midjourney生图次数已达上限")
	}
	prompt, err := s.promptHandler(requestModel)
	if err != nil {
		return nil, err
	}
	glog.Line(true).Debug(prompt)
	newTopicId := snowflake.GenerateID()
	re = &response.ConversationMidjourneySpeak{}
	re.TopicId = requestModel.TopicId
	re.TopicType = constant.TopicTypeMidjourney
	if gconv.Int64(re.TopicId) == 0 {
		re.TopicId = gconv.String(newTopicId)
		re.Title = "Midjourney生图: " + gstr.SubStrRune(requestModel.Content, 0, 50)
	}
	nowTime := xtime.GetNowTime()
	if err := g.DB().Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) (err error) {
		// 如果是新话题，话题入库
		if requestModel.TopicId == "" || requestModel.TopicId == "0" {
			if _, err = dao.Topic.Ctx(ctx).TX(tx).Data(g.Map{
				"id":         re.TopicId,
				"user_id":    userId,
				"title":      re.Title,
				"type":       constant.TopicTypeMidjourney,
				"created_at": nowTime,
			}).Insert(); err != nil {
				return err
			}
		}
		// 提问入库
		qConversationId := snowflake.GenerateID()
		re.QuestionContent = requestModel.Content
		if _, err = dao.Conversation.Ctx(ctx).TX(tx).Data(g.Map{
			"id":         qConversationId,
			"user_id":    userId,
			"topic_id":   re.TopicId,
			"role":       "user",
			"content":    re.QuestionContent,
			"created_at": nowTime,
		}).Insert(); err != nil {
			return err
		}
		re.QuestionId = gconv.String(qConversationId)
		// 回答入库
		aConversationId := snowflake.GenerateID()
		if _, err = dao.Conversation.Ctx(ctx).TX(tx).Data(g.Map{
			"id":         aConversationId,
			"user_id":    userId,
			"topic_id":   re.TopicId,
			"role":       "assistant",
			"created_at": nowTime,
		}).Insert(); err != nil {
			return err
		}
		re.AnswerId = gconv.String(aConversationId)
		// 扣除次数
		if err = libservice.Wallet.ChangeWalletBalance(ctx, tx, &libservice.ChangeWalletParam{
			UserId:     userId,
			WalletType: walletType,
			Amount:     -amount,
			Remark:     fmt.Sprintf("生成图片【%s】扣除", gstr.SubStrRune(re.QuestionContent, 0, 50)),
			TargetType: constant.WalletChangeTargetTypeConversationMidjourney,
			TargetID:   qConversationId,
		}); err != nil {
			glog.Line(true).Println("扣除提问次数失败", qConversationId, err)
			return err
		}
		// 请求生成图片的接口
		err = midjourney.GenerateImage(ctx, tx, &midjourney.GenerateImageParams{
			ConversationId:  gconv.Int64(re.AnswerId),
			ApplicationType: requestModel.ApplicationType,
			Prompt:          prompt,
		})

		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		glog.Line(true).Debug("生图失败", err)
		return nil, err
	}
	return re, nil
}

func (s *conversationMidjourneyService) Custom(r *ghttp.Request) (re *response.ConversationMidjourneySpeak, err error) {
	requestModel := &request.ConversationMidjourneyCustom{}
	if err = r.Parse(requestModel); err != nil {
		return nil, err
	}
	userId := auth.GetUserId(r)
	walletType := constant.WalletTypeMidjourney
	amount := 100
	// 如果用户次数不足直接报错
	walletData := libservice.Wallet.GetAllBalance(userId)
	if gconv.Int(walletData.Midjourney) < amount {
		midjourneyUseBalance, err := helper.GetConfig("midjourneyUseBalance")
		if err != nil {
			return nil, err
		}
		walletType = constant.WalletTypeBalance
		amount = gconv.Int(midjourneyUseBalance)
		if gconv.Int(walletData.Balance) < gconv.Int(midjourneyUseBalance) {
			notice.Write(r, notice.ShowDialog, &notice.HttpShowDialogMessage{
				Data:        "您的" + helper.GetWalletName(constant.WalletTypeMidjourney) + "或" + helper.GetWalletName(constant.WalletTypeBalance) + "不足请充值",
				ConfirmText: "去购买",
				ConfirmJump: "/purchase/goods-list",
			})
			return nil, nil
		}
	}
	// 如果达到每日最高上限，则提示报错
	midjourneyDailyLimit, err := helper.GetConfig("midjourneyDailyLimit")
	if err != nil {
		return nil, errors.New("获取系统配置参数失败")
	}
	todayStart := xtime.GetTodayBegin()
	todayEnd := xtime.GetTodayEnd()
	todayTimes, err := dao.Conversation.As("c").LeftJoin(dao.Topic.Table+" t", "c.topic_id=t.id").Where("t.type=? AND c.created_at>=? AND c.created_at<=?", constant.TopicTypeMidjourney, todayStart, todayEnd).Count()
	if todayTimes/2 >= gconv.Int(midjourneyDailyLimit) {
		return nil, errors.New("今日Midjourney生图次数已达上限")
	}
	// 查找提及的数据
	referData := &entity.Conversation{}
	err = dao.Conversation.Where("id=?", requestModel.ReferConversationId).Scan(referData)
	if err != nil {
		return nil, errors.New("获取提及对话失败")
	}
	referMidjourneyData := &entity.ConversationMidjourney{}
	err = dao.ConversationMidjourney.Where("conversation_id=?", requestModel.ReferConversationId).Scan(referMidjourneyData)
	if err != nil {
		return nil, errors.New("获取提及对话失败")
	}
	actionContent := ""
	switch requestModel.ActionType {
	case constant.ActionTypeUpscale:
		actionContent = "对" + requestModel.ReferConversationId + "的第" + gconv.String(requestModel.Index) + "张图片进行放大操作"
	case constant.ActionTypeVariate:
		actionContent = "对" + requestModel.ReferConversationId + "的第" + gconv.String(requestModel.Index) + "张图片进行变化操作"
	case constant.ActionTypeReRoll:
		actionContent = "对" + requestModel.ReferConversationId + "进行刷新操作"
	case constant.ActionTypeVary:
		varyStr := ""
		if requestModel.Index == 1 {
			varyStr = "较大变化"
		} else if requestModel.Index == 2 {
			varyStr = "细微变化"
		}
		actionContent = "对" + requestModel.ReferConversationId + "进行" + varyStr + "操作"
	case constant.ActionTypeZoomOut:
		zoomStr := ""
		if requestModel.Index == 1 {
			zoomStr = "放大2倍"
		} else if requestModel.Index == 2 {
			zoomStr = "放大1.5倍"
		}
		actionContent = "对" + requestModel.ReferConversationId + "进行" + zoomStr + "操作"
	case constant.ActionTypePan:
		panStr := ""
		if requestModel.Index == 1 {
			panStr = "向左延展"
		} else if requestModel.Index == 2 {
			panStr = "向右延展"
		} else if requestModel.Index == 3 {
			panStr = "向上延展"
		} else if requestModel.Index == 4 {
			panStr = "向下延展"
		}
		actionContent = "对" + requestModel.ReferConversationId + "进行" + panStr + "操作"
	}
	re = &response.ConversationMidjourneySpeak{}
	re.TopicId = gconv.String(referData.TopicId)
	re.TopicType = constant.TopicTypeMidjourney
	nowTime := xtime.GetNowTime()
	if err := g.DB().Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) (err error) {
		// 提问入库
		qConversationId := snowflake.GenerateID()
		re.QuestionContent = actionContent
		if _, err = dao.Conversation.Ctx(ctx).TX(tx).Data(g.Map{
			"id":         qConversationId,
			"user_id":    userId,
			"topic_id":   re.TopicId,
			"role":       "user",
			"content":    re.QuestionContent,
			"created_at": nowTime,
		}).Insert(); err != nil {
			return err
		}
		re.QuestionId = gconv.String(qConversationId)
		// 回答入库
		aConversationId := snowflake.GenerateID()
		if _, err = dao.Conversation.Ctx(ctx).TX(tx).Data(g.Map{
			"id":         aConversationId,
			"user_id":    userId,
			"topic_id":   re.TopicId,
			"role":       "assistant",
			"created_at": nowTime,
		}).Insert(); err != nil {
			return err
		}
		re.AnswerId = gconv.String(aConversationId)
		// 扣除次数
		if err = libservice.Wallet.ChangeWalletBalance(ctx, tx, &libservice.ChangeWalletParam{
			UserId:     userId,
			WalletType: walletType,
			Amount:     -amount,
			Remark:     fmt.Sprintf("生成图片【%s】扣除", gstr.SubStrRune(re.QuestionContent, 0, 50)),
			TargetType: constant.WalletChangeTargetTypeConversationMidjourney,
			TargetID:   qConversationId,
		}); err != nil {
			glog.Line(true).Println("扣除提问次数失败", qConversationId, err)
			return err
		}
		// 请求生成图片的接口
		//err = midjourney.CustomIdImage(ctx, tx, &midjourney.CustomIdImageParams{
		//	ActionType:          requestModel.ActionType,
		//	ConversationId:      gconv.Int64(re.AnswerId),
		//	ReferConversationId: gconv.Int64(requestModel.ReferConversationId),
		//	Index:               requestModel.Index,
		//	CustomId:            requestModel.CustomId,
		//})

		err = midjourney.ModalImage(ctx, tx, &midjourney.CustomIdImageParams{
			ActionType:          requestModel.ActionType,
			ConversationId:      gconv.Int64(re.AnswerId),
			ReferConversationId: gconv.Int64(requestModel.ReferConversationId),
			Index:               requestModel.Index,
			CustomId:            requestModel.CustomId,
		})

		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		glog.Line(true).Debug("生图失败", err)
		return nil, err
	}
	return re, nil
}

// Listener Bot的监听方法在项目启动时调用
func (s *conversationMidjourneyService) Listener() (err error) {
	// 开启队列消费方法
	midjourney.QueueInstance().Run()
	go func() {
		for {
			eventData := <-midjourney.QueueInstance().Event
			switch eventData.EventType {
			case constant.QueueMidjourneyEventInsertQueue:
				s.ListenerEventInsertQueue(eventData)
			//case constant.QueueMidjourneyEventInteractionCreate:
			//	s.ListenerEventInteractionCreate(eventData)
			case constant.QueueMidjourneyEventInteractionSuccess:
				s.ListenerEventInteractionSuccess(eventData)
			//case constant.QueueMidjourneyEventWriteMessageContent:
			//	s.ListenerEventWriteMessageContent(eventData)
			case constant.QueueMidjourneyEventEnded:
				s.ListenerEventEnd(eventData)
			case constant.QueueMidjourneyEventProgress:
				s.ListenerEventProgress(eventData)
			//case constant.QueueMidjourneyEventContinue:
			//	s.ListenerEventContinue(eventData)
			//case constant.QueueMidjourneyEventVerifyHuman:
			//	s.ListenerEventVerifyHuman(eventData)
			case constant.QueueMidjourneyEventError:
				s.ListenerEventError(eventData)

			}
		}
	}()
	// 循环所有midjourney的配置并开启监听
	configList := &[]*entity.ConfigMidjourney{}
	err = dao.ConfigMidjourney.Where("status=1").Scan(configList)
	if err != nil {
		return err
	}
	// 根据配置监听
	for _, config := range *configList {
		if config.ListenModel == constant.ConfigMidjourneyListenModelBot {
			go func(config *entity.ConfigMidjourney) {
				glog.Line().Println(config.Title + "开始启动监听")
				err = midjourney.BotRun(config)
				if err != nil {
					glog.Line(true).Println("Bot监听失败:", err)
				}
			}(config)
		}
		// 取消一开始就监听的逻辑防止长时间连接discord的ws封号，改到提交任务的时候再监听
		//} else if config.ListenModel == constant.ConfigMidjourneyListenModelUserWss {
		//	go func(config *entity.ConfigMidjourney) {
		//		err = midjourney.WsRun(config)
		//		if err != nil {
		//			glog.Line(true).Println("Ws监听失败:", err)
		//		}
		//	}(config)
		//}
	}
	return nil
}

// ListenerEventInsertQueue 进度队列
func (s *conversationMidjourneyService) ListenerEventInsertQueue(params *midjourney.QueueEvent) {
	glog.Line(true).Debug("进入队列")
	queueData := &entity.QueueMidjourney{}
	err := dao.QueueMidjourney.Where("id=?", params.QueueData.Id).Scan(queueData)
	if err != nil {
		glog.Line(true).Println("midjourney生成图片获取队列数据失败", params, err)
		return
	}
	// 获取队列对应的conversation数据
	conversationData := &entity.Conversation{}
	err = dao.Conversation.Where("id=?", queueData.ConversationId).Scan(conversationData)
	if err != nil {
		glog.Line(true).Println("midjourney生成图片获取对话数据失败", params, err)
		return
	}
	// 获取话题数据
	topicData := &entity.Topic{}
	err = dao.Topic.Where("id=?", conversationData.TopicId).Scan(topicData)
	if err != nil {
		glog.Line(true).Println("midjourney生成图片获取话题数据失败", params, err)
		return
	}
	noticeData := &response.WebsocketConversationMidjourneyListenerEvent{
		ConversationId: gconv.String(conversationData.Id),
		UserId:         gconv.String(conversationData.UserId),
		TopicId:        gconv.String(conversationData.TopicId),
		TopicTitle:     topicData.Title,
		TopicType:      topicData.Type,
		Role:           conversationData.Role,
		ActionType:     queueData.ActionType,
	}
	notice.WsSendJsonMessageToAll(&notice.WSMsg{
		Type: constant.WSMsgResponseTypeMidjourneyInsertQueue,
		Data: noticeData,
	})
}

// ListenerEventInteractionSuccess 交互命令发送成功
func (s *conversationMidjourneyService) ListenerEventInteractionSuccess(params *midjourney.QueueEvent) {
	glog.Line(true).Debug("创建开始")
	queueData := &entity.QueueMidjourney{}
	err := dao.QueueMidjourney.Where("id=?", params.QueueData.Id).Scan(queueData)
	if err != nil {
		glog.Line(true).Println("midjourney生成图片获取队列数据失败", params, err)
		return
	}
	// 获取队列对应的conversation数据
	conversationData := &entity.Conversation{}
	err = dao.Conversation.Where("id=?", queueData.ConversationId).Scan(conversationData)
	if err != nil {
		glog.Line(true).Println("midjourney生成图片获取对话数据失败", params, err)
		return
	}
	// 获取话题数据
	topicData := &entity.Topic{}
	err = dao.Topic.Where("id=?", conversationData.TopicId).Scan(topicData)
	if err != nil {
		glog.Line(true).Println("midjourney生成图片获取话题数据失败", params, err)
		return
	}
	noticeData := &response.WebsocketConversationMidjourneyListenerEvent{
		ConversationId: gconv.String(conversationData.Id),
		UserId:         gconv.String(conversationData.UserId),
		TopicId:        gconv.String(conversationData.TopicId),
		TopicTitle:     topicData.Title,
		TopicType:      topicData.Type,
		Role:           conversationData.Role,
		ActionType:     queueData.ActionType,
	}
	notice.WsSendJsonMessageToAll(&notice.WSMsg{
		Type: constant.WSMsgResponseTypeMidjourneyCreate,
		Data: noticeData,
	})
}

// ListenerEventEnd 生成完成
func (s *conversationMidjourneyService) ListenerEventEnd(params *midjourney.QueueEvent) {
	glog.Line(true).Debug("创建结束")
	queueData := &entity.QueueMidjourney{}
	err := dao.QueueMidjourney.Where("id=?", params.QueueData.Id).Scan(queueData)
	if err != nil {
		glog.Line(true).Println("midjourney生成图片获取队列数据失败", params, err)
		return
	}
	// 获取队列对应的conversation数据
	conversationData := &entity.Conversation{}
	err = dao.Conversation.Where("id=?", queueData.ConversationId).Scan(conversationData)
	if err != nil {
		glog.Line(true).Println("midjourney生成图片获取对话数据失败", params, err)
		return
	}
	// 获取话题数据
	topicData := &entity.Topic{}
	err = dao.Topic.Where("id=?", conversationData.TopicId).Scan(topicData)
	if err != nil {
		glog.Line(true).Println("midjourney生成图片获取话题数据失败", params, err)
		return
	}
	// 图片生成完成，拿到图片地址，把图片存入到本地，并写入数据库
	if params.Message == nil {
		glog.Line(true).Println("midjourney生成图片保存到本地失败", params.QueueData, err)
		return
	}
	// 如果设置中图片不保存到本地则不进行保存
	midjourneySaveImage, err := helper.GetConfig("midjourneySaveImage")
	if err != nil {
		glog.Line(true).Println("midjourney生成图片获取配置失败", params, err)
	}
	fileName := params.Message.Attachments[0].Filename
	relativePath := ""
	thumbnailRelativePath := ""
	oriFileName := params.Message.Attachments[0].Filename
	oriUrl := params.Message.Attachments[0].URL
	fileSize := gconv.Int64(params.Message.Attachments[0].Size)
	if midjourneySaveImage == "1" {
		saveData, err := file.RemoteFileSave(params.Message.Attachments[0].URL, params.Config.Proxy)
		if err == nil {
			fileName = saveData.FileName
			relativePath = saveData.RelativePath
			thumbnailRelativePath = saveData.ThumbnailRelativePath
			oriFileName = saveData.OriFileName
			oriUrl = saveData.OriUrl
			fileSize = saveData.FileSize
		} else {
			glog.Line(true).Println("midjourney生成图片保存到本地失败", params, err)
		}
	}

	fileId := snowflake.GenerateID()
	fileInsertData := g.Map{
		"id":           fileId,
		"file_name":    fileName,
		"user_id":      conversationData.UserId,
		"queue_id":     params.QueueData.Id,
		"path":         relativePath,
		"thumbnail":    thumbnailRelativePath,
		"prompt":       params.QueueData.MessageContent,
		"mj_file_name": oriFileName,
		"mj_url":       oriUrl,
		"width":        params.Message.Attachments[0].Width,
		"height":       params.Message.Attachments[0].Height,
		"size":         fileSize,
		"created_at":   xtime.GetNowTime(),
	}
	if _, err = dao.FileMidjourney.Data(fileInsertData).Insert(); err != nil {
		glog.Line(true).Println("midjourney生成图片写入数据失败", params, err)
		return
	}
	// 如果右附加组件，则把附加组件也写入到数据库中
	components := ""
	componentsSlice := make([]*midjourney.WsReceiveMessageDComponentsItem, 0)
	if params.Message.Components != nil {
		componentsSlice = params.Message.Components
		componentsJson, err := gjson.Encode(params.Message.Components)
		if err == nil {
			components = gconv.String(componentsJson)
		} else {
			glog.Line(true).Println("midjourney生成图片zh转换Components失败", params.Message, err)
		}
	}

	// 把midjourney的数据写入到conversation_midjourney表中
	if _, err = dao.ConversationMidjourney.Data(g.Map{
		"conversation_id": conversationData.Id,
		"action_type":     params.QueueData.ActionType,
		"file_id":         fileId,
		"components":      components,
	}).Insert(); err != nil {
		glog.Line(true).Println("midjourney生成图片写入数据失败", params, err)
		return
	}

	referencedConversationId := ""
	referencedComponentsSlice := make([]*midjourney.WsReceiveMessageDComponentsItem, 0)
	if params.Message.ReferencedMessage != nil && params.Message.ReferencedMessage.Components != nil {
		referencedComponentsSlice = params.Message.ReferencedMessage.Components
		referencedComponentsJson, err := gjson.Encode(params.Message.ReferencedMessage.Components)
		if err == nil {
			referencedComponents := gconv.String(referencedComponentsJson)
			if referencedComponents != "" {
				referencedQueueData := &entity.QueueMidjourney{}
				err = dao.QueueMidjourney.Where("message_id=?", params.Message.ReferencedMessage.Id).Scan(referencedQueueData)
				if err == nil {
					referencedConversationId = gconv.String(referencedQueueData.ConversationId)
					if _, err = dao.ConversationMidjourney.Data(g.Map{
						"components": referencedComponents,
					}).Where("conversation_id=?", referencedQueueData.ConversationId).Update(); err != nil {
						glog.Line(true).Println("midjourney生成图片写入提及数据失败", params, err)
					}
				}

			}
		} else {
			glog.Line(true).Println("midjourney生成图片zh转换Components失败", params.Message, err)
		}
	}

	// 修改对话数据中的内容 这里要改成能访问到的url地址
	hostUrl, err := security.HostUrl()
	if err != nil {
		glog.Println("域名授权失败，您可能正在使用盗版程序，请购买正版")
		return
	}
	imgUrl := hostUrl + "/file/midjourney-image?id=" + gconv.String(fileInsertData["id"])
	thumbnailImgUrl := hostUrl + "/file/midjourney-image?id=" + gconv.String(fileInsertData["id"]) + "&thumbnail=1"
	if _, err = dao.Conversation.Data(g.Map{
		"content": "![image](" + imgUrl + ")",
	}).Where("id=?", queueData.ConversationId).Update(); err != nil {
		glog.Line(true).Println("midjourney生成图片修改对话数据失败", params, err)
		return
	}
	// websocket 通知前端图片生成完毕
	noticeData := &response.WebsocketConversationMidjourneyListenerEvent{
		ConversationId:           gconv.String(conversationData.Id),
		UserId:                   gconv.String(conversationData.UserId),
		TopicId:                  gconv.String(conversationData.TopicId),
		TopicTitle:               topicData.Title,
		TopicType:                topicData.Type,
		Role:                     conversationData.Role,
		ActionType:               queueData.ActionType,
		Content:                  conversationData.Content,
		ImgUrl:                   imgUrl,
		ThumbnailImgUrl:          thumbnailImgUrl,
		Progress:                 100,
		Components:               componentsSlice,
		ReferencedConversationId: referencedConversationId,
		ReferencedComponents:     referencedComponentsSlice,
	}
	notice.WsSendJsonMessageToAll(&notice.WSMsg{
		Type: constant.WSMsgResponseTypeMidjourneyEnd,
		Data: noticeData,
	})
}

// ListenerEventError 生成错误
func (s *conversationMidjourneyService) ListenerEventError(params *midjourney.QueueEvent) {
	glog.Line(true).Debug("创建出错")
	queueData := &entity.QueueMidjourney{}
	err := dao.QueueMidjourney.Where("id=?", params.QueueData.Id).Scan(queueData)
	if err != nil {
		glog.Line(true).Println("midjourney生成图片获取队列数据失败", params, err)
		return
	}
	// 获取队列对应的conversation数据
	conversationData := &entity.Conversation{}
	err = dao.Conversation.Where("id=?", queueData.ConversationId).Scan(conversationData)
	if err != nil {
		glog.Line(true).Println("midjourney生成图片获取对话数据失败", params, err)
		return
	}
	// 获取话题数据
	topicData := &entity.Topic{}
	err = dao.Topic.Where("id=?", conversationData.TopicId).Scan(topicData)
	if err != nil {
		glog.Line(true).Println("midjourney生成图片获取话题数据失败", params, err)
		return
	}
	// 把midjourney的数据写入到conversation_midjourney表中
	if _, err = dao.ConversationMidjourney.Data(g.Map{
		"conversation_id": conversationData.Id,
		"action_type":     params.QueueData.ActionType,
		"error_data":      params.QueueData.ErrorData,
	}).Insert(); err != nil {
		glog.Line(true).Println("midjourney生成图片写入数据失败", params, err)
		return
	}
	noticeData := &response.WebsocketConversationMidjourneyListenerEvent{
		ConversationId: gconv.String(conversationData.Id),
		UserId:         gconv.String(conversationData.UserId),
		TopicId:        gconv.String(conversationData.TopicId),
		TopicTitle:     topicData.Title,
		TopicType:      topicData.Type,
		Role:           conversationData.Role,
		ActionType:     queueData.ActionType,
		Error:          params.QueueData.ErrorData,
	}
	notice.WsSendJsonMessageToAll(&notice.WSMsg{
		Type: constant.WSMsgResponseTypeMidjourneyError,
		Data: noticeData,
	})
}

// ListenerEventProgress 生成进度
func (s *conversationMidjourneyService) ListenerEventProgress(params *midjourney.QueueEvent) {
	glog.Line(true).Debug("创建进度")
	queueData := &entity.QueueMidjourney{}
	err := dao.QueueMidjourney.Where("id=?", params.QueueData.Id).Scan(queueData)
	if err != nil {
		glog.Line(true).Println("midjourney生成图片获取队列数据失败", params, err)
		return
	}
	// 获取队列对应的conversation数据
	conversationData := &entity.Conversation{}
	err = dao.Conversation.Where("id=?", queueData.ConversationId).Scan(conversationData)
	if err != nil {
		glog.Line(true).Println("midjourney生成图片获取对话数据失败", params, err)
		return
	}
	// 获取话题数据
	topicData := &entity.Topic{}
	err = dao.Topic.Where("id=?", conversationData.TopicId).Scan(topicData)
	if err != nil {
		glog.Line(true).Println("midjourney生成图片获取话题数据失败", params, err)
		return
	}
	noticeData := &response.WebsocketConversationMidjourneyListenerEvent{
		ConversationId: gconv.String(conversationData.Id),
		UserId:         gconv.String(conversationData.UserId),
		TopicId:        gconv.String(conversationData.TopicId),
		TopicTitle:     topicData.Title,
		TopicType:      topicData.Type,
		Role:           conversationData.Role,
		ActionType:     queueData.ActionType,
		Progress:       params.QueueData.Progress,
	}
	notice.WsSendJsonMessageToAll(&notice.WSMsg{
		Type: constant.WSMsgResponseTypeMidjourneyProgress,
		Data: noticeData,
	})
}

func (s *conversationMidjourneyService) promptHandler(requestModel *request.ConversationMidjourneySpeak) (output string, err error) {
	prompt := requestModel.Content
	// 过滤掉所有参数
	output = s.removeParams(prompt)
	if output == "" {
		return "", errors.New("无法解析提示词内容")
	}
	// 提示词中文翻译成英文
	transRe, err := trans.Text(&baidu.TransTextParams{
		From: "zh",
		To:   "en",
		Q:    output,
	})
	if err != nil {
		return "", err
	}
	if transRe != nil {
		output = ""
		if len(transRe.Result.TransResult) > 0 {
			for _, transData := range transRe.Result.TransResult {
				output += transData.Dst
			}
		}
		if output == "" {
			return "", errors.New("无法解析提示词内容")
		}
	}
	// 追加参数
	requestMap := gconv.MapStrStr(requestModel)
	delete(requestMap, "content")
	delete(requestMap, "application_type")
	delete(requestMap, "images")
	delete(requestMap, "iw")
	delete(requestMap, "topic_id")
	for key, item := range requestMap {
		if item != "" {
			if key == "no" {
				noOutput := gstr.SubStrRune(item, 5)
				noOutput = s.removeParams(noOutput)
				// 翻译no参数
				noTransRe, err := trans.Text(&baidu.TransTextParams{
					From: "zh",
					To:   "en",
					Q:    noOutput,
				})
				if err != nil {
					return "", err
				}
				if noTransRe != nil {
					noOutput = ""
					if len(noTransRe.Result.TransResult) > 0 {
						for _, noTransData := range noTransRe.Result.TransResult {
							noOutput += noTransData.Dst
						}
					}
				}
				output += " --no " + noOutput
			} else {
				output += " " + requestMap[key]
			}
		}
	}
	if requestModel.Images != "" {
		output = requestModel.Images + " " + output + " " + requestModel.Iw
	}
	return output, nil
}

func (s *conversationMidjourneyService) removeParams(input string) (output string) {
	// 使用正则表达式匹配 -- 参数及其内容
	re := regexp.MustCompile(`--\w+\s*\w*`)

	// 替换匹配到的参数及内容为空字符串
	output = re.ReplaceAllString(input, "")
	output = gstr.Replace(output, "，", ", ")

	// 移除多余的空格和逗号
	output = strings.TrimSpace(output)
	outputArr := gstr.Explode(",", output)
	output = ""
	for index, item := range outputArr {
		item = gstr.Trim(item)
		if item != "" {
			if index == 0 {
				output += gstr.Trim(item)
			} else {
				output += ", " + gstr.Trim(item)
			}
		}
	}
	// 过滤掉所有网址
	re1 := regexp.MustCompile(`http[s]?://[^\s]*`)
	output = re1.ReplaceAllString(output, "")
	return output
}
