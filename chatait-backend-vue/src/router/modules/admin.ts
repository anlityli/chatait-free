/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import { shallowRef } from 'vue'
import { UserAddIcon } from 'tdesign-icons-vue-next'

export default [
  {
    path: '/admin',
    name: 'admin',
    component: () => import('@/layouts/index.vue'),
    redirect: '/admin/list',
    meta: { title: '管理员管理', icon: shallowRef(UserAddIcon) },
    children: [
      {
        path: 'list',
        name: 'adminList',
        component: () => import('@/pages/admin/list.vue'),
        meta: { title: '管理员列表' },
      },
      {
        path: 'add',
        name: 'adminAdd',
        component: () => import('@/pages/admin/edit.vue'),
        meta: { title: '添加管理员', highlight: 'adminList' },
      },
      {
        path: 'edit',
        name: 'adminEdit',
        component: () => import('@/pages/admin/edit.vue'),
        meta: { title: '编辑管理员', highlight: 'adminList' },
      },
      {
        path: 'role-list',
        name: 'adminRoleList',
        component: () => import('@/pages/admin/roleList.vue'),
        meta: { title: '管理员角色列表' },
      },
      {
        path: 'role-add',
        name: 'adminRoleAdd',
        component: () => import('@/pages/admin/roleEdit.vue'),
        meta: { title: '添加管理员角色', highlight: 'adminRoleList' },
      },
      {
        path: 'role-edit',
        name: 'adminRoleEdit',
        component: () => import('@/pages/admin/roleEdit.vue'),
        meta: { title: '编辑管理员角色', highlight: 'adminRoleList' },
      },
      {
        path: 'role-permission',
        name: 'adminRolePermission',
        component: () => import('@/pages/admin/rolePermission.vue'),
        meta: { title: '角色权限', highlight: 'adminRoleList' },
      },
      {
        path: 'role-column',
        name: 'adminRoleColumn',
        component: () => import('@/pages/admin/roleColumn.vue'),
        meta: { title: '列表字段', highlight: 'adminRoleList' },
      },
      {
        path: 'reset-password',
        name: 'adminResetPassword',
        component: () => import('@/pages/admin/resetPassword.vue'),
        meta: { title: '重置密码' },
      },
    ],
  },
]
