// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package constant

const (
	ConfigLevelMember = 1
	ConfigLevelPlus   = 2
)

const (
	ConfigMidjourneyListenModelBot     = 0
	ConfigMidjourneyListenModelUserWss = 1
)

const (
	ConfigMidjourneyCreateModelFast  = "fast"
	ConfigMidjourneyCreateModelRelax = "relax"
	ConfigMidjourneyCreateModelTurbo = "turbo"
)

const (
	ConfigBaiduFeatureTranslate = "translate" // 翻译
	ConfigBaiduFeatureCensor    = "censor"    // 内容审核
)

const (
	ConfigSensitiveWordValidateTypeConversation = 1 // 对话
	ConfigSensitiveWordValidateTypeNickname     = 2 // 昵称
)
