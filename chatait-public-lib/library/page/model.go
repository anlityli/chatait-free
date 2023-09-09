// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package page

import "github.com/gogf/gf/frame/g"

type Response struct {
	ListData   interface{} `json:"list_data"`
	TotalCount int         `json:"total_count"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalPage  int         `json:"total_page"`
	PageCount  int         `json:"page_count"`
}

type Param struct {
	Page        int
	PageSize    int
	TableName   string
	Where       string
	WhereParams g.Slice
	Join        ParamJoin
	OrderBy     string
	Field       string
}

type ParamJoin []*ParamJoinItem

type ParamJoinItem struct {
	JoinType  string
	JoinTable string
	On        string
}
