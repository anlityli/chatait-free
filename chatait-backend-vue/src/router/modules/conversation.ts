/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import { shallowRef } from 'vue'
import { ChatIcon } from 'tdesign-icons-vue-next'

export default [
  {
    path: '/conversation',
    name: 'conversation',
    component: () => import('@/layouts/index.vue'),
    redirect: '/conversation/topic-list',
    meta: { title: '对话管理', icon: shallowRef(ChatIcon) },
    children: [
      {
        path: 'topic-list',
        name: 'conversationTopicList',
        component: () => import('@/pages/conversation/topicList.vue'),
        meta: { title: '话题列表' },
      },
      {
        path: 'list',
        name: 'conversationList',
        component: () => import('@/pages/conversation/list.vue'),
        meta: { title: '对话列表' },
      },
    ],
  },
]
