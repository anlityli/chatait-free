<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <div>
    <top-panel class="row-container" />
    <order-list class="row-container"></order-list>
  </div>
</template>

<script setup lang="ts">
import TopPanel from '@/pages/dashboard/components/TopPanel.vue'
import OrderList from '@/pages/dashboard/components/OrderList.vue'
</script>

<script lang="ts">
export default {
  name: 'DashboardBase',
}
</script>

<style scoped>
.row-container:not(:last-child) {
  margin-bottom: 16px;
}
</style>
