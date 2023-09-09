// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package datalist

import "github.com/gogf/gf/frame/g"

// ColumnItem 列
type ColumnItem struct {
	Field             string                                `json:"field"`
	FieldName         string                                `json:"field_name"`
	FieldAttr         *FieldAttr                            `json:"field_attr"`
	Hidden            bool                                  `json:"-"`
	ValueCallBack     func(rowData interface{}) string      `json:"-"`
	OriValueCallBack  func(rowData interface{}) interface{} `json:"-"`
	ValueAttrCallBack func(rowData interface{}) string      `json:"-"`
	CanFilter         bool                                  `json:"-"`
	FilterType        *FilterType                           `json:"-"`
	FilterField       string                                `json:"-"`
}

type FilterType struct {
	Attr       string  `json:"attr"`
	SelectData g.Slice `json:"attr_data"`
}

// ColumnList 列的表
type ColumnList []*ColumnItem

// Columns 列表信息
type Columns struct {
	ListName       string `json:"list_name"`
	ListID         string `json:"list_id"`
	RelevantListID string `json:"relevant_list_id"` // 此字段主要用于一个列表方法用在多处路由上，防止没有权限的情况出现，具体用法可以参照奖金流水列表
	ColumnList     ColumnList
}

// FieldAttr 字段的属性
type FieldAttr struct {
	Width int `json:"width"`
}

// Result 返回的结果
type Result struct {
	Columns     ColumnList               `json:"columns"`
	List        ResultList               `json:"list"`
	ListID      string                   `json:"list_id"`
	TotalCount  int                      `json:"total_count"`
	Page        int                      `json:"page"`
	PageSize    int                      `json:"page_size"`
	TotalPage   int                      `json:"total_page"`
	PageCount   int                      `json:"page_count"`
	FilterTypes []*ResultFilterTypesItem `json:"filter_types"`
}

type ResultFilterTypesItem struct {
	Field     string  `json:"field"`
	FieldName string  `json:"field_name"`
	Attr      string  `json:"attr"`
	AttrData  g.Slice `json:"attr_data"`
}

// ResultList 结果列表
type ResultList []map[string]*ResultListValue

// ResultListValue 结果列表集
type ResultListValue struct {
	Value     string      `json:"value"`
	OriValue  interface{} `json:"ori_value"`
	ValueAttr interface{} `json:"value_attr"`
}

// FilterWhereParam 筛选参数
type FilterWhereParam struct {
	Where  string
	Params g.Slice
}

// ExportFile 导出文件
type ExportFile struct {
	Path     string `json:"path"`
	FileSize string `json:"file_size"`
}
