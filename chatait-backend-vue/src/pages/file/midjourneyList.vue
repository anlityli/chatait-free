<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <div>
    <common-table
      ref="tableEle"
      :request-path="`file/midjourney-list`"
      :primary-filter-field="'file_name'"
      enable-expand
    >
      <template #path="slotProps">
        <t-image-viewer :images="[`${host}/file/midjourney-image?id=${slotProps.params.row.id.ori_value}`]">
          <template #trigger="{ open }">
            <t-image
              class="conversation-content-midjourney-image"
              :src="`${host}/file/midjourney-image?id=${slotProps.params.row.id.ori_value}`"
              shape="round"
              fit="contain"
              position="left"
              @click="open"
            />
          </template>
        </t-image-viewer>
      </template>
      <template #expandedRow="slotProps"
        ><span style="font-weight: bold">prompt:</span>
        {{ slotProps.params.row.prompt.value }}
      </template>
    </common-table>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import CommonTable from '@/components/common-table/CommonTable.vue'

const router = useRouter()

const tableEle = ref(null)
const host = import.meta.env.VITE_HTTP_HOST

onMounted(async () => {})
</script>
