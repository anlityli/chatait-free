// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package midjourney

// ApplicationId Midjourney在Discord的应用Id
const ApplicationId = "936929561302675456"
const ApiUrl = "https://discord.com/api/v9/"
const WsUrl = "wss://gateway.discord.gg/?encoding=json&v=9"

const (
	MessageTypeGenerate = 0  // 生图
	MessageTypeCustomId = 19 // CustomId
)

const (
	RequestTypeGenerate = 2 // 生图
	RequestTypeCustomId = 3 // CustomId
)
