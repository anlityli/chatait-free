/*
 * Copyright 2024 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import { shallowRef } from 'vue'
import { LogoutIcon, UserIcon } from 'tdesign-icons-vue-next'

export default [
  {
    path: '/user',
    name: 'user',
    component: () => import('@/layouts/index.vue'),
    redirect: '/user/index',
    meta: { title: '会员管理', icon: shallowRef(UserIcon) },
    children: [
      {
        path: 'list',
        name: 'userList',
        component: () => import('@/pages/user/list.vue'),
        meta: { title: '会员列表' },
      },
      {
        path: 'index',
        name: 'UserIndex',
        component: () => import('@/pages/user/index.vue'),
        meta: { title: '个人中心' },
      },
      {
        path: 'sensitive-word-list',
        name: 'userSensitiveWordList',
        component: () => import('@/pages/user/sensitiveWordList.vue'),
        meta: { title: '敏感词触发列表' },
      },
      {
        path: 'sensitive-word-one',
        name: 'userSensitiveWordOne',
        component: () => import('@/pages/user/sensitiveWordOne.vue'),
        meta: { title: '敏感词触发详情', highlight: 'userSensitiveWordList' },
      },
    ],
  },
  {
    path: '/loginRedirect',
    name: 'loginRedirect',
    redirect: '/login',
    meta: { title: '登录页', icon: shallowRef(LogoutIcon) },
    component: () => import('@/layouts/blank.vue'),
    children: [
      {
        path: 'index',
        redirect: '/login',
        component: () => import('@/layouts/blank.vue'),
        meta: { title: '登录中心' },
      },
    ],
  },
]
