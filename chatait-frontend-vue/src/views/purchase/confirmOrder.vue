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
import { ResponseShopGoodsItem, ResponseShopOrder, ResponseShopOrderCalcAmount } from '@/utils/model/response/shop'
import { ShopOrderItem } from '@/utils/model/app/shop'

const route = useRoute()
const router = useRouter()

const orderItemList = ref<ResponseShopGoodsItem[]>([])
const orderAmount = ref<number>(0)
const orderListData = ref<ShopOrderItem[]>([])
const orderSubmitForm = ref({
  goods_id: route.params.goodsId,
  goods_num: 1,
})

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
  orderItemList.value = []
  const goodsDetail = (await http.get('shop/goods-detail', {
    goods_id: route.params.goodsId,
  })) as ResponseShopGoodsItem

  orderListData.value.push({
    description: goodsDetail.title,
    subLabel: '',
    realPrice: tool.centToYuan(goodsDetail.real_price),
  })

  const orderAmountResponse = (await http.post(
    'shop/order-calc-amount',
    orderSubmitForm.value,
  )) as ResponseShopOrderCalcAmount
  orderAmount.value = orderAmountResponse.order_amount

  orderListData.value.push({
    description: '',
    subLabel: '合计: ',
    realPrice: tool.centToYuan(orderAmount.value),
  })
}

const handleToPayOrder = async () => {
  const orderSubmitResponse = (await http.post('shop/generate-order', orderSubmitForm.value)) as ResponseShopOrder
  console.log(orderSubmitResponse)
  await router.push(`/purchase/pay-order/${orderSubmitResponse.id}`)
}

onMounted(() => {
  handleGetData()
})
</script>
<template>
  <div class="purchase-confirm-order-wrap">
    <t-card>
      <div class="order-header">
        <div class="order-title">确认订单</div>
        <div class="order-sn"></div>
      </div>
      <t-divider></t-divider>
      <div class="order-content">
        <t-base-table row-key="id" :data="orderListData" :columns="columns"></t-base-table>
        <div class="order-submit">
          <t-button @click="handleToPayOrder">确认订单</t-button>
        </div>
      </div>
    </t-card>
  </div>
</template>

<style lang="scss">
.purchase-confirm-order-wrap {
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
