// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package libservice

import (
	"context"
	"database/sql"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/constant"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/dao"
	"github.com/anlityli/chatait-free/chatait-public-lib/app/model/entity"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/helper"
	"github.com/anlityli/chatait-free/chatait-public-lib/library/xtime"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
)

var Shop = &shopService{}

type shopService struct {
}

// SetOrderPaid 设置订单为已支付
func (s *shopService) SetOrderPaid(ctx context.Context, tx *gdb.TX, orderId int64, payAmount int) (err error) {
	orderData := &entity.ShopOrder{}
	err = dao.ShopOrder.Ctx(ctx).TX(tx).Where("id=?", orderId).Scan(orderData)
	if err != nil {
		return err
	}
	if orderData.Id > 0 {
		// 订单设为已完成
		if _, err = dao.ShopOrder.Ctx(ctx).TX(tx).Data(g.Map{
			"status":     constant.ShopOrderStatusFinish,
			"pay_amount": payAmount,
			"updated_at": xtime.GetNowTime(),
		}).Where("id=?", orderData.Id).Update(); err != nil {
			return err
		}
		// 订单内的商品详情
		orderGoodsList := &[]*entity.ShopOrderGoods{}
		err = dao.ShopOrderGoods.Ctx(ctx).TX(tx).Where("order_id=?", orderData.Id).Scan(orderGoodsList)
		if err != nil && err != sql.ErrNoRows {
			return err
		}
		if len(*orderGoodsList) > 0 {
			for _, item := range *orderGoodsList {
				if item.GoodsSnapshot == "" {
					continue
				}
				goodsDecode, err := gjson.Decode(item.GoodsSnapshot)
				if err != nil {
					return err
				}
				goodsData := &entity.ShopGoods{}
				err = gconv.Scan(goodsDecode, goodsData)
				if err != nil {
					return err
				}
				// 购买级别
				if goodsData.BuyType == constant.ShopGoodsBuyTypeLevel {
					if goodsData.ActiveLevelId <= 0 {
						continue
					}
					// 会员信息
					userData := &entity.User{}
					err = dao.User.Ctx(ctx).TX(tx).Where("id=?", orderData.UserId).Scan(userData)
					if err != nil {
						return err
					}
					var oriExpireDate *gtime.Time
					if userData.LevelExpireDate != nil && xtime.GetNow().Timestamp() < userData.LevelExpireDate.Timestamp() {
						oriExpireDate = userData.LevelExpireDate
					} else {
						oriExpireDate = xtime.GetNow()
					}
					expireDate := ""
					switch goodsData.ActiveExpireType {
					case constant.ShopGoodsActiveExpireTypeNone:
						expireDate = oriExpireDate.Format("Y-m-d")
					case constant.ShopGoodsActiveExpireTypeDay:
						expireDate = oriExpireDate.AddDate(0, 0, 1*item.GoodsNum).Format("Y-m-d")
					case constant.ShopGoodsActiveExpireTypeMonth:
						expireDate = oriExpireDate.AddDate(0, 1*item.GoodsNum, 0).Format("Y-m-d")
					case constant.ShopGoodsActiveExpireTypeYear:
						expireDate = oriExpireDate.AddDate(1*item.GoodsNum, 0, 0).Format("Y-m-d")
					}
					// 会员升级
					var levelExpireDate *gtime.Time
					if expireDate != "" {
						levelExpireDate = gtime.NewFromStr(expireDate)
					}
					if err = User.ChangeLevel(ctx, tx, &ChangeLevelParams{
						UserId:          orderData.UserId,
						NewLevelId:      goodsData.ActiveLevelId,
						LevelExpireDate: levelExpireDate,
						Remark:          "购买[" + goodsData.Title + "]升级",
					}); err != nil {
						return err
					}
					// 如果会员的原级别比新级别低，则把会员当前的提问次数按照级别设置
					if userData.LevelId < goodsData.ActiveLevelId {
						balance := Wallet.GetAllBalance(orderData.UserId)
						// 如果之前有余额，先扣除
						if balance.Gpt3 > 0 {
							err = Wallet.ChangeWalletBalance(ctx, tx, &ChangeWalletParam{
								UserId:     orderData.UserId,
								WalletType: constant.WalletTypeGpt3,
								Amount:     -gconv.Int(balance.Gpt3),
								Remark:     "续费或购买计划扣除之前余额",
								TargetType: constant.WalletChangeTargetTypeShopOrderGoods,
								TargetID:   item.Id,
							})
							if err != nil {
								return err
							}
						}
						if balance.Gpt4 > 0 {
							err = Wallet.ChangeWalletBalance(ctx, tx, &ChangeWalletParam{
								UserId:     orderData.UserId,
								WalletType: constant.WalletTypeGpt4,
								Amount:     -gconv.Int(balance.Gpt4),
								Remark:     "续费或购买计划扣除之前余额",
								TargetType: constant.WalletChangeTargetTypeShopOrderGoods,
								TargetID:   item.Id,
							})
							if err != nil {
								return err
							}
						}
						if balance.Midjourney > 0 {
							err = Wallet.ChangeWalletBalance(ctx, tx, &ChangeWalletParam{
								UserId:     orderData.UserId,
								WalletType: constant.WalletTypeMidjourney,
								Amount:     -gconv.Int(balance.Midjourney),
								Remark:     "续费或购买计划扣除之前余额",
								TargetType: constant.WalletChangeTargetTypeShopOrderGoods,
								TargetID:   item.Id,
							})
							if err != nil {
								return err
							}
						}
						configLevelData, err := helper.GetConfigLevel(goodsData.ActiveLevelId)
						if err != nil {
							return err
						}
						// 增加余额
						if configLevelData.MonthGpt3 > 0 {
							err = Wallet.ChangeWalletBalance(ctx, tx, &ChangeWalletParam{
								UserId:     orderData.UserId,
								WalletType: constant.WalletTypeGpt3,
								Amount:     configLevelData.MonthGpt3,
								Remark:     "续费或购买计划增加",
								TargetType: constant.WalletChangeTargetTypeShopOrderGoods,
								TargetID:   item.Id,
							})
							if err != nil {
								return err
							}
						}
						if configLevelData.MonthGpt4 > 0 {
							err = Wallet.ChangeWalletBalance(ctx, tx, &ChangeWalletParam{
								UserId:     orderData.UserId,
								WalletType: constant.WalletTypeGpt4,
								Amount:     configLevelData.MonthGpt4,
								Remark:     "续费或购买计划增加",
								TargetType: constant.WalletChangeTargetTypeShopOrderGoods,
								TargetID:   item.Id,
							})
							if err != nil {
								return err
							}
						}
						if configLevelData.MonthMidjourney > 0 {
							err = Wallet.ChangeWalletBalance(ctx, tx, &ChangeWalletParam{
								UserId:     orderData.UserId,
								WalletType: constant.WalletTypeMidjourney,
								Amount:     configLevelData.MonthMidjourney,
								Remark:     "续费或购买计划增加",
								TargetType: constant.WalletChangeTargetTypeShopOrderGoods,
								TargetID:   item.Id,
							})
							if err != nil {
								return err
							}
						}
					}
				} else if goodsData.BuyType == constant.ShopGoodsBuyTypeBalance || goodsData.BuyType == constant.ShopGoodsBuyTypeGpt3 || goodsData.BuyType == constant.ShopGoodsBuyTypeGpt4 || goodsData.BuyType == constant.ShopGoodsBuyTypeMidjourney {
					// 购买钱包余额
					if goodsData.BuyValue <= 0 {
						continue
					}
					walletType := ""
					switch goodsData.BuyType {
					case constant.ShopGoodsBuyTypeBalance:
						walletType = constant.WalletTypeBalance
					case constant.ShopGoodsBuyTypeGpt3:
						walletType = constant.WalletTypeGpt3
					case constant.ShopGoodsBuyTypeGpt4:
						walletType = constant.WalletTypeGpt4
					case constant.ShopGoodsBuyTypeMidjourney:
						walletType = constant.WalletTypeMidjourney
					}
					err = Wallet.ChangeWalletBalance(ctx, tx, &ChangeWalletParam{
						UserId:     orderData.UserId,
						WalletType: walletType,
						Amount:     goodsData.BuyValue,
						Remark:     "购买计划增加",
						TargetType: constant.WalletChangeTargetTypeShopOrderGoods,
						TargetID:   item.Id,
					})
					if err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}
