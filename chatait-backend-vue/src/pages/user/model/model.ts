/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

export interface UserChangeLevelForm {
  user_id: string
  level_id: number
  level_expire_date: string
  remark: string
}

export interface UserBanForm {
  selected: string[]
  is_ban: number
}

export interface UserChangePasswordForm {
  user_id: string
  password: string
}
