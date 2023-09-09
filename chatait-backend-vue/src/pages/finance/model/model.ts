/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import { WalletType } from '@/constants/config'

export interface FormFinanceWalletChange {
  user_id: string
  wallet_type: WalletType
  amount: string
  remark: string
}
