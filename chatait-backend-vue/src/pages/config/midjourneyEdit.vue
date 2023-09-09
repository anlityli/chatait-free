<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <div class="main-content-container">
    <t-form class="edit-form" :data="formData" :colon="true" @submit="handleSubmit">
      <t-form-item label="配置标题" name="title">
        <t-input v-model="formData.title" class="edit-input"></t-input>
      </t-form-item>
      <t-form-item label="服务ID" name="guild_id" help="即: guild_id">
        <t-input v-model="formData.guild_id" class="edit-input"></t-input>
      </t-form-item>
      <t-form-item label="频道ID" name="channel_id" help="即: channel_id">
        <t-input v-model="formData.channel_id" class="edit-input"></t-input>
      </t-form-item>
      <t-form-item label="账户Token" name="user_token">
        <t-input v-model="formData.user_token" class="edit-input"></t-input>
      </t-form-item>
      <t-form-item label="MidjourneyBotID" name="mj_bot_id" help="注意：是Midjourney的bot_id">
        <t-input v-model="formData.mj_bot_id" class="edit-input"></t-input>
      </t-form-item>
      <t-form-item label="BotToken" name="bot_token">
        <t-input v-model="formData.bot_token" class="edit-input"></t-input>
      </t-form-item>
      <t-form-item label="会话ID" name="session_id" help="即: session_id">
        <t-input v-model="formData.session_id" class="edit-input"></t-input>
      </t-form-item>
      <t-form-item label="UserAgent" name="user_agent">
        <t-input v-model="formData.user_agent" class="edit-input"></t-input>
      </t-form-item>
      <t-form-item label="HuggingFaceToken" name="hugging_face_token">
        <t-input v-model="formData.hugging_face_token" class="edit-input"></t-input>
      </t-form-item>
      <t-form-item label="代理地址" name="proxy">
        <t-input v-model="formData.proxy" class="edit-input"></t-input>
      </t-form-item>
      <t-form-item label="是否启用" name="status">
        <t-switch v-model="formData.status" :custom-value="[1, 0]"></t-switch>
      </t-form-item>
      <t-form-item label="监听模式" name="listen_model">
        <t-select v-model="formData.listen_model" class="edit-input">
          <t-option :value="1" label="UserWss"></t-option>
          <t-option :value="0" label="Bot"></t-option>
        </t-select>
      </t-form-item>
      <t-form-item label="生图模式" name="create_model">
        <t-select v-model="formData.create_model" class="edit-input">
          <t-option :value="MidjourneyCreateModel.fast" label="fast"></t-option>
          <t-option :value="MidjourneyCreateModel.relax" label="relax"></t-option>
          <t-option :value="MidjourneyCreateModel.turbo" label="turbo"></t-option>
        </t-select>
      </t-form-item>
      <t-form-item label="Websocket闲置时长" name="proxy" help="闲置该时长后断开websocket连接，有任务后重连">
        <t-input-number v-model="formData.ws_idle_time" class="edit-input" theme="normal">
          <template #suffix>
            {{ '秒' }}
          </template>
        </t-input-number>
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
import { FormMidjourneyEdit } from '@/pages/config/model/midjourney'
import { MidjourneyCreateModel } from '@/constants/config'

const router = useRouter()
const route = useRoute()

const formData = ref<FormMidjourneyEdit>({
  id: '',
  title: '',
  guild_id: '',
  channel_id: '',
  user_token: '',
  mj_bot_id: '',
  bot_token: '',
  session_id: '',
  user_agent:
    'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36',
  hugging_face_token: '',
  proxy: '',
  status: 1,
  listen_model: 1,
  create_model: MidjourneyCreateModel.fast,
  ws_idle_time: 3600,
})

const handleSubmit = async () => {
  try {
    if (route.name === 'configMidjourneyEdit') {
      await http.post('config/midjourney-edit', formData.value)
    } else {
      await http.post('config/midjourney-add', formData.value)
    }
    await MessagePlugin.success('操作成功')
    await router.back()
  } catch (e) {
    await MessagePlugin.success(`发生错误: ${e.toString()}`)
  }
}

onMounted(async () => {
  if (route.name === 'configMidjourneyEdit') {
    const responseData = (await http.get(`config/midjourney-one?id=${route.query.id}`)) as any
    Object.keys(responseData).forEach((key) => {
      const formKey = key as keyof FormMidjourneyEdit
      if (formKey in formData.value) {
        formData.value[formKey] = responseData[key] as never
      }
    })
  }
})
</script>
