// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package serv

import (
	"github.com/anlityli/chatait-free/chatait-frontend-server/app/service"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/libservice"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/xtime"
	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/os/glog"
)

// OnRun 运行时执行
func OnRun() {
	// 执行Bot的监听器
	go func() {
		err := service.ConversationMidjourney.Listener()
		if err != nil {
			glog.Line(true).Println("监听器启动失败", err)
		}
	}()
	// 每天给高级会员发放余额
	go func() {
		_, err := gcron.Add("0 0 2 * * *", func() {
			libservice.HandOutExecTime = xtime.GetNowTime()
			libservice.User.DailyMemberHandler()
		})
		if err != nil {
			glog.Line(true).Println("每日给高级会员发放余额定时器启动失败", err)
		}
	}()
}
