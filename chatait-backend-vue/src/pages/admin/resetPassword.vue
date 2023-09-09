<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <div class="main-content-container">
    <t-form
      class="edit-form"
      :data="formData"
      :colon="true"
      :rules="rules"
      @submit="handleSubmit"
      @validate="onValidate"
    >
      <t-form-item label="密码" name="password">
        <t-input v-model="formData.password" type="password" class="edit-input"></t-input>
      </t-form-item>
      <t-form-item label="确认密码" name="re_password">
        <t-input v-model="formData.re_password" type="password" class="edit-input"></t-input>
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
import { FormRule, MessagePlugin } from 'tdesign-vue-next'
import { useRoute, useRouter } from 'vue-router'
import { CustomValidateResolveType } from 'tdesign-vue-next/es/form/type'
import { FormAdminResetPassword } from '@/pages/admin/model/model'
import http from '@/utils/network/http'

const router = useRouter()
const route = useRoute()

const formData = ref<FormAdminResetPassword>({
  password: '',
  re_password: '',
})

const rePassword = (val: string): Promise<CustomValidateResolveType> => {
  return new Promise((resolve) => {
    const timer = setTimeout(() => {
      resolve(formData.value.password === val)
      clearTimeout(timer)
    })
  })
}

const onValidate = (params: any) => {
  if (params.validateResult === true) {
    console.log('Validate Success')
  } else {
    console.log('Validate Errors: ', params.firstError, params.validateResult)
  }
}

const rules: Record<string, FormRule[]> = {
  password: [{ required: true, message: '密码必填', type: 'error' }],
  re_password: [
    { required: true, message: '邮箱必填', type: 'error' },
    { validator: rePassword, message: '两次密码不一致', type: 'error' },
  ],
}

const handleSubmit = async (params: any) => {
  params.e.preventDefault()
  if (params.validateResult === true) {
    try {
      await http.post('admin/reset-password', formData.value)
      await MessagePlugin.success('操作成功')
      await router.back()
    } catch (e) {
      await MessagePlugin.success(`发生错误: ${e.toString()}`)
    }
  } else {
    await MessagePlugin.warning(params.firstError)
  }
}

onMounted(async () => {})
</script>
