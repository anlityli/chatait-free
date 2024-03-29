// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// QueueMidjourneyDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type QueueMidjourneyDao struct {
	gmvc.M                         // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB                 // DB is the raw underlying database management object.
	Table   string                 // Table is the table name of the DAO.
	Columns queueMidjourneyColumns // Columns contains all the columns of Table that for convenient usage.
}

// QueueMidjourneyColumns defines and stores column names for table c_queue_midjourney.
type queueMidjourneyColumns struct {
	Id              string // ID
	ConversationId  string // 对话ID
	ConfigId        string // 接口配置ID
	ActionType      string // 行为类型  1生图 2Upscale
	ApplicationType string // 应用机器人类型 1MJ 2Niji
	Nonce           string // nonceID
	MessageId       string // 消息结束ID(生成图片完成时的消息ID)
	ReferMessageId  string // 提到的消息ID(生图动作为0)
	InteractionId   string // 交互ID
	ReferIndex      string // 处理提到的消息的索引
	MessageHash     string // 消息hash
	MessageType     string // 消息type
	MessageContent  string // 消息内容(提示词内容用于匹配任务)
	RequestType     string // 请求消息时用到的类型 2生图 3Upscale 3variation
	RequestUrl      string // 请求接口的url
	RequestData     string // 请求接口的数据内容
	ResponseData    string // 接口返回的数据内容
	ErrorData       string // 错误数据内容
	Status          string // 状态 0任务进入队列 1 任务开始 2任务正常结束 3任务出错
	Progress        string // 任务执行进度
	CreatedAt       string // 创建时间
	StartedAt       string // 任务开始时间
	EndedAt         string // 任务结束时间
	ErrorAt         string // 任务发生错误时间
}

func NewQueueMidjourneyDao() *QueueMidjourneyDao {
	return &QueueMidjourneyDao{
		M:     g.DB("default").Model("c_queue_midjourney").Safe(),
		DB:    g.DB("default"),
		Table: "c_queue_midjourney",
		Columns: queueMidjourneyColumns{
			Id:              "id",
			ConversationId:  "conversation_id",
			ConfigId:        "config_id",
			ActionType:      "action_type",
			ApplicationType: "application_type",
			Nonce:           "nonce",
			MessageId:       "message_id",
			ReferMessageId:  "refer_message_id",
			InteractionId:   "interaction_id",
			ReferIndex:      "refer_index",
			MessageHash:     "message_hash",
			MessageType:     "message_type",
			MessageContent:  "message_content",
			RequestType:     "request_type",
			RequestUrl:      "request_url",
			RequestData:     "request_data",
			ResponseData:    "response_data",
			ErrorData:       "error_data",
			Status:          "status",
			Progress:        "progress",
			CreatedAt:       "created_at",
			StartedAt:       "started_at",
			EndedAt:         "ended_at",
			ErrorAt:         "error_at",
		},
	}
}
