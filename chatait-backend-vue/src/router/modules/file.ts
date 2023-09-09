/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import { shallowRef } from 'vue'
import { FileIcon } from 'tdesign-icons-vue-next'

export default [
  {
    path: '/file',
    name: 'file',
    component: () => import('@/layouts/index.vue'),
    redirect: '/file/midjourney-list',
    meta: { title: '对话管理', icon: shallowRef(FileIcon) },
    children: [
      {
        path: 'midjourney-list',
        name: 'fileMidjourneyList',
        component: () => import('@/pages/file/midjourneyList.vue'),
        meta: { title: 'Midjourney文件列表' },
      },
    ],
  },
]
