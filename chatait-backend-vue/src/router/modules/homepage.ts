/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import { DashboardIcon } from 'tdesign-icons-vue-next'
import { shallowRef } from 'vue'

export default [
  {
    path: '/dashboard',
    component: () => import('@/layouts/index.vue'),
    redirect: '/dashboard/index',
    name: 'dashboard',
    meta: {
      title: '仪表盘',
      icon: shallowRef(DashboardIcon),
      orderNo: 0,
    },
    children: [
      {
        path: 'index',
        name: 'DashboardIndex',
        component: () => import('@/pages/dashboard/index.vue'),
        meta: {
          title: '概览仪表盘',
        },
      },
    ],
  },
]
