// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package datalist

import (
	"github.com/anlityli/chatait-free/chatait-public-lib/library/helper"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
)

// FilterWhereAndParams 筛选用到的where 和 params
func FilterWhereAndParams(r *ghttp.Request, columns *Columns) (re *FilterWhereParam, err error) {
	columnList := columns.ColumnList
	where := "1=1"
	params := g.Slice{}
	getData := r.GetMap()
	// 循环出所有可以筛选的列
	for _, column := range columnList {
		if !column.CanFilter {
			continue
		}
		getValueInterface, ok := getData[column.Field]
		if !ok {
			continue
		}
		getValue := gstr.Trim(gconv.String(getValueInterface), ", \t\n\r")
		if gstr.Pos(getValue, "|") > 0 {
			where += " AND ("
			childValueArr := gstr.Explode("|", getValue)
			for k, value := range childValueArr {
				relation := ""
				if k != 0 {
					relation = "AND"
				}
				whereAndParams, err := getWhereAndParams(value, column.FilterField, relation)
				if err != nil {
					return nil, err
				}
				where += whereAndParams.Where
				params = append(params, whereAndParams.Params...)
			}
			where += ")"
		} else {
			whereAndParams, err := getWhereAndParams(getValue, column.FilterField, "AND")
			if err != nil {
				return nil, err
			}
			where += whereAndParams.Where
			params = append(params, whereAndParams.Params...)
		}
	}
	re = &FilterWhereParam{
		Where:  where,
		Params: params,
	}
	return re, nil
}

func getWhereAndParams(getValue string, tableField string, relation string) (re *FilterWhereParam, err error) {
	where := ""
	params := g.Slice{}
	isDate := false
	filterModel := ""
	getSymbol := ""
	bindValue := ""
	areaBindValue := g.Slice{}
	if gstr.Pos(getValue, ",") > 0 {
		getValueArr := gstr.Explode(",", getValue)
		getSymbol = gstr.ToUpper(getValueArr[0])
		if getSymbol == "IN" {
			bindValueArr := getValueArr
			bindValueArr = bindValueArr[1:]
			bindValue = gstr.Implode("','", bindValueArr)
			bindValue = "'" + bindValue + "'"
		} else {
			bindValue = getValueArr[1]
			filterModel = getValueArr[len(getValueArr)-1]
			// 日期筛选
			if filterModel == "date" {
				valueTime, err := gtime.StrToTime(bindValue)
				if err != nil {
					return nil, err
				}
				isDate = true
				bindValue = gconv.String(valueTime.Timestamp())
			} else if filterModel == "area" {
				getValueArrEnd := len(getValueArr) - 1
				tempArr := getValueArr[1:getValueArrEnd]
				for _, tempItem := range tempArr {
					areaBindValue = append(areaBindValue, tempItem)
				}
			} else if filterModel == "amount" || filterModel == "money" {
				bindValue = gconv.String(helper.YuanToCent(gconv.Int(bindValue)))
			}
		}
	} else {
		getSymbol = "="
		bindValue = getValue
	}
	if getSymbol == "LIKE" {
		where += " " + relation + " INSTR(" + tableField + ",?" + ")>0"
		params = append(params, bindValue)
	} else if getSymbol == "NOTLIKE" {
		where += " " + relation + " INSTR(" + tableField + ",?" + ")=0"
		params = append(params, bindValue)
	} else if getSymbol == "IN" {
		where += " " + relation + " " + tableField + " IN (" + bindValue + ")"
	} else {
		if isDate && getSymbol == "=" {
			where += " " + relation + " " + tableField + " >= ?"
			where += " AND " + tableField + " <= ?"
			params = append(params, bindValue, gconv.Int(bindValue)+86399)
		} else if isDate && getSymbol == "<=" {
			where += " " + relation + " " + tableField + " <= ?"
			params = append(params, gconv.Int(bindValue)+86399)
		} else if isDate && getSymbol == ">" {
			where += " " + relation + " " + tableField + " > ?"
			params = append(params, gconv.Int(bindValue)+86399)
		} else if filterModel == "area" {
			tableAliasArr := gstr.Explode(".", tableField)
			tableAlias := ""
			if len(tableAliasArr) > 1 {
				tableAlias = tableAliasArr[0] + "."
			}
			if len(areaBindValue) == 1 {
				where += " " + relation + " " + tableAlias + "province_code=?"
			} else if len(areaBindValue) == 2 {
				where += " " + relation + " (" + tableAlias + "province_code=?" + " AND " + tableAlias + "city_code=?)"
			} else if len(areaBindValue) == 3 {
				where += " " + relation + " (" + tableAlias + "province_code=?" + " AND " + tableAlias + "city_code=? AND " + tableAlias + "county_code=?)"
			}
			params = append(params, areaBindValue...)
		} else {
			where += " " + relation + " " + tableField + getSymbol + "?"
			params = append(params, bindValue)
		}
	}
	re = &FilterWhereParam{
		Where:  where,
		Params: params,
	}
	return re, nil
}

// YesNoFilterType 是或否的筛选条件
func YesNoFilterType(yesNoLabel ...string) (re *FilterType) {
	yesLabel := "是"
	noLabel := "否"
	if len(yesNoLabel) > 0 && yesNoLabel[0] != "" {
		yesLabel = yesNoLabel[0]
	}
	if len(yesNoLabel) > 1 && yesNoLabel[1] != "" {
		noLabel = yesNoLabel[1]
	}
	// 获取全部会员级别
	selectData := g.Slice{
		g.Map{
			"label": yesLabel,
			"value": 1,
		},
		g.Map{
			"label": noLabel,
			"value": 0,
		},
	}
	return &FilterType{
		Attr:       "select",
		SelectData: selectData,
	}
}

// YesNoValue 是或否的中文值
func YesNoValue(value int, yesNoLabel ...string) string {
	yesLabel := "是"
	noLabel := "否"
	if len(yesNoLabel) > 0 && yesNoLabel[0] != "" {
		yesLabel = yesNoLabel[0]
	}
	if len(yesNoLabel) > 1 && yesNoLabel[1] != "" {
		noLabel = yesNoLabel[1]
	}
	if value == 0 {
		return noLabel
	} else {
		return yesLabel
	}
}
