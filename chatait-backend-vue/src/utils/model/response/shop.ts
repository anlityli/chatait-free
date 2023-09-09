/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

export interface ResponseShopGoods {
  id: string
  title: string
  content: string
  feat_items: string
  buy_type: number
  active_level_id: number
  active_expire_type: number
  active_expire_value: number
  buy_value: number
  market_price: number
  real_price: number
  status: number
  sort: number
  created_at: number
  updated_at: number
}
