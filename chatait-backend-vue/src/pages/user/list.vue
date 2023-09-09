<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <div>
    <common-table
      ref="tableEle"
      :request-path="`user/list`"
      :operation-column="true"
      :primary-filter-field="'username'"
      :show-export-button="true"
      :fixed-columns="[{ field: 'operation', direction: 'right' }]"
      @on-export="handleExport"
    >
      <template #selected="{ selectedRows }">
        <t-dropdown-item @click="handleBanDialog(selectedRows, 1)">禁用</t-dropdown-item>
      </template>
      <template #avatar="slotProps">
        <t-avatar>
          <template #content>
            <avatar-image :url="slotProps.params.row.avatar.ori_value"></avatar-image>
          </template>
        </t-avatar>
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
            <t-dropdown-item
              v-if="slotProps.params.row['is_ban'].ori_value === 1"
              @click="handleBanDialog([slotProps.params.row.id], 0)"
            >
              解除禁用
            </t-dropdown-item>
            <t-dropdown-item
              v-if="slotProps.params.row['is_ban'].ori_value === 0"
              @click="handleBanDialog([slotProps.params.row.id], 1)"
            >
              设为禁用
            </t-dropdown-item>
            <t-dropdown-item @click="handleChangeLevelDialog(slotProps.params.row)">修改会员级别</t-dropdown-item>
            <t-dropdown-item @click="handleChangePasswordDialog(slotProps.params.row)">重置会员密码</t-dropdown-item>
          </t-dropdown-menu>
        </t-dropdown>
      </template>
    </common-table>
    <t-dialog
      v-model:visible="banDialogVisible"
      header="禁用会员"
      body="禁用所选会员?"
      attach="body"
      :confirm-on-enter="true"
      :on-confirm="handleBan"
    ></t-dialog>
    <t-dialog
      v-model:visible="changeLevelDialogVisible"
      header="修改会员级别"
      :confirm-on-enter="true"
      :on-confirm="handleChangeLevel"
    >
      <t-form class="edit-form" :data="changeLevelForm" :colon="true" @submit="handleChangeLevel">
        <t-form-item label="会员级别" name="level_id">
          <t-select v-model="changeLevelForm.level_id">
            <t-option v-for="(item, index) in allLevel" :key="index" :label="item.level_name" :value="item.id" />
          </t-select>
        </t-form-item>
        <t-form-item v-show="changeLevelForm.level_id !== 1" label="有效期" name="level_expire_date">
          <t-date-picker v-model="changeLevelForm.level_expire_date" format="YYYY-MM-DD"></t-date-picker>
        </t-form-item>
        <t-form-item label="备注" name="remark">
          <t-input v-model="changeLevelForm.remark"></t-input>
        </t-form-item>
      </t-form>
    </t-dialog>
    <t-dialog
      v-model:visible="changePasswordDialogVisible"
      header="重置会员密码"
      :confirm-on-enter="true"
      :on-confirm="handleChangePassword"
    >
      <t-form class="edit-form" :data="changePasswordForm" :colon="true" @submit="handleChangePassword">
        <t-form-item label="密码" name="password">
          <t-input v-model="changePasswordForm.password" type="password"></t-input>
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
import { ResponseTableDataListValue } from '@/utils/model/response/tableData'
import http from '@/utils/network/http'
import { UserBanForm, UserChangeLevelForm, UserChangePasswordForm } from '@/pages/user/model/model'
import { ResponseConfigLevelListItem } from '@/utils/model/response/config'
import AvatarImage from '@/components/avatar-image/AvatarImage.vue'

const router = useRouter()

const tableEle = ref(null)

const banDialogVisible = ref(false)

const banForm = ref<UserBanForm>({
  selected: [],
  is_ban: 0,
})
const handleBanDialog = (selectedRows: ResponseTableDataListValue[], isBan: number) => {
  banForm.value.selected = []
  for (let i = 0; i < selectedRows.length; i++) {
    banForm.value.selected.push(selectedRows[i].ori_value)
  }
  banForm.value.is_ban = isBan
  banDialogVisible.value = true
}
const handleBan = async () => {
  if (banForm.value.selected.length > 0) {
    await http.post('user/ban', banForm.value)
    tableEle.value.handleRefresh()
  }
  banDialogVisible.value = false
}

const changeLevelDialogVisible = ref(false)
const changeLevelForm = ref<UserChangeLevelForm>({
  user_id: '',
  level_id: 1,
  level_expire_date: '',
  remark: '',
})
const allLevel = ref<ResponseConfigLevelListItem[]>([])
const handleChangeLevelDialog = (row: any) => {
  changeLevelForm.value.user_id = row.id.ori_value
  changeLevelForm.value.level_id = row.level_id.ori_value
  changeLevelForm.value.level_expire_date = row.level_expire_date.ori_value
  changeLevelForm.value.remark = ''
  changeLevelDialogVisible.value = true
}

const handleChangeLevel = async () => {
  await http.post('user/change-level', changeLevelForm.value)
  tableEle.value.handleRefresh()
  changeLevelDialogVisible.value = false
}

const changePasswordDialogVisible = ref(false)
const changePasswordForm = ref<UserChangePasswordForm>({
  user_id: '',
  password: '',
})

const handleChangePasswordDialog = (row: any) => {
  changePasswordForm.value.user_id = row.id.ori_value
  changePasswordForm.value.password = ''
  changePasswordDialogVisible.value = true
}

const handleChangePassword = async () => {
  await http.post('user/reset-password', changePasswordForm.value)
  tableEle.value.handleRefresh()
  changePasswordDialogVisible.value = false
}

const handleExport = () => {
  console.log('导出')
}

onMounted(async () => {
  allLevel.value = (await http.get(`config/level-list`)) as ResponseConfigLevelListItem[]
})
</script>
