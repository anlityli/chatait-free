<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <result title="系统维护中" tip="系统维护中，请稍后再试" type="maintenance">
    <t-button theme="primary" @click="() => $router.push('/')">返回首页</t-button>
  </result>
</template>

<script setup lang="ts">
import Result from '@/components/result/index.vue'
</script>

<script lang="ts">
export default {
  name: 'ResultMaintenance',
}
</script>
