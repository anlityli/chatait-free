// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package openai

type ChatCompletionParams struct {
	Model    string
	Messages RequestChatParamsMessages
}

type CreateChatCompletionCallbackFunc func(originContent string, contentObj *ResponseChat) error

type RequestChatParams struct {
	Model     string                    `json:"model"`
	Messages  RequestChatParamsMessages `json:"messages"`
	MaxTokens int                       `json:"max_tokens"`
	Stream    bool                      `json:"stream"`
}

type RequestChatParamsMessages []*RequestChatParamsMessageItem

type RequestChatParamsMessageItem struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ResponseChat struct {
	Id      string              `json:"id"`
	Object  string              `json:"object"`
	Created int                 `json:"created"`
	Choices ResponseChatChoices `json:"choices"`
	Usage   *ResponseChatUsage  `json:"usage"`
}

type ResponseChatChoices []*ResponseChatChoicesItem

type ResponseChatChoicesItem struct {
	Index        int                             `json:"index"`
	Message      *ResponseChatChoicesItemMessage `json:"message"`
	Delta        *ResponseChatChoicesItemMessage `json:"delta"`
	FinishReason string                          `json:"finish_reason"`
}

type ResponseChatChoicesItemMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ResponseChatUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type ResponseChatError struct {
	Error *ResponseChatErrorData `json:"error"`
}

type ResponseChatErrorData struct {
	Message string      `json:"message"`
	Type    string      `json:"type"`
	Param   interface{} `json:"param"`
	Code    string      `json:"code"`
}
