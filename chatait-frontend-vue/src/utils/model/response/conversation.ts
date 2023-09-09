/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

export interface ResponseConversationTopic {
  id: string
  title: string
  type: number
  created_at: number
  updated_at: number
}

export interface ResponseConversationSpeakItemMjDataComponentsComponents {
  type: number
  style: number
  label: string
  emoji: Record<string, string>
  custom_id: string
}

export interface ResponseConversationSpeakItemMjDataComponents {
  type: number
  components: ResponseConversationSpeakItemMjDataComponentsComponents[]
}

export interface ResponseConversationSpeakItemMjData {
  action_type: number
  img_url: string
  thumbnail_img_url: string
  progress: number
  components: ResponseConversationSpeakItemMjDataComponents[]
  error: string
  progress_text?: string
}

export interface ResponseConversationSpeakItem {
  id: string
  topic_id: string
  role: string
  content: string
  mj_data: ResponseConversationSpeakItemMjData
  created_at: number
}

export interface ResponseConversationStreamUuid {
  uuid: string
}

export interface ResponseConversationSpeak {
  topic_id: string
  title: string
  topic_type: number
}

export interface ResponseConversationMidjourneySpeak {
  topic_id: string
  title: string
  topic_type: number
  question_id: string
  question_content: string
  answer_id: string
}
