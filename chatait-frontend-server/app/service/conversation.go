// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package service

import (
	"context"
	"github.com/anlityli/chatait-free/chatait-frontend-server/app/model/request"
	"github.com/anlityli/chatait-free/chatait-frontend-server/app/model/response"
	"github.com/anlityli/chatait-free/chatait-frontend-server/library/auth"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/page"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"sync"
)

var Conversation = &conversationService{}

type conversationService struct {
}

var streamMap sync.Map

// TopicList 话题列表
func (s *conversationService) TopicList(r *ghttp.Request) (re *page.Response, err error) {
	requestModel := &request.ConversationTopicList{}
	if err = r.Parse(requestModel); err != nil {
		return nil, err
	}
	userId := auth.GetUserId(r)
	listData := &response.ConversationTopicList{}
	return page.Data(r, &page.Param{
		TableName:   dao.Topic.Table,
		Where:       "user_id=?",
		WhereParams: g.Slice{userId},
		OrderBy:     "id DESC",
	}, listData)
}

func (s *conversationService) TopicDetail(r *ghttp.Request) (re *response.ConversationTopic, err error) {
	requestModel := &request.ConversationTopicDetail{}
	if err = r.Parse(requestModel); err != nil {
		return nil, err
	}
	re = &response.ConversationTopic{}
	err = dao.Topic.Where("id=?", requestModel.TopicId).Scan(re)
	if err != nil {
		return nil, err
	}
	return re, nil
}

func (s *conversationService) TopicListByLastId(r *ghttp.Request) (re *response.ConversationTopicList, err error) {
	requestModel := &request.ConversationTopicListByLastId{}
	if err = r.Parse(requestModel); err != nil {
		return nil, err
	}
	if requestModel.Limit == 0 {
		requestModel.Limit = 10
	}
	re = &response.ConversationTopicList{}
	userId := auth.GetUserId(r)
	if userId == 0 {
		*re = make(response.ConversationTopicList, 0)
		return re, nil
	}
	where := "user_id=?"
	params := g.Slice{userId}
	if gconv.Int64(requestModel.LastId) != 0 {
		where += " AND id<?"
		params = append(params, requestModel.LastId)
	}
	err = dao.Topic.Where(where, params).Order("id DESC").Limit(requestModel.Limit).Scan(re)
	if err != nil {
		return nil, err
	}
	return re, nil
}

// TopicEdit 话题编辑
func (s *conversationService) TopicEdit(r *ghttp.Request) (err error) {
	requestModel := &request.ConversationTopicEdit{}
	if err = r.Parse(requestModel); err != nil {
		return err
	}
	userId := auth.GetUserId(r)
	if _, err = dao.Topic.Data(g.Map{
		"title": requestModel.Title,
	}).Where("id=? AND user_id=?", requestModel.TopicId, userId).Update(); err != nil {
		return err
	}
	return nil
}

// TopicDel 删除话题
func (s *conversationService) TopicDel(r *ghttp.Request) (err error) {
	requestModel := &request.ConversationTopicDel{}
	if err = r.Parse(requestModel); err != nil {
		return err
	}
	userId := auth.GetUserId(r)
	if err := g.DB().Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) (err error) {
		if _, err = dao.Topic.Ctx(ctx).TX(tx).Where("id=? AND user_id=?", requestModel.TopicId, userId).Delete(); err != nil {
			return err
		}
		if _, err = dao.Conversation.Ctx(ctx).TX(tx).Where("user_id=? AND topic_id=?", userId, requestModel.TopicId).Delete(); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// SpeakList 聊天记录列表
func (s *conversationService) SpeakList(r *ghttp.Request) (re *page.Response, err error) {
	requestModel := &request.ConversationSpeakList{}
	if err = r.Parse(requestModel); err != nil {
		return nil, err
	}
	userId := auth.GetUserId(r)
	listData := &response.ConversationSpeakList{}
	re, err = page.Data(r, &page.Param{
		TableName:   dao.Conversation.Table,
		Where:       "user_id=? AND topic_id=? AND role<>?",
		WhereParams: g.Slice{userId, requestModel.TopicId, "system"},
		OrderBy:     "id DESC",
	}, listData)
	if err != nil {
		return nil, err
	}
	return re, nil
}
