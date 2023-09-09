/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

export interface FormShopGoodsEditFeatItem {
  icon: string
  text: string
}

export interface FormShopGoodsEdit {
  id: string
  title: string
  content: string
  feat_items: FormShopGoodsEditFeatItem[]
  buy_type: number
  active_level_id: number
  active_expire_type: number
  active_expire_value: number
  buy_value: number
  buy_value_yuan: string
  market_price: number
  market_price_yuan: string
  real_price: number
  real_price_yuan: string
  status: number
  sort: number
}
