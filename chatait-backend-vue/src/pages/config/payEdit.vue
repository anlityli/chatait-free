<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <div class="main-content-container">
    <t-form class="edit-form" :data="formData" :colon="true" @submit="handleSubmit">
      <t-form-item v-for="(item, index) in formData.params" :key="index" :label="item.param_name" :name="item.param">
        <t-input v-model="item.value" class="edit-input"></t-input>
      </t-form-item>
      <t-form-item label="是否启用" name="status">
        <t-switch v-model="formData.status" :custom-value="[1, 0]"></t-switch>
      </t-form-item>
      <t-divider>支付渠道</t-divider>
      <t-form-item
        v-for="(item, index) in formData.pay_channel"
        :key="index"
        :label="item.channel_name"
        :name="item.channel"
      >
        <t-switch v-model="item.status" :custom-value="[1, 0]"></t-switch>
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
import http from '@/utils/network/http'
import { ResponseConfigPayListItem } from '@/utils/model/response/config'
import { FormPayEdit } from '@/pages/config/model/pay'

const router = useRouter()
const route = useRoute()

const formData = ref<FormPayEdit>({
  id: '',
  params: [],
  pay_channel: [],
  frontend_description: '',
  status: 1,
})

const handleSubmit = async () => {
  try {
    await http.post('config/pay-edit', formData.value)
    await MessagePlugin.success('操作成功')
    await router.back()
  } catch (e) {
    await MessagePlugin.success(`发生错误: ${e.toString()}`)
  }
}

onMounted(async () => {
  const responseData = (await http.get(`config/pay-one?id=${route.query.id}`)) as ResponseConfigPayListItem
  formData.value.id = route.query.id as string
  formData.value.params = responseData.params
  formData.value.pay_channel = responseData.pay_channel
  formData.value.frontend_description = responseData.frontend_description
  formData.value.status = responseData.status
})
</script>
