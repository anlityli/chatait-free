// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package snowflake

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"sync"
)

// IDCreator ID生成器
type IDCreator struct {
	Node *Node
}

var idCreatorMap sync.Map

func Instance(name string, params *NewParams) *IDCreator {
	if instance, ok := idCreatorMap.Load(name); ok {
		return instance.(*IDCreator)
	}

	// 创建新的单例对象
	instance := &IDCreator{}
	node, err := NewNode(params)
	if err != nil {
		glog.Line(true).Println(err.Error())
		return nil
	}
	instance.Node = node

	// 将新的单例对象存储到 sync.Map 中
	idCreatorMap.Store(name, instance)

	return instance
}

// GenerateID 生成
func GenerateID() (id int64) {
	workID := g.Config().GetInt64("distributedConf.workID")
	return Instance("default", &NewParams{
		Node:  workID,
		Epoch: EpochAit,
	}).Node.Generate().Int64()
}

// GenerateDiscordId 生成discord的id
func GenerateDiscordId() (id int64) {
	workID := g.Config().GetInt64("distributedConf.workID")
	return Instance("discord", &NewParams{
		Node:  workID,
		Epoch: EpochDiscord,
	}).Node.Generate().Int64()
}
