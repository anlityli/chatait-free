/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import { Token, UserInfo } from '@/utils/storage/model'

export default {
  setToken(token: Token) {
    localStorage.setItem('token', JSON.stringify(token))
  },
  getToken(): Token | null {
    const token = localStorage.getItem('token')
    if (token === null) {
      return null
    }
    return JSON.parse(token) as Token
  },
  clearToken() {
    localStorage.removeItem('token')
  },
  setUserInfo(info: UserInfo) {
    localStorage.setItem('userInfo', JSON.stringify(info))
  },
  getUserInfo(): UserInfo | null {
    const userInfo = localStorage.getItem('userInfo')
    if (userInfo === null) {
      return null
    }
    return JSON.parse(userInfo) as UserInfo
  },
  clearUserInfo() {
    localStorage.removeItem('userInfo')
  },
}
