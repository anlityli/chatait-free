// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/anlityli/chatait-free/chatait-backend-server/app/service"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/notice"
	"github.com/gogf/gf/net/ghttp"
)

type Config struct {
}

func (c *Config) AllOption(r *ghttp.Request) {
	re, err := service.Config.AllOption(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, re)
}

func (c *Config) OptionEdit(r *ghttp.Request) {
	err := service.Config.OptionEdit(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, "操作成功")
}

func (c *Config) LevelList(r *ghttp.Request) {
	re, err := service.Config.LevelList(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, re)
}

func (c *Config) LevelEdit(r *ghttp.Request) {
	err := service.Config.LevelEdit(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, "操作成功")
}

func (c *Config) WalletList(r *ghttp.Request) {
	re, err := service.Config.WalletList(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, re)
}
func (c *Config) WalletEdit(r *ghttp.Request) {
	err := service.Config.WalletEdit(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, "操作成功")
}

func (c *Config) WalletOne(r *ghttp.Request) {
	re, err := service.Config.WalletOne(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, re)
}

func (c *Config) PayList(r *ghttp.Request) {
	re, err := service.Config.PayList(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, re)
}

func (c *Config) PayOne(r *ghttp.Request) {
	re, err := service.Config.PayOne(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, re)
}

func (c *Config) PayEdit(r *ghttp.Request) {
	err := service.Config.PayEdit(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, "操作成功")
}

func (c *Config) MidjourneyList(r *ghttp.Request) {
	re, err := service.Config.MidjourneyList(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, re)
}

func (c *Config) MidjourneyOne(r *ghttp.Request) {
	re, err := service.Config.MidjourneyOne(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, re)
}

func (c *Config) MidjourneyAdd(r *ghttp.Request) {
	err := service.Config.MidjourneyAdd(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, "操作成功")
}

func (c *Config) MidjourneyEdit(r *ghttp.Request) {
	err := service.Config.MidjourneyEdit(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, "操作成功")
}

func (c *Config) MidjourneyDelete(r *ghttp.Request) {
	err := service.Config.MidjourneyDelete(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, "操作成功")
}

func (c *Config) OpenaiList(r *ghttp.Request) {
	re, err := service.Config.OpenaiList(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, re)
}

func (c *Config) OpenaiOne(r *ghttp.Request) {
	re, err := service.Config.OpenaiOne(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, re)
}

func (c *Config) OpenaiAdd(r *ghttp.Request) {
	err := service.Config.OpenaiAdd(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, "操作成功")
}

func (c *Config) OpenaiEdit(r *ghttp.Request) {
	err := service.Config.OpenaiEdit(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, "操作成功")
}

func (c *Config) OpenaiDelete(r *ghttp.Request) {
	err := service.Config.OpenaiDelete(r)
	if err != nil {
		notice.Write(r, notice.OtherError, err.Error())
	}
	notice.Write(r, notice.NoError, "操作成功")
}
