/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import { ResponseConversationSpeakItemMjDataComponents } from '@/utils/model/response/conversation'

export interface WSMsg {
  type: string
  data: any
}

export interface WSConversationMidjourneyListenerEvent {
  conversation_id: string
  user_id: string
  topic_id: string
  topic_title: string
  topic_type: number
  role: string
  action_type: number
  content: string
  img_url: string
  thumbnail_img_url: string
  progress: number
  components: ResponseConversationSpeakItemMjDataComponents[]
  referenced_conversation_id: string
  referenced_components: ResponseConversationSpeakItemMjDataComponents[]
  error: string
}
