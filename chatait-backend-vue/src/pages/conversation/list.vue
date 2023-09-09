<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <div>
    <common-table ref="tableEle" :request-path="`conversation/list`" :primary-filter-field="'username'" enable-expand>
      <template #expandedRow="slotProps">
        <template
          v-if="
            slotProps.params.row.topic_type.ori_value === TopicTypeMidjourney &&
            slotProps.params.row.role.ori_value === 'assistant'
          "
        >
          <div style="display: flex; width: 100%">
            <div style="flex: 0 0 auto; width: 100px; height: 100px">
              <t-image-viewer
                v-if="slotProps.params.row.mj_data.ori_value.img_url !== ''"
                :images="[slotProps.params.row.mj_data.ori_value.img_url]"
              >
                <template #trigger="{ open }">
                  <t-image
                    class="conversation-content-midjourney-image"
                    :src="slotProps.params.row.mj_data.ori_value.img_url"
                    shape="round"
                    fit="contain"
                    position="left"
                    @click="open"
                  />
                </template>
              </t-image-viewer>
            </div>
            <div style="flex: 1 1 auto; width: 100%; padding: 5px; box-sizing: border-box">
              <span style="font-weight: bold">prompt:</span>
              {{ slotProps.params.row.mj_data.ori_value.prompt }}
            </div>
          </div>
        </template>
        <template v-else>
          <span style="font-weight: bold">content:</span>
          {{ slotProps.params.row.content.ori_value }}
        </template>
      </template>
    </common-table>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import CommonTable from '@/components/common-table/CommonTable.vue'
import { TopicTypeMidjourney } from '@/constants/topic'

const router = useRouter()

const tableEle = ref(null)

onMounted(async () => {})
</script>
