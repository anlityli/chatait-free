<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<script setup lang="ts">
import '@/assets/style/main.scss'
import { onMounted } from 'vue'
import { getAppStore, getConfigStore } from '@/store'

const appStore = getAppStore()
const configStore = getConfigStore()

appStore.updateMode()
onMounted(() => {
  // document.documentElement.setAttribute('theme-mode', appStore.theme)
  configStore.getWalletList()
})
</script>

<template>
  <div class="app">
    <router-view />
  </div>
</template>

<style lang="scss" scoped>
.app {
  height: 100%;
}
</style>
