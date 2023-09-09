// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package libservice

import (
	"database/sql"
	"errors"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/model/entity"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
)

var Pay = &payService{}

type payService struct {
}

type ConfigPay struct {
	Id                  int                     `json:"id"`
	ApiName             string                  `json:"api_name"`
	Params              []*ConfigPayParamsItem  `json:"params"`
	PayChannel          []*ConfigPayChannelItem `json:"pay_channel"`
	FrontendDescription string                  `json:"frontend_description"`
	BackendDescription  string                  `json:"backend_description"`
	Status              int                     `json:"status"`
	CreatedAt           int                     `json:"created_at"`
	UpdatedAt           int                     `json:"updated_at"`
}

// ConfigPayParamsItem 配置参数
type ConfigPayParamsItem struct {
	Param     string `json:"param"`
	ParamName string `json:"param_name"`
	Value     string `json:"value"`
}

// ConfigPayChannelItem 渠道类型
type ConfigPayChannelItem struct {
	Id          int    `json:"id"`
	Channel     string `json:"channel"`
	ChannelName string `json:"channel_name"`
	Status      int    `json:"status"`
}

// AllConfigPay 所有支付方式
func (s *payService) AllConfigPay() (re *[]*ConfigPay, err error) {
	payList := &[]*entity.ConfigPay{}
	err = dao.ConfigPay.Where("status=1").Scan(payList)
	if err != nil {
		return nil, err
	}
	*re = make([]*ConfigPay, 0)
	for _, item := range *payList {
		tempItem := &ConfigPay{
			Id:                  item.Id,
			ApiName:             item.ApiName,
			FrontendDescription: item.FrontendDescription,
			BackendDescription:  item.BackendDescription,
			Status:              item.Status,
			CreatedAt:           item.CreatedAt,
			UpdatedAt:           item.UpdatedAt,
		}
		if item.Params != "" {
			params, err := gjson.Decode(item.Params)
			if err != nil {
				return nil, err
			}
			err = gconv.Scan(params, &tempItem.Params)
			if err != nil {
				return nil, err
			}
		}
		if item.PayChannel != "" {
			payChannel, err := gjson.Decode(item.PayChannel)
			if err != nil {
				return nil, err
			}
			err = gconv.Scan(payChannel, &tempItem.PayChannel)
			if err != nil {
				return nil, err
			}
		}
		*re = append(*re, tempItem)
	}
	return re, nil
}

// OneConfigPay 获取一条支付方式
func (s *payService) OneConfigPay(id int) (re *ConfigPay, err error) {
	payData := &entity.ConfigPay{}
	err = dao.ConfigPay.Where("id=?", id).Scan(payData)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if payData.Id <= 0 {
		return nil, errors.New("支付方式不不存在")
	}
	re = &ConfigPay{
		Id:                  payData.Id,
		ApiName:             payData.ApiName,
		FrontendDescription: payData.FrontendDescription,
		BackendDescription:  payData.BackendDescription,
		Status:              payData.Status,
		CreatedAt:           payData.CreatedAt,
		UpdatedAt:           payData.UpdatedAt,
	}
	if payData.Params != "" {
		params, err := gjson.Decode(payData.Params)
		if err != nil {
			return nil, err
		}
		err = gconv.Scan(params, &re.Params)
		if err != nil {
			return nil, err
		}
	}
	if payData.PayChannel != "" {
		payChannel, err := gjson.Decode(payData.PayChannel)
		if err != nil {
			return nil, err
		}
		err = gconv.Scan(payChannel, &re.PayChannel)
		if err != nil {
			return nil, err
		}
	}
	glog.Line().Debug(re.PayChannel)
	return re, nil
}

// OneConfigPayParamValue 一个支付方式的参数的值
func (s *payService) OneConfigPayParamValue(id int, param string) (re string, err error) {
	configPay, err := Pay.OneConfigPay(id)
	if err != nil {
		return "", err
	}
	configPayParams := configPay.Params
	for _, configPayParamsItem := range configPayParams {
		if configPayParamsItem.Param == param {
			re = configPayParamsItem.Value
			return re, nil
		}
	}
	return "", errors.New("参数不存在")
}
