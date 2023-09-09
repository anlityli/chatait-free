// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package menu

type ItemModel struct {
	Key               string    `json:"key"`
	Title             string    `json:"title"`
	RoutePath         string    `json:"route_path"`
	RelevantRoutePath []string  `json:"relevant_route_path"` // route_path没有权限的话，关联的路由也没有权限
	Show              bool      `json:"show"`
	Leaf              bool      `json:"leaf"`
	Children          ListModel `json:"children"`
}

type ListModel []*ItemModel
