// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package page

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// DefaultPageSize 默认分页大小
const DefaultPageSize = 10

// TotalPage 总页数
func TotalPage(totalCount int, pageSize int) int {
	if totalCount%pageSize == 0 {
		return gconv.Int(totalCount / pageSize)
	}
	return gconv.Int(totalCount/pageSize) + 1
}

// Data 以model形式返回分页形式数据
func Data(r *ghttp.Request, param *Param, pointStruct interface{}) (pageRe *Response, err error) {
	pageRe = &Response{}
	// 分页处理
	tempPage := r.GetInt("page")
	tempPageSize := r.GetInt("page_size")
	if tempPage > 0 {
		param.Page = tempPage
	}
	if tempPageSize > 0 {
		param.PageSize = tempPageSize
	}
	if param.Page == 0 {
		param.Page = 1
	}
	if param.PageSize == 0 {
		param.PageSize = DefaultPageSize
	}
	pageRe.Page = param.Page
	pageRe.PageSize = param.PageSize

	model := g.DB().Model(param.TableName)
	// join处理
	if param.Join != nil {
		for _, v := range param.Join {
			if v.JoinType == "leftJoin" {
				model = model.LeftJoin(v.JoinTable, v.On)
			} else if v.JoinType == "rightJoin" {
				model = model.RightJoin(v.JoinTable, v.On)
			} else if v.JoinType == "innerJoin" {
				model = model.InnerJoin(v.JoinTable, v.On)
			}
		}
	}
	// order by处理
	if param.OrderBy != "" {
		model = model.Order(param.OrderBy)
	}
	// 查询条件
	if param.Where == "" {
		model = model.Where("1=1")
	} else {
		model = model.Where(param.Where, param.WhereParams)
	}
	// 计算总条数
	modelClone := model.Clone()
	pageRe.TotalCount, err = modelClone.Count()
	if err != nil {
		return nil, err
	}
	// 筛选字段
	if param.Field != "" {
		model = model.Fields(param.Field)
	}
	// 总页数
	pageRe.TotalPage = TotalPage(pageRe.TotalCount, pageRe.PageSize)
	// 分页查询
	listData, err := model.Page(pageRe.Page, pageRe.PageSize).All()
	if err != nil {
		return nil, err
	}
	pageRe.PageCount = len(listData)
	if listData.IsEmpty() {
		pageRe.ListData = g.Slice{}
	} else {
		if err := gconv.SliceStruct(listData, pointStruct); err != nil {
			return nil, err
		}
		pageRe.ListData = pointStruct
	}
	return pageRe, nil
}
