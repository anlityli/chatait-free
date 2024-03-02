// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package response

type ConfigOptionItem struct {
	ConfigName string                    `json:"config_name"`
	Title      string                    `json:"title"`
	Unit       string                    `json:"unit"`
	InputType  int                       `json:"input_type"`
	Options    []*ConfigOptionItemOption `json:"options"`
	Value      interface{}               `json:"value"`
	Type       string                    `json:"type"`
	Sort       int                       `json:"sort"`
	CreatedAt  int                       `json:"created_at"`
	UpdatedAt  int                       `json:"updated_at"`
}

type ConfigOptionItemOption struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type ConfigOptionList []*ConfigOptionItem

type ConfigLevel struct {
	Id              int    `json:"id"`
	LevelName       string `json:"level_name"`
	MonthGpt3       int    `json:"month_gpt3"`
	MonthGpt4       int    `json:"month_gpt4"`
	MonthMidjourney int    `json:"month_midjourney"`
}

type ConfigLevelList []*ConfigLevel

type ConfigWallet struct {
	Field      string `json:"field"`
	WalletName string `json:"wallet_name"`
}

type ConfigWalletList []*ConfigWallet

type ConfigPay struct {
	Id                  string                  `json:"id"`
	ApiName             string                  `json:"api_name"`
	Params              []*ConfigPayParamsItem  `json:"params"`
	PayChannel          []*ConfigPayChannelItem `json:"pay_channel"`
	FrontendDescription string                  `json:"frontend_description"`
	BackendDescription  string                  `json:"backend_description"`
	Status              int                     `json:"status"`
	CreatedAt           int                     `json:"created_at"`
	UpdatedAt           int                     `json:"updated_at"`
}

type ConfigPayList []*ConfigPay

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

type ConfigMidjourney struct {
	Id               string `json:"id"`
	Title            string `json:"title"`
	GuildId          string `json:"guild_id"`
	ChannelId        string `json:"channel_id"`
	UserToken        string `json:"user_token"`
	MjBotId          string `json:"mj_bot_id"`
	BotToken         string `json:"bot_token"`
	SessionId        string `json:"session_id"`
	UserAgent        string `json:"user_agent"`
	HuggingFaceToken string `json:"hugging_face_token"`
	Proxy            string `json:"proxy"`
	Status           int    `json:"status"`
	ListenModel      int    `json:"listen_model"`
	CreateModel      string `json:"create_model"`
	WsIdleTime       int    `json:"ws_idle_time"`
	CallNum          int    `json:"call_num"`
	CreatedAt        int    `json:"created_at"`
	UpdatedAt        int    `json:"updated_at"`
}

type ConfigMidjourneyList []*ConfigMidjourney

type ConfigOpenai struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	ApiUrl    string `json:"api_url"`
	ApiKey    string `json:"api_key"`
	Proxy     string `json:"proxy"`
	MaxTokens int    `json:"max_tokens"`
	Gpt3Model string `json:"gpt3_model"`
	Gpt4Model string `json:"gpt4_model"`
	Status    int    `json:"status"`
	CallNum   int    `json:"call_num"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
}

type ConfigOpenaiList []*ConfigOpenai

type ConfigBaidu struct {
	Id        string   `json:"id"`
	Title     string   `json:"title"`
	ApiKey    string   `json:"api_key"`
	SecretKey string   `json:"secret_key"`
	Status    int      `json:"status"`
	Features  []string `json:"features"`
	CallNum   int      `json:"call_num"`
	CreatedAt int      `json:"created_at"`
	UpdatedAt int      `json:"updated_at"`
}

type ConfigBaiduList []*ConfigBaidu

type ConfigSensitiveWord struct {
	Id        string `json:"id"`
	Content   string `json:"content"`
	CreatedAt int    `json:"created_at"`
}

type ConfigSensitiveWordList []*ConfigSensitiveWord
