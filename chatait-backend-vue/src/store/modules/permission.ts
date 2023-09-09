/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import { defineStore } from 'pinia'
import { RouteRecordRaw } from 'vue-router'

import router, { fixedRouterList, homepageRouterList } from '@/router'
import { store, useAdminStore } from '@/store'

/**
 * 权限菜单处理
 * @param routes
 * @param apiMenu
 */
const permissionMenuHandle = (routes: RouteRecordRaw[], apiMenu: any): RouteRecordRaw[] => {
  const reRoutes = <RouteRecordRaw[]>[]
  for (let apiIndex = 0; apiIndex < apiMenu.length; apiIndex++) {
    for (let i = 0; i < routes.length; i++) {
      if (apiMenu[apiIndex].key === routes[i].name) {
        const tempMenuItem = <RouteRecordRaw>{
          path: routes[i].path,
          redirect: routes[i].redirect,
          alias: routes[i].alias,
          name: routes[i].name,
          meta: routes[i].meta,
          children: [],
        }
        tempMenuItem.meta.title = apiMenu[apiIndex].title
        if (
          routes[i].children !== undefined &&
          routes[i].children.length > 0 &&
          apiMenu[apiIndex].children !== null &&
          apiMenu[apiIndex].children.length > 0
        ) {
          tempMenuItem.children = permissionMenuHandle(routes[i].children, apiMenu[apiIndex].children)
        }
        reRoutes.push(tempMenuItem)
      }
    }
  }
  return reRoutes
}

export const usePermissionStore = defineStore('permission', {
  state: () => ({
    whiteListRouters: ['/login'],
    routers: [],
    removeRoutes: [],
    asyncRoutes: [],
    apiMenu: [],
  }),
  actions: {
    async initRoutes() {
      // 全部路由
      const allRoutes = [...homepageRouterList, ...fixedRouterList]
      // api返回的可访问的菜单
      this.routers = permissionMenuHandle(allRoutes, this.apiMenu)

      // 在菜单展示全部路由
      // this.routers = [...homepageRouterList, ...fixedRouterList]
      // 在菜单只展示动态路由和首页
      // this.routers = [...homepageRouterList, ...accessedRouters];
      // 在菜单只展示动态路由
      // this.routers = [...accessedRouters];
      this.asyncRoutes = allRoutes
    },
    async buildAsyncRoutes() {
      try {
        // 拿到菜单
        const adminStore = useAdminStore()
        const adminInfo = await adminStore.getAdminInfo()
        this.apiMenu = adminInfo.menu
        await this.initRoutes()
        return this.asyncRoutes
      } catch (error) {
        throw new Error("Can't build routes")
      }
    },
    async restoreRoutes() {
      this.removeRoutes.forEach((item: RouteRecordRaw) => {
        router.addRoute(item)
      })
    },
  },
})

export function getPermissionStore() {
  return usePermissionStore(store)
}
