<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <div class="main-content-container">
    <t-form class="edit-form" :data="formData" :colon="true" @submit="handleSubmit">
      <t-form-item label="管理员登录名" name="admin_name">
        <t-input v-model="formData.admin_name" class="edit-input"></t-input>
      </t-form-item>
      <t-form-item label="管理员姓名" name="real_name" help="管理员的真实姓名">
        <t-input v-model="formData.real_name" class="edit-input"></t-input>
      </t-form-item>
      <t-form-item label="备注" name="remark">
        <t-input v-model="formData.remark" class="edit-input"></t-input>
      </t-form-item>
      <t-form-item label="角色" name="role_id">
        <t-select v-model="formData.role_id" class="edit-input">
          <t-option v-for="(item, index) in allRole" :key="index" :label="item.role_name" :value="item.id" />
        </t-select>
      </t-form-item>
      <t-form-item label="是否启用" name="is_enable">
        <t-switch v-model="formData.is_enable" :custom-value="[1, 0]"></t-switch>
      </t-form-item>
      <t-form-item label="密码" name="password" help="如不修改不要填写">
        <t-input v-model="formData.password" type="password" class="edit-input"></t-input>
      </t-form-item>
      <t-form-item>
        <t-space size="small">
          <t-button theme="primary" type="submit">提交</t-button>
        </t-space>
      </t-form-item>
    </t-form>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { MessagePlugin } from 'tdesign-vue-next'
import { useRoute, useRouter } from 'vue-router'
import { FormAdminEdit } from '@/pages/admin/model/model'
import http from '@/utils/network/http'
import { ResponseAdminItem, ResponseAdminRoleItem } from '@/utils/model/response/admin'

const router = useRouter()
const route = useRoute()

const formData = ref<FormAdminEdit>({
  id: '',
  user_id: '0',
  admin_name: '',
  real_name: '',
  remark: '',
  role_id: '1',
  is_enable: 1,
  password: '',
  bind_ip: [],
})

const allRole = ref<ResponseAdminRoleItem[]>([])

const handleSubmit = async () => {
  try {
    if (route.name === 'adminEdit') {
      await http.post('admin/edit', formData.value)
    } else {
      await http.post('admin/add', formData.value)
    }
    await MessagePlugin.success('操作成功')
    await router.back()
  } catch (e) {
    await MessagePlugin.success(`发生错误: ${e.toString()}`)
  }
}

const handleAllRole = async () => {
  allRole.value = (await http.get(`admin/all-role`)) as ResponseAdminRoleItem[]
  console.log(allRole)
}

onMounted(async () => {
  await handleAllRole()
  if (route.name === 'adminEdit') {
    const adminData = (await http.get(`admin/one?id=${route.query.id}`)) as ResponseAdminItem
    formData.value.id = route.query.id as string
    formData.value.user_id = adminData.user_id
    formData.value.admin_name = adminData.admin_name
    formData.value.real_name = adminData.real_name
    formData.value.remark = adminData.remark
    formData.value.role_id = adminData.role_id
    formData.value.is_enable = adminData.is_enable
  }
})
</script>
