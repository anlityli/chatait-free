// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package midjourney

// ApplicationId Midjourney在Discord的应用Id
const ApplicationId = "936929561302675456"

// ID
const ApiUrl = "https://discord.com/api/v9/"
const WsUrl = "wss://gateway.discord.gg/?encoding=json&v=9"

const (
	MJApplicationId                   = "936929561302675456"
	MJVersionId                       = "1166847114203123795"
	MJCommandGenerateImageId          = "938956540159881230"
	MJCommandGenerateImageDescription = "Create images with Midjourney"
)

const (
	NJApplicationId                   = "1022952195194359889"
	NJVersionId                       = "1166842163141816443"
	NJCommandGenerateImageId          = "1023054140580057099"
	NJCommandGenerateImageDescription = "Create images with Niji journey"
)

const (
	MessageTypeGenerate = 0  // 生图
	MessageTypeCustomId = 19 // CustomId
)

const (
	RequestTypeGenerate = 2 // 生图
	RequestTypeCustomId = 3 // CustomId
)
