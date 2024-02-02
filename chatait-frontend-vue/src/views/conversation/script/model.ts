/*
 * Copyright 2024 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

export interface SpeakForm {
  topic_type: number
  stream_uuid: string
  topic_id: string
  content: string
}

export interface MidjourneySpeakForm {
  topic_id: string
  content: string
  application_type: number
  no: string
  images: string
  seed: string
  ar: string
  chaos: string
  quality: string
  model: string
  style: string
  stylize: string
  tile: string
  iw: string
}

export type MidjourneySpeakFormKey = keyof MidjourneySpeakForm

export interface MidjourneyCustomForm {
  refer_conversation_id: string
  action_type: number
  index: number
  custom_id: string
}
