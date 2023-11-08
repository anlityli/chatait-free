// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package midjourney

import "github.com/bwmarrin/discordgo"

type WsRunConnCallback func()

type WsMessage struct {
	Op int         `json:"op"`
	D  interface{} `json:"d"`
}

type WsMessageAuth struct {
	Token        string                   `json:"token"`
	Capabilities int                      `json:"capabilities"`
	Properties   *WsMessageAuthProperties `json:"properties"`
	Compress     bool                     `json:"compress"`
}

type WsMessageAuthProperties struct {
	Os      string `json:"os"`
	Browser string `json:"browser"`
	Device  string `json:"device"`
}

type WsReceiveMessage struct {
	T  string      `json:"t"`
	S  int         `json:"s"`
	Op int         `json:"op"`
	D  interface{} `json:"d"`
}

type WsReceiveMessageDReady struct {
	User      *WsReceiveMessageDReadyUser `json:"user"`
	SessionId string                      `json:"session_id"`
}

type WsReceiveMessageDReadyUser struct {
	Id string `json:"id"`
}

type WsReceiveMessageDCommon struct {
	Id                string                              `json:"id"`
	ChannelId         string                              `json:"channel_id"`
	GuildId           string                              `json:"guild_id"`
	Type              int                                 `json:"type"`
	Author            *WsReceiveMessageDAuthor            `json:"author"`
	Nonce             string                              `json:"nonce"`
	Content           string                              `json:"content"`
	Flags             int                                 `json:"flags"`
	Components        []*WsReceiveMessageDComponentsItem  `json:"components"`
	Attachments       []*discordgo.MessageAttachment      `json:"attachments"`
	Embeds            []*WsReceiveMessageDEmbedsItem      `json:"embeds"`
	Interaction       *WsReceiveMessageDInteraction       `json:"interaction"`
	ReferencedMessage *WsReceiveMessageDReferencedMessage `json:"referenced_message"`
	MessageReference  *discordgo.MessageReference         `json:"message_reference"`
}

type WsReceiveMessageDAuthor struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

type WsReceiveMessageDComponentsItem struct {
	Type       int                                              `json:"type"`
	Components []*WsReceiveMessageDComponentsItemComponentsItem `json:"components"`
}

type WsReceiveMessageDComponentsItemComponentsItem struct {
	Type     int                                                 `json:"type"`
	Style    int                                                 `json:"style"`
	Label    string                                              `json:"label"`
	Emoji    *WsReceiveMessageDComponentsItemComponentsItemEmoji `json:"emoji"`
	CustomId string                                              `json:"custom_id"`
}

type WsReceiveMessageDComponentsItemComponentsItemEmoji struct {
	Name string `json:"name"`
}

type WsReceiveMessageDAttachmentsItem struct {
	Id          string `json:"id"`
	Filename    string `json:"filename"`
	ContentType string `json:"content_type"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	Url         string `json:"url"`
	ProxyUrl    string `json:"proxy_url"`
	Size        int    `json:"size"`
}

type WsReceiveMessageDEmbedsItem struct {
	Color       int                               `json:"color"`
	Description string                            `json:"description"`
	Title       string                            `json:"title"`
	Image       *WsReceiveMessageDEmbedsItemImage `json:"image"`
}

type WsReceiveMessageDEmbedsItemImage struct {
	Url string `json:"url"`
}

type WsReceiveMessageDReferencedMessage struct {
	Id         string                             `json:"id"`
	ChannelId  string                             `json:"channel_id"`
	Content    string                             `json:"content"`
	Timestamp  string                             `json:"timestamp"`
	Components []*WsReceiveMessageDComponentsItem `json:"components"`
}

type WsReceiveMessageDInteraction struct {
	Id   string                      `json:"id"`
	Name string                      `json:"name"`
	Type int                         `json:"type"`
	User *WsReceiveMessageDReadyUser `json:"user"`
}
