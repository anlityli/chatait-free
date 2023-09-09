<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <div class="main-content-container">
    <t-form class="edit-form" :data="formData" :colon="true" @submit="handleSubmit">
      <t-form-item label="角色名称" name="role_name">
        <t-input v-model="formData.role_name" class="edit-input"></t-input>
      </t-form-item>
      <t-form-item label="描述" name="remark">
        <t-input v-model="formData.remark" class="edit-input"></t-input>
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
import { FormAdminRoleEdit } from '@/pages/admin/model/model'
import http from '@/utils/network/http'

const router = useRouter()
const route = useRoute()

const formData = ref<FormAdminRoleEdit>({
  id: '',
  role_name: '',
  remark: '',
})

const handleSubmit = async () => {
  try {
    if (route.name === 'adminRoleEdit') {
      await http.post('admin/role-edit', formData.value)
    } else {
      await http.post('admin/role-add', formData.value)
    }
    await MessagePlugin.success('操作成功')
    await router.back()
  } catch (e) {
    await MessagePlugin.success(`发生错误: ${e.toString()}`)
  }
}

onMounted(async () => {
  if (route.name === 'adminRoleEdit') {
    const adminData = await http.get(`admin/role-one?id=${route.query.id}`)
    formData.value.id = route.query.id as string
    formData.value.role_name = adminData.role_name
    formData.value.remark = adminData.remark
  }
})
</script>
