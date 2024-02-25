// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package baidu

import (
	"database/sql"
	"errors"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/helper"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/xtime"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"sync"
)

// Client 客户端
type Client struct {
	configMap sync.Map // 存config的map
}

var client *Client
var clientOnce sync.Once

// Instance 单例
func Instance() *Client {
	clientOnce.Do(func() {
		client = &Client{}
		if err := client.init(); err != nil {
			glog.Line().Println("百度init失败:" + err.Error())
		}
	})
	return client
}

// init 初始化
func (c *Client) init() (err error) {
	return nil
}

// GetConfig 获取GetConfig
func (c *Client) GetConfig(feature string) (configData *Config, err error) {
	configData = &Config{}
	err = dao.ConfigBaidu.Where("status=1 AND ( features='[]' OR features='' OR features LIKE ?)", "%"+feature+"%").Limit(1).Order("call_num ASC").Scan(configData)
	if err != nil && err != sql.ErrNoRows {
		glog.Line(true).Println(configData, err)
		return nil, err
	}
	if configData.Id <= 0 {
		glog.Line(true).Println("为获取到百度配置信息")
		return nil, nil
	}
	err = c.getAccessToken(configData)
	if err != nil {
		glog.Line(true).Println(configData, err)
		return nil, err
	}
	return configData, nil
}

func (c *Client) getAccessToken(configData *Config) (err error) {
	storeConfig, ok := c.configMap.Load(configData.Id)
	if ok {
		storeConfigData := storeConfig.(*Config)
		if xtime.GetNowTime() < storeConfigData.AccessTokenExpireIn && storeConfigData.ApiKey == configData.ApiKey && storeConfigData.SecretKey == configData.SecretKey && configData.AccessToken != "" {
			glog.Line(true).Debug("百度accessToken未过期，直接返回")
			return nil
		}
		c.configMap.Delete(configData.Id)
	}
	// 获取新的AccessToken
	err = c.requestAccessToken(configData)
	if err != nil {
		glog.Line(true).Println(configData, err)
		return err
	}
	c.configMap.Store(configData.Id, configData)
	return nil
}

func (c *Client) requestAccessToken(configData *Config) (err error) {
	httpClient := ghttp.NewClient()
	requestData := g.Map{
		"grant_type":    "client_credentials",
		"client_id":     configData.ApiKey,
		"client_secret": configData.SecretKey,
	}
	resp, err := httpClient.Post(AccessTokenUrl, requestData)
	if err != nil {
		glog.Line(true).Println(AccessTokenUrl, requestData, configData, err)
		return err
	}
	reString := resp.ReadAllString()
	reMap, err := helper.JSONToMap(reString)
	if err != nil {
		glog.Line(true).Println(AccessTokenUrl, requestData, reString, configData, err)
		return err
	}
	if _, ok := reMap["error"]; ok {
		glog.Line(true).Println(AccessTokenUrl, requestData, reString, configData, err)
		return errors.New(gconv.String(reMap["error"]) + gconv.String(reMap["error_description"]))
	}
	configData.AccessToken = gconv.String(reMap["access_token"])
	configData.AccessTokenExpireIn = xtime.GetNowTime() + gconv.Int64(reMap["expires_in"]) - 30
	return nil
}
