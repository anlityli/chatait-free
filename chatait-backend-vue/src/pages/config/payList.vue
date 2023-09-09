<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <div class="main-content-container">
    <t-table :data="tableData" :columns="column" row-key="field">
      <template #operation="slotProps">
        <t-dropdown :min-column-width="88" trigger="click">
          <t-button variant="text">
            操作数据
            <template #suffix>
              <chevron-down-icon></chevron-down-icon>
            </template>
          </t-button>
          <t-dropdown-menu>
            <t-dropdown-item @click="handleEdit(slotProps.row.id)">编辑</t-dropdown-item>
          </t-dropdown-menu>
        </t-dropdown>
      </template>
    </t-table>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { ChevronDownIcon } from 'tdesign-icons-vue-next'
import { TableRowData } from 'tdesign-vue-next'
import http from '@/utils/network/http'
import { ResponseConfigPayListItem } from '@/utils/model/response/config'

const router = useRouter()

const column = ref<TableRowData[]>([
  {
    title: '支付方式',
    align: 'center',
    colKey: 'api_name',
  },
  {
    title: '是否启用',
    align: 'center',
    colKey: 'status',
  },
  {
    title: '操作',
    width: 150,
    align: 'center',
    colKey: 'operation',
  },
])

const tableData = ref([])

const handleEdit = async (id: string) => {
  await router.push(`/config/pay-edit?id=${id}`)
}

onMounted(async () => {
  const responseData = (await http.get('config/pay-list')) as ResponseConfigPayListItem[]
  for (let i = 0; i < responseData.length; i++) {
    const tempObj = {
      id: responseData[i].id,
      api_name: responseData[i].api_name,
      status: responseData[i].status === 1 ? '已启用' : '已禁用',
    }
    tableData.value.push(tempObj)
  }
})
</script>
