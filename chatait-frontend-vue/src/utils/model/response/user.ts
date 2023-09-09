/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

export interface ResponseUser {
  id: string
  username: string
  nickname: string
  avatar: string
  level_id: number
  level_name: string
  level_expire_date: string
  created_at: number
  last_login_at: number
}
