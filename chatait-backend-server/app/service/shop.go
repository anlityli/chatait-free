// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/anlityli/chatait-free/chatait-backend-server/app/model/request"
	"github.com/anlityli/chatait-free/chatait-backend-server/app/model/response"
	"github.com/anlityli/chatait-free/chatait-backend-server/app/service/column"
	"github.com/anlityli/chatait-free/chatait-backend-server/library/datalist"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/constant"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/libservice"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/model/entity"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/page"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/snowflake"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/xtime"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var Shop = &shopService{}

type shopService struct {
}

// GoodsList 商品列表
func (s *shopService) GoodsList(r *ghttp.Request) (re *datalist.Result, err error) {
	columnsModel := &column.Shop{}
	listColumns := columnsModel.GoodsListColumns()
	// 筛选
	whereAndParams, err := datalist.FilterWhereAndParams(r, listColumns)
	if err != nil {
		return nil, err
	}
	listModel := &response.ShopGoodsList{}
	data, err := page.Data(r, &page.Param{
		TableName:   dao.ShopGoods.Table,
		Where:       whereAndParams.Where,
		WhereParams: whereAndParams.Params,
		OrderBy:     "sort ASC",
	}, listModel)
	if err != nil {
		return nil, err
	}
	return datalist.List(r, data, listColumns)
}

func (s *shopService) GoodsOne(r *ghttp.Request) (re *response.ShopGoods, err error) {
	requestModel := &request.ShopId{}
	if err := r.Parse(requestModel); err != nil {
		return nil, err
	}
	re = &response.ShopGoods{}
	err = dao.ShopGoods.Where("id=?", g.Slice{requestModel.Id}).Scan(re)
	if err != nil {
		return nil, err
	}
	return re, nil
}

// GoodsAdd 添加商品
func (s *shopService) GoodsAdd(r *ghttp.Request) (err error) {
	requestModel := &request.ShopGoodsAdd{}
	if err := r.Parse(requestModel); err != nil {
		return err
	}
	goodsData := &entity.ShopGoods{}
	err = dao.ShopGoods.Where("title=?", requestModel.Title).Scan(goodsData)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if goodsData.Id > 0 {
		return errors.New("标题已存在")
	}
	insertData := gconv.Map(requestModel)
	id := snowflake.GenerateID()
	insertData["id"] = id
	insertData["created_at"] = xtime.GetNowTime()
	if _, err = dao.ShopGoods.Data(insertData).Insert(); err != nil {
		return err
	}
	return nil
}

// GoodsEdit 商品编辑
func (s *shopService) GoodsEdit(r *ghttp.Request) (err error) {
	requestModel := &request.ShopGoodsEdit{}
	if err := r.Parse(requestModel); err != nil {
		return err
	}
	goodsData := &entity.ShopGoods{}
	err = dao.ShopGoods.Where("id<>? AND title=?", requestModel.Id, requestModel.Title).Scan(goodsData)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if goodsData.Id > 0 {
		return errors.New("标题已存在")
	}
	updateData := gconv.Map(requestModel)
	delete(updateData, "id")
	updateData["updated_at"] = xtime.GetNowTime()
	if _, err = dao.ShopGoods.Data(updateData).Where("id=?", requestModel.Id).Update(); err != nil {
		return err
	}
	return nil
}

// GoodsSort 商品排序
func (s *shopService) GoodsSort(r *ghttp.Request) (err error) {
	requestModel := &request.ShopGoodsSort{}
	if err = r.Parse(requestModel); err != nil {
		return err
	}
	if err = g.DB().Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) error {
		for index, id := range requestModel.Sort {
			if _, err = dao.ShopGoods.Ctx(ctx).TX(tx).Data(g.Map{
				"sort": (requestModel.Page-1)*requestModel.PageSize + index + 1,
			}).Where("id=?", id).Update(); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// GoodsDelete 商品删除
func (s *shopService) GoodsDelete(r *ghttp.Request) (err error) {
	requestModel := &request.ShopIds{}
	if err = r.Parse(requestModel); err != nil {
		return err
	}
	if err = g.DB().Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) error {
		for _, id := range requestModel.Selected {
			if _, err := dao.ShopGoods.Ctx(ctx).TX(tx).Where("id=?", id).Delete(); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (s *shopService) OrderList(r *ghttp.Request) (re *datalist.Result, err error) {
	columnsModel := &column.Shop{}
	listColumns := columnsModel.OrderListColumns()
	// 筛选
	whereAndParams, err := datalist.FilterWhereAndParams(r, listColumns)
	if err != nil {
		return nil, err
	}
	listModel := &response.ShopOrderList{}
	data, err := page.Data(r, &page.Param{
		TableName:   dao.ShopOrder.Table + " o",
		Where:       whereAndParams.Where,
		WhereParams: whereAndParams.Params,
		Join: page.ParamJoin{
			&page.ParamJoinItem{
				JoinType:  "leftJoin",
				JoinTable: dao.User.Table + " u",
				On:        "u.id=o.user_id",
			},
			&page.ParamJoinItem{
				JoinType:  "leftJoin",
				JoinTable: dao.UserInfo.Table + " ui",
				On:        "ui.user_id=o.user_id",
			},
		},
		Field:   "o.*,u.username,ui.nickname",
		OrderBy: "o.id DESC",
	}, listModel)
	if err != nil {
		return nil, err
	}
	return datalist.List(r, data, listColumns)
}

func (s *shopService) OrderStatus(r *ghttp.Request) (err error) {
	requestModel := &request.ShopOrderStatus{}
	if err := r.Parse(requestModel); err != nil {
		return err
	}
	orderData := &entity.ShopOrder{}
	err = dao.ShopOrder.Where("id=?", requestModel.Id).Scan(orderData)
	if err != nil {
		return err
	}
	if requestModel.Status == constant.ShopOrderStatusInit {
		return errors.New("订单状态不正确")
	} else if requestModel.Status == constant.ShopOrderStatusPaid {
		if orderData.Status != constant.ShopOrderStatusInit {
			return errors.New("只有未支付订单可以设置为已支付")
		}
	} else if requestModel.Status == constant.ShopOrderStatusFinish {
		if orderData.Status != constant.ShopOrderStatusPaid {
			return errors.New("只有已支付订单可以设置为已完成")
		}
	}
	nowTime := xtime.GetNowTime()
	if err = g.DB().Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) (err error) {
		updateData := g.Map{
			"status": requestModel.Status,
		}
		if requestModel.Status == constant.ShopOrderStatusPaid {
			updateData["paid_at"] = nowTime
		}
		if _, err = dao.ShopOrder.Ctx(ctx).TX(tx).Data(updateData).Where("id=?", requestModel.Id).Update(); err != nil {
			return err
		}
		if requestModel.Status == constant.ShopOrderStatusPaid {
			// 处理已支付逻辑
			err = libservice.Shop.SetOrderPaid(ctx, tx, orderData.Id, orderData.OrderAmount)
			if err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
