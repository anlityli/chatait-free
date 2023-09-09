<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <result title="网络异常" tip="网络异常，请稍后再试" type="wifi">
    <div>
      <t-button theme="default" @click="() => $router.push('/')">返回首页</t-button>
      <t-button @click="() => $router.push('/')">重新加载</t-button>
    </div>
  </result>
</template>

<script setup lang="ts">
import Result from '@/components/result/index.vue'
</script>

<script lang="ts">
export default {
  name: 'ResultNetworkError',
}
</script>
