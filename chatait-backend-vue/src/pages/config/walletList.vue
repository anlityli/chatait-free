<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <div class="main-content-container">
    <t-table :data="tableData" :columns="column" row-key="field">
      <template #wallet_name="slotProps">
        <t-input
          v-model="slotProps.row.wallet_name"
          @change="handleEdit(slotProps.row.field, slotProps.row.wallet_name)"
        ></t-input>
      </template>
    </t-table>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { TableRowData } from 'tdesign-vue-next'
import http from '@/utils/network/http'
import { ResponseConfigWalletListItem } from '@/utils/model/response/config'
import { useConfigStore } from '@/store'

const configStore = useConfigStore()

const column = ref<TableRowData[]>([
  {
    title: '字段名',
    align: 'center',
    colKey: 'field',
  },
  {
    title: '钱包名称',
    align: 'center',
    colKey: 'wallet_name',
  },
])

const tableData = ref<ResponseConfigWalletListItem[]>([])

const handleEdit = async (field: string, walletName: string) => {
  await http.post('config/wallet-edit', {
    field,
    wallet_name: walletName,
  })
  await configStore.getWalletList()
}

onMounted(async () => {
  tableData.value = (await http.get('config/wallet-list')) as ResponseConfigWalletListItem[]
})
</script>
