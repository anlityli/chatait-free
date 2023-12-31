/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

export interface ResponseConfigWalletListItem {
  field: string
  wallet_name: string
}

export interface ResponseConfigPayChannelItem {
  id: number
  channel_name: string
  channel: string
  status: number
}

export interface ResponseConfigPayListItem {
  id: number
  api_name: string
  pay_channel: ResponseConfigPayChannelItem[]
  frontend_description: string
}
