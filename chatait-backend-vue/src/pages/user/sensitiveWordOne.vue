<!--
  - Copyright 2024 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <div class="main-content-container">
    <t-form class="edit-form" :colon="true">
      <t-form-item label="会员名" name="title">
        {{ sensitiveWordData.username }}
      </t-form-item>
      <t-form-item label="昵称" name="title">
        {{ sensitiveWordData.nickname }}
      </t-form-item>
      <t-form-item label="提交原文" name="title">
        {{ sensitiveWordData.content }}
      </t-form-item>
      <t-form-item label="触发规则Json" name="title">
        {{ sensitiveWordData.validate_result }}
      </t-form-item>
    </t-form>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import http from '@/utils/network/http'
import { ResponseUserSensitiveWord } from '@/utils/model/response/user'

const router = useRouter()
const route = useRoute()

const sensitiveWordData = ref<ResponseUserSensitiveWord>({
  id: '',
  username: '',
  nickname: '',
  type: 1,
  topic_type: 1,
  content: '',
  validate_result: '',
  created_at: 0,
})

onMounted(async () => {
  sensitiveWordData.value = (await http.get(
    `user/sensitive-word-one?id=${route.query.id}`,
  )) as ResponseUserSensitiveWord
})
</script>
