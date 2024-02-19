// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package midjourney

import (
	"context"
	"database/sql"
	"errors"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/constant"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/model/entity"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/snowflake"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/xtime"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"strings"
)

// GenerateImage 生成图片
func GenerateImage(ctx context.Context, tx *gdb.TX, params *GenerateImageParams) (err error) {
	glog.Line(true).Debug("开始执行生图")
	config, err := Instance().GetConfig()
	if err != nil {
		glog.Line(true).Debug(err)
		return err
	}
	applicationId := MJApplicationId
	versionId := MJVersionId
	commandGenerateImageId := MJCommandGenerateImageId
	commandGenerateImageDescription := MJCommandGenerateImageDescription
	if params.ApplicationType == constant.QueueMidjourneyApplicationTypeNJ {
		applicationId = NJApplicationId
		versionId = NJVersionId
		commandGenerateImageId = NJCommandGenerateImageId
		commandGenerateImageDescription = NJCommandGenerateImageDescription
	}
	prompt := trimPrompt(config, params.Prompt)
	nonce := snowflake.GenerateDiscordId()
	requestData := &ReqTriggerDiscord{
		Type:          RequestTypeGenerate,
		GuildId:       config.GuildId,
		ChannelId:     config.ChannelId,
		ApplicationId: applicationId,
		SessionId:     config.SessionId,
		Data: &DSCommand{
			Version: versionId,
			Id:      commandGenerateImageId,
			Name:    "imagine",
			Type:    1,
			Options: []*DSOption{
				{
					Type:  3,
					Name:  "prompt",
					Value: prompt,
				},
			},
			ApplicationCommand: &DSApplicationCommand{
				Id:                       commandGenerateImageId,
				ApplicationId:            applicationId,
				Version:                  versionId,
				DefaultPermission:        true,
				DefaultMemberPermissions: nil,
				Type:                     1,
				Nsfw:                     false,
				Name:                     "imagine",
				Description:              commandGenerateImageDescription,
				DmPermission:             true,
				Options:                  []*DSCommandOption{{Type: 3, Name: "prompt", Description: "The prompt to imagine", Required: true}},
			},
			Attachments: []*ReqCommandAttachments{},
		},
		Nonce: gconv.String(nonce),
	}
	// 构造请求数据，把请求内容写入到队列中
	requestDataJson, err := gjson.Encode(requestData)
	if err != nil {
		glog.Line(true).Debug(err)
		return err
	}
	id := snowflake.GenerateID()
	queueData := &entity.QueueMidjourney{
		Id:              id,
		ConversationId:  params.ConversationId,
		ConfigId:        config.Id,
		ActionType:      constant.ActionTypeGenerate,
		ApplicationType: params.ApplicationType,
		Nonce:           nonce,
		MessageType:     MessageTypeGenerate,
		MessageContent:  prompt,
		RequestType:     RequestTypeGenerate,
		RequestUrl:      ApiUrl + "interactions",
		RequestData:     gconv.String(requestDataJson),
		Status:          constant.QueueMidjourneyStatusInit,
		CreatedAt:       gconv.Int(xtime.GetNowTime()),
	}
	err = QueueInstance().InsertTask(queueData)
	if err != nil {
		glog.Line(true).Debug(err)
		return err
	}
	// 调用接口，接口调用次数增加
	if _, err = dao.ConfigMidjourney.Ctx(ctx).TX(tx).Data(g.Map{
		"call_num": gdb.Raw("call_num+1"),
	}).Where("id=?", config.Id).Update(); err != nil {
		glog.Line(true).Debug(err)
		return err
	}
	return nil
}

// CustomIdImage 组件处理图片
func CustomIdImage(ctx context.Context, tx *gdb.TX, params *CustomIdImageParams) (err error) {
	config, err := Instance().GetConfig()
	if err != nil {
		glog.Line(true).Debug(err)
		return err
	}
	referQueueData := &entity.QueueMidjourney{}
	err = dao.QueueMidjourney.Ctx(ctx).TX(tx).Where("conversation_id=?", params.ReferConversationId).Scan(referQueueData)
	if err != nil && err != sql.ErrNoRows {
		glog.Line(true).Debug(err)
		return err
	}
	if referQueueData.Id <= 0 {
		return errors.New("对话相应的队列信息不存在")
	}
	applicationId := MJApplicationId
	if referQueueData.ApplicationType == constant.QueueMidjourneyApplicationTypeNJ {
		applicationId = NJApplicationId
	}
	nonce := snowflake.GenerateDiscordId()
	requestData := &ReqCustomIdDiscord{
		Type:          RequestTypeCustomId,
		GuildId:       config.GuildId,
		ChannelId:     config.ChannelId,
		MessageFlags:  0,
		MessageId:     gconv.String(referQueueData.MessageId),
		ApplicationId: applicationId,
		SessionId:     config.SessionId,
		Data: &CustomIdData{
			ComponentType: 2,
			CustomId:      params.CustomId,
		},
		Nonce: gconv.String(nonce),
	}
	// 构造请求数据，把请求内容写入到队列中
	requestDataJson, err := gjson.Encode(requestData)
	if err != nil {
		glog.Line(true).Debug(err)
		return err
	}
	id := snowflake.GenerateID()
	queueData := &entity.QueueMidjourney{
		Id:              id,
		ConversationId:  params.ConversationId,
		ConfigId:        config.Id,
		ActionType:      params.ActionType,
		ApplicationType: referQueueData.ApplicationType,
		Nonce:           nonce,
		ReferMessageId:  referQueueData.MessageId,
		ReferIndex:      params.Index,
		MessageType:     MessageTypeCustomId,
		RequestType:     RequestTypeCustomId,
		RequestUrl:      ApiUrl + "interactions",
		RequestData:     gconv.String(requestDataJson),
		Status:          constant.QueueMidjourneyStatusInit,
		CreatedAt:       gconv.Int(xtime.GetNowTime()),
	}
	err = QueueInstance().InsertTask(queueData)
	if err != nil {
		glog.Line(true).Debug(err)
		return err
	}
	// 调用接口，接口调用次数增加
	if _, err = dao.ConfigMidjourney.Ctx(ctx).TX(tx).Data(g.Map{
		"call_num": gdb.Raw("call_num+1"),
	}).Where("id=?", config.Id).Update(); err != nil {
		glog.Line(true).Debug(err)
		return err
	}
	return nil
}

func ModalImage(ctx context.Context, tx *gdb.TX, params *CustomIdImageParams) (err error) {
	config, err := Instance().GetConfig()
	if err != nil {
		glog.Line(true).Debug(err)
		return err
	}
	referQueueData := &entity.QueueMidjourney{}
	err = dao.QueueMidjourney.Ctx(ctx).TX(tx).Where("conversation_id=?", params.ReferConversationId).Scan(referQueueData)
	if err != nil && err != sql.ErrNoRows {
		glog.Line(true).Debug(err)
		return err
	}
	if referQueueData.Id <= 0 {
		return errors.New("对话相应的队列信息不存在")
	}
	applicationId := MJApplicationId
	if referQueueData.ApplicationType == constant.QueueMidjourneyApplicationTypeNJ {
		applicationId = NJApplicationId
	}
	nonce := snowflake.GenerateDiscordId()
	//dataId := snowflake.GenerateDiscordId()
	requestData := &ReqModalDiscord{
		Type:          RequestTypeModal,
		ApplicationId: applicationId,
		ChannelId:     config.ChannelId,
		GuildId:       config.GuildId,
		Data: &ModalData{
			Id:       "1209150918260432896",
			CustomId: "MJ::RemixModal::70a89e80-20d8-4e7e-b24f-ae83b0a7d663::1::1",
			Components: []*ModalDataComponentsItem{
				{
					Type: 1,
					Components: []*ModalDataComponentsItemComponentsItem{
						{
							Type:     4,
							CustomId: "MJ::RemixModal::new_prompt",
							Value:    "1girl, red hair, red clothes --ar 1:1 --relax",
						},
					},
				},
			},
		},
		SessionId: config.SessionId,
		Nonce:     gconv.String(nonce),
	}
	// 构造请求数据，把请求内容写入到队列中
	requestDataJson, err := gjson.Encode(requestData)
	if err != nil {
		glog.Line(true).Debug(err)
		return err
	}
	id := snowflake.GenerateID()
	queueData := &entity.QueueMidjourney{
		Id:              id,
		ConversationId:  params.ConversationId,
		ConfigId:        config.Id,
		ActionType:      params.ActionType,
		ApplicationType: referQueueData.ApplicationType,
		Nonce:           nonce,
		ReferMessageId:  referQueueData.MessageId,
		ReferIndex:      params.Index,
		MessageType:     MessageTypeModal,
		RequestType:     RequestTypeModal,
		RequestUrl:      ApiUrl + "interactions",
		RequestData:     gconv.String(requestDataJson),
		Status:          constant.QueueMidjourneyStatusInit,
		CreatedAt:       gconv.Int(xtime.GetNowTime()),
	}
	err = QueueInstance().InsertTask(queueData)
	if err != nil {
		glog.Line(true).Debug(err)
		return err
	}
	// 调用接口，接口调用次数增加
	if _, err = dao.ConfigMidjourney.Ctx(ctx).TX(tx).Data(g.Map{
		"call_num": gdb.Raw("call_num+1"),
	}).Where("id=?", config.Id).Update(); err != nil {
		glog.Line(true).Debug(err)
		return err
	}
	return nil
}

// trimPrompt 清理提示词
func trimPrompt(config *entity.ConfigMidjourney, prompt string) string {
	words := strings.Split(prompt, " ")
	cleanedWords := make([]string, 0)
	for _, word := range words {
		if word != "" && word != "--"+constant.ConfigMidjourneyCreateModelFast && word != "--"+constant.ConfigMidjourneyCreateModelRelax && word != "--"+constant.ConfigMidjourneyCreateModelTurbo {
			cleanedWords = append(cleanedWords, word)
		}
	}
	cleanedWords = append(cleanedWords, "--"+config.CreateModel)
	cleanedStr := strings.Join(cleanedWords, " ")
	return cleanedStr
}
