// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package router

import (
	"github.com/anlityli/chatait-free/chatait-backend-server/app/api/conversation"
	"github.com/anlityli/chatait-free/chatait-backend-server/router/utils"
)

var conversationApi = &conversation.Conversation{}

var conversationRouter = []*utils.RouterItem{
	{Method: "GET", Pattern: "/topic-list", Object: conversationApi.TopicList},
	{Method: "GET", Pattern: "/list", Object: conversationApi.List},
}
