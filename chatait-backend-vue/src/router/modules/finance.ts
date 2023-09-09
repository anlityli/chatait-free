/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import { shallowRef } from 'vue'
import { MoneyCircleIcon } from 'tdesign-icons-vue-next'

export default [
  {
    path: '/finance',
    name: 'finance',
    component: () => import('@/layouts/index.vue'),
    redirect: '/finance/wallet-list',
    meta: { title: '财务管理', icon: shallowRef(MoneyCircleIcon) },
    children: [
      {
        path: 'wallet-list',
        name: 'financeWalletList',
        component: () => import('@/pages/finance/walletList.vue'),
        meta: { title: '钱包列表' },
      },
      {
        path: 'wallet-flow-list-balance',
        name: 'financeWalletFlowListBalance',
        component: () => import('@/pages/finance/walletFlowList.vue'),
        meta: { title: '余额流水' },
      },
      {
        path: 'wallet-flow-list-gpt3',
        name: 'financeWalletFlowListGpt3',
        component: () => import('@/pages/finance/walletFlowList.vue'),
        meta: { title: `Gpt3流水` },
      },
      {
        path: 'wallet-flow-list-gpt4',
        name: 'financeWalletFlowListGpt4',
        component: () => import('@/pages/finance/walletFlowList.vue'),
        meta: { title: `Gpt4流水` },
      },
      {
        path: 'wallet-flow-list-midjourney',
        name: 'financeWalletFlowListMidjourney',
        component: () => import('@/pages/finance/walletFlowList.vue'),
        meta: { title: `Midjourney流水` },
      },
    ],
  },
]
