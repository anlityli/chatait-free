// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package service

import (
	"github.com/anlityli/chatait-free/chatait-backend-server/app/model/request"
	"github.com/anlityli/chatait-free/chatait-backend-server/app/model/response"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/constant"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/xtime"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var Dashboard = &dashboardService{}

type dashboardService struct {
}

// UserStatistic 用户统计
func (s *dashboardService) UserStatistic(r *ghttp.Request) (re *response.DashboardStatisticCount, err error) {
	requestModel := &request.DashboardStatisticCount{}
	if err := r.Parse(requestModel); err != nil {
		return nil, err
	}
	re = &response.DashboardStatisticCount{}
	statisticWhere := s.statisticWhere(requestModel)
	re.ThisCount, err = dao.User.Where(statisticWhere.thisWhere, statisticWhere.thisParams).Count()
	if err != nil {
		return nil, err
	}
	re.LastCount, err = dao.User.Where(statisticWhere.lastWhere, statisticWhere.lastParams).Count()
	if err != nil {
		return nil, err
	}
	return re, nil
}

// OrderStatistic 订单统计
func (s *dashboardService) OrderStatistic(r *ghttp.Request) (re *response.DashboardStatisticCount, err error) {
	requestModel := &request.DashboardStatisticCount{}
	if err := r.Parse(requestModel); err != nil {
		return nil, err
	}
	re = &response.DashboardStatisticCount{}
	statisticWhere := s.statisticWhere(requestModel)
	re.ThisCount, err = dao.ShopOrder.Where(statisticWhere.thisWhere, statisticWhere.thisParams).Count()
	if err != nil {
		return nil, err
	}
	re.LastCount, err = dao.ShopOrder.Where(statisticWhere.lastWhere, statisticWhere.lastParams).Count()
	if err != nil {
		return nil, err
	}
	return re, nil
}

// AmountStatistic 金额统计
func (s *dashboardService) AmountStatistic(r *ghttp.Request) (re *response.DashboardStatisticCount, err error) {
	requestModel := &request.DashboardStatisticCount{}
	if err := r.Parse(requestModel); err != nil {
		return nil, err
	}
	re = &response.DashboardStatisticCount{}
	statisticWhere := s.statisticWhere(requestModel)
	statisticWhere.thisWhere += " AND status=?"
	statisticWhere.thisParams = append(statisticWhere.thisParams, constant.ShopOrderStatusFinish)
	thisOneSum, err := dao.ShopOrder.Where(statisticWhere.thisWhere, statisticWhere.thisParams).Fields("SUM(pay_amount) pay_amount_sum").One()
	if err != nil {
		return nil, err
	}
	statisticWhere.lastWhere += " AND status=?"
	statisticWhere.lastParams = append(statisticWhere.lastParams, constant.ShopOrderStatusFinish)
	lastOneSum, err := dao.ShopOrder.Where(statisticWhere.lastWhere, statisticWhere.lastParams).Fields("SUM(pay_amount) pay_amount_sum").One()
	if err != nil {
		return nil, err
	}
	re.ThisCount = gconv.Int(thisOneSum["pay_amount_sum"])
	re.LastCount = gconv.Int(lastOneSum["pay_amount_sum"])
	return re, nil
}

// ConversationStatistic 总提问数量统计
func (s *dashboardService) ConversationStatistic(r *ghttp.Request) (re *response.DashboardStatisticCount, err error) {
	requestModel := &request.DashboardStatisticCount{}
	if err := r.Parse(requestModel); err != nil {
		return nil, err
	}
	re = &response.DashboardStatisticCount{}
	statisticWhere := s.statisticWhere(requestModel)
	re.ThisCount, err = dao.Conversation.Where(statisticWhere.thisWhere, statisticWhere.thisParams).Count()
	if err != nil {
		return nil, err
	}
	re.ThisCount = re.ThisCount / 2
	re.LastCount, err = dao.Conversation.Where(statisticWhere.lastWhere, statisticWhere.lastParams).Count()
	if err != nil {
		return nil, err
	}
	re.LastCount = re.LastCount / 2
	return re, nil
}

type statisticWhereRe struct {
	thisWhere  string
	thisParams g.Slice
	lastWhere  string
	lastParams g.Slice
}

func (s *dashboardService) statisticWhere(requestModel *request.DashboardStatisticCount) (re *statisticWhereRe) {
	re = &statisticWhereRe{}
	re.thisWhere = ""
	re.thisParams = g.Slice{}
	re.lastWhere = ""
	re.lastParams = g.Slice{}
	if requestModel.Type == constant.DashboardStatisticIntervalDaily {
		thisBegin := xtime.GetTodayBegin()
		thisEnd := xtime.GetTodayEnd()
		re.thisWhere = "created_at>? AND created_at<?"
		re.thisParams = g.Slice{thisBegin, thisEnd}
		lastBegin := xtime.GetYesterdayBegin()
		lastEnd := xtime.GetYesterdayEnd()
		re.lastWhere = "created_at>? AND created_at<?"
		re.lastParams = g.Slice{lastBegin, lastEnd}
	} else if requestModel.Type == constant.DashboardStatisticIntervalWeekly {
		thisBegin := xtime.GetWeekBegin()
		thisEnd := xtime.GetWeekEnd()
		re.thisWhere = "created_at>? AND created_at<?"
		re.thisParams = g.Slice{thisBegin, thisEnd}
		lastBegin := xtime.GetLastWeekBegin()
		lastEnd := xtime.GetLastWeekEnd()
		re.lastWhere = "created_at>? AND created_at<?"
		re.lastParams = g.Slice{lastBegin, lastEnd}
	} else if requestModel.Type == constant.DashboardStatisticIntervalMonthly {
		thisBegin := xtime.GetMonthBegin()
		thisEnd := xtime.GetMonthEnd()
		re.thisWhere = "created_at>? AND created_at<?"
		re.thisParams = g.Slice{thisBegin, thisEnd}
		lastBegin := xtime.GetLastMonthBegin()
		lastEnd := xtime.GetLastMonthEnd()
		re.lastWhere = "created_at>? AND created_at<?"
		re.lastParams = g.Slice{lastBegin, lastEnd}
	}
	return re
}
