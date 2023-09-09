<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <div>
    <common-table ref="tableEle" :request-path="`conversation/topic-list`" :primary-filter-field="'title'">
    </common-table>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import CommonTable from '@/components/common-table/CommonTable.vue'

const router = useRouter()

const tableEle = ref(null)

onMounted(async () => {})
</script>
