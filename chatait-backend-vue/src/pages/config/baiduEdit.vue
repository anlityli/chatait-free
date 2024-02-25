<!--
  - Copyright 2024 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <div class="main-content-container">
    <t-form class="edit-form" :data="formData" :colon="true" @submit="handleSubmit">
      <t-form-item label="配置标题" name="title">
        <t-input v-model="formData.title" class="edit-input"></t-input>
      </t-form-item>
      <t-form-item label="Api Key" name="api_key">
        <t-input v-model="formData.api_key" class="edit-input"></t-input>
      </t-form-item>
      <t-form-item label="Secret Key" name="secret_key">
        <t-input v-model="formData.secret_key" class="edit-input"></t-input>
      </t-form-item>
      <t-form-item label="是否启用" name="status">
        <t-switch v-model="formData.status" :custom-value="[1, 0]"></t-switch>
      </t-form-item>
      <t-form-item label="开通功能" name="status">
        <t-checkbox-group v-model="formData.features" :options="featuresOption"></t-checkbox-group>
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
import { CheckboxOption, MessagePlugin } from 'tdesign-vue-next'
import { useRoute, useRouter } from 'vue-router'
import http from '@/utils/network/http'
import { FormBaiduEdit } from '@/pages/config/model/baidu'
import { BaiduFeature } from '@/constants/config'

const router = useRouter()
const route = useRoute()

const featuresOption = ref<CheckboxOption[]>([
  { label: '全选', checkAll: true },
  {
    value: BaiduFeature.translate,
    label: '翻译',
  },
  { value: BaiduFeature.censor, label: '内容审核' },
])
const formData = ref<FormBaiduEdit>({
  id: '',
  title: '',
  api_key: '',
  secret_key: '',
  status: 1,
  features: [],
})

const handleSubmit = async () => {
  if (formData.value.features.length <= 0) {
    await MessagePlugin.error(`发生错误: 开通功能必选`)
    return
  }
  try {
    if (route.name === 'configBaiduEdit') {
      await http.post('config/baidu-edit', formData.value)
    } else {
      await http.post('config/baidu-add', formData.value)
    }
    await MessagePlugin.success('操作成功')
    await router.back()
  } catch (e) {
    await MessagePlugin.error(`发生错误: ${e.toString()}`)
  }
}

onMounted(async () => {
  if (route.name === 'configBaiduEdit') {
    const responseData = (await http.get(`config/baidu-one?id=${route.query.id}`)) as any
    Object.keys(responseData).forEach((key) => {
      const formKey = key as keyof FormBaiduEdit
      if (formKey in formData.value) {
        formData.value[formKey] = responseData[key] as never
      }
      if (formKey === 'features' && formData.value[formKey].length <= 0) {
        formData.value[formKey].push(BaiduFeature.translate)
        formData.value[formKey].push(BaiduFeature.censor)
      }
    })
  }
})
</script>
