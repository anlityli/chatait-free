/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import 'nprogress/nprogress.css' // progress bar style
import NProgress from 'nprogress' // progress bar
import { MessagePlugin } from 'tdesign-vue-next'
import { RouteRecordRaw } from 'vue-router'

import router from '@/router'
import { getPermissionStore, useAdminStore } from '@/store'
import { PAGE_NOT_FOUND_ROUTE } from '@/utils/route/constant'

NProgress.configure({ showSpinner: false })

router.beforeEach(async (to, from, next) => {
  NProgress.start()

  const permissionStore = getPermissionStore()
  const { whiteListRouters } = permissionStore

  const adminStore = useAdminStore()

  if (adminStore.isLogin()) {
    if (to.path === '/login') {
      next()
      return
    }
    try {
      await adminStore.getAdminInfo()

      const { asyncRoutes } = permissionStore

      if (asyncRoutes && asyncRoutes.length === 0) {
        const routeList = await permissionStore.buildAsyncRoutes()
        routeList.forEach((item: RouteRecordRaw) => {
          router.addRoute(item)
        })

        if (to.name === PAGE_NOT_FOUND_ROUTE.name) {
          // 动态添加路由后，此处应当重定向到fullPath，否则会加载404页面内容
          next({ path: to.fullPath, replace: true, query: to.query })
        } else {
          const redirect = decodeURIComponent((from.query.redirect || to.path) as string)
          next(to.path === redirect ? { ...to, replace: true } : { path: redirect })
          return
        }
      }
      if (router.hasRoute(to.name)) {
        next()
      } else {
        next(`/`)
      }
    } catch (error) {
      await MessagePlugin.error(error.message)
      next({
        path: '/login',
        query: { redirect: encodeURIComponent(to.fullPath) },
      })
      NProgress.done()
    }
  } else {
    /* white list router */
    if (whiteListRouters.indexOf(to.path) !== -1) {
      next()
    } else {
      next({
        path: '/login',
        query: { redirect: encodeURIComponent(to.fullPath) },
      })
    }
    NProgress.done()
  }
})

router.afterEach(async (to) => {
  if (to.path === '/login') {
    const adminStore = useAdminStore()
    const permissionStore = getPermissionStore()

    adminStore.logout()
    await permissionStore.restoreRoutes()
  }
  NProgress.done()
})
