// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package trans

import (
	"errors"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/api/baidu"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
)

// Text 文本翻译
func Text(params *baidu.TransTextParams) (re *baidu.TransTextResponse, err error) {
	configData, err := baidu.Instance().GetConfig()
	if err != nil {
		glog.Line(true).Println(params, err)
		return nil, err
	}
	// 配置为空，则不翻译
	if configData == nil {
		return nil, nil
	}
	// 调用接口次数增加
	if _, err = dao.ConfigBaidu.Data(g.Map{
		"call_num": gdb.Raw("call_num+1"),
	}).Where("id=?", configData.Id).Update(); err != nil {
		glog.Line(true).Println(params, err)
		return nil, err
	}
	httpClient := ghttp.NewClient()
	httpClient.SetHeader("Content-Type", "application/json;charset=utf-8")
	requestData := g.Map{
		"from": params.From,
		"to":   params.To,
		"q":    params.Q,
	}
	requestDataJson, err := gjson.Encode(requestData)
	if err != nil {
		glog.Line(true).Println(params, requestData, err)
		return nil, err
	}
	response, err := httpClient.Post(baidu.TransTextURL+"?access_token="+configData.AccessToken, requestDataJson)
	if err != nil {
		glog.Line(true).Println(baidu.TransTextURL+"?access_token="+configData.AccessToken, requestData, err)
		return
	}
	defer response.Close()
	reString := response.ReadAllString()
	reJson, err := gjson.Decode(reString)
	if err != nil {
		glog.Line(true).Println(baidu.TransTextURL+"?access_token="+configData.AccessToken, requestData, reString, err)
		return nil, err
	}
	re = &baidu.TransTextResponse{}
	err = gconv.Scan(reJson, re)
	if err != nil {
		glog.Line(true).Println(baidu.TransTextURL+"?access_token="+configData.AccessToken, requestData, reString, err)
		return nil, err
	}
	if re.ErrorCode != "" && re.ErrorCode != "0" {
		glog.Line(true).Println(baidu.TransTextURL+"?access_token="+configData.AccessToken, requestData, reString)
		return nil, errors.New(re.ErrorCode + "" + re.ErrorMsg)
	}
	return re, nil
}
