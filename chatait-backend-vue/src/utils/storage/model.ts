/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

export interface Token {
  accessToken: string
  accessTokenExpire: number
  accessTokenExpireIn: number
  refreshToken: string
  refreshTokenExpire: number
  refreshTokenExpireIn: number
}

export interface AdminInfo {
  id: string
  user_id: string
  username: string
  admin_name: string
  real_name: string
  role_id: string
  role_name: string
  menu: any
  admin_permission: any
}
