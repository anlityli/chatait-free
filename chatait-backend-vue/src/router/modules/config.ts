/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import { shallowRef } from 'vue'
import { SettingIcon } from 'tdesign-icons-vue-next'

export default [
  {
    path: '/config',
    name: 'config',
    component: () => import('@/layouts/index.vue'),
    redirect: '/config/all-option',
    meta: { title: '系统设置', icon: shallowRef(SettingIcon) },
    children: [
      {
        path: 'all-option',
        name: 'configAllOption',
        component: () => import('@/pages/config/allOption.vue'),
        meta: { title: '系统选项' },
      },
      {
        path: 'wallet-list',
        name: 'configWalletList',
        component: () => import('@/pages/config/walletList.vue'),
        meta: { title: '钱包配置列表' },
      },
      {
        path: 'pay-list',
        name: 'configPayList',
        component: () => import('@/pages/config/payList.vue'),
        meta: { title: '支付方式列表' },
      },
      {
        path: 'pay-edit',
        name: 'configPayEdit',
        component: () => import('@/pages/config/payEdit.vue'),
        meta: { title: '编辑支付方式', highlight: 'configPayList' },
      },
      {
        path: 'level-list',
        name: 'configLevelList',
        component: () => import('@/pages/config/levelList.vue'),
        meta: { title: '级别配置列表' },
      },
      {
        path: 'midjourney-list',
        name: 'configMidjourneyList',
        component: () => import('@/pages/config/midjourneyList.vue'),
        meta: { title: 'Midjourney配置列表' },
      },
      {
        path: 'midjourney-add',
        name: 'configMidjourneyAdd',
        component: () => import('@/pages/config/midjourneyEdit.vue'),
        meta: { title: '添加Midjourney配置', highlight: 'configMidjourneyList' },
      },
      {
        path: 'midjourney-edit',
        name: 'configMidjourneyEdit',
        component: () => import('@/pages/config/midjourneyEdit.vue'),
        meta: { title: '编辑Midjourney配置', highlight: 'configMidjourneyList' },
      },
      {
        path: 'openai-list',
        name: 'configOpenaiList',
        component: () => import('@/pages/config/openaiList.vue'),
        meta: { title: 'Openai配置列表' },
      },
      {
        path: 'openai-add',
        name: 'configOpenaiAdd',
        component: () => import('@/pages/config/openaiEdit.vue'),
        meta: { title: '添加Openai配置', highlight: 'configOpenaiList' },
      },
      {
        path: 'openai-edit',
        name: 'configOpenaiEdit',
        component: () => import('@/pages/config/openaiEdit.vue'),
        meta: { title: '编辑Openai配置', highlight: 'configOpenaiList' },
      },
    ],
  },
]
