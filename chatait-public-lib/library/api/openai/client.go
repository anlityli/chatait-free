// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package openai

import (
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/model/entity"
	"github.com/gogf/gf/os/glog"
	"sync"
)

// OpenaiClient openai推送
type OpenaiClient struct {
}

var openaiClient *OpenaiClient
var openaiClientOnce sync.Once

// Instance 单例
func Instance() *OpenaiClient {
	openaiClientOnce.Do(func() {
		openaiClient = &OpenaiClient{}
		if err := openaiClient.init(); err != nil {
			glog.Line().Println("OpenaiClient init失败:" + err.Error())
		}
	})
	return openaiClient
}

// init 初始化
func (c *OpenaiClient) init() (err error) {
	return
}

// GetConfig 获取配置
func (c *OpenaiClient) GetConfig() (configData *entity.ConfigOpenai, err error) {
	configData = &entity.ConfigOpenai{}
	err = dao.ConfigOpenai.Where("1=1").Limit(1).Order("call_num ASC").Scan(configData)
	if err != nil {
		return nil, err
	}
	return configData, nil
}
