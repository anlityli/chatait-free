// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package response

import "github.com/anlityli/chatait-free/chatait-public-lib/library/api/midjourney"

type WebsocketConversationMidjourneyListenerEvent struct {
	ConversationId           string                                        `json:"conversation_id"`
	UserId                   string                                        `json:"user_id"`
	TopicId                  string                                        `json:"topic_id"`
	TopicTitle               string                                        `json:"topic_title"`
	TopicType                int                                           `json:"topic_type"`
	Role                     string                                        `json:"role"`
	ActionType               int                                           `json:"action_type"`
	Content                  string                                        `json:"content"`
	ImgUrl                   string                                        `json:"img_url"`
	ThumbnailImgUrl          string                                        `json:"thumbnail_img_url"`
	Progress                 int                                           `json:"progress"`
	Components               []*midjourney.WsReceiveMessageDComponentsItem `json:"components"`
	ReferencedConversationId string                                        `json:"referenced_conversation_id"`
	ReferencedComponents     []*midjourney.WsReceiveMessageDComponentsItem `json:"referenced_components"`
	Error                    string                                        `json:"error"`
}
