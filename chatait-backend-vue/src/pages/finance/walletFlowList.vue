<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <div v-if="route.name === 'financeWalletFlowListBalance' || route.name === 'financeWalletFlowListGpt3'">
    <common-table
      ref="tableEle"
      :request-path="
        route.name === 'financeWalletFlowListBalance'
          ? 'finance/wallet-flow-list-balance'
          : route.name === 'financeWalletFlowListGpt3'
          ? 'finance/wallet-flow-list-gpt3'
          : route.name === 'financeWalletFlowListGpt4'
          ? 'finance/wallet-flow-list-gpt4'
          : 'finance/wallet-flow-list-midjourney'
      "
      :primary-filter-field="'username'"
      :show-export-button="true"
      :row-select="{
        enable: false,
        type: 'single',
      }"
      @on-export="handleExport"
    >
    </common-table>
  </div>
  <div v-else>
    <t-card :bordered="false"
      ><div style="text-align: center">您当前的版本为免费版，购买正式版开通此功能！</div></t-card
    >
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import CommonTable from '@/components/common-table/CommonTable.vue'

const route = useRoute()
const router = useRouter()

const tableEle = ref(null)

const handleExport = () => {
  console.log('导出')
}

onMounted(async () => {
})
</script>
