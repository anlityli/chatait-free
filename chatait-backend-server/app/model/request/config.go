// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package request

type ConfigId struct {
	Id string `json:"id" v:"required"`
}

type ConfigIds struct {
	Selected []string `json:"selected" v:"required"`
}

type ConfigOptionEdit struct {
	ConfigName string `json:"config_name" v:"required"`
	Value      string `json:"value" v:"required"`
}

type ConfigWalletEdit struct {
	Field      string `json:"field" v:"required"`
	WalletName string `json:"wallet_name" v:"required"`
}

type ConfigWalletOne struct {
	Field string `json:"field" v:"required"`
}

type ConfigPayEdit struct {
	Id                  string                  `json:"id" v:"required"`
	Params              []*ConfigPayParamsItem  `json:"params" v:"required"`
	PayChannel          []*ConfigPayChannelItem `json:"pay_channel" v:"required"`
	FrontendDescription string                  `json:"frontend_description"`
	Status              int                     `json:"status"`
}

type ConfigPayParamsItem struct {
	Param     string `json:"param"`
	ParamName string `json:"param_name"`
	Value     string `json:"value"`
}

type ConfigPayChannelItem struct {
	Id          int    `json:"id"`
	ChannelName string `json:"channel_name"`
	Channel     string `json:"channel"`
	Status      int    `json:"status"`
}

type ConfigLevelEdit struct {
	Id    int    `json:"id" v:"required"`
	Field string `json:"field" v:"required"`
	Value string `json:"value" v:"required"`
}

type ConfigMidjourneyAdd struct {
	Title            string `json:"title" v:"required"`
	GuildId          string `json:"guild_id" v:"required"`
	ChannelId        string `json:"channel_id" v:"required"`
	UserToken        string `json:"user_token" v:"required"`
	MjBotId          string `json:"mj_bot_id" v:"required"`
	BotToken         string `json:"bot_token" v:"required"`
	SessionId        string `json:"session_id" v:"required"`
	UserAgent        string `json:"user_agent" v:"required"`
	HuggingFaceToken string `json:"hugging_face_token"`
	Proxy            string `json:"proxy"`
	Status           int    `json:"status" v:"required"`
	ListenModel      int    `json:"listen_model" v:"required"`
	CreateModel      string `json:"create_model" v:"required"`
	WsIdleTime       int    `json:"ws_idle_time"`
}

type ConfigMidjourneyEdit struct {
	Id string `json:"id" v:"required"`
	ConfigMidjourneyAdd
}

type ConfigOpenaiAdd struct {
	Title     string `json:"title" v:"required"`
	ApiKey    string `json:"api_key" v:"required"`
	Proxy     string `json:"proxy"`
	MaxTokens int    `json:"max_tokens"`
	Status    int    `json:"status" v:"required"`
}

type ConfigOpenaiEdit struct {
	Id string `json:"id" v:"required"`
	ConfigOpenaiAdd
}

type ConfigBaiduAdd struct {
	Title     string `json:"title" v:"required"`
	ApiKey    string `json:"api_key" v:"required"`
	SecretKey string `json:"secret_key" v:"required"`
	Status    int    `json:"status" v:"required"`
}

type ConfigBaiduEdit struct {
	Id string `json:"id" v:"required"`
	ConfigBaiduAdd
}
