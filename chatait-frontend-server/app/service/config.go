// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package service

import (
	"database/sql"
	"github.com/anlityli/chatait-free/chatait-frontend-server/app/model/response"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/model/entity"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/helper"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var Config = &configService{}

type configService struct {
}

func (s *configService) WalletList(r *ghttp.Request) (re *response.ConfigWalletList, err error) {
	re = &response.ConfigWalletList{}
	err = dao.ConfigWallet.Where("1=1").Scan(re)
	if err != nil {
		return nil, err
	}
	return re, nil
}

func (s *configService) Options(r *ghttp.Request) (re map[string]interface{}, err error) {
	emailCodeEnable, err := helper.GetConfig("emailCodeEnable")
	if err != nil {
		return nil, err
	}
	re = g.Map{
		"emailCodeEnable": emailCodeEnable,
	}
	return re, nil
}

func (s *configService) PayList(r *ghttp.Request) (re *response.ConfigPayList, err error) {
	re = &response.ConfigPayList{}
	*re = make(response.ConfigPayList, 0)
	list := &[]*entity.ConfigPay{}
	err = dao.ConfigPay.Where("status=1").Scan(list)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if len(*list) > 0 {
		for _, item := range *list {
			reItem := &response.ConfigPay{}
			reItem.Id = item.Id
			reItem.ApiName = item.ApiName
			reItem.FrontendDescription = item.FrontendDescription
			reItem.PayChannel = make([]*response.ConfigPayChannelItem, 0)
			if item.PayChannel != "" {
				payChannelDecode, err := gjson.Decode(item.PayChannel)
				if err != nil {
					return nil, err
				}
				payChannelList := &[]*response.ConfigPayChannelItem{}
				err = gconv.Scan(payChannelDecode, payChannelList)
				if err != nil {
					return nil, err
				}
				for _, channelItem := range *payChannelList {
					if channelItem.Status == 1 {
						reItem.PayChannel = append(reItem.PayChannel, channelItem)
					}
				}
			}
			*re = append(*re, reItem)
		}
	}
	return re, nil
}
