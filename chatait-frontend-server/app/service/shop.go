// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/anlityli/chatait-free/chatait-frontend-server/app/model/request"
	"github.com/anlityli/chatait-free/chatait-frontend-server/app/model/response"
	"github.com/anlityli/chatait-free/chatait-frontend-server/library/auth"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/constant"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/model/entity"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/helper"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/page"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/snowflake"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/xtime"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtimer"
	"github.com/gogf/gf/util/gconv"
	"time"
)

var Shop = &shopService{}

type shopService struct {
}

// GoodsList 商品列表
func (s *shopService) GoodsList(r *ghttp.Request) (re *response.ShopGoodsList, err error) {
	requestModel := &request.ShopGoodsList{}
	if err = r.Parse(requestModel); err != nil {
		return nil, err
	}
	re = &response.ShopGoodsList{}
	err = dao.ShopGoods.Where("status=1").Order("sort ASC").Scan(re)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	for _, item := range *re {
		if item.FeatItems != "" {
			featItemsDecode, err := gjson.Decode(item.FeatItems)
			if err != nil {
				return nil, err
			}
			err = gconv.Scan(featItemsDecode, &item.FeatItemsSlice)
			if err != nil {
				return nil, err
			}
		}
	}
	return re, nil
}

// GoodsDetail 商品明细
func (s *shopService) GoodsDetail(r *ghttp.Request) (re *response.ShopGoods, err error) {
	requestModel := &request.ShopGoodsDetail{}
	if err = r.Parse(requestModel); err != nil {
		return nil, err
	}
	re = &response.ShopGoods{}
	err = dao.ShopGoods.Where("id=? AND status=1", requestModel.GoodsId).Scan(re)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return re, nil
}

// OrderList 订单列表
func (s *shopService) OrderList(r *ghttp.Request) (re *page.Response, err error) {
	requestModel := &request.ShopOrderList{}
	if err = r.Parse(requestModel); err != nil {
		return nil, err
	}
	userId := auth.GetUserId(r)
	where := "user_id=?"
	whereParams := g.Slice{userId}
	if requestModel.Status != "" {
		where += " AND status=?"
		whereParams = append(whereParams, requestModel.Status)
	}
	listData := &response.ShopOrderList{}
	re, err = page.Data(r, &page.Param{
		TableName:   dao.ShopOrder.Table + " o",
		Where:       where,
		WhereParams: whereParams,
		OrderBy:     "o.id DESC",
	}, listData)
	if len(*listData) > 0 {
		for _, orderItem := range *listData {
			orderGoodsList := &response.ShopOrderGoodsList{}
			err = dao.ShopOrderGoods.Where("order_id=?", orderItem.Id).Scan(orderGoodsList)
			if err != nil && err != sql.ErrNoRows {
				return nil, err
			}
			orderItem.OrderGoodsList = *orderGoodsList
		}
	}
	return re, nil
}

// OrderDetail 订单详情
func (s *shopService) OrderDetail(r *ghttp.Request) (re *response.ShopOrder, err error) {
	requestModel := &request.ShopOrderDetail{}
	if err = r.Parse(requestModel); err != nil {
		return nil, err
	}
	userId := auth.GetUserId(r)
	re = &response.ShopOrder{}
	err = dao.ShopOrder.Where("id=? AND user_id=?", requestModel.OrderId, userId).Scan(re)
	if err != nil {
		return nil, err
	}
	orderGoodsList := &response.ShopOrderGoodsList{}
	err = dao.ShopOrderGoods.Where("order_id=?", requestModel.OrderId).Scan(orderGoodsList)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	re.OrderGoodsList = *orderGoodsList
	return re, nil
}

// OrderCalcAmount 订单计算金额
func (s *shopService) OrderCalcAmount(r *ghttp.Request) (re *response.ShopCalcOrderAmount, err error) {
	requestModel := &request.ShopGenerateOrder{}
	if err = r.Parse(requestModel); err != nil {
		return nil, err
	}
	goodsData := &entity.ShopGoods{}
	err = dao.ShopGoods.Where("id=?", requestModel.GoodsId).Scan(goodsData)
	if err != nil {
		return nil, err
	}
	orderAmount := goodsData.RealPrice * requestModel.GoodsNum
	re = &response.ShopCalcOrderAmount{}
	re.OrderAmount = orderAmount
	return re, nil
}

// GenerateOrder 生成订单
func (s *shopService) GenerateOrder(r *ghttp.Request) (re *response.ShopOrder, err error) {
	requestModel := &request.ShopGenerateOrder{}
	if err = r.Parse(requestModel); err != nil {
		return nil, err
	}
	nowTime := xtime.GetNowTime()
	goodsData := &entity.ShopGoods{}
	err = dao.ShopGoods.Where("id=?", requestModel.GoodsId).Scan(goodsData)
	if err != nil {
		return nil, err
	}
	if goodsData.Status != 1 {
		return nil, errors.New("商品已下架")
	}
	orderAmount := goodsData.RealPrice * requestModel.GoodsNum
	userId := auth.GetUserId(r)
	shopOrderExpireIn, err := helper.GetConfig("shopOrderExpireIn")
	if err != nil {
		return nil, err
	}
	shopOrderExpireInTime := gconv.Int64(shopOrderExpireIn)
	re = &response.ShopOrder{}
	orderInsertId := snowflake.GenerateID()
	if err := g.DB().Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) (err error) {
		orderSn, err := s.generateOrderSn(orderInsertId)
		if err != nil {
			return err
		}
		orderInsertData := g.Map{
			"id":            orderInsertId,
			"order_sn":      orderSn,
			"user_id":       userId,
			"order_amount":  orderAmount,
			"created_at":    nowTime,
			"due_expire_at": nowTime + shopOrderExpireInTime*60,
		}
		if _, err = dao.ShopOrder.Ctx(ctx).TX(tx).Data(orderInsertData).Insert(); err != nil {
			return err
		}
		orderGoodsInsertId := snowflake.GenerateID()
		goodsSnapshot, err := gjson.Encode(goodsData)
		if err != nil {
			return err
		}
		orderGoodsInsertData := g.Map{
			"id":             orderGoodsInsertId,
			"order_id":       orderInsertId,
			"user_id":        userId,
			"goods_id":       goodsData.Id,
			"goods_num":      requestModel.GoodsNum,
			"goods_snapshot": goodsSnapshot,
			"created_at":     nowTime,
		}
		if _, err = dao.ShopOrderGoods.Ctx(ctx).TX(tx).Data(orderGoodsInsertData).Insert(); err != nil {
			return err
		}

		err = gconv.Scan(orderInsertData, re)
		if err != nil {
			return err
		}
		orderGoodsInsertData["goods_snapshot"] = goodsData
		err = gconv.Scan(g.Slice{orderGoodsInsertData}, &re.OrderGoodsList)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	// 超时自动取消订单
	go func() {
		gtimer.AddOnce(time.Duration(shopOrderExpireInTime)*time.Second, func() {
			orderData := &entity.ShopOrder{}
			err = dao.ShopOrder.Where("id=?", orderInsertId).Scan(orderData)
			if err != nil {
				glog.Line(true).Println("订单取消失败", err)
				return
			}
			if orderData.Id > 0 && orderData.Status == constant.ShopOrderStatusInit {
				if _, err := dao.ShopOrder.Data(g.Map{
					"status":     constant.ShopOrderStatusCancel,
					"expired_at": xtime.GetNowTime(),
				}).Where("id=?", orderInsertId).Update(); err != nil {
					glog.Line(true).Println("订单取消失败", err)
				}
			}
		})
	}()
	return re, nil
}

// generateOrderSn 生成订单号
func (s *shopService) generateOrderSn(orderId ...int64) (re string, err error) {
	var id int64
	if len(orderId) > 0 {
		id = orderId[0]
	} else {
		id = snowflake.GenerateID()
	}
	date := xtime.GetNowFormat("Ymd")
	return date + gconv.String(id), nil
}

// PayOrder 支付订单
func (s *shopService) PayOrder(r *ghttp.Request) (re *response.ShopPayOrder, err error) {
	requestModel := &request.ShopPayOrder{}
	if err = r.Parse(requestModel); err != nil {
		return nil, err
	}
	userId := auth.GetUserId(r)
	// 订单是否是可以支付状态
	orderData := &entity.ShopOrder{}
	err = dao.ShopOrder.Where("id=?", requestModel.OrderId).Scan(orderData)
	if err != nil {
		return nil, err
	}
	if orderData.UserId != userId {
		return nil, errors.New("订单不存在")
	}
	if orderData.Status == constant.ShopOrderStatusCancel {
		return nil, errors.New("订单已取消")
	} else if orderData.Status != constant.ShopOrderStatusInit {
		return nil, errors.New("订单已支付")
	}
	payment, err := Pay.AddPayFlow(&AddPayFlowParams{
		R:           r,
		FlowType:    constant.PayFlowTypeShopOrder,
		TargetId:    orderData.Id,
		ConfigPayId: requestModel.ConfigPayId,
		OrderAmount: orderData.OrderAmount,
		PayChannel:  requestModel.PayChannel,
	})
	if err != nil {
		return nil, err
	}
	re = &response.ShopPayOrder{}
	// todo 不同的支付方式返回不同的payment
	return re, nil
}
