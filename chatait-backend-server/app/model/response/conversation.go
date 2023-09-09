// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package response

type ConversationTopic struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Title     string `json:"title"`
	Type      int    `json:"type"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
}

type ConversationTopicList []*ConversationTopic

type Conversation struct {
	Id         string                       `json:"id"`
	UserId     string                       `json:"user_id"`
	Username   string                       `json:"username"`
	Nickname   string                       `json:"nickname"`
	TopicId    string                       `json:"topic_id"`
	TopicType  int                          `json:"topic_type"`
	TopicTitle string                       `json:"topic_title"`
	Role       string                       `json:"role"`
	Content    string                       `json:"content"`
	MjData     *ConversationSpeakItemMjData `json:"mj_data"`
	CreatedAt  int                          `json:"created_at"`
}

type ConversationList []*Conversation

type ConversationSpeakItemMjData struct {
	ActionType int    `json:"action_type"`
	ImgUrl     string `json:"img_url"`
	Prompt     string `json:"prompt"`
	Progress   int    `json:"progress"`
	Error      string `json:"error"`
}
