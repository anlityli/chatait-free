<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <router-view :class="[mode]" />
</template>
<script setup lang="ts">
import { computed, onMounted } from 'vue'

import { useConfigStore, useSettingStore } from '@/store'

const store = useSettingStore()
const configStore = useConfigStore()

const mode = computed(() => {
  return store.displayMode
})

onMounted(() => {
  configStore.getWalletList()
})
</script>
<style lang="less" scoped>
#nprogress .bar {
  background: var(--td-brand-color) !important;
}
</style>
