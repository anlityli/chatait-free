<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <div>
    <common-table
      ref="tableEle"
      :request-path="`config/openai-list`"
      :operation-column="true"
      :primary-filter-field="'title'"
      :show-add-button="true"
      :fixed-columns="[{ field: 'operation', direction: 'right' }]"
      :row-select="{ enable: true, type: 'multiple' }"
      @on-add="handleAdd"
    >
      <template #selected="{ selectedRows }">
        <t-dropdown-item @click="handleDeleteDialog(selectedRows)">删除</t-dropdown-item>
      </template>
      <template #status="slotProps">
        <t-tag v-if="slotProps.params.row.status.ori_value === 1" theme="success">
          {{ slotProps.params.row.status.value }}
        </t-tag>
        <t-tag v-else theme="danger">{{ slotProps.params.row.status.value }}</t-tag>
      </template>
      <template #operation="slotProps">
        <t-dropdown :min-column-width="88" trigger="click">
          <t-button variant="text">
            操作数据
            <template #suffix>
              <chevron-down-icon></chevron-down-icon>
            </template>
          </t-button>
          <t-dropdown-menu>
            <t-dropdown-item @click="handleEdit(slotProps.params.row.id.ori_value)">编辑</t-dropdown-item>
            <t-dropdown-item @click="handleDeleteDialog([slotProps.params.row.id])">删除</t-dropdown-item>
          </t-dropdown-menu>
        </t-dropdown>
      </template>
    </common-table>
    <t-dialog
      v-model:visible="deleteDialogVisible"
      header="删除数据"
      body="删除所选数据?"
      attach="body"
      :confirm-on-enter="true"
      :on-confirm="handleDelete"
    ></t-dialog>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { ChevronDownIcon } from 'tdesign-icons-vue-next'
import CommonTable from '@/components/common-table/CommonTable.vue'
import { ResponseTableDataListValue } from '@/utils/model/response/tableData'
import http from '@/utils/network/http'

const router = useRouter()

const tableEle = ref(null)

const handleAdd = () => {
  router.push('/config/openai-add')
}

const handleEdit = (id: string) => {
  router.push(`/config/openai-edit?id=${id}`)
}

const deleteDialogVisible = ref(false)

const deleteRows = ref<ResponseTableDataListValue[]>([])
const handleDeleteDialog = (selectedRows: ResponseTableDataListValue[]) => {
  deleteRows.value = selectedRows
  deleteDialogVisible.value = true
}
const handleDelete = async () => {
  const selectedRows = deleteRows.value
  const requestData = {
    selected: <string[]>[],
  }
  for (let i = 0; i < selectedRows.length; i++) {
    requestData.selected.push(selectedRows[i].ori_value)
  }
  if (requestData.selected.length > 0) {
    await http.post('config/openai-delete', requestData)
    tableEle.value.handleRefresh()
  }
  deleteDialogVisible.value = false
}

onMounted(async () => {})
</script>
