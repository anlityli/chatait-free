<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <div class="main-content-container">
    <t-table :data="tableData" :columns="column" row-key="field">
      <template #level_name="slotProps">
        <t-input
          v-model="slotProps.row.level_name"
          @change="(e: any) => {
            handleEdit(slotProps.row.id, 'level_name', e.toString())
          }"
        ></t-input>
      </template>
      <template #month_gpt3="slotProps">
        <t-input
          :value="tool.centToYuan(slotProps.row.month_gpt3)"
          @change="(e: any) => {
            handleEdit(slotProps.row.id, 'month_gpt3', tool.yuanToCent(e).toString())
          }"
        ></t-input>
      </template>
      <template #month_gpt4="slotProps">
        <t-input
          :value="tool.centToYuan(slotProps.row.month_gpt4)"
          @change="(e: any) => {
            handleEdit(slotProps.row.id, 'month_gpt4', tool.yuanToCent(e).toString())
          }"
        ></t-input>
      </template>
      <template #month_midjourney="slotProps">
        <t-input
          :value="tool.centToYuan(slotProps.row.month_midjourney)"
          @change="(e: any) => {
            handleEdit(slotProps.row.id, 'month_midjourney', tool.yuanToCent(e).toString())
          }"
        ></t-input>
      </template>
    </t-table>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { TableRowData } from 'tdesign-vue-next'
import http from '@/utils/network/http'
import { ResponseConfigLevelListItem } from '@/utils/model/response/config'
import { useConfigStore } from '@/store'
import { WalletType } from '@/constants/config'
import tool from '@/utils/tool/tool'

const configStore = useConfigStore()

const column = ref<TableRowData[]>([
  {
    title: '级别ID',
    align: 'center',
    colKey: 'id',
  },
  {
    title: '级别名称',
    align: 'center',
    colKey: 'level_name',
  },
  {
    title: `每月赠送${configStore.walletName(WalletType.gpt3)}`,
    align: 'center',
    colKey: 'month_gpt3',
  },
  {
    title: `每月赠送${configStore.walletName(WalletType.gpt4)}`,
    align: 'center',
    colKey: 'month_gpt4',
  },
  {
    title: `每月赠送${configStore.walletName(WalletType.midjourney)}`,
    align: 'center',
    colKey: 'month_midjourney',
  },
])

const tableData = ref<ResponseConfigLevelListItem[]>([])

const handleGetData = async () => {
  tableData.value = (await http.get('config/level-list')) as ResponseConfigLevelListItem[]
}
const handleEdit = async (id: number, field: string, value: string) => {
  console.log(value)
  await http.post('config/level-edit', {
    id,
    field,
    value,
  })
  await handleGetData()
}

onMounted(async () => {
  await handleGetData()
})
</script>
