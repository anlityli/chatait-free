/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import conversation from '@/router/modules/conversation'
import findPassword from '@/router/modules/findPassword'
import signup from '@/router/modules/signup'
import site from '@/router/modules/site'
import login from '@/router/modules/login'
import purchase from '@/router/modules/purchase'
import user from '@/router/modules/user'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    redirect: import.meta.env.VITE_DEFAULT_HOMEPAGE === 'login' ? '/login' : '/conversation/0',
  },
  conversation,
  findPassword,
  purchase,
  login,
  signup,
  site,
  user,
]

export default createRouter({
  history: createWebHistory('/'),
  routes,
})
