// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package baidu

const AccessTokenUrl = "https://aip.baidubce.com/oauth/2.0/token"
const TransTextURL = "https://aip.baidubce.com/rpc/2.0/mt/texttrans/v1"
const CensorTextURL = "https://aip.baidubce.com/rest/2.0/solution/v1/text_censor/v2/user_defined"

// 审核结果
const (
	CensorTextConclusionTypePass      = 1 // 合规
	CensorTextConclusionTypeIllegal   = 2 // 不合规
	CensorTextConclusionTypeSuspected = 3 // 疑似
	CensorTextConclusionTypeFail      = 4 // 失败
)
