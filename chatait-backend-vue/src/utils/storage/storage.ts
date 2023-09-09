/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import { AdminInfo, Token } from '@/utils/storage/model'

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
  setAminInfo(info: AdminInfo) {
    localStorage.setItem('adminInfo', JSON.stringify(info))
  },
  getAdminInfo(): AdminInfo | null {
    const adminInfo = localStorage.getItem('adminInfo')
    if (adminInfo === null) {
      return null
    }
    return JSON.parse(adminInfo) as AdminInfo
  },
  clearAdminInfo() {
    localStorage.removeItem('admin')
    localStorage.removeItem('adminInfo')
  },
}
