// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package console

import (
	"fmt"
	"github.com/anlityli/chatait-free/chatait-backend-server/library/console/server"
	"github.com/gogf/gf/os/glog"
	"log"
	"sync"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcmd"
)

// Console 客户端
type Console struct {
	parser        *gcmd.Parser
	supportedArgs map[string]bool
}

var console *Console
var consoleOnce sync.Once

// Instance 单例
func Instance() *Console {
	consoleOnce.Do(func() {
		console = &Console{}
		if err := console.init(); err != nil {
			log.Fatalf("终端命名初始化失败:%s\n", err.Error())
		}
	})
	return console
}

// init 终端初始
func (c *Console) init() (err error) {
	c.supportedArgs = map[string]bool{
		"c,config":  true,
		"v,version": false,
	}
	simulationArgs := g.Config().GetStrings("backendConf.simulationArgs")
	if g.Config().GetBool("backendConf.debugStatus") && len(simulationArgs) > 0 {
		args := []string{"main"}
		args = append(args, simulationArgs...)
		c.parser, err = gcmd.ParseWithArgs(args, c.supportedArgs)
	} else {
		c.parser, err = gcmd.Parse(c.supportedArgs)
	}
	return
}

// GetParser 获取匹配对象
func (c *Console) GetParser() *gcmd.Parser {
	return c.parser
}

// Run 运行
func (c *Console) Run() {
	allArgs := c.parser.GetArgAll()
	allOpts := c.parser.GetOptAll()
	// 没有参数没有选项的情况，直接运行启动
	if len(allArgs) <= 1 && len(allOpts) == 0 {
		server.RunServer()
	} else {
		// 指定配置文件路径
		if c.parser.ContainsOpt("c") {
			if c.parser.GetOpt("c") != "" {
				if err := g.Config().SetPath(c.parser.GetOpt("c")); err != nil {
					glog.Line().Fatalf("指定配置文件失败:", err.Error())
					return
				}
				server.RunServer()
			}
			return
		}
		if c.parser.ContainsOpt("v") {
			fmt.Printf("v%s\r\n", "")
			return
		}
	}
}
