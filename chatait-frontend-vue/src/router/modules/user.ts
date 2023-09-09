/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

export default {
  path: '/user',
  component: () => import('@/views/layout/index.vue'),
  redirect: '/user/profile',
  children: [
    {
      path: '/user/index',
      component: () => import('@/views/user/index.vue'),
      redirect: '/user/profile',
      children: [
        {
          path: '/user/profile',
          component: () => import('@/views/user/profile.vue'),
          name: 'userProfile',
        },
        {
          path: '/user/finance/flow/:type',
          component: () => import('@/views/user/financeFlow.vue'),
          name: 'userFinanceFlow',
        },
      ],
    },
  ],
}
