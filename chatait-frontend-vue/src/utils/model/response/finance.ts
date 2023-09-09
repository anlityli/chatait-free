/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

export interface ResponseFinanceWalletFlowItem {
  id: string
  user_id: string
  amount: number
  total: number
  is_incr: number
  target_type: number
  target_id: number
  year: number
  month: number
  day: number
  created_at: number
  updated_at: number
}

export interface ResponseFinanceWalletInfo {
  user_id: string
  balance: number
  gpt3: number
  gpt4: number
  midjourney: number
}
