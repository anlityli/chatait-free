<!--
  - Copyright 2024 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <div>
    <common-table
      ref="tableEle"
      :request-path="`user/sensitive-word-list`"
      :operation-column="true"
      :primary-filter-field="'username'"
      :fixed-columns="[{ field: 'operation', direction: 'right' }]"
    >
      <template #operation="slotProps">
        <t-dropdown :min-column-width="88" trigger="click">
          <t-button variant="text">
            操作数据
            <template #suffix>
              <chevron-down-icon></chevron-down-icon>
            </template>
          </t-button>
          <t-dropdown-menu>
            <t-dropdown-item @click="handleOne(slotProps.params.row.id.ori_value)">查看详情</t-dropdown-item>
          </t-dropdown-menu>
        </t-dropdown>
      </template>
    </common-table>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { ChevronDownIcon } from 'tdesign-icons-vue-next'
import CommonTable from '@/components/common-table/CommonTable.vue'

const router = useRouter()

const tableEle = ref(null)

const handleOne = (id: string) => {
  router.push(`/user/sensitive-word-one?id=${id}`)
}
</script>
