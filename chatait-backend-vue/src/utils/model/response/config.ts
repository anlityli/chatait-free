/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

export interface ResponseConfigAllOptionOptionsItem {
  label: string
  value: string
}

export interface ResponseConfigAllOption {
  config_name: string
  created_at: number
  input_type: number
  options: ResponseConfigAllOptionOptionsItem[]
  sort: number
  title: string
  type: string
  unit: string
  updated_at: number
  value: string
}

export interface ResponseConfigLevelListItem {
  id: number
  level_name: string
  month_gpt3: number
  month_gpt4: number
  month_midjourney: number
}

export interface ResponseConfigWalletListItem {
  field: string
  wallet_name: string
}

export interface ResponseConfigPayListItemParams {
  param: string
  param_name: string
  value: string
}

export interface ResponseConfigPayListItemPayChannel {
  id: number
  channel: string
  channel_name: string
  status: number
}

export interface ResponseConfigPayListItem {
  id: string
  api_name: string
  params: ResponseConfigPayListItemParams[]
  pay_channel: ResponseConfigPayListItemPayChannel[]
  frontend_description: string
  backend_description: string
  status: number
}

export interface ResponseConfigMidjourney {
  id: string
  title: string
  guild_id: string
  channel_id: string
  user_token: string
  mj_bot_id: string
  bot_token: string
  session_id: string
  user_agent: string
  proxy: string
  status: number
  listen_model: number
  created_at: number
  updated_at: number
}
