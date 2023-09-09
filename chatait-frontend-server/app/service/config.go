// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package service

import (
	"github.com/anlityli/chatait-free/chatait-frontend-server/app/model/response"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/helper"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
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
