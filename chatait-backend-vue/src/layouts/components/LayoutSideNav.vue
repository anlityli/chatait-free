<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <l-side-nav
    v-if="settingStore.showSidebar"
    :show-logo="settingStore.showSidebarLogo"
    :layout="settingStore.layout"
    :is-fixed="settingStore.isSidebarFixed"
    :menu="sideMenu"
    :theme="settingStore.displayMode"
    :is-compact="settingStore.isSidebarCompact"
  />
</template>

<script setup lang="ts">
import { storeToRefs } from 'pinia'
import { computed } from 'vue'
import { useRoute } from 'vue-router'

import { usePermissionStore, useSettingStore } from '@/store'
import type { MenuRoute } from '@/types/interface'

import LSideNav from './SideNav.vue'

const route = useRoute()
const permissionStore = usePermissionStore()
const settingStore = useSettingStore()
const { routers: menuRouters } = storeToRefs(permissionStore)

const sideMenu = computed(() => {
  const { layout, splitMenu } = settingStore
  let newMenuRouters = menuRouters.value as Array<MenuRoute>
  if (layout === 'mix' && splitMenu) {
    newMenuRouters.forEach((menu) => {
      if (route.path.indexOf(menu.path) === 0) {
        newMenuRouters = menu.children.map((subMenu) => ({ ...subMenu, path: `${menu.path}/${subMenu.path}` }))
      }
    })
  }
  return newMenuRouters
})
</script>
