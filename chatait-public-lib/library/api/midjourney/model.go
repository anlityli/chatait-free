// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package midjourney

type GenerateImageParams struct {
	ConversationId  int64  `json:"conversation_id"`
	ApplicationType int    `json:"application_type"`
	Prompt          string `json:"prompt"`
}

type CustomIdImageParams struct {
	ActionType          int
	ConversationId      int64
	ReferConversationId int64
	Index               int
	CustomId            string
}

type ReqTriggerDiscord struct {
	Type          int64      `json:"type"`
	GuildId       string     `json:"guild_id"`
	ChannelId     string     `json:"channel_id"`
	ApplicationId string     `json:"application_id"`
	SessionId     string     `json:"session_id"`
	Data          *DSCommand `json:"data"`
	Nonce         string     `json:"nonce"`
}

type DSCommand struct {
	Version            string                   `json:"version"`
	Id                 string                   `json:"id"`
	Name               string                   `json:"name"`
	Type               int64                    `json:"type"`
	Options            []*DSOption              `json:"options"`
	ApplicationCommand *DSApplicationCommand    `json:"application_command"`
	Attachments        []*ReqCommandAttachments `json:"attachments"`
}

type DSOption struct {
	Type  int64       `json:"type"`
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

type DSApplicationCommand struct {
	Id                       string             `json:"id"`
	ApplicationId            string             `json:"application_id"`
	Version                  string             `json:"version"`
	DefaultPermission        bool               `json:"default_permission"`
	DefaultMemberPermissions map[string]int     `json:"default_member_permissions"`
	Type                     int64              `json:"type"`
	Nsfw                     bool               `json:"nsfw"`
	Name                     string             `json:"name"`
	Description              string             `json:"description"`
	DmPermission             bool               `json:"dm_permission"`
	Options                  []*DSCommandOption `json:"options"`
}

type DSCommandOption struct {
	Type        int64  `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
}

type ReqCustomIdDiscord struct {
	Type          int64         `json:"type"`
	GuildId       string        `json:"guild_id"`
	ChannelId     string        `json:"channel_id"`
	MessageFlags  int64         `json:"message_flags"`
	MessageId     string        `json:"message_id"`
	ApplicationId string        `json:"application_id"`
	SessionId     string        `json:"session_id"`
	Data          *CustomIdData `json:"data"`
	Nonce         string        `json:"nonce"`
}

type CustomIdData struct {
	ComponentType int64  `json:"component_type"`
	CustomId      string `json:"custom_id"`
}

type ReqCommandAttachments struct {
	Id             string `json:"id"`
	Filename       string `json:"filename"`
	UploadFilename string `json:"uploaded_filename"`
}

type ReqModalDiscord struct {
	Type          int64      `json:"type"`
	ApplicationId string     `json:"application_id"`
	ChannelId     string     `json:"channel_id"`
	GuildId       string     `json:"guild_id"`
	Data          *ModalData `json:"data"`
	SessionId     string     `json:"session_id"`
	Nonce         string     `json:"nonce"`
}

type ModalData struct {
	Id         string                     `json:"id"`
	CustomId   string                     `json:"custom_id"`
	Components []*ModalDataComponentsItem `json:"components"`
}

type ModalDataComponentsItem struct {
	Type       int64                                    `json:"type"`
	Components []*ModalDataComponentsItemComponentsItem `json:"components"`
}

type ModalDataComponentsItemComponentsItem struct {
	Type     int64  `json:"type"`
	CustomId string `json:"custom_id"`
	Value    string `json:"value"`
}
