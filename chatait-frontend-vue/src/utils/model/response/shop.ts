/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

export interface ResponseShopGoodsItemFeatItems {
  icon: string
  text: string
}

export interface ResponseShopGoodsItem {
  id: string
  title: string
  content: string
  feat_items: string
  feat_items_slice: ResponseShopGoodsItemFeatItems[]
  active_level_id: number
  active_expire_type: number
  market_price: number
  real_price: number
  status: number
  sort: number
  created_at: number
  updated_at: number
}

export interface ResponseShopOrderCalcAmount {
  order_amount: number
}

export interface ResponseShopOrder {
  id: string
  order_sn: string
  user_id: string
  order_amount: number
  pay_amount: number
  status: number
  created_at: number
  updated_at: number
  expire_at: number
  order_goods_list: any
}
