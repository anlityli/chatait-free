/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import { MidjourneyCreateModel } from '@/constants/config'

export interface FormMidjourneyEdit {
  id: string
  title: string
  guild_id: string
  channel_id: string
  user_token: string
  mj_bot_id: string
  bot_token: string
  session_id: string
  user_agent: string
  hugging_face_token: string
  proxy: string
  status: number
  listen_model: number
  create_model: MidjourneyCreateModel
  ws_idle_time: number
}
