<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import tool from '@/utils/tool/tool'
import http from '@/utils/network/http'
import { ResponseShopOrder } from '@/utils/model/response/shop'
import { ShopOrderItem } from '@/utils/model/app/shop'

const route = useRoute()
const router = useRouter()

const orderDetail = ref<ResponseShopOrder | null>(null)
const orderListData = ref<ShopOrderItem[]>([])

const columns = ref([
  {
    colKey: 'description',
    title: '购买计划',
    cell: (h: any, { row }: any) => {
      return row.description
    },
  },
  {
    colKey: 'subLabel',
    title: '',
    width: '80',
    cell: (h: any, { row }: any) => {
      return row.subLabel
    },
  },
  {
    colKey: 'realPrice',
    title: '金额',
    width: '100',
    align: 'right',
    cell: (h: any, { row }: any) => {
      return row.realPrice
    },
  },
])

const handleGetData = async () => {
  orderDetail.value = (await http.get('shop/order-detail', {
    order_id: route.params.orderId,
  })) as ResponseShopOrder

  orderListData.value.push({
    description: orderDetail.value.order_goods_list[0].goods_snapshot.title,
    subLabel: '',
    realPrice: tool.centToYuan(orderDetail.value.order_goods_list[0].goods_snapshot.real_price),
  })

  orderListData.value.push({
    description: '',
    subLabel: '合计: ',
    realPrice: tool.centToYuan(orderDetail.value.order_amount),
  })
}

const handleToPayOrder = async () => {
  await router.push(`/purchase/pay-order/${orderDetail.value?.id}`)
}

onMounted(() => {
  handleGetData()
})
</script>
<template>
  <div class="purchase-order-detail-wrap">
    <t-card>
      <div class="order-header">
        <div class="order-title">订单详情</div>
        <div class="order-sn">SN: {{ orderDetail?.order_sn }}</div>
      </div>
      <t-divider></t-divider>
      <div class="order-content">
        <t-base-table row-key="id" :data="orderListData" :columns="columns"></t-base-table>
        <div v-if="orderDetail?.status === 0" class="order-submit">
          <t-button @click="handleToPayOrder">支付订单</t-button>
        </div>
      </div>
    </t-card>
  </div>
</template>

<style lang="scss">
.purchase-order-detail-wrap {
  .order-header {
    width: 100%;
    display: flex;
    overflow: hidden;

    .order-title {
      flex: 1 1 auto;
      width: 100%;
      font-weight: bold;
      overflow: hidden;
    }

    .order-sn {
      flex: 1 1 auto;
      width: 100%;
      text-align: right;
      overflow: hidden;
    }
  }

  .order-submit {
    padding: 10px;
    box-sizing: border-box;
    text-align: right;
  }
}
</style>
