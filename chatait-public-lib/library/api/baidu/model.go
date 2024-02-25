// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package baidu

import "github.com/anlityli/chatait-free/chatait-public-lib/app/model/entity"

type Config struct {
	entity.ConfigBaidu
	AccessToken         string
	AccessTokenExpireIn int64
}

type TransTextParams struct {
	From string `json:"from"`
	To   string `json:"to"`
	Q    string `json:"q"`
}

type TransTextResponse struct {
	LogId     string                   `json:"log_id"`
	Result    *TransTextResponseResult `json:"result"`
	ErrorMsg  string                   `json:"error_msg"`
	ErrorCode string                   `json:"error_code"`
}

type TransTextResponseResult struct {
	From        string                          `json:"from"`
	To          string                          `json:"to"`
	TransResult []*TransTextResponseTransResult `json:"trans_result"`
}

type TransTextResponseTransResult struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}

type CensorTextParams struct {
	Text   string `json:"text"`
	UserId string `json:"userId"`
}

type CensorTextResponse struct {
	LogId          string                        `json:"log_id"`
	ErrorMsg       string                        `json:"error_msg"`
	ErrorCode      string                        `json:"error_code"`
	Conclusion     string                        `json:"conclusion"`
	ConclusionType int                           `json:"conclusionType"` // 1.合规，2.不合规，3.疑似，4.审核失败
	Data           []*CensorTextResponseDataItem `json:"data"`
}

type CensorTextResponseDataItem struct {
	Type    int                               `json:"type"`
	SubType int                               `json:"subType"`
	Msg     string                            `json:"msg"`
	Hits    []*CensorTextResponseDataHitsItem `json:"hits"`
}

type CensorTextResponseDataHitsItem struct {
	Probability       string      `json:"probability"`
	DatasetName       string      `json:"datasetName"`
	Words             []string    `json:"words"`
	Details           interface{} `json:"details"`
	ModelHitPositions interface{} `json:"modelHitPositions"`
	WordHitPositions  interface{} `json:"wordHitPositions"`
}
