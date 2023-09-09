/*
 * Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
 * Use of this source code is governed by a AGPL v3.0 style
 * license that can be found in the LICENSE file.
 */

export default {
  path: '/purchase',
  component: () => import('@/views/layout/index.vue'),
  redirect: '/purchase/goods-list',
  children: [
    {
      path: '/purchase/index',
      component: () => import('@/views/purchase/index.vue'),
      redirect: '/purchase/goods-list',
      children: [
        {
          path: '/purchase/goods-list',
          component: () => import('@/views/purchase/goodsList.vue'),
          name: 'purchaseGoodsList',
        },
        {
          path: '/purchase/confirm-order/:goodsId',
          component: () => import('@/views/purchase/confirmOrder.vue'),
          name: 'purchaseConfirmOrder',
        },
        {
          path: '/purchase/pay-order/:orderId',
          component: () => import('@/views/purchase/payOrder.vue'),
          name: 'purchasePayOrder',
        },
        {
          path: '/purchase/order-list',
          component: () => import('@/views/purchase/orderList.vue'),
          name: 'purchaseOrderList',
        },
        {
          path: '/purchase/order-detail/:orderId',
          component: () => import('@/views/purchase/orderDetail.vue'),
          name: 'purchaseOrderDetail',
        },
      ],
    },
  ],
}
