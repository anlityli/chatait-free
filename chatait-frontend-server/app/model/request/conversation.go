// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package request

type ConversationTopicList struct {
}

type ConversationTopicDetail struct {
	TopicId string `json:"topic_id" v:"required"`
}

type ConversationTopicListByLastId struct {
	LastId string `json:"last_id" v:"required"`
	Limit  int    `json:"limit"`
}

type ConversationSpeakList struct {
	TopicId string `json:"topic_id" v:"required"`
}

type ConversationSpeak struct {
	TopicType  int    `json:"topic_type" v:"required|in:1,2"`
	StreamUuid string `json:"stream_uuid" v:"required"`
	TopicId    string `json:"topic_id" v:"required"`
	Content    string `json:"content" v:"required|max-length:500#内容必填|最大长度500个字符"`
}

type ConversationSpeakStream struct {
	StreamUuid string `json:"stream_uuid" v:"required"`
}

type ConversationTopicDel struct {
	TopicId string `json:"topic_id" v:"required"`
}

type ConversationTopicEdit struct {
	TopicId string `json:"topic_id" v:"required"`
	Title   string `json:"title" v:"required|max-length:50"`
}

type ConversationMidjourneySpeak struct {
	TopicId         string `json:"topic_id" v:"required"`
	Content         string `json:"content" v:"required"`
	ApplicationType int    `json:"application_type" v:"required"`
	No              string `json:"no"`
	Images          string `json:"images"`
	Seed            string `json:"seed"`
	Ar              string `json:"ar"`
	Chaos           string `json:"chaos"`
	Quality         string `json:"quality"`
	Model           string `json:"model"`
	Style           string `json:"style"`
	Stylize         string `json:"stylize"`
	Tile            string `json:"tile"`
	Iw              string `json:"iw"`
}

type ConversationMidjourneyCustom struct {
	ReferConversationId string `json:"refer_conversation_id" v:"required"`
	ActionType          int    `json:"action_type" v:"required|in:2,3,4,5,6,7"`
	Index               int    `json:"index"`
	CustomId            string `json:"custom_id" v:"required"`
}
