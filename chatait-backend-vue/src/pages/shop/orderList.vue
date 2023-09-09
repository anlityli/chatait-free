<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <div>
    <common-table
      ref="tableEle"
      :request-path="`shop/order-list`"
      :operation-column="true"
      :primary-filter-field="'order_sn'"
      :show-export-button="true"
      :fixed-columns="[{ field: 'operation', direction: 'right' }]"
      @on-export="handleExport"
    >
      <template #operation="slotProps">
        <t-dropdown :min-column-width="88" trigger="click">
          <t-button variant="text">
            操作数据
            <template #suffix>
              <chevron-down-icon></chevron-down-icon>
            </template>
          </t-button>
          <t-dropdown-menu>
            <t-dropdown-item
              v-if="slotProps.params.row.status.ori_value === 0"
              @click="handleSetStatusDialog(slotProps.params.row.id.ori_value, 1)"
              >设为已支付
            </t-dropdown-item>
            <t-dropdown-item
              v-if="slotProps.params.row.status.ori_value === 1"
              @click="handleSetStatusDialog(slotProps.params.row.id.ori_value, 4)"
              >设为已完成
            </t-dropdown-item>
          </t-dropdown-menu>
        </t-dropdown>
      </template>
    </common-table>
    <t-dialog
      v-model:visible="setStatusDialogVisible"
      header="修改数据"
      body="修改所选数据?"
      attach="body"
      :confirm-on-enter="true"
      :on-confirm="handleSetStatus"
    ></t-dialog>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { ChevronDownIcon } from 'tdesign-icons-vue-next'
import CommonTable from '@/components/common-table/CommonTable.vue'
import http from '@/utils/network/http'

const router = useRouter()

const tableEle = ref(null)

const setStatusDialogVisible = ref(false)

const setStatusParam = ref({
  id: '',
  status: 0,
})
const handleSetStatusDialog = (id: string, status: number) => {
  setStatusParam.value.id = id
  setStatusParam.value.status = status
  setStatusDialogVisible.value = true
}
const handleSetStatus = async () => {
  await http.post('shop/order-status', setStatusParam.value)
  tableEle.value.handleRefresh()
  setStatusDialogVisible.value = false
}

const handleExport = () => {
  console.log('导出')
}

onMounted(async () => {})
</script>
