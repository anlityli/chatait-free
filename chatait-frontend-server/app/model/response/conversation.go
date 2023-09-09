// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package response

type ConversationTopic struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Type      int    `json:"type"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
}

type ConversationTopicList []*ConversationTopic

type ConversationSpeakItem struct {
	Id        string                       `json:"id"`
	TopicId   string                       `json:"topic_id"`
	Role      string                       `json:"role"`
	Content   string                       `json:"content"`
	MjData    *ConversationSpeakItemMjData `json:"mj_data"`
	CreatedAt int                          `json:"created_at"`
}

type ConversationSpeakItemMjData struct {
	ActionType int    `json:"action_type"`
	ImgUrl     string `json:"img_url"`
	Progress   int    `json:"progress"`
	Error      string `json:"error"`
}

type ConversationSpeakList []*ConversationSpeakItem

type ConversationStreamUuid struct {
	Uuid string `json:"uuid"`
}

type ConversationSpeak struct {
	TopicId   string `json:"topic_id"`
	Title     string `json:"title"`
	TopicType int    `json:"topic_type"`
}

type ConversationMidjourneySpeak struct {
	TopicId         string `json:"topic_id"`
	Title           string `json:"title"`
	TopicType       int    `json:"topic_type"`
	QuestionId      string `json:"question_id"`
	QuestionContent string `json:"question_content"`
	AnswerId        string `json:"answer_id"`
}
