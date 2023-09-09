<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <div :class="sideNavCls">
    <t-menu :class="menuCls" :theme="theme" :value="active" :collapsed="collapsed" :default-expanded="defaultExpanded">
      <template #logo>
        <span v-if="showLogo" :class="`${prefix}-side-nav-logo-wrapper`" @click="goHome">
          <img v-if="collapsed" src="@/assets/logo.png" class="logo" style="object-fit: cover; width: 30px" />
          <img v-else src="@/assets/logo-full.png" class="logo" style="object-fit: cover; width: 130px" />
        </span>
      </template>
      <menu-content :nav-data="menu" />
      <template #operations>
        <span class="version-container"> {{ !collapsed ? 'Chat AIT' : '' }} {{ pgk.version }} </span>
      </template>
    </t-menu>
    <div :class="`${prefix}-side-nav-placeholder${collapsed ? '-hidden' : ''}`"></div>
  </div>
</template>

<script setup lang="ts">
import union from 'lodash/union'
import type { PropType } from 'vue'
import { computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { prefix } from '@/config/global'
import { getActive, getRoutesExpanded } from '@/router'
import { useSettingStore } from '@/store'
import type { MenuRoute } from '@/types/interface'

import pgk from '../../../package.json'
import MenuContent from './MenuContent.vue'

const MIN_POINT = 992 - 1

const props = defineProps({
  menu: {
    type: Array as PropType<MenuRoute[]>,
    default: () => [],
  },
  showLogo: {
    type: Boolean as PropType<boolean>,
    default: true,
  },
  isFixed: {
    type: Boolean as PropType<boolean>,
    default: true,
  },
  layout: {
    type: String as PropType<string>,
    default: '',
  },
  headerHeight: {
    type: String as PropType<string>,
    default: '64px',
  },
  theme: {
    type: String as PropType<'light' | 'dark'>,
    default: 'light',
  },
  isCompact: {
    type: Boolean as PropType<boolean>,
    default: false,
  },
})

const collapsed = computed(() => useSettingStore().isSidebarCompact)

const active = computed(() => getActive())

const defaultExpanded = computed(() => {
  const path = getActive()
  const parentPath = path.substring(0, path.lastIndexOf('/'))
  const expanded = getRoutesExpanded()
  return union(expanded, parentPath === '' ? [] : [parentPath])
})

const sideNavCls = computed(() => {
  const { isCompact } = props
  return [
    `${prefix}-sidebar-layout`,
    {
      [`${prefix}-sidebar-compact`]: isCompact,
    },
  ]
})

const menuCls = computed(() => {
  const { showLogo, isFixed, layout } = props
  return [
    `${prefix}-side-nav`,
    {
      [`${prefix}-side-nav-no-logo`]: !showLogo,
      [`${prefix}-side-nav-no-fixed`]: !isFixed,
      [`${prefix}-side-nav-mix-fixed`]: layout === 'mix' && isFixed,
    },
  ]
})

const router = useRouter()
const settingStore = useSettingStore()

const autoCollapsed = () => {
  const isCompact = window.innerWidth <= MIN_POINT
  settingStore.updateConfig({
    isSidebarCompact: isCompact,
  })
}

onMounted(() => {
  autoCollapsed()
  window.onresize = () => {
    autoCollapsed()
  }
})

const goHome = () => {
  router.push('/dashboard/index')
}
</script>

<style lang="less" scoped></style>
