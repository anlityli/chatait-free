<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <l-header
    v-if="settingStore.showHeader"
    :show-logo="settingStore.showHeaderLogo"
    :theme="settingStore.displayMode"
    :layout="settingStore.layout"
    :is-fixed="settingStore.isHeaderFixed"
    :menu="headerMenu"
    :is-compact="settingStore.isSidebarCompact"
  />
</template>

<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { computed } from 'vue'

import { usePermissionStore, useSettingStore } from '@/store'

import LHeader from './Header.vue'

const permissionStore = usePermissionStore()
const settingStore = useSettingStore()
const { routers: menuRouters } = storeToRefs(permissionStore)
const headerMenu = computed(() => {
  if (settingStore.layout === 'mix') {
    if (settingStore.splitMenu) {
      return menuRouters.value.map((menu) => ({
        ...menu,
        children: [],
      }))
    }
    return []
  }
  return menuRouters.value
})
</script>
