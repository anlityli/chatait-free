<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import http from '@/utils/network/http'
import { ResponseShopOrder } from '@/utils/model/response/shop'
import { ResponsePage } from '@/utils/model/response/page'
import tool from '@/utils/tool/tool'

const route = useRoute()
const router = useRouter()
const currentStatus = ref('')
const orderList = ref<ResponsePage>({
  list_data: <ResponseShopOrder[]>[],
  total_count: 0,
  page: 1,
  page_size: 20,
  total_page: 0,
  page_count: 0,
})

const pagination = {
  current: 1,
  pageSize: 20,
  defaultCurrent: orderList.value.page,
  defaultPageSize: orderList.value.page_size,
  total: orderList.value.total_count,
}

const columns = ref([
  {
    colKey: 'order_sn',
    title: '订单编号',
    width: '260',
    cell: (h: any, { row }: any) => {
      return row.order_sn
    },
  },
  {
    colKey: 'status',
    title: '订单状态',
    width: '100',
    cell: (h: any, { row }: any) => {
      if (row.status === 0) {
        return '待支付'
      }
      if (row.status === 4) {
        return '已完成'
      }
      if (row.status === 9) {
        return '已取消'
      }
      return ''
    },
  },
  {
    colKey: 'order_amount',
    title: '订单金额',
    align: 'right',
    cell: (h: any, { row }: any) => {
      return tool.centToYuan(row.order_amount)
    },
  },
  {
    colKey: 'pay_amount',
    title: '实付金额',
    align: 'right',
    cell: (h: any, { row }: any) => {
      return tool.centToYuan(row.pay_amount)
    },
  },
  {
    colKey: 'created_at',
    title: '下单时间',
    cell: (h: any, { row }: any) => {
      return tool.formatDate(row.created_at, true)
    },
  },
])

const handleGetData = async () => {
  orderList.value = (await http.get('shop/order-list', {
    status: currentStatus.value,
    page: pagination.current,
    page_size: pagination.pageSize,
  })) as ResponsePage
  pagination.total = orderList.value.total_count
}

const handlePageChange = async (pageInfo: any) => {
  pagination.current = pageInfo.current
  pagination.pageSize = pageInfo.pageSize
  await handleGetData()
}

const handleStatusChange = async () => {
  handleOrderListInit()
  await handleGetData()
}

const handleOrderListInit = () => {
  orderList.value = {
    list_data: <ResponseShopOrder[]>[],
    total_count: 0,
    page: 1,
    page_size: 0,
    total_page: 0,
    page_count: 0,
  }
}

const handleRowClick = async ({ row }: any) => {
  await router.push(`/purchase/order-detail/${row.id}`)
}

onMounted(async () => {
  handleOrderListInit()
  await handleGetData()
})
</script>
<template>
  <div class="purchase-order-list-wrap">
    <t-card>
      <div class="order-list-header">
        <div class="order-list-title">订单列表</div>
        <div class="order-list-tool">
          <t-radio-group v-model="currentStatus" variant="primary-filled" default-value="" @change="handleStatusChange">
            <t-radio-button :value="''">全部订单</t-radio-button>
            <t-radio-button :value="'0'">待付款</t-radio-button>
            <t-radio-button :value="'4'">已完成</t-radio-button>
            <t-radio-button :value="'9'">已取消</t-radio-button>
          </t-radio-group>
        </div>
      </div>
      <div class="order-list-table">
        <t-base-table
          row-key="id"
          :data="orderList.list_data"
          :columns="columns"
          :pagination="pagination"
          stripe
          hover
          @page-change="handlePageChange"
          @row-click="handleRowClick"
        ></t-base-table>
      </div>
    </t-card>
  </div>
</template>

<style lang="scss">
.purchase-order-list-wrap {
  .order-list-header {
    width: 100%;
    display: flex;
    overflow: hidden;

    .order-list-title {
      flex: 1 1 auto;
      width: 100%;
      font-weight: bold;
      overflow: hidden;
    }

    .order-list-tool {
      flex: 0 0 auto;
      width: 320px;
      text-align: right;
      overflow: hidden;
    }
  }

  .order-list-table {
    box-sizing: border-box;
    padding: 10px 0;
  }
}
</style>
