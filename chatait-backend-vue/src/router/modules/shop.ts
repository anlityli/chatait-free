/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

import { shallowRef } from 'vue'
import { ShopIcon } from 'tdesign-icons-vue-next'

export default [
  {
    path: '/shop',
    name: 'shop',
    component: () => import('@/layouts/index.vue'),
    redirect: '/shop/goods-list',
    meta: { title: '商城管理', icon: shallowRef(ShopIcon) },
    children: [
      {
        path: 'goods-list',
        name: 'shopGoodsList',
        component: () => import('@/pages/shop/goodsList.vue'),
        meta: { title: '商品列表' },
      },
      {
        path: 'goods-add',
        name: 'shopGoodsAdd',
        component: () => import('@/pages/shop/goodsEdit.vue'),
        meta: { title: '添加商品', highlight: 'shopGoodsList' },
      },
      {
        path: 'goods-edit',
        name: 'shopGoodsEdit',
        component: () => import('@/pages/shop/goodsEdit.vue'),
        meta: { title: '编辑商品', highlight: 'shopGoodsList' },
      },
      {
        path: 'order-list',
        name: 'shopOrderList',
        component: () => import('@/pages/shop/orderList.vue'),
        meta: { title: '订单列表' },
      },
    ],
  },
]
