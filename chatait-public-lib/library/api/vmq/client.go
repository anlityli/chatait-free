// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package vmq

import (
	"github.com/anlityli/chatait-free/chatait-public-lib/app/libservice"
	"github.com/gogf/gf/os/glog"
	"sync"
)

// VmqClient VmqClient
type VmqClient struct {
	config *libservice.ConfigPay
	apiKey string
	host   string
}

var vmqClient *VmqClient
var vmqClientOnce sync.Once

// Instance 单例
func Instance() *VmqClient {
	vmqClientOnce.Do(func() {
		vmqClient = &VmqClient{}
		if err := vmqClient.init(); err != nil {
			glog.Line().Println("vmqClient init失败:" + err.Error())
		}
	})
	return vmqClient
}

// init 初始化
func (c *VmqClient) init() (err error) {
	c.config, err = libservice.Pay.OneConfigPay(1)
	if err != nil {
		return err
	}
	for _, item := range c.config.Params {
		if item.Param == "apiKey" {
			c.apiKey = item.Value
		} else if item.Param == "host" {
			lastStr := item.Value[len(item.Value)-1:]
			if lastStr != "/" {
				item.Value = item.Value + "/"
			}
			c.host = item.Value
		}
	}
	return
}

// GetApiKey 获取通讯密钥
func (c *VmqClient) GetApiKey() string {
	return c.apiKey
}

// GetHost 获取host
func (c *VmqClient) GetHost() string {
	return c.host
}
