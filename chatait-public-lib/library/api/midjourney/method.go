// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package midjourney

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/constant"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/model/entity"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/snowflake"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/xtime"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/text/gstr"
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
	err = QueueInstance().InsertTask(queueData, func(signal int) {

	})
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
	referQueueData := &entity.QueueMidjourney{}
	err = dao.QueueMidjourney.Ctx(ctx).TX(tx).Where("conversation_id=?", params.ReferConversationId).Scan(referQueueData)
	if err != nil && err != sql.ErrNoRows {
		glog.Line(true).Debug(err)
		return err
	}
	if referQueueData.Id <= 0 {
		return errors.New("对话相应的队列信息不存在")
	}
	// 为保障新的任务和引用的任务是同一配置，从引用的队列中获取配置
	config := &entity.ConfigMidjourney{}
	err = dao.ConfigMidjourney.Where("id=? AND status=1", referQueueData.ConfigId).Scan(config)
	if err != nil {
		return errors.New("绘画配置不存在")
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
	err = QueueInstance().InsertTask(queueData, func(signal int) {
		glog.Line(true).Debug("队列任务完成", queueData, signal)
		// 如果是弹模态框的任务则再调用模态方法
		if isModalCustomId(params.CustomId) {
			err = CustomIdModalImage(&CustomIdModalImageParams{
				ActionType:          params.ActionType,
				ConversationId:      params.ConversationId,
				ReferConversationId: params.ReferConversationId,
				Index:               params.Index,
				OriCustomId:         params.CustomId,
				DataId:              gconv.String(queueData.InteractionId),
				NewPrompt:           params.NewPrompt,
			})
			if err != nil {
				glog.Line(true).Debug(err)
			}
		}
	})
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

func CustomIdModalImage(params *CustomIdModalImageParams) (err error) {
	referQueueData := &entity.QueueMidjourney{}
	err = dao.QueueMidjourney.Where("conversation_id=?", params.ReferConversationId).Scan(referQueueData)
	if err != nil && err != sql.ErrNoRows {
		glog.Line(true).Debug(err)
		return err
	}
	if referQueueData.Id <= 0 {
		return errors.New("对话响应的队列信息不存在")
	}
	// 为保障新的任务和引用的任务是同一配置，从引用的队列中获取配置
	config := &entity.ConfigMidjourney{}
	err = dao.ConfigMidjourney.Where("id=? AND status=1", referQueueData.ConfigId).Scan(config)
	if err != nil {
		return errors.New("绘画配置不存在")
	}
	applicationId := MJApplicationId
	if referQueueData.ApplicationType == constant.QueueMidjourneyApplicationTypeNJ {
		applicationId = NJApplicationId
	}
	nonce := snowflake.GenerateDiscordId()
	if params.NewPrompt == "" {
		params.NewPrompt = referQueueData.MessageContent
	}
	if params.NewPrompt == "" {
		return errors.New("新提示词不能为空")
	}
	newPrompt := trimPrompt(config, params.NewPrompt)
	// 替换CustomId
	dataCustomId, componentsCustomId, err := replaceModalCustomId(params.OriCustomId, gconv.String(referQueueData.MessageId))
	if err != nil {
		return err
	}
	requestData := &ReqCustomIdModalDiscord{
		Type:          RequestTypeModal,
		ApplicationId: applicationId,
		ChannelId:     config.ChannelId,
		GuildId:       config.GuildId,
		Data: &CustomIdModalData{
			Id:       params.DataId,
			CustomId: dataCustomId,
			Components: []*CustomIdModalDataComponentsItem{
				{
					Type: 1,
					Components: []*CustomIdModalDataComponentsItemComponentsItem{
						{
							Type:     4,
							CustomId: componentsCustomId,
							Value:    newPrompt,
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
	err = QueueInstance().InsertTask(queueData, func(signal int) {

	})
	if err != nil {
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

func isModalCustomId(CustomId string) bool {
	dataCustomId, _, _ := replaceModalCustomId(CustomId, "")
	return dataCustomId != ""
}

func replaceModalCustomId(oriCustomId string, oriMessageId string) (dataCustomId string, componentsCustomId string, err error) {
	if gstr.Contains(oriCustomId, "MJ::JOB::reroll") {
		dataCustomId = "MJ::ImagineModal::" + oriMessageId
		componentsCustomId = "MJ::ImagineModal::new_prompt"
	} else if gstr.Contains(oriCustomId, "MJ::JOB::variation") {
		customIdArr := gstr.Explode("::", oriCustomId)
		dataCustomId = fmt.Sprintf("MJ::RemixModal::%s::%s::1", customIdArr[4], customIdArr[3])
		componentsCustomId = "MJ::RemixModal::new_prompt"
	} else if gstr.Contains(oriCustomId, "MJ::JOB::low_variation") {
		customIdArr := gstr.Explode("::", oriCustomId)
		dataCustomId = fmt.Sprintf("MJ::RemixModal::%s::%s::0", customIdArr[4], customIdArr[3])
		componentsCustomId = "MJ::RemixModal::new_prompt"
	} else if gstr.Contains(oriCustomId, "MJ::JOB::high_variation") {
		customIdArr := gstr.Explode("::", oriCustomId)
		dataCustomId = fmt.Sprintf("MJ::RemixModal::%s::%s::1", customIdArr[4], customIdArr[3])
		componentsCustomId = "MJ::RemixModal::new_prompt"
	} else if gstr.Contains(oriCustomId, "MJ::JOB::pan_left") {
		customIdArr := gstr.Explode("::", oriCustomId)
		dataCustomId = fmt.Sprintf("MJ::PanModal::left::%s::%s", customIdArr[4], customIdArr[3])
		componentsCustomId = "MJ::PanModal::prompt"
	} else if gstr.Contains(oriCustomId, "MJ::JOB::pan_right") {
		customIdArr := gstr.Explode("::", oriCustomId)
		dataCustomId = fmt.Sprintf("MJ::PanModal::right::%s::%s", customIdArr[4], customIdArr[3])
		componentsCustomId = "MJ::PanModal::prompt"
	} else if gstr.Contains(oriCustomId, "MJ::JOB::pan_up") {
		customIdArr := gstr.Explode("::", oriCustomId)
		dataCustomId = fmt.Sprintf("MJ::PanModal::up::%s::%s", customIdArr[4], customIdArr[3])
		componentsCustomId = "MJ::PanModal::prompt"
	} else if gstr.Contains(oriCustomId, "MJ::JOB::pan_down") {
		customIdArr := gstr.Explode("::", oriCustomId)
		dataCustomId = fmt.Sprintf("MJ::PanModal::down::%s::%s", customIdArr[4], customIdArr[3])
		componentsCustomId = "MJ::PanModal::prompt"
	} else {
		return "", "", errors.New("不支持的格式")
	}
	return dataCustomId, componentsCustomId, nil
}
