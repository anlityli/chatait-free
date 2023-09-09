/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import { defineStore } from 'pinia'
import http from '@/utils/network/http'
import storage from '@/utils/storage/storage'
import tool from '@/utils/tool/tool'
import { ResponseAdminInfo } from '@/utils/model/response/admin'
import { AdminInfo } from '@/utils/storage/model'

export const useAdminStore = defineStore('admin', {
  state: () => ({
    adminInfo: <AdminInfo | null>null,
  }),
  getters: {},
  actions: {
    async login(adminForm: Record<string, unknown>) {
      const loginResponseData = await http.postWithoutToken('oauth/login', {
        admin_name: adminForm.admin_name,
        password: adminForm.password,
      })
      // 登录成功后写入token
      storage.setToken({
        accessToken: loginResponseData.access_token,
        accessTokenExpire: loginResponseData.access_token_expire_in + tool.getTimestamp(),
        accessTokenExpireIn: loginResponseData.access_token_expire_in,
        refreshToken: loginResponseData.refresh_token,
        refreshTokenExpire: loginResponseData.refresh_token_expire_in + tool.getTimestamp(),
        refreshTokenExpireIn: loginResponseData.refresh_token_expire_in,
      })
      // 获取管理员信息
      this.adminInfo = await this.getAdminInfo()
    },
    async setAdminInfo() {
      // 获取管理员信息
      const adminInfoData = (await http.get('admin/info')) as ResponseAdminInfo
      storage.setAminInfo({
        id: adminInfoData.id,
        user_id: adminInfoData.user_id,
        username: adminInfoData.username,
        admin_name: adminInfoData.admin_name,
        real_name: adminInfoData.real_name,
        role_id: adminInfoData.role_id,
        role_name: adminInfoData.role_name,
        menu: adminInfoData.menu,
        admin_permission: adminInfoData.admin_permission,
      })
    },
    async getAdminInfo(): Promise<AdminInfo> {
      let adminInfo = storage.getAdminInfo()
      if (adminInfo === null || adminInfo.id === '') {
        await this.setAdminInfo()
        adminInfo = storage.getAdminInfo()
      }
      return adminInfo as AdminInfo
    },
    logout() {
      console.log('退出登录')
      storage.clearToken()
      storage.clearAdminInfo()
    },
    isLogin(): boolean {
      const token = storage.getToken()
      if (token === null) {
        return false
      }
      if (token.accessToken !== '' && token.accessTokenExpire > tool.getTimestamp()) {
        return true
      }
      return token.refreshToken !== '' && token.refreshTokenExpire > tool.getTimestamp()
    },
  },
  persist: {
    // afterRestore: () => {
    //   const permissionStore = usePermissionStore()
    //   permissionStore.initRoutes()
    // },
    // key: 'admin',
    // paths: ['token'],
  },
})
