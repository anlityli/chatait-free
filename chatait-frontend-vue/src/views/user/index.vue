<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<script setup lang="ts">
import { useRoute, useRouter } from 'vue-router'
import { onMounted, ref } from 'vue'
import { WalletType } from '@/utils/constant/config'
import { useConfigStore } from '@/store'

const route = useRoute()
const router = useRouter()
const configStore = useConfigStore()

const activeTab = ref('profile')

const handleTo = (value: string) => {
  activeTab.value = value
  router.push(`${value}`)
}

onMounted(() => {
  activeTab.value = route.path
})
</script>
<template>
  <div class="user-wrap">
    <t-tabs v-model="activeTab" @change="handleTo">
      <t-tab-panel value="/user/profile" label="个人信息"></t-tab-panel>
      <t-tab-panel
        value="/user/finance/flow/balance"
        :label="`${configStore.walletName(WalletType.balance)}`"
      ></t-tab-panel>
      <t-tab-panel value="/user/finance/flow/gpt3" :label="`${configStore.walletName(WalletType.gpt3)}`"></t-tab-panel>
      <t-tab-panel value="/user/finance/flow/gpt4" :label="`${configStore.walletName(WalletType.gpt4)}`"></t-tab-panel>
      <t-tab-panel
        value="/user/finance/flow/midjourney"
        :label="`${configStore.walletName(WalletType.midjourney)}`"
      ></t-tab-panel>
    </t-tabs>
    <div class="user-content">
      <router-view :key="route.fullPath"></router-view>
    </div>
  </div>
</template>

<style lang="scss">
.user-wrap {
  display: flex;
  flex-direction: column;
  width: 100%;
  height: 100%;
  overflow: hidden;

  .t-tabs {
    flex: 0 0 auto;
  }

  .user-content {
    flex: 1 1 auto;
    width: 100%;
    height: 100%;
    overflow: auto;
  }
}
</style>
