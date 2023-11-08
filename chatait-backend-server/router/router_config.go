// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package router

import (
	"github.com/anlityli/chatait-free/chatait-backend-server/app/api/config"
	"github.com/anlityli/chatait-free/chatait-backend-server/router/utils"
)

var configApi = &config.Config{}

var configRouter = []*utils.RouterItem{
	{Method: "GET", Pattern: "/all-option", Object: configApi.AllOption},
	{Method: "POST", Pattern: "/option-edit", Object: configApi.OptionEdit},
	{Method: "GET", Pattern: "/level-list", Object: configApi.LevelList, NoPermission: true},
	{Method: "POST", Pattern: "/level-edit", Object: configApi.LevelEdit},
	{Method: "GET", Pattern: "/wallet-list", Object: configApi.WalletList, NoPermission: true, NoLogin: true},
	{Method: "POST", Pattern: "/wallet-edit", Object: configApi.WalletEdit},
	{Method: "GET", Pattern: "/wallet-one", Object: configApi.WalletOne},
	{Method: "GET", Pattern: "/pay-list", Object: configApi.PayList},
	{Method: "GET", Pattern: "/pay-one", Object: configApi.PayOne},
	{Method: "POST", Pattern: "/pay-edit", Object: configApi.PayEdit},
	{Method: "GET", Pattern: "/midjourney-list", Object: configApi.MidjourneyList},
	{Method: "GET", Pattern: "/midjourney-one", Object: configApi.MidjourneyOne},
	{Method: "POST", Pattern: "/midjourney-add", Object: configApi.MidjourneyAdd},
	{Method: "POST", Pattern: "/midjourney-edit", Object: configApi.MidjourneyEdit},
	{Method: "POST", Pattern: "/midjourney-delete", Object: configApi.MidjourneyDelete},
	{Method: "GET", Pattern: "/openai-list", Object: configApi.OpenaiList},
	{Method: "GET", Pattern: "/openai-one", Object: configApi.OpenaiOne},
	{Method: "POST", Pattern: "/openai-add", Object: configApi.OpenaiAdd},
	{Method: "POST", Pattern: "/openai-edit", Object: configApi.OpenaiEdit},
	{Method: "POST", Pattern: "/openai-delete", Object: configApi.OpenaiDelete},
	{Method: "GET", Pattern: "/baidu-list", Object: configApi.BaiduList},
	{Method: "GET", Pattern: "/baidu-one", Object: configApi.BaiduOne},
	{Method: "POST", Pattern: "/baidu-add", Object: configApi.BaiduAdd},
	{Method: "POST", Pattern: "/baidu-edit", Object: configApi.BaiduEdit},
	{Method: "POST", Pattern: "/baidu-delete", Object: configApi.BaiduDelete},
}
