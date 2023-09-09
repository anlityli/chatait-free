<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<script setup lang="ts">
import { useRoute, useRouter } from 'vue-router'
import { onMounted, ref } from 'vue'

const route = useRoute()
const router = useRouter()

const activeTab = ref('goodsList')

const handleTo = (value: string) => {
  activeTab.value = value
  router.push(`${value}`)
}

onMounted(() => {
  activeTab.value = route.path
  if (route.name === 'purchaseConfirmOrder') {
    activeTab.value = '/purchase/goods-list'
  } else if (route.name === 'purchasePayOrder') {
    activeTab.value = '/purchase/goods-list'
  } else if (route.name === 'purchaseOrderDetail') {
    activeTab.value = '/purchase/order-list'
  }
})
</script>
<template>
  <div class="purchase-wrap">
    <t-tabs v-model="activeTab" @change="handleTo">
      <t-tab-panel value="/purchase/goods-list" label="计划列表"></t-tab-panel>
      <t-tab-panel value="/purchase/order-list" label="订单列表"></t-tab-panel>
    </t-tabs>
    <div class="purchase-content">
      <router-view :key="route.fullPath"></router-view>
    </div>
  </div>
</template>

<style lang="scss">
.purchase-wrap {
  display: flex;
  flex-direction: column;
  width: 100%;
  height: 100%;
  overflow: hidden;

  .t-tabs {
    flex: 0 0 auto;
  }

  .purchase-content {
    flex: 1 1 auto;
    width: 100%;
    height: 100%;
    overflow: auto;
    box-sizing: border-box;
    padding: 10px;
  }
}
</style>
