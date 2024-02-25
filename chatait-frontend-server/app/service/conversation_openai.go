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
	"github.com/anlityli/chatait-free/chatait-public-lib/library/api/openai"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/helper"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/notice"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/snowflake"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/xtime"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"net/http"
)

var ConversationOpenai = &conversationOpenaiService{}

type conversationOpenaiService struct {
}

// Speak 对话
func (s *conversationOpenaiService) Speak(r *ghttp.Request) (re *response.ConversationSpeak, err error) {
	requestModel := &request.ConversationSpeak{}
	if err = r.Parse(requestModel); err != nil {
		return nil, err
	}
	userId := auth.GetUserId(r)
	walletType := constant.WalletTypeGpt3
	amount := 100
	model := openai.ModelGPT35Turbo
	// 敏感词过滤
	wordsValidateRe, err := helper.SensitiveWordsValidate(&helper.SensitiveWordsValidateParams{
		UserId:       userId,
		ValidateType: constant.ConfigSensitiveWordValidateTypeConversation,
		TopicType:    requestModel.TopicType,
		Content:      requestModel.Content,
	})
	if err != nil {
		return nil, err
	}
	if !wordsValidateRe {
		return nil, errors.New("您提交的内容存在不合规内容，请检查后重新提交")
	}
	// 如果用户次数不足直接报错
	walletData := libservice.Wallet.GetAllBalance(userId)
	if requestModel.TopicType == constant.TopicTypeOpenaiGPT3 {
		if gconv.Int(walletData.Gpt3) < amount {
			walletType = constant.WalletTypeBalance
			gpt3UseBalance, err := helper.GetConfig("gpt3UseBalance")
			if err != nil {
				return nil, err
			}
			amount = gconv.Int(gpt3UseBalance)
			if gconv.Int(walletData.Balance) < gconv.Int(gpt3UseBalance) {
				notice.Write(r, notice.ShowDialog, &notice.HttpShowDialogMessage{
					Data:        "您的" + helper.GetWalletName(constant.WalletTypeGpt3) + "次数或" + helper.GetWalletName(constant.WalletTypeBalance) + "不足请充值",
					ConfirmText: "去购买",
					ConfirmJump: "/purchase/goods-list",
				})
				return nil, nil
			}
		}
	} else if requestModel.TopicType == constant.TopicTypeOpenaiGPT4 {
		walletType = constant.WalletTypeGpt4
		model = openai.ModelGPT4
		if gconv.Int(walletData.Gpt4) < amount {
			walletType = constant.WalletTypeBalance
			gpt4UseBalance, err := helper.GetConfig("gpt4UseBalance")
			if err != nil {
				return nil, err
			}
			amount = gconv.Int(gpt4UseBalance)
			if gconv.Int(walletData.Balance) < gconv.Int(gpt4UseBalance) {
				notice.Write(r, notice.ShowDialog, &notice.HttpShowDialogMessage{
					Data:        "您的" + helper.GetWalletName(constant.WalletTypeGpt4) + "次数或" + helper.GetWalletName(constant.WalletTypeBalance) + "不足请充值",
					ConfirmText: "去购买",
					ConfirmJump: "/purchase/goods-list",
				})
				return nil, nil
			}
		}
	} else {
		return nil, errors.New("话题类型不正确")
	}
	streamItem := s.getFromStreamMap(requestModel.StreamUuid)
	if streamItem == nil {
		return nil, errors.New("stream_uuid 不存在")
	}
	newTopicId := snowflake.GenerateID()
	re = &response.ConversationSpeak{}
	re.TopicId = requestModel.TopicId
	re.TopicType = requestModel.TopicType
	if gconv.Int64(re.TopicId) == 0 {
		re.TopicId = gconv.String(newTopicId)
		re.Title = gstr.SubStrRune(requestModel.Content, 0, 50)
	}
	go func() {
		err = s.speakLogic(streamItem, r, requestModel, newTopicId, walletType, amount, model)
		if err != nil {
			glog.Line(true).Debug("对话逻辑发生错误", err)
			_ = s.writeResponse(streamItem.wPointer, notice.OtherError, err.Error())
			_ = s.writeResponseClose(streamItem.wPointer, err.Error())
			streamItem.doneChan <- 2
		}
	}()

	return re, nil
}

func (s *conversationOpenaiService) SpeakStream(r *ghttp.Request) (err error) {
	requestModel := &request.ConversationSpeakStream{}
	if err = r.Parse(requestModel); err != nil {
		return err
	}
	// 利用go原始库来响应结果，因为用gf的w响应结果不会实时以流的形式响应
	w := r.Response.Writer.RawWriter()
	_, ok := w.(http.Flusher)
	if !ok {
		return errors.New("server not support")
	}
	// 规定响应头为流模式
	r.Response.Header().Set("Content-Type", "text/event-stream")
	r.Response.Header().Set("Connection", "keep-alive")
	r.Response.Header().Set("Cache-Control", "no-cache")
	// 下面两句防止业务结束以后，gf底层还会再调一下WriteHeader函数报错
	r.Response.Writer.WriteHeader(http.StatusOK) // 写入状态码为200
	r.Response.Writer.Flush()                    // 清一下缓冲区
	// 把UUID为key写入到流map中，以供其他业务逻辑调用
	streamItem := &StreamMapItem{
		wPointer: &w,
		doneChan: make(chan int),
	}
	s.setToStreamMap(requestModel.StreamUuid, streamItem)

	responseObj := &response.ConversationSpeakItem{}
	_ = s.writeResponse(streamItem.wPointer, notice.NoError, responseObj)

	//go func() {
	//	for i := 0; i < 2; i++ {
	//		time.Sleep(1 * time.Second)
	//		_ = s.writeResponse(streamItem.wPointer, notice.NoError, i)
	//	}
	//	//_ = s.writeResponseClose(streamItem.wPointer)
	//	//streamItem.doneChan <- 1
	//}()

	isDone := <-streamItem.doneChan
	if isDone > 0 {
		glog.Line().Debug("拿到结束", isDone)
		s.delFromStreamMap(requestModel.StreamUuid)
		return nil
	}

	return nil
}

type StreamMapItem struct {
	wPointer *http.ResponseWriter
	doneChan chan int
}

// setToStreamMap 把w存入map中
func (s *conversationOpenaiService) setToStreamMap(streamUuid string, item *StreamMapItem) {
	streamMap.Store(streamUuid, item)
}

// getFromStreamMap 从map中拿到w
func (s *conversationOpenaiService) getFromStreamMap(streamUuid string) *StreamMapItem {
	item, ok := streamMap.Load(streamUuid)
	if ok {
		return item.(*StreamMapItem)
	}
	return nil
}

// delFromStreamMap  从map中删掉w
func (s *conversationOpenaiService) delFromStreamMap(streamUuid string) {
	streamMap.Delete(streamUuid)
}

// speakLogic 对话逻辑
func (s *conversationOpenaiService) speakLogic(streamItem *StreamMapItem, r *ghttp.Request, requestModel *request.ConversationSpeak, newTopicId int64, walletType string, amount int, model string) (err error) {
	nowTime := xtime.GetNowTime()
	userId := auth.GetUserId(r)
	topicId := gconv.Int64(requestModel.TopicId)
	if topicId == 0 && newTopicId == 0 {
		return errors.New("未知话题")
	}

	// 构造请求内容
	openaiMessages := make(openai.RequestChatParamsMessages, 0)
	// 获取系统身份的内容
	systemContent, err := helper.GetConfig("gptSystemContent")
	if err != nil {
		return err
	}
	// 系统每次都参与对话
	if systemContent != "" {
		openaiMessages = append(openaiMessages, &openai.RequestChatParamsMessageItem{
			Role:    "system",
			Content: systemContent,
		})
	}
	// 拿到会员的前几次聊天内容
	conversationList := &[]*entity.Conversation{}
	err = dao.Conversation.Where("user_id=? AND topic_id=?", userId, topicId).Order("id DESC").Limit(6).Scan(conversationList)
	if err != nil {
		glog.Line(true).Println("查找前几次聊天内容失败", err)
		return err
	}

	if len(*conversationList) > 0 {
		for _, conversationItem := range *conversationList {
			openaiMessages = append(openai.RequestChatParamsMessages{
				&openai.RequestChatParamsMessageItem{
					Role:    conversationItem.Role,
					Content: conversationItem.Content,
				},
			}, openaiMessages...)
		}
	}
	openaiMessages = append(openaiMessages, &openai.RequestChatParamsMessageItem{
		Role:    "user",
		Content: requestModel.Content,
	})
	speakInsertId := snowflake.GenerateID()
	speakInsertData := g.Map{
		"id":         speakInsertId,
		"user_id":    userId,
		"topic_id":   topicId,
		"role":       "user",
		"content":    requestModel.Content,
		"created_at": nowTime,
	}
	responseInsertId := snowflake.GenerateID()
	responseInsertData := g.Map{
		"id":         responseInsertId,
		"user_id":    userId,
		"topic_id":   topicId,
		"role":       "",
		"content":    "",
		"created_at": 0,
	}
	err = openai.ChatCompletion(&openai.ChatCompletionParams{
		Model:    model,
		Messages: openaiMessages,
	}, func(originContent string, contentObj *openai.ResponseChat) error {
		//if text != "" {
		//	_, err = fmt.Fprintf(w, text)
		//	if err == nil {
		//		w.Flush()
		//	}
		//}
		if contentObj != nil {
			if len(contentObj.Choices) > 0 && contentObj.Choices[0].Delta.Role != "" {
				responseInsertData["role"] = contentObj.Choices[0].Delta.Role
			}
			if gconv.Int64(responseInsertData["created_at"]) == 0 {
				responseInsertData["created_at"] = xtime.GetNowTime()
			}
			responseInsertData["content"] = gconv.String(responseInsertData["content"]) + contentObj.Choices[0].Delta.Content

			responseObj := &response.ConversationSpeakItem{}
			err = gconv.Scan(responseInsertData, responseObj)
			if err != nil {
				glog.Line(true).Println("转换响应结构失败", responseInsertData, err)
				return err
			}
			responseObj.Content = contentObj.Choices[0].Delta.Content
			responseObj.CreatedAt = gconv.Int(xtime.GetNowTime())
			_ = s.writeResponse(streamItem.wPointer, notice.NoError, responseObj)
		}
		return nil
	})
	if err != nil {
		glog.Line(true).Println("聊天接口逻辑失败", err)
		return err
	}
	_ = s.writeResponseClose(streamItem.wPointer, "")
	streamItem.doneChan <- 1
	if err := g.DB().Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) (err error) {
		// 新建话题
		if topicId == 0 {
			topicId = newTopicId
			topicInsertData := g.Map{
				"id":         topicId,
				"user_id":    userId,
				"title":      gstr.SubStrRune(requestModel.Content, 0, 50),
				"type":       requestModel.TopicType,
				"created_at": nowTime,
			}
			if _, err = dao.Topic.Ctx(ctx).TX(tx).Data(topicInsertData).Insert(); err != nil {
				glog.Line(true).Println("话题入库失败", topicInsertData, err)
				return err
			}
		}
		// 把对话内容入库
		speakInsertData["topic_id"] = topicId
		if _, err = dao.Conversation.Ctx(ctx).TX(tx).Data(speakInsertData).Insert(); err != nil {
			glog.Line(true).Println("对话内容入库失败", speakInsertData, err)
			return err
		}
		responseInsertData["topic_id"] = topicId
		if _, err = dao.Conversation.Ctx(ctx).TX(tx).Data(responseInsertData).Insert(); err != nil {
			glog.Line(true).Println("对话响应入库失败", responseInsertData, err)
			return err
		}
		// 扣除token次数
		targetType := constant.WalletChangeTargetTypeConversationGpt3
		if requestModel.TopicType == constant.TopicTypeOpenaiGPT4 {
			targetType = constant.WalletChangeTargetTypeConversationGpt4
		}
		if err = libservice.Wallet.ChangeWalletBalance(ctx, tx, &libservice.ChangeWalletParam{
			UserId:     userId,
			WalletType: walletType,
			Amount:     -amount,
			Remark:     fmt.Sprintf("提交问题【%s】扣除", gstr.SubStrRune(requestModel.Content, 0, 50)),
			TargetType: targetType,
			TargetID:   speakInsertId,
		}); err != nil {
			glog.Line(true).Println("扣除提问次数失败", speakInsertData, err)
			return err
		}
		return nil
	}); err != nil {
		glog.Line(true).Debug("入库失败", err)
		return err
	}
	return nil
}

// writeResponse 写入响应
func (s *conversationOpenaiService) writeResponse(wPointer *http.ResponseWriter, errorCode notice.ErrorCode, data interface{}) (err error) {
	noticeData := &notice.HttpModel{}
	noticeData.Error = errorCode
	noticeData.Message = data
	wData, err := gjson.Encode(noticeData)
	if err != nil {
		return err
	}
	w := *wPointer
	flusher, ok := w.(http.Flusher)
	if !ok {
		return errors.New("server not support")
	}
	//glog.Line().Debug("data: ", wData)
	_, err = fmt.Fprintf(w, "data: %s\n\n", wData)
	if err != nil {
		return err
	}
	flusher.Flush()
	return nil
}

// writeResponseClose 写入结束
func (s *conversationOpenaiService) writeResponseClose(wPointer *http.ResponseWriter, errorMessage string) (err error) {
	w := *wPointer
	flusher, ok := w.(http.Flusher)
	if !ok {
		return errors.New("server not support")
	}
	data := "data: close"
	if errorMessage != "" {
		data = "data: " + errorMessage
	}
	_, err = fmt.Fprintf(w, "event: close\n%s\n\n", data)
	if err != nil {
		return err
	}
	flusher.Flush()
	return nil
}

// StreamUuid 生成UUID
func (s *conversationOpenaiService) StreamUuid() (re *response.ConversationStreamUuid) {
	return &response.ConversationStreamUuid{
		Uuid: helper.GenerateUuid(),
	}
}
