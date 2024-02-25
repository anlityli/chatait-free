/*
 * Copyright 2024 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

export interface ResponseUserSensitiveWord {
  id: string
  username: string
  nickname: string
  type: number
  topic_type: number
  content: string
  validate_result: string
  created_at: number
}
