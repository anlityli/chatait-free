<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <div>
    <common-table
      ref="tableEle"
      :request-path="`finance/wallet-list`"
      :operation-column="true"
      :primary-filter-field="'username'"
      :show-export-button="true"
      :fixed-columns="[{ field: 'operation', direction: 'right' }]"
      :row-select="{
        enable: false,
        type: 'single',
      }"
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
            <t-dropdown-item @click="handleChangeDialog(slotProps.params.row.user_id.ori_value)"
              >变更余额
            </t-dropdown-item>
          </t-dropdown-menu>
        </t-dropdown>
      </template>
    </common-table>
    <t-dialog
      v-model:visible="changeDialogVisible"
      header="变更余额"
      body="变更所选会员?"
      attach="body"
      :confirm-on-enter="true"
      :on-confirm="handleChange"
    >
      <t-form class="edit-form" :data="formData" :colon="true" @submit="handleChange">
        <t-form-item label="钱包类型" name="wallet_type">
          <t-select v-model="formData.wallet_type" class="edit-input">
            <t-option
              v-for="(item, index) in configStore.walletList"
              :key="index"
              :label="item.wallet_name"
              :value="item.field"
            />
          </t-select>
        </t-form-item>
        <t-form-item label="变更金额" name="amount">
          <t-input v-model="formData.amount" class="edit-input"></t-input>
        </t-form-item>
        <t-form-item label="备注" name="remark">
          <t-input v-model="formData.remark" class="edit-input"></t-input>
        </t-form-item>
      </t-form>
    </t-dialog>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { ChevronDownIcon } from 'tdesign-icons-vue-next'
import CommonTable from '@/components/common-table/CommonTable.vue'
import http from '@/utils/network/http'
import { FormFinanceWalletChange } from '@/pages/finance/model/model'
import { WalletType } from '@/constants/config'
import tool from '@/utils/tool/tool'
import { useConfigStore } from '@/store'

const router = useRouter()
const configStore = useConfigStore()

const tableEle = ref(null)

const formData = ref<FormFinanceWalletChange>({
  user_id: '',
  wallet_type: WalletType.balance,
  amount: '0.00',
  remark: '',
})

const changeDialogVisible = ref(false)

const handleChangeDialog = (userId: string) => {
  formData.value.user_id = userId
  formData.value.wallet_type = WalletType.balance
  formData.value.amount = '0.00'
  formData.value.remark = ''
  changeDialogVisible.value = true
}
const handleChange = async () => {
  const requestData = {
    user_id: formData.value.user_id,
    wallet_type: formData.value.wallet_type,
    amount: tool.yuanToCent(formData.value.amount),
    remark: formData.value.remark,
  }
  await http.post('finance/wallet-change', requestData)
  tableEle.value.handleRefresh()
  changeDialogVisible.value = false
}

const handleExport = () => {
  console.log('导出')
}

onMounted(async () => {})
</script>
