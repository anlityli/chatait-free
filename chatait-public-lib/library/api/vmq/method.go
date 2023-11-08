// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package vmq

import (
	"errors"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/helper"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/security"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/xtime"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
)

// CreateOrder 创建订单
func CreateOrder(params *CreateOrderParams) (re *CreateOrderResponse, err error) {
	nowTime := xtime.GetNowTime()
	// 先校验当前支付接口状态是否正常
	state, err := GetState(&GetStateParams{
		R: params.R,
	})
	if err != nil {
		return nil, err
	}
	if nowTime-state.LastHeart > 20*60 {
		glog.Line(true).Println("支付接口状态异常", params, state)
		return nil, errors.New("支付接口状态异常")
	}
	if state.JKState != 1 {
		glog.Line(true).Println("支付接口状态异常", params, state)
		return nil, errors.New("支付接口状态异常")
	}
	// 开始请求支付
	hostUrl, err := security.HostUrl(params.R)
	if err != nil {
		return nil, err
	}
	notifyUrl := hostUrl + "/notify/vmq"
	requestData := g.Map{
		"payId":     params.PayFlowId,
		"type":      params.PayChannel,
		"price":     helper.CentToYuan(params.OrderAmount),
		"sign":      "",
		"param":     params.Param,
		"isHtml":    0,
		"notifyUrl": notifyUrl,
	}
	sign, err := GetSign(requestData)
	if err != nil {
		return nil, err
	}
	requestData["sign"] = sign
	httpClient := ghttp.NewClient()
	response, err := httpClient.Post(Instance().GetHost()+"createOrder", requestData)
	if err != nil {
		return nil, err
	}
	defer response.Close()
	reString := response.ReadAllString()
	reMap, err := helper.JSONToMap(reString)
	if err != nil {
		glog.Line(true).Println(params, reString, err)
		return nil, err
	}
	glog.Line(true).Debug(reMap)
	if gconv.Int(reMap["code"]) != 1 {
		glog.Line(true).Println(params, reMap)
		return nil, errors.New(gconv.String(reMap["msg"]))
	}
	re = &CreateOrderResponse{}
	err = gconv.Scan(reMap["data"], re)
	if err != nil {
		glog.Line(true).Println(params, reMap)
		return nil, err
	}
	re.ResponseData = reString
	return re, nil
}

// GetState 获取当前支付接口状态
func GetState(params *GetStateParams) (re *GetStateResponse, err error) {
	nowTime := xtime.GetNowTime()
	requestData := g.Map{
		"t":    nowTime,
		"sign": "",
	}
	glog.Line(true).Debug(gconv.String(nowTime) + Instance().GetApiKey())
	sign, err := gmd5.EncryptString(gconv.String(nowTime) + Instance().GetApiKey())
	if err != nil {
		return nil, err
	}
	requestData["sign"] = sign
	glog.Line(true).Debug(Instance().GetHost() + "getState")
	glog.Line(true).Debug(requestData)
	httpClient := ghttp.NewClient()
	response, err := httpClient.Post(Instance().GetHost()+"getState", requestData)
	if err != nil {
		return nil, err
	}
	defer response.Close()
	reString := response.ReadAllString()
	reMap, err := helper.JSONToMap(reString)
	if err != nil {
		glog.Line(true).Println(params, reString, err)
		return nil, err
	}
	glog.Line(true).Debug(reMap)
	if gconv.Int(reMap["code"]) != 1 {
		glog.Line(true).Println(params, reMap)
		return nil, errors.New(gconv.String(reMap["msg"]))
	}
	re = &GetStateResponse{}
	err = gconv.Scan(reMap["data"], re)
	if err != nil {
		glog.Line(true).Println(params, reMap)
		return nil, err
	}
	re.ResponseData = reString
	glog.Line(true).Println(reString)
	return re, nil
}

// GetSign 拿到签名
func GetSign(data interface{}) (sign string, err error) {
	dataMap := gconv.Map(data)
	if len(dataMap) <= 0 {
		return "", errors.New("验签失败: 验签参数不正确")
	}
	str := ""
	if payId, ok := dataMap["payId"]; ok {
		str += gconv.String(payId)
	}
	if param, ok := dataMap["param"]; ok {
		str += gconv.String(param)
	}
	if vType, ok := dataMap["type"]; ok {
		str += gconv.String(vType)
	}
	if price, ok := dataMap["price"]; ok {
		str += gconv.String(price)
	}
	if reallyPrice, ok := dataMap["reallyPrice"]; ok {
		str += gconv.String(reallyPrice)
	}
	apiKey := Instance().GetApiKey()
	str += apiKey
	return gmd5.EncryptString(str)
}

// ValidateSign 验签
func ValidateSign(data interface{}) (err error) {
	dataMap := gconv.Map(data)
	if len(dataMap) <= 0 {
		return errors.New("验签失败: 验签参数不正确")
	}
	sign, err := GetSign(data)
	if err != nil {
		return err
	}
	rSign, ok := dataMap["sign"]
	if !ok {
		return errors.New("验签失败: 缺少必要参数")
	}
	if sign != gconv.String(rSign) {
		return errors.New("验签失败: 签名不匹配")
	}
	return nil
}
