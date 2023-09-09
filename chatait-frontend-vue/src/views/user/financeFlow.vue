<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import http from '@/utils/network/http'
import { ResponseFinanceWalletFlowItem } from '@/utils/model/response/finance'
import tool from '@/utils/tool/tool'
import { ResponsePage } from '@/utils/model/response/page'
import { WalletType } from '@/utils/constant/config'
import { useConfigStore } from '@/store'

const route = useRoute()
const router = useRouter()
const configStore = useConfigStore()
const cardTitle = ref('')
const flowList = ref<ResponsePage>({
  list_data: <ResponseFinanceWalletFlowItem[]>[],
  total_count: 0,
  page: 1,
  page_size: 20,
  total_page: 0,
  page_count: 0,
})
const pagination = {
  current: 1,
  pageSize: 20,
  defaultCurrent: flowList.value.page,
  defaultPageSize: flowList.value.page_size,
  total: flowList.value.total_count,
}

const columns = ref([
  {
    colKey: 'created_at',
    title: '时间',
    width: 120,
    cell: (h: any, { row }: any) => {
      return tool.formatDate(row.created_at, false)
    },
  },
  { colKey: 'id', title: '流水号', width: 200 },
  {
    colKey: 'amount',
    title: '交易额',
    width: 150,
    align: 'right',
    cell: (h: any, { row }: any) => {
      if (row.amount > 0) {
        return `+${tool.centToYuan(row.amount)}`
      }
      return tool.centToYuan(row.amount)
    },
  },
  {
    colKey: 'total',
    title: '余额',
    width: 150,
    align: 'right',
    cell: (h: any, { row }: any) => {
      return tool.centToYuan(row.total)
    },
  },
  {
    colKey: 'remark',
    title: '备注',
    cell: (h: any, { row }: any) => {
      return row.remark
    },
  },
])

const handleGetData = async () => {
  flowList.value = (await http.get('finance/wallet-flow-list', {
    wallet_type: route.params.type,
    page: pagination.current,
    page_size: pagination.pageSize,
  })) as ResponsePage
  pagination.total = flowList.value.total_count
}

const handlePageChange = async (pageInfo: any) => {
  pagination.current = pageInfo.current
  pagination.pageSize = pageInfo.pageSize
  await handleGetData()
}

onMounted(async () => {
  if (route.params.type === 'balance') {
    cardTitle.value = `${configStore.walletName(WalletType.balance)}记录`
  } else if (route.params.type === 'gpt3') {
    cardTitle.value = `${configStore.walletName(WalletType.gpt3)}记录`
  } else if (route.params.type === 'gpt4') {
    cardTitle.value = `${configStore.walletName(WalletType.gpt4)}记录`
  } else if (route.params.type === 'midjourney') {
    cardTitle.value = `${configStore.walletName(WalletType.midjourney)}记录`
  }
  flowList.value = {
    list_data: <ResponseFinanceWalletFlowItem[]>[],
    total_count: 0,
    page: 1,
    page_size: 0,
    total_page: 0,
    page_count: 0,
  }
  await handleGetData()
})
</script>
<template>
  <div class="finance-profile-wrap">
    <t-card :title="cardTitle">
      <t-base-table
        row-key="id"
        :data="flowList.list_data"
        :columns="columns"
        :pagination="pagination"
        @page-change="handlePageChange"
      ></t-base-table>
    </t-card>
  </div>
</template>

<style lang="scss" scoped>
.finance-profile-wrap {
  box-sizing: border-box;
  padding: 10px;
}
</style>
