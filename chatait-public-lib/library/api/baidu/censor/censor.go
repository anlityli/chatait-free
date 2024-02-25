package censor

import (
	"errors"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/constant"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/api/baidu"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
)

func Text(params *baidu.CensorTextParams) (re *baidu.CensorTextResponse, err error) {
	configData, err := baidu.Instance().GetConfig(constant.ConfigBaiduFeatureCensor)
	if err != nil {
		glog.Line(true).Println(params, err)
		return nil, err
	}
	// 配置为空，则不审核
	if configData == nil {
		re = &baidu.CensorTextResponse{
			ConclusionType: 1,
		}
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
		"text": params.Text,
	}
	requestDataJson, err := gjson.Encode(requestData)
	if err != nil {
		glog.Line(true).Println(params, requestData, err)
		return nil, err
	}
	response, err := httpClient.Post(baidu.CensorTextURL+"?access_token="+configData.AccessToken, requestDataJson)
	if err != nil {
		glog.Line(true).Println(baidu.CensorTextURL+"?access_token="+configData.AccessToken, requestData, err)
		return
	}
	defer response.Close()
	reString := response.ReadAllString()
	reJson, err := gjson.Decode(reString)
	if err != nil {
		glog.Line(true).Println(baidu.CensorTextURL+"?access_token="+configData.AccessToken, requestData, reString, err)
		return nil, err
	}
	re = &baidu.CensorTextResponse{}
	err = gconv.Scan(reJson, re)
	if err != nil {
		glog.Line(true).Println(baidu.CensorTextURL+"?access_token="+configData.AccessToken, requestData, reString, err)
		return nil, err
	}
	if re.ErrorCode != "" && re.ErrorCode != "0" {
		glog.Line(true).Println(baidu.CensorTextURL+"?access_token="+configData.AccessToken, requestData, reString)
		return nil, errors.New(re.ErrorCode + "" + re.ErrorMsg)
	}
	return re, nil
}
