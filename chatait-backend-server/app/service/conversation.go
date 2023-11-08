// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package service

import (
	"database/sql"
	"github.com/anlityli/chatait-free/chatait-backend-server/app/model/response"
	"github.com/anlityli/chatait-free/chatait-backend-server/app/service/column"
	"github.com/anlityli/chatait-free/chatait-backend-server/library/datalist"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/constant"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/model/entity"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/notice"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/page"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/security"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var Conversation = &conversationService{}

type conversationService struct {
}

// TopicList 话题列表
func (s *conversationService) TopicList(r *ghttp.Request) (re *datalist.Result, err error) {
	columnsModel := &column.Conversation{}
	listColumns := columnsModel.TopicListColumns()
	// 筛选
	whereAndParams, err := datalist.FilterWhereAndParams(r, listColumns)
	if err != nil {
		return nil, err
	}
	listModel := &response.ConversationTopicList{}
	data, err := page.Data(r, &page.Param{
		TableName:   dao.Topic.Table + " t",
		Where:       whereAndParams.Where,
		WhereParams: whereAndParams.Params,
		Join: page.ParamJoin{
			&page.ParamJoinItem{
				JoinType:  "leftJoin",
				JoinTable: dao.User.Table + " u",
				On:        "u.id=t.user_id",
			},
			&page.ParamJoinItem{
				JoinType:  "leftJoin",
				JoinTable: dao.UserInfo.Table + " ui",
				On:        "ui.user_id=t.user_id",
			},
		},
		Field:   "t.*,u.username,ui.nickname",
		OrderBy: "t.id DESC",
	}, listModel)
	if err != nil {
		return nil, err
	}
	return datalist.List(r, data, listColumns)
}

// List 对话列表
func (s *conversationService) List(r *ghttp.Request) (re *datalist.Result, err error) {
	columnsModel := &column.Conversation{}
	listColumns := columnsModel.ListColumns()
	// 筛选
	whereAndParams, err := datalist.FilterWhereAndParams(r, listColumns)
	if err != nil {
		return nil, err
	}
	listModel := &response.ConversationList{}
	data, err := page.Data(r, &page.Param{
		TableName:   dao.Conversation.Table + " c",
		Where:       whereAndParams.Where,
		WhereParams: whereAndParams.Params,
		Join: page.ParamJoin{
			&page.ParamJoinItem{
				JoinType:  "leftJoin",
				JoinTable: dao.Topic.Table + " t",
				On:        "t.id=c.topic_id",
			},
			&page.ParamJoinItem{
				JoinType:  "leftJoin",
				JoinTable: dao.User.Table + " u",
				On:        "u.id=t.user_id",
			},
			&page.ParamJoinItem{
				JoinType:  "leftJoin",
				JoinTable: dao.UserInfo.Table + " ui",
				On:        "ui.user_id=t.user_id",
			},
		},
		Field:   "c.*,t.title topic_title,t.type topic_type, u.username,ui.nickname",
		OrderBy: "c.id DESC",
	}, listModel)
	if err != nil {
		return nil, err
	}
	re, err = datalist.List(r, data, listColumns)
	if err != nil {
		return nil, err
	}
	hostUrl, err := security.HostUrl(r)
	if err != nil {
		notice.Write(r, notice.OtherError, "域名授权失败，您可能正在使用盗版程序，请购买正版")
		return nil, err
	}
	for _, item := range re.List {
		if gconv.Int(item["topic_type"].OriValue) == constant.TopicTypeMidjourney {
			itemMjData := &response.ConversationSpeakItemMjData{}
			mjData := &entity.ConversationMidjourney{}
			err = dao.ConversationMidjourney.Where("conversation_id=?", gconv.String(item["id"].OriValue)).Scan(mjData)
			if err != nil && err != sql.ErrNoRows {
				return nil, err
			}
			if mjData.ConversationId > 0 {
				itemMjData.ActionType = mjData.ActionType
				itemMjData.Error = mjData.ErrorData
				fileData := &entity.FileMidjourney{}
				err = dao.FileMidjourney.Where("id=?", mjData.FileId).Scan(fileData)
				if err != nil && err != sql.ErrNoRows {
					return nil, err
				}
				if fileData.Id > 0 {
					itemMjData.ImgUrl = hostUrl + "/file/midjourney-image?id=" + gconv.String(fileData.Id)
					itemMjData.Prompt = fileData.Prompt
				}
			} else {
				queueData := &entity.QueueMidjourney{}
				err = dao.QueueMidjourney.Where("conversation_id=?", gconv.String(item["id"].OriValue)).Scan(queueData)
				if err != nil && err != sql.ErrNoRows {
					return nil, err
				}
				if queueData.Id > 0 {
					itemMjData.ActionType = queueData.ActionType
					itemMjData.Progress = queueData.Progress
				}
			}
			item["mj_data"].OriValue = itemMjData
		}
	}
	return re, nil
}
