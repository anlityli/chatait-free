<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <result title="500 Internal Server Error" type="500" tip="抱歉，服务器出错啦">
    <t-button @click="() => $router.push('/')">返回首页</t-button>
  </result>
</template>

<script setup lang="ts">
import Result from '@/components/result/index.vue'
</script>

<script lang="ts">
export default {
  name: 'Result500',
}
</script>
