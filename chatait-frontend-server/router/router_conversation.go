// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package router

import (
	"github.com/anlityli/chatait-free/chatait-frontend-server/app/api/conversation"
	"github.com/anlityli/chatait-free/chatait-frontend-server/router/utils"
)

var conversationApi = &conversation.Conversation{}

var conversationRouter = []*utils.RouterItem{
	{Method: "GET", Pattern: "/topic-list", Object: conversationApi.TopicList},
	{Method: "GET", Pattern: "/topic-detail", Object: conversationApi.TopicDetail},
	{Method: "GET", Pattern: "/topic-list-by-last-id", Object: conversationApi.TopicListByLastId},
	{Method: "POST", Pattern: "/topic-edit", Object: conversationApi.TopicEdit},
	{Method: "POST", Pattern: "/topic-del", Object: conversationApi.TopicDel},
	{Method: "GET", Pattern: "/speak-list", Object: conversationApi.SpeakList},
	{Method: "POST", Pattern: "/speak", Object: conversationApi.Speak},
	{Method: "GET", Pattern: "/stream-uuid", Object: conversationApi.StreamUuid},
	{Method: "GET", Pattern: "/es/speak-stream", Object: conversationApi.SpeakStream, NoSign: true, NoLogin: true},
}
