// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package router

import (
	"github.com/anlityli/chatait-free/chatait-backend-server/app/api/dashboard"
	"github.com/anlityli/chatait-free/chatait-backend-server/router/utils"
)

var dashboardApi = &dashboard.Dashboard{}

var dashboardRouter = []*utils.RouterItem{
	{Method: "GET", Pattern: "/user-statistic", Object: dashboardApi.UserStatistic},
	{Method: "GET", Pattern: "/amount-statistic", Object: dashboardApi.AmountStatistic},
	{Method: "GET", Pattern: "/order-statistic", Object: dashboardApi.OrderStatistic},
	{Method: "GET", Pattern: "/conversation-statistic", Object: dashboardApi.ConversationStatistic},
}
