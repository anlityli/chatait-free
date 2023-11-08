// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package conversation

import (
	"github.com/anlityli/chatait-free/chatait-frontend-server/app/service"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/notice"
	"github.com/gogf/gf/net/ghttp"
)

type Conversation struct {
}

func (c *Conversation) TopicList(r *ghttp.Request) {
	if re, err := service.Conversation.TopicList(r); err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	} else {
		notice.Write(r, notice.NoError, re)
	}
}

func (c *Conversation) TopicDetail(r *ghttp.Request) {
	if re, err := service.Conversation.TopicDetail(r); err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	} else {
		notice.Write(r, notice.NoError, re)
	}
}

func (c *Conversation) TopicListByLastId(r *ghttp.Request) {
	if re, err := service.Conversation.TopicListByLastId(r); err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	} else {
		notice.Write(r, notice.NoError, re)
	}
}

func (c *Conversation) TopicEdit(r *ghttp.Request) {
	if err := service.Conversation.TopicEdit(r); err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	} else {
		notice.Write(r, notice.NoError, "操作成功")
	}
}

func (c *Conversation) TopicDel(r *ghttp.Request) {
	if err := service.Conversation.TopicDel(r); err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	} else {
		notice.Write(r, notice.NoError, "操作成功")
	}
}

func (c *Conversation) SpeakList(r *ghttp.Request) {
	if re, err := service.Conversation.SpeakList(r); err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	} else {
		notice.Write(r, notice.NoError, re)
	}
}

func (c *Conversation) Speak(r *ghttp.Request) {
	if re, err := service.ConversationOpenai.Speak(r); err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	} else {
		notice.Write(r, notice.NoError, re)
	}
}

// SpeakStream 对话流
func (c *Conversation) SpeakStream(r *ghttp.Request) {
	if err := service.ConversationOpenai.SpeakStream(r); err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
}

// StreamUuid 生成
func (c *Conversation) StreamUuid(r *ghttp.Request) {
	notice.Write(r, notice.NoError, service.ConversationOpenai.StreamUuid())
}

// MidjourneySpeak mj的对话
func (c *Conversation) MidjourneySpeak(r *ghttp.Request) {
	if re, err := service.ConversationMidjourney.Speak(r); err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	} else {
		notice.Write(r, notice.NoError, re)
	}
}

func (c *Conversation) MidjourneyCustom(r *ghttp.Request) {
	if re, err := service.ConversationMidjourney.Custom(r); err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	} else {
		notice.Write(r, notice.NoError, re)
	}
}
