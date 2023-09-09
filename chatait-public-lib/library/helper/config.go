// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package helper

import (
	"errors"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/model/entity"
)

// GetConfig 获取配置的值
func GetConfig(configName string) (re string, err error) {
	configData := &entity.Config{}
	err = dao.Config.Where("config_name=?", configName).Scan(configData)
	if err != nil {
		return "", errors.New(configName + "配置参数不存在")
	}
	return configData.Value, nil
}

// GetConfigLevel 获取级别配置信息
func GetConfigLevel(levelId int) (re *entity.ConfigLevel, err error) {
	re = &entity.ConfigLevel{}
	err = dao.ConfigLevel.Where("id=?", levelId).Scan(re)
	if err != nil {
		return nil, err
	}
	return re, nil
}

// GetWalletName 获取钱包名称
func GetWalletName(field string) (re string) {
	walletData := &entity.ConfigWallet{}
	err := dao.ConfigWallet.Where("field=?", field).Scan(walletData)
	if err != nil {
		return ""
	} else {
		return walletData.WalletName
	}
}

func GetConfigMidjourney(id int64) (re *entity.ConfigMidjourney, err error) {
	re = &entity.ConfigMidjourney{}
	err = dao.ConfigMidjourney.Where("id=?", id).Scan(re)
	if err != nil {
		return nil, err
	}
	return re, nil
}
