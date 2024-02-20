// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package constant

const (
	QueueMidjourneyStatusInit       = 0
	QueueMidjourneyStatusProceeding = 1
	QueueMidjourneyStatusEnded      = 2
	QueueMidjourneyStatusError      = 3
	QueueMidjourneyStatusInterrupt  = 4 // 中断主要是一些模态框的任务不能认定为全部完成，认为是中断也不用再继续了
)

const (
	QueueMidjourneyEventInsertQueue         = "insertQueue"
	QueueMidjourneyEventInteractionCreate   = "interactionCreate"
	QueueMidjourneyEventInteractionSuccess  = "interactionSuccess"
	QueueMidjourneyEventWriteMessageContent = "writeMessageContent"
	QueueMidjourneyEventEnded               = "ended"
	QueueMidjourneyEventProgress            = "progress"
	QueueMidjourneyEventContinue            = "continue"
	QueueMidjourneyEventVerifyHuman         = "verifyHuman"
	QueueMidjourneyEventError               = "error"
)

const (
	QueueMidjourneyApplicationTypeMJ = 1 // MJ机器人
	QueueMidjourneyApplicationTypeNJ = 2 // Niji机器人
)
