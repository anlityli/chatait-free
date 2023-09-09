<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { ArrowLeftIcon, ViewListIcon } from 'tdesign-icons-vue-next'
import NMenu from '@/components/menu/NMenu.vue'
import tool from '@/utils/tool/tool'
import { useWebsocketStore } from '@/store/modules/websocket'

const router = useRouter()
const route = useRoute()
const websocket = useWebsocketStore()

const currentYear = new Date().getFullYear()

const pageLoading = ref(true)
const drawerVisible = ref(false)

const handleDrawer = () => {
  drawerVisible.value = !drawerVisible.value
}

const handleClickItem = () => {
  drawerVisible.value = false
}

onMounted(async () => {
  if (import.meta.env.VITE_DEFAULT_HOMEPAGE === 'login' && !tool.isLogin()) {
    await router.push('/login')
  }
  await websocket.connect()
  pageLoading.value = false
  // console.log(route.name)
})
</script>

<template>
  <t-layout class="t-wrap">
    <t-drawer
      v-model:visible="drawerVisible"
      class="t-drawer"
      :close-btn="true"
      :footer="false"
      :header="false"
      :placement="'left'"
      :size="'232px'"
    >
      <n-menu :is-drawer="true" @click-item="handleClickItem"></n-menu>
    </t-drawer>
    <t-aside class="t-aside">
      <n-menu></n-menu>
    </t-aside>
    <t-layout>
      <t-header class="t-header">
        <div class="header-left">
          <t-button shape="square" variant="outline" @click="handleDrawer">
            <view-list-icon></view-list-icon>
          </t-button>
        </div>
        <div class="header-title">
          <arrow-left-icon></arrow-left-icon>
          <span>点击菜单查看话题</span>
        </div>
        <div class="header-right"></div>
      </t-header>
      <t-content class="t-content">
        <router-view :key="route.fullPath" />
      </t-content>
      <t-footer class="t-footer" :style="route.name === 'conversationTopic' ? 'background-color: #f6f6fd' : ''"
        >Copyright © 2022-{{ currentYear.toString() }}
        ChatAIT. All Rights Reserved
      </t-footer>
    </t-layout>
  </t-layout>
</template>

<style lang="scss">
.t-layout {
  background: none !important;
  overflow: hidden !important;
}

.t-wrap {
  .t-drawer .t-drawer__body {
    padding: 0 !important;
  }

  .t-drawer .t-drawer__close-btn {
    z-index: 1;
  }
}
</style>

<style lang="scss" scoped>
@media screen and (min-width: 768px) {
  .t-header {
    display: none;
  }
}

@media screen and (max-width: 768px) {
  .t-aside {
    display: none;
  }
  .t-header {
    display: flex;
  }
}

.t-wrap {
  width: 100%;
  height: 100%;

  .t-header {
    width: 100%;
    height: 47px !important;
    border-bottom: 1px solid var(--td-component-stroke);

    .header-left,
    .header-right {
      flex: 0 0 auto;
      display: flex;
      width: 47px;
      height: 47px;
      box-sizing: border-box;
      justify-content: center;
      align-items: center;
    }

    .header-title {
      flex: 1 1 auto;
      display: flex;
      width: 100%;
      box-sizing: border-box;
      justify-content: left;
      align-items: center;
    }
  }

  .t-aside {
    border-right: 1px solid var(--td-component-stroke);
  }

  .t-content {
    flex: 1 1 auto;
    height: 100%;
    overflow: hidden;
  }

  .t-footer {
    flex: 0 0 auto;
    height: 30px;
    padding: 5px;
    text-align: center;
    line-height: 30px;
    overflow: hidden;
  }
}
</style>
