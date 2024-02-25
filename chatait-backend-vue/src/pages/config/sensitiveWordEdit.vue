<!--
  - Copyright 2024 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <div class="main-content-container">
    <t-form class="edit-form" :data="formData" :colon="true" @submit="handleSubmit">
      <t-form-item label="敏感词" name="title">
        <t-input v-model="formData.content" class="edit-input"></t-input>
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
import { ref } from 'vue'
import { MessagePlugin } from 'tdesign-vue-next'
import { useRouter } from 'vue-router'
import http from '@/utils/network/http'
import { FormSensitiveWordEdit } from '@/pages/config/model/sensitiveWord'

const router = useRouter()

const formData = ref<FormSensitiveWordEdit>({
  id: '',
  content: '',
})

const handleSubmit = async () => {
  try {
    await http.post('config/sensitive-word-add', formData.value)
    await MessagePlugin.success('操作成功')
    await router.back()
  } catch (e) {
    await MessagePlugin.error(`发生错误: ${e.toString()}`)
  }
}
</script>
